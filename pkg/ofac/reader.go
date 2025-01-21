// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/moov-io/watchman/pkg/download"
)

// Read will consume the file at path and attempt to parse it was a CSV OFAC file.
//
// For more details on the raw OFAC files see https://moov-io.github.io/watchman/file-structure.html
func Read(files download.Files) (*Results, error) {
	results := &Results{
		SDNs:                []SDN{},
		Addresses:           make(map[string][]Address, 0),
		AlternateIdentities: make(map[string][]AlternateIdentity, 0),
		SDNComments:         make(map[string][]SDNComments, 0),
	}

	hashes := make([]string, 4)

	for filename, file := range files {
		switch strings.ToLower(filepath.Base(filename)) {
		case "add.csv":
			res, hash, err := csvAddressFile(file)
			if err != nil {
				return nil, fmt.Errorf("add.csv: %v", err)
			}
			hashes[0] = hash
			results.append(res)

		case "alt.csv":
			res, hash, err := csvAlternateIdentityFile(file)
			if err != nil {
				return nil, fmt.Errorf("alt.csv: %v", err)
			}
			hashes[1] = hash
			results.append(res)

		case "sdn.csv":
			res, hash, err := csvSDNFile(file)
			if err != nil {
				return nil, fmt.Errorf("sdn.csv: %v", err)
			}
			hashes[2] = hash
			results.append(res)

		case "sdn_comments.csv":
			res, hash, err := csvSDNCommentsFile(file)
			if err != nil {
				return nil, fmt.Errorf("sdn_comments.csv: %v", err)
			}
			hashes[3] = hash
			results.append(res)

		default:
			file.Close()
			return nil, fmt.Errorf("error: file %s does not have a handler for processing", filename)
		}
	}

	// Join the hashes together
	var buf bytes.Buffer
	for _, h := range hashes {
		buf.WriteString(h)
	}
	results.ListHash = calculateHash(buf.Bytes())

	// Merge extended comments into SDN
	results.SDNs = mergeSpilloverRecords(results.SDNs, results.SDNComments)

	return results, nil
}

type Results struct {
	// SDNs returns an array of OFAC Specially Designated Nationals
	SDNs []SDN `json:"sdn"`

	// Addresses returns an array of OFAC Specially Designated National Addresses
	Addresses map[string][]Address `json:"address"`

	// AlternateIdentities returns an array of OFAC Specially Designated National Alternate Identity
	AlternateIdentities map[string][]AlternateIdentity `json:"alternateIdentity"`

	// SDNComments returns an array of OFAC Specially Designated National Comments
	SDNComments map[string][]SDNComments `json:"sdnComments"`

	ListHash string
}

func (r *Results) append(rr *Results) {
	r.SDNs = append(r.SDNs, rr.SDNs...)

	if rr.Addresses != nil {
		for _, addresses := range rr.Addresses {
			for _, addr := range addresses {
				r.Addresses[addr.EntityID] = append(r.Addresses[addr.EntityID], addr)
			}
		}
	}

	if rr.AlternateIdentities != nil {
		for _, alts := range rr.AlternateIdentities {
			for _, alt := range alts {
				r.AlternateIdentities[alt.EntityID] = append(r.AlternateIdentities[alt.EntityID], alt)
			}
		}
	}

	if rr.SDNComments != nil {
		for _, comments := range rr.SDNComments {
			for _, comment := range comments {
				r.SDNComments[comment.EntityID] = append(r.SDNComments[comment.EntityID], comment)
			}
		}
	}
}

func calculateHash(input []byte) string {
	h := sha256.Sum256(input)
	return hex.EncodeToString(h[:])
}

func hashWriter(rc io.ReadCloser) (io.Reader, *bytes.Buffer) {
	var buf bytes.Buffer
	r := io.TeeReader(rc, &buf)
	return r, &buf
}

