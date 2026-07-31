package corpusfile

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"sync"
	"time"

	"github.com/klauspost/compress/zstd"
	"github.com/moov-io/watchman/pkg/search"
)

func init() {
	// SourceData is interface{}; register concrete list types used in tests/prod mappers.
	// Additional source types can register in their packages via gob.Register.
	gob.Register(map[string]interface{}{})
	gob.Register(map[string]string{})
}

// DefaultColdCodec is used when EncodeCold is called without an explicit codec.
// Gob+zstd is the production default: smaller cold files, modest hydrate CPU.
var DefaultColdCodec uint8 = CodecGobZstd

// ColdPayload is the per-entity blob stored outside the hot working set.
// Only fields that scoreSimilarityFast does not need on the index side belong here.
//
// Hot path keeps PreparedFields (names, tokens, prepared addresses, normalized phones)
// plus emails (exact match still reads Contact.EmailAddresses), crypto, IDs, etc.
type ColdPayload struct {
	// SourceData is the original list record (OFAC SDN, CSL rows, etc.).
	SourceData any

	// Display / API fields not required for scoring once PreparedFields exist.
	Addresses []search.Address
	// Raw phone/fax/websites; scoring uses PreparedFields.Contact for phone/fax.
	// Emails stay hot (exact match reads Contact.EmailAddresses).
	PhoneNumbers []string
	FaxNumbers   []string
	Websites     []string

	// Supporting text stripped from hot sanctions/affiliation/historical rows.
	SanctionsDescription string
	AffiliationDetails   []string // parallel to hot Affiliations order
	HistoricalDates      []time.Time

	// Type-level name mirrors (Entity.Name + PreparedFields cover scoring).
	PersonName         string
	PersonAltNames     []string
	BusinessName       string
	BusinessAltNames   []string
	OrganizationName   string
	OrganizationAltNames []string
	AircraftName       string
	AircraftAltNames   []string
	VesselName         string
	VesselAltNames     []string
}

// SplitResult is one entity divided into hot (scoring) and cold (response) parts.
type SplitResult struct {
	// Hot is a full Entity suitable for Similarity, with cold fields cleared.
	Hot search.Entity[search.Value]

	// Cold holds fields needed to rehydrate API responses.
	Cold ColdPayload
}

// SplitEntity moves non-scoring fields into ColdPayload.
// The returned Hot entity is safe for pkg/search.Similarity* paths.
func SplitEntity(e search.Entity[search.Value]) SplitResult {
	cold := ColdPayload{
		SourceData:   e.SourceData,
		Addresses:    e.Addresses,
		PhoneNumbers: e.Contact.PhoneNumbers,
		FaxNumbers:   e.Contact.FaxNumbers,
		Websites:     e.Contact.Websites,
	}

	hot := e
	hot.SourceData = nil
	hot.Addresses = nil
	hot.Contact.PhoneNumbers = nil
	hot.Contact.FaxNumbers = nil
	hot.Contact.Websites = nil
	// Emails remain on hot.Contact for exact match.

	if e.SanctionsInfo != nil {
		cold.SanctionsDescription = e.SanctionsInfo.Description
		if e.SanctionsInfo.Description != "" {
			cp := *e.SanctionsInfo
			cp.Description = ""
			hot.SanctionsInfo = &cp
		}
	}

	if n := len(e.Affiliations); n > 0 {
		cold.AffiliationDetails = make([]string, n)
		hot.Affiliations = make([]search.Affiliation, n)
		for i, a := range e.Affiliations {
			cold.AffiliationDetails[i] = a.Details
			hot.Affiliations[i] = search.Affiliation{
				EntityName: a.EntityName,
				Type:       a.Type,
				// Details cold
			}
		}
	}

	if n := len(e.HistoricalInfo); n > 0 {
		cold.HistoricalDates = make([]time.Time, n)
		hot.HistoricalInfo = make([]search.HistoricalInfo, n)
		for i, h := range e.HistoricalInfo {
			cold.HistoricalDates[i] = h.Date
			hot.HistoricalInfo[i] = search.HistoricalInfo{
				Type:  h.Type,
				Value: h.Value,
			}
		}
	}

	// Type-level Name/AltNames are mirrored in Entity.Name + PreparedFields.
	if e.Person != nil {
		cold.PersonName = e.Person.Name
		cold.PersonAltNames = e.Person.AltNames
		p := *e.Person
		p.Name = ""
		p.AltNames = nil
		hot.Person = &p
	}
	if e.Business != nil {
		cold.BusinessName = e.Business.Name
		cold.BusinessAltNames = e.Business.AltNames
		b := *e.Business
		b.Name = ""
		b.AltNames = nil
		hot.Business = &b
	}
	if e.Organization != nil {
		cold.OrganizationName = e.Organization.Name
		cold.OrganizationAltNames = e.Organization.AltNames
		o := *e.Organization
		o.Name = ""
		o.AltNames = nil
		hot.Organization = &o
	}
	if e.Aircraft != nil {
		cold.AircraftName = e.Aircraft.Name
		cold.AircraftAltNames = e.Aircraft.AltNames
		a := *e.Aircraft
		a.Name = ""
		a.AltNames = nil
		hot.Aircraft = &a
	}
	if e.Vessel != nil {
		cold.VesselName = e.Vessel.Name
		cold.VesselAltNames = e.Vessel.AltNames
		v := *e.Vessel
		v.Name = ""
		v.AltNames = nil
		hot.Vessel = &v
	}

	return SplitResult{Hot: hot, Cold: cold}
}

