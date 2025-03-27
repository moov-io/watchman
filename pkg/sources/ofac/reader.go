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
	"sync"

	"github.com/moov-io/watchman/pkg/download"
)

const (
	estimatedSDNs      = 20000
	estimatedAddresses = 20000
	estimatedAlts      = 10000
	estimatedComments  = 100
)

// Read will consume the file at path and attempt to parse it was a CSV OFAC file.
//
// For more details on the raw OFAC files see https://moov-io.github.io/watchman/file-structure.html
func Read(files download.Files) (*Results, error) {
	results := &Results{
		SDNs:                make([]SDN, 0, estimatedSDNs),
		Addresses:           make(map[string][]Address, estimatedAddresses),
		AlternateIdentities: make(map[string][]AlternateIdentity, estimatedAlts),
		SDNComments:         make(map[string][]SDNComments, estimatedComments),
	}

	hashes := make([]string, 4)

	for filename, file := range files {
		switch strings.ToLower(filepath.Base(filename)) {
		case "add.csv":
			hash, err := csvAddressFile(results, file)
			if err != nil {
				return nil, fmt.Errorf("add.csv: %v", err)
			}
			hashes[0] = hash

		case "alt.csv":
			hash, err := csvAlternateIdentityFile(results, file)
			if err != nil {
				return nil, fmt.Errorf("alt.csv: %v", err)
			}
			hashes[1] = hash

		case "sdn.csv":
			hash, err := csvSDNFile(results, file)
			if err != nil {
				return nil, fmt.Errorf("sdn.csv: %v", err)
			}
			hashes[2] = hash

		case "sdn_comments.csv":
			hash, err := csvSDNCommentsFile(results, file)
			if err != nil {
				return nil, fmt.Errorf("sdn_comments.csv: %v", err)
			}
			hashes[3] = hash

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

func calculateHash(input []byte) string {
	h := sha256.Sum256(input)
	return hex.EncodeToString(h[:])
}

func hashWriter(rc io.ReadCloser) (io.Reader, *bytes.Buffer) {
	var buf bytes.Buffer
	r := io.TeeReader(rc, &buf)
	return r, &buf
}

var (
	// Pool for CSV record slices
	csvRecordPool = sync.Pool{
		New: func() interface{} {
			return make([]string, 12) // Max size needed (SDN file)
		},
	}
)

func csvAddressFile(results *Results, f io.ReadCloser) (string, error) {
	defer f.Close()

	rc, hashbuf := hashWriter(f)

	// Read File into a Variable
	reader := csv.NewReader(rc)
	reader.ReuseRecord = true

	// Pre-allocate record with known field count
	recordCopy := make([]string, 6)

	for {
		readRecord, err := reader.Read()
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
			return "", err
		}

		copy(recordCopy, readRecord)
		record := replaceNull(recordCopy)

		if len(record) != 6 {
			continue
		}

		entityID := record[0]
		if _, ok := results.Addresses[entityID]; !ok {
			results.Addresses[entityID] = make([]Address, 0, 1) // Most entities have ~1 addresses // TODO(adam):
		}
		results.Addresses[entityID] = append(results.Addresses[entityID], Address{
			EntityID:                    entityID,
			AddressID:                   record[1],
			Address:                     record[2],
			CityStateProvincePostalCode: record[3],
			Country:                     record[4],
			AddressRemarks:              record[5],
		})
	}

	hash := calculateHash(hashbuf.Bytes())

	return hash, nil
}

func csvAlternateIdentityFile(results *Results, f io.ReadCloser) (string, error) {
	defer f.Close()

	rc, hashbuf := hashWriter(f)

	// Read File into a Variable
	reader := csv.NewReader(rc)
	reader.ReuseRecord = true

	// Pre-allocate with exact sizes needed
	recordCopy := make([]string, 5)

	for {
		readRecord, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return "", err
		}
		if len(readRecord) != 5 {
			continue
		}

		// Make a copy since we'll store this data
		copy(recordCopy, readRecord)
		record := replaceNull(recordCopy)

		entityID := record[0]
		// Pre-allocate slice if not exists with typical capacity
		if _, ok := results.AlternateIdentities[entityID]; !ok {
			results.AlternateIdentities[entityID] = make([]AlternateIdentity, 0, 2)
		}

		results.AlternateIdentities[entityID] = append(results.AlternateIdentities[entityID], AlternateIdentity{
			EntityID:         record[0],
			AlternateID:      record[1],
			AlternateType:    record[2],
			AlternateName:    record[3],
			AlternateRemarks: record[4],
		})
	}

	hash := calculateHash(hashbuf.Bytes())
	return hash, nil
}

func csvSDNFile(results *Results, f io.ReadCloser) (string, error) {
	defer f.Close()

	rc, hashbuf := hashWriter(f)

	// Create reader and enable record reuse
	reader := csv.NewReader(rc)
	reader.ReuseRecord = true

	// Grab pre-allocated record slice with known field count
	record, ok := csvRecordPool.Get().([]string)
	if !ok {
		return "", fmt.Errorf("unexpected %T from csvRecordPool", record)
	}
	defer csvRecordPool.Put(record)

	recordCopy, ok := csvRecordPool.Get().([]string)
	if !ok {
		return "", fmt.Errorf("unexpected %T from csvRecordPool (as copy)", recordCopy)
	}
	defer csvRecordPool.Put(recordCopy)

	for {
		readRecord, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return "", err
		}

		if len(readRecord) != 12 {
			continue
		}

		// Make a copy since we'll store this data
		copy(recordCopy, readRecord)
		record = replaceNull(recordCopy)

		// Pre-allocate programs slice for common case
		programs := make([]string, 0, 2) // Most entries have 1-2 programs
		programs = append(programs, splitPrograms(record[3])...)

		results.SDNs = append(results.SDNs, SDN{
			EntityID:               record[0],
			SDNName:                record[1],
			SDNType:                record[2],
			Programs:               programs,
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

	return hash, nil
}

func csvSDNCommentsFile(results *Results, f io.ReadCloser) (string, error) {
	defer f.Close()
	rc, hashbuf := hashWriter(f)

	// Create reader and enable record reuse
	reader := csv.NewReader(rc)
	reader.ReuseRecord = true
	reader.LazyQuotes = true

	// Pre-allocate record slices
	recordCopy := make([]string, 2)

	for {
		readRecord, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return "", err
		}

		if len(readRecord) != 2 {
			continue
		}

		// Make a copy since we'll store this data
		copy(recordCopy, readRecord)
		record := replaceNull(recordCopy)

		entityID := record[0]
		// Pre-allocate slice if not exists
		if _, ok := results.SDNComments[entityID]; !ok {
			results.SDNComments[entityID] = make([]SDNComments, 0, 2) // Most entities have 1-2 comments
		}

		results.SDNComments[entityID] = append(results.SDNComments[entityID], SDNComments{
			EntityID:                 entityID,
			RemarksExtended:          record[1],
			DigitalCurrencyAddresses: readDigitalCurrencyAddresses(record[1]),
		})
	}

	hash := calculateHash(hashbuf.Bytes())
	return hash, nil
}

// replaceNull replaces a CSV field that contain -0- with "".  Null values for all four formats consist of "-0-"
// (ASCII characters 45, 48, 45).
func replaceNull(s []string) []string {
	const null = "-0-"
	for i := 0; i < len(s); i++ {
		// Fast path - if it doesn't contain -0-, just trim
		if !strings.Contains(s[i], null) {
			s[i] = strings.TrimSpace(s[i])
			continue
		}
		// Handle null case in place
		b := []byte(s[i])
		j := 0 // write index
		for k := 0; k < len(b); {
			if k+3 <= len(b) && b[k] == '-' && b[k+1] == '0' && b[k+2] == '-' {
				k += 3 // skip -0-
				continue
			}
			b[j] = b[k]
			j++
			k++
		}
		s[i] = strings.TrimSpace(string(b[:j]))
	}
	return s
}

// Some entries in the SDN have a malformed programs list.
// Fields containing lists are supposed to be semicolon delimited, but the programs list doesn't always follow this convention.
// Ex: "SDGT] [IFSR" => "SDGT; IFSR", "SDNTK] [FTO] [SDGT" => "SDNTK; FTO; SDGT"
var prgmReplacer = strings.NewReplacer("] [", "; ", "]", "", "[", "")

func splitPrograms(in string) []string {
	if in == "" {
		return nil
	}

	// Fast path for common case of single program
	if !strings.ContainsAny(in, "[];") {
		return []string{strings.TrimSpace(in)}
	}

	// For complex cases, clean once then split
	cleaned := prgmReplacer.Replace(in)
	if cleaned == "" {
		return nil
	}

	// Split and trim in place
	parts := strings.Split(cleaned, "; ")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func splitRemarks(input string) []string {
	if input == "" {
		return nil
	}

	// Trim suffix once
	if strings.HasSuffix(input, ".") {
		input = input[:len(input)-1]
	}

	// Pre-count separators to allocate exact size
	count := 1
	for i := 0; i < len(input); i++ {
		if input[i] == ';' {
			count++
		}
	}

	// Allocate exact size needed
	parts := make([]string, 0, count)
	start := 0
	for i := 0; i < len(input); i++ {
		if input[i] == ';' {
			if i > start {
				parts = append(parts, input[start:i])
			}
			start = i + 1
		}
	}
	if start < len(input) {
		parts = append(parts, input[start:])
	}
	return parts
}

type remark struct {
	matchedName string
	fullName    string
	value       string
}

func findMatchingRemarks(remarks []string, suffix string) []remark {
	// Pre-allocate with typical size
	out := make([]remark, 0, 1) // Usually only one match
	if suffix == "" {
		return out
	}

	for i := range remarks {
		idx := strings.Index(remarks[i], suffix)
		if idx == -1 {
			continue
		}

		// Use bytes operations instead of multiple string ops
		value := remarks[i][idx+len(suffix):]
		if len(value) > 0 && value[0] == ':' {
			value = value[1:]
		}

		end := len(value)
		if end > 0 {
			if value[end-1] == '.' {
				end--
			}
			if value[end-1] == ';' {
				end--
			}
		}

		out = append(out, remark{
			matchedName: suffix,
			fullName:    remarks[i][:idx+len(suffix)],
			value:       strings.TrimSpace(value[:end]),
		})
	}
	return out
}

func firstValue(values []remark) string {
	if len(values) == 0 {
		return ""
	}
	return values[0].value
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
	// Fast path - check if we even need to parse
	if !strings.Contains(remarks, "Currency Address") {
		return nil
	}

	// Pre-allocate with typical size based on real-world data
	out := make([]DigitalCurrencyAddress, 0, 2)
	if remarks == "" {
		return out
	}

	// Split remarks only once and reuse
	parts := strings.FieldsFunc(remarks, func(r rune) bool {
		return r == ';' || r == '.'
	})

	for i := range parts {
		part := strings.TrimSpace(parts[i])
		if part == "" {
			continue
		}

		// Check each currency
		for _, currency := range digitalCurrencies {
			if !strings.Contains(part, currency) {
				continue
			}

			// Find exact currency match with spaces
			marker := " " + currency + " "
			idx := strings.Index(part, marker)
			if idx == -1 {
				continue
			}

			// Extract the address by taking everything after currency until space/semicolon
			addrStart := idx + len(marker)
			if addrStart >= len(part) {
				continue
			}

			// Find end of address
			addrEnd := addrStart
			for addrEnd < len(part) {
				if part[addrEnd] == ' ' || part[addrEnd] == ';' || part[addrEnd] == '.' {
					break
				}
				addrEnd++
			}

			if addrEnd > addrStart {
				out = append(out, DigitalCurrencyAddress{
					Currency: currency,
					Address:  part[addrStart:addrEnd],
				})
				break // Found currency for this part
			}
		}
	}

	if len(out) == 0 {
		return nil
	}
	return out
}

func mergeSpilloverRecords(sdns []SDN, allComments map[string][]SDNComments) []SDN {
	var b strings.Builder
	for i := range sdns {
		comments := allComments[sdns[i].EntityID]
		if len(comments) == 0 {
			continue
		}

		b.WriteString(sdns[i].Remarks)
		for _, comment := range comments {
			if sdns[i].EntityID == comment.EntityID {
				b.WriteString(comment.RemarksExtended)
			}
		}
		sdns[i].Remarks = b.String()
		b.Reset()
	}
	return sdns
}