func csvAddressFile(f io.ReadCloser) (*Results, string, error) {
	defer f.Close()

	rc, hashbuf := hashWriter(f)

	out := make(map[string][]Address, 0)

	// Read File into a Variable
	reader := csv.NewReader(rc)
	for {
		record, err := reader.Read()
		if err != nil {
			// reached the last line
			if errors.Is(err, io.EOF) {
				break
			}
			// malformed row
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return nil, "", err
		}
		if len(record) != 6 {
			continue
		}

		record = replaceNull(record)

		entityID := record[0]
		out[entityID] = append(out[entityID], Address{
			EntityID:                    entityID,
			AddressID:                   record[1],
			Address:                     record[2],
			CityStateProvincePostalCode: record[3],
			Country:                     record[4],
			AddressRemarks:              record[5],
		})
	}

	hash := calculateHash(hashbuf.Bytes())

	return &Results{Addresses: out}, hash, nil
}

func csvAlternateIdentityFile(f io.ReadCloser) (*Results, string, error) {
	defer f.Close()

	rc, hashbuf := hashWriter(f)

	out := make(map[string][]AlternateIdentity, 0)

	// Read File into a Variable
	reader := csv.NewReader(rc)
	for {
		record, err := reader.Read()
		if err != nil {
			// reached the last line
			if errors.Is(err, io.EOF) {
				break
			}
			// malformed row
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return nil, "", err
		}
		if len(record) != 5 {
			continue
		}
		record = replaceNull(record)

		entityID := record[0]
		out[entityID] = append(out[entityID], AlternateIdentity{
			EntityID:         record[0],
			AlternateID:      record[1],
			AlternateType:    record[2],
			AlternateName:    record[3],
			AlternateRemarks: record[4],
		})
	}

	hash := calculateHash(hashbuf.Bytes())

	return &Results{AlternateIdentities: out}, hash, nil
}

func csvSDNFile(f io.ReadCloser) (*Results, string, error) {
	defer f.Close()

	rc, hashbuf := hashWriter(f)

	var out []SDN

	// Read File into a Variable
	reader := csv.NewReader(rc)
	for {
		record, err := reader.Read()
		if err != nil {
			// reached the last line
			if errors.Is(err, io.EOF) {
				break
			}
			// malformed row
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return nil, "", err
		}
		if len(record) != 12 {
			continue
		}
		record = replaceNull(record)
		out = append(out, SDN{
			EntityID:               record[0],
			SDNName:                record[1],
			SDNType:                record[2],
			Programs:               splitPrograms(record[3]),
			Title:                  record[4],
			CallSign:               record[5],
			VesselType:             record[6],
			Tonnage:                record[7],
			GrossRegisteredTonnage: record[8],
			VesselFlag:             record[9],
			VesselOwner:            record[10],
			Remarks:                record[11],
		})
	}

	hash := calculateHash(hashbuf.Bytes())

	return &Results{SDNs: out}, hash, nil
}

func csvSDNCommentsFile(f io.ReadCloser) (*Results, string, error) {
	defer f.Close()

	rc, hashbuf := hashWriter(f)

	// Read File into a Variable
	r := csv.NewReader(rc)
	r.LazyQuotes = true

	// Loop through lines & turn into object
	out := make(map[string][]SDNComments, 0)
	for {
		line, err := r.Read()
		if err != nil {
			// reached the last line
			if errors.Is(err, io.EOF) {
				break
			}
			// malformed row
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return nil, "", err
		}
		if len(line) != 2 {
			continue
		}
		line = replaceNull(line)

		entityID := line[0]
		out[entityID] = append(out[entityID], SDNComments{
			EntityID:                 entityID,
			RemarksExtended:          line[1],
			DigitalCurrencyAddresses: readDigitalCurrencyAddresses(line[1]),
		})
	}

	hash := calculateHash(hashbuf.Bytes())

	return &Results{SDNComments: out}, hash, nil
}

// replaceNull replaces a CSV field that contain -0- with "".  Null values for all four formats consist of "-0-"
// (ASCII characters 45, 48, 45).
func replaceNull(s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.TrimSpace(strings.Replace(s[i], "-0-", "", -1))
	}
	return s
}