// Reattach merges a cold payload back onto a hot entity for API responses.
func Reattach(hot search.Entity[search.Value], cold ColdPayload) search.Entity[search.Value] {
	out := hot
	out.SourceData = cold.SourceData
	out.Addresses = cold.Addresses
	out.Contact.PhoneNumbers = cold.PhoneNumbers
	out.Contact.FaxNumbers = cold.FaxNumbers
	out.Contact.Websites = cold.Websites

	if out.SanctionsInfo != nil && cold.SanctionsDescription != "" {
		cp := *out.SanctionsInfo
		cp.Description = cold.SanctionsDescription
		out.SanctionsInfo = &cp
	} else if out.SanctionsInfo == nil && cold.SanctionsDescription != "" {
		out.SanctionsInfo = &search.SanctionsInfo{Description: cold.SanctionsDescription}
	}

	if len(out.Affiliations) > 0 && len(cold.AffiliationDetails) > 0 {
		affs := make([]search.Affiliation, len(out.Affiliations))
		copy(affs, out.Affiliations)
		for i := range affs {
			if i < len(cold.AffiliationDetails) {
				affs[i].Details = cold.AffiliationDetails[i]
			}
		}
		out.Affiliations = affs
	}

	if len(out.HistoricalInfo) > 0 && len(cold.HistoricalDates) > 0 {
		hist := make([]search.HistoricalInfo, len(out.HistoricalInfo))
		copy(hist, out.HistoricalInfo)
		for i := range hist {
			if i < len(cold.HistoricalDates) {
				hist[i].Date = cold.HistoricalDates[i]
			}
		}
		out.HistoricalInfo = hist
	}

	if out.Person != nil {
		p := *out.Person
		p.Name = cold.PersonName
		p.AltNames = cold.PersonAltNames
		out.Person = &p
	}
	if out.Business != nil {
		b := *out.Business
		b.Name = cold.BusinessName
		b.AltNames = cold.BusinessAltNames
		out.Business = &b
	}
	if out.Organization != nil {
		o := *out.Organization
		o.Name = cold.OrganizationName
		o.AltNames = cold.OrganizationAltNames
		out.Organization = &o
	}
	if out.Aircraft != nil {
		a := *out.Aircraft
		a.Name = cold.AircraftName
		a.AltNames = cold.AircraftAltNames
		out.Aircraft = &a
	}
	if out.Vessel != nil {
		v := *out.Vessel
		v.Name = cold.VesselName
		v.AltNames = cold.VesselAltNames
		out.Vessel = &v
	}

	return out
}

// EncodeCold serializes a cold payload using DefaultColdCodec.
func EncodeCold(c ColdPayload) ([]byte, error) {
	return EncodeColdCodec(c, DefaultColdCodec)
}

// EncodeColdCodec serializes with an explicit codec byte.
func EncodeColdCodec(c ColdPayload, codec uint8) ([]byte, error) {
	var body bytes.Buffer
	if err := gob.NewEncoder(&body).Encode(c); err != nil {
		return nil, fmt.Errorf("gob encode cold: %w", err)
	}

	switch codec {
	case CodecGob:
		out := make([]byte, 1+body.Len())
		out[0] = CodecGob
		copy(out[1:], body.Bytes())
		return out, nil

	case CodecGobZstd:
		enc := getZstdEncoder()
		compressed := enc.EncodeAll(body.Bytes(), make([]byte, 0, body.Len()/2))
		putZstdEncoder(enc)
		out := make([]byte, 1+len(compressed))
		out[0] = CodecGobZstd
		copy(out[1:], compressed)
		return out, nil

	default:
		return nil, errCodec(codec)
	}
}

// DecodeCold parses a cold blob produced by EncodeCold / EncodeColdCodec.
func DecodeCold(b []byte) (ColdPayload, error) {
	var c ColdPayload
	if len(b) < 1 {
		return c, errShort("cold blob", 1, len(b))
	}

	payload := b[1:]
	switch b[0] {
	case CodecGob:
		// raw gob
	case CodecGobZstd:
		dec := getZstdDecoder()
		var err error
		payload, err = dec.DecodeAll(payload, make([]byte, 0, len(payload)*2))
		putZstdDecoder(dec)
		if err != nil {
			return c, fmt.Errorf("zstd decode: %w", err)
		}
	default:
		return c, errCodec(b[0])
	}

	if err := gob.NewDecoder(bytes.NewReader(payload)).Decode(&c); err != nil {
		return c, fmt.Errorf("gob decode cold: %w", err)
	}
	return c, nil
}

var (
	zstdEncPool = sync.Pool{New: func() any {
		enc, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))
		if err != nil {
			panic(err) //nolint:forbidigo // constructor failure is programmer/env error
		}
		return enc
	}}
	zstdDecPool = sync.Pool{New: func() any {
		dec, err := zstd.NewReader(nil)
		if err != nil {
			panic(err) //nolint:forbidigo
		}
		return dec
	}}
)

func getZstdEncoder() *zstd.Encoder { return zstdEncPool.Get().(*zstd.Encoder) }
func putZstdEncoder(e *zstd.Encoder) { zstdEncPool.Put(e) }
func getZstdDecoder() *zstd.Decoder  { return zstdDecPool.Get().(*zstd.Decoder) }
func putZstdDecoder(d *zstd.Decoder) { zstdDecPool.Put(d) }

// crc of blob bytes (for directory integrity checks).
func blobCRC(b []byte) uint32 {
	return crc32.ChecksumIEEE(b)
}

// nowUnixNano is broken out for tests.
var nowUnixNano = func() int64 { return time.Now().UnixNano() }