// Some entries in the SDN have a malformed programs list.
// Fields containing lists are supposed to be semicolon delimited, but the programs list doesn't always follow this convention.
// Ex: "SDGT] [IFSR" => "SDGT; IFSR", "SDNTK] [FTO] [SDGT" => "SDNTK; FTO; SDGT"
var prgmReplacer = strings.NewReplacer("] [", "; ", "]", "", "[", "")

func cleanPrgmsList(s string) string {
	return strings.TrimSpace(prgmReplacer.Replace(s))
}

func splitPrograms(in string) []string {
	norm := cleanPrgmsList(in)
	return strings.Split(norm, "; ")
}

func splitRemarks(input string) []string {
	return strings.Split(strings.TrimSuffix(input, "."), ";")
}

type remark struct {
	matchedName string
	fullName    string
	value       string
}

func findMatchingRemarks(remarks []string, suffix string) []remark {
	var out []remark
	if suffix == "" {
		return out
	}
	for i := range remarks {
		idx := strings.Index(remarks[i], suffix)
		if idx == -1 {
			continue // not found
		}

		value := remarks[i][idx+len(suffix):]
		value = strings.TrimPrefix(value, ":") // identifiers can end with a colon
		value = strings.TrimSuffix(value, ";")
		value = strings.TrimSuffix(value, ".")

		out = append(out, remark{
			matchedName: suffix,
			fullName:    remarks[i][:idx+len(suffix)],
			value:       strings.TrimSpace(value),
		})
	}
	return out
}

func findRemarkValues(remarks []string, suffix string) []string {
	found := findMatchingRemarks(remarks, suffix)
	var out []string
	for i := range found {
		out = append(out, found[i].value)
	}
	return out
}

func firstValue(values []remark) string {
	if len(values) == 0 {
		return ""
	}
	return values[0].value
}

func withFirstF[T any](values []remark, f func(remark) T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}
	return f(values[0])
}

func withFirstP[T any](values []remark, f func(remark) *T) *T {
	if len(values) == 0 {
		return nil
	}
	return f(values[0])
}

var (
	digitalCurrencies = []string{
		"XBT",  // Bitcoin
		"ETH",  // Ethereum
		"XMR",  // Monero
		"LTC",  // Litecoin
		"ZEC",  // ZCash
		"DASH", // Dash
		"BTG",  // Bitcoin Gold
		"ETC",  // Ethereum Classic
		"BSV",  // Bitcoin Satoshi Vision
		"BCH",  // Bitcoin Cash
		"XVG",  // Verge
		"USDC", // USD Coin
		"USDT", // USD Tether
		"XRP",  // Ripple
		"TRX",  // Tron
		"ARB",  // Arbitrum
		"BSC",  // Binance Smart Chain
	}
)

func readDigitalCurrencyAddresses(remarks string) []DigitalCurrencyAddress {
	var out []DigitalCurrencyAddress

	// The format is semicolon delineated, but "Digital Currency Address" is sometimes truncated badly
	//
	//   alt. Digital Currency Address - XBT 12jVCWW1ZhTLA5yVnroEJswqKwsfiZKsax;
	//
	parts := splitRemarks(remarks)
	for i := range parts {
		// Check if the currency is in the remark
		var addressIndex int
		for j := range digitalCurrencies {
			idx := strings.Index(parts[i], fmt.Sprintf(" %s ", digitalCurrencies[j]))
			if idx > -1 {
				addressIndex = idx
				break
			}
		}
		if addressIndex > 0 {
			fields := strings.Fields(parts[i][addressIndex:])
			if len(fields) < 2 {
				break // bad parsing
			}
			out = append(out, DigitalCurrencyAddress{
				Currency: fields[0],
				Address:  fields[1],
			})
			continue
		}
	}

	return out
}

func mergeSpilloverRecords(sdns []SDN, allComments map[string][]SDNComments) []SDN {
	for i := range sdns {
		comments := allComments[sdns[i].EntityID]

		for _, comment := range comments {
			if sdns[i].EntityID == comment.EntityID {
				sdns[i].Remarks += comment.RemarksExtended // has to be index to update
			}
		}
	}
	return sdns
}
