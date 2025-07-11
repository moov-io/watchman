package linksim

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/dgryski/go-minhash"
	"github.com/spaolacci/murmur3"
	"gopkg.in/go-dedup/simhash.v1"
)

// Hashes represents the output for database storage, mapping field paths to their hash values.
type Hashes map[string]interface{}

// GenerateHashes computes hashes for each field in the Entity, suitable for database storage and sorting.
func GenerateHashes(e search.Entity[search.Value]) Hashes {
	hashes := make(Hashes)

	// Use normalized fields from PreparedFields where available
	hashes["Name"] = calculateSimhash(e.PreparedFields.Name)
	if len(e.PreparedFields.AltNames) > 0 {
		hashes["AltNames"] = calculateMinhash(e.PreparedFields.AltNames, 128) // 128 hash functions for minhash
	}
	hashes["EntityType"] = hashString(string(e.Type))
	hashes["Source"] = hashString(string(e.Source))
	hashes["SourceID"] = calculateSimhash(e.SourceID)

	// Person
	if e.Person != nil {
		hashes["Person.Name"] = calculateSimhash(e.Person.Name)
		if len(e.Person.AltNames) > 0 {
			hashes["Person.AltNames"] = calculateMinhash(normalizeNames(e.Person.AltNames), 128)
		}
		hashes["Person.Gender"] = hashString(string(e.Person.Gender))
		if e.Person.BirthDate != nil {
			hashes["Person.BirthDate"] = hashDate(e.Person.BirthDate)
		}
		hashes["Person.PlaceOfBirth"] = calculateSimhash(e.Person.PlaceOfBirth)
		if e.Person.DeathDate != nil {
			hashes["Person.DeathDate"] = hashDate(e.Person.DeathDate)
		}
		if len(e.Person.Titles) > 0 {
			hashes["Person.Titles"] = calculateMinhash(normalizeNames(e.Person.Titles), 128)
		}
		if len(e.Person.GovernmentIDs) > 0 {
			for i, id := range e.Person.GovernmentIDs {
				prefix := fmt.Sprintf("Person.GovernmentIDs[%d]", i)
				hashes[prefix+".Type"] = hashString(string(id.Type))
				hashes[prefix+".Country"] = calculateSimhash(id.Country)
				hashes[prefix+".Identifier"] = calculateSimhash(id.Identifier)
			}
		}
	}

	// Business
	if e.Business != nil {
		hashes["Business.Name"] = calculateSimhash(e.Business.Name)
		if len(e.Business.AltNames) > 0 {
			hashes["Business.AltNames"] = calculateMinhash(normalizeNames(e.Business.AltNames), 128)
		}
		if e.Business.Created != nil {
			hashes["Business.Created"] = hashDate(e.Business.Created)
		}
		if e.Business.Dissolved != nil {
			hashes["Business.Dissolved"] = hashDate(e.Business.Dissolved)
		}
		if len(e.Business.GovernmentIDs) > 0 {
			for i, id := range e.Business.GovernmentIDs {
				prefix := fmt.Sprintf("Business.GovernmentIDs[%d]", i)
				hashes[prefix+".Type"] = hashString(string(id.Type))
				hashes[prefix+".Country"] = calculateSimhash(id.Country)
				hashes[prefix+".Identifier"] = calculateSimhash(id.Identifier)
			}
		}
	}

	// Organization (similar to Business)
	if e.Organization != nil {
		hashes["Organization.Name"] = calculateSimhash(e.Organization.Name)
		if len(e.Organization.AltNames) > 0 {
			hashes["Organization.AltNames"] = calculateMinhash(normalizeNames(e.Organization.AltNames), 128)
		}
		if e.Organization.Created != nil {
			hashes["Organization.Created"] = hashDate(e.Organization.Created)
		}
		if e.Organization.Dissolved != nil {
			hashes["Organization.Dissolved"] = hashDate(e.Organization.Dissolved)
		}
		if len(e.Organization.GovernmentIDs) > 0 {
			for i, id := range e.Organization.GovernmentIDs {
				prefix := fmt.Sprintf("Organization.GovernmentIDs[%d]", i)
				hashes[prefix+".Type"] = hashString(string(id.Type))
				hashes[prefix+".Country"] = calculateSimhash(id.Country)
				hashes[prefix+".Identifier"] = calculateSimhash(id.Identifier)
			}
		}
	}

	// Aircraft
	if e.Aircraft != nil {
		hashes["Aircraft.Name"] = calculateSimhash(e.Aircraft.Name)
		if len(e.Aircraft.AltNames) > 0 {
			hashes["Aircraft.AltNames"] = calculateMinhash(normalizeNames(e.Aircraft.AltNames), 128)
		}
		hashes["Aircraft.Type"] = hashString(string(e.Aircraft.Type))
		hashes["Aircraft.Flag"] = calculateSimhash(e.Aircraft.Flag)
		if e.Aircraft.Built != nil {
			hashes["Aircraft.Built"] = hashDate(e.Aircraft.Built)
		}
		hashes["Aircraft.ICAOCode"] = calculateSimhash(e.Aircraft.ICAOCode)
		hashes["Aircraft.Model"] = calculateSimhash(e.Aircraft.Model)
		hashes["Aircraft.SerialNumber"] = calculateSimhash(e.Aircraft.SerialNumber)
	}

	// Vessel
	if e.Vessel != nil {
		hashes["Vessel.Name"] = calculateSimhash(e.Vessel.Name)
		if len(e.Vessel.AltNames) > 0 {
			hashes["Vessel.AltNames"] = calculateMinhash(normalizeNames(e.Vessel.AltNames), 128)
		}
		hashes["Vessel.IMONumber"] = calculateSimhash(e.Vessel.IMONumber)
		hashes["Vessel.Type"] = hashString(string(e.Vessel.Type))
		hashes["Vessel.Flag"] = calculateSimhash(e.Vessel.Flag)
		if e.Vessel.Built != nil {
			hashes["Vessel.Built"] = hashDate(e.Vessel.Built)
		}
		hashes["Vessel.Model"] = calculateSimhash(e.Vessel.Model)
		hashes["Vessel.MMSI"] = calculateSimhash(e.Vessel.MMSI)
		hashes["Vessel.CallSign"] = calculateSimhash(e.Vessel.CallSign)
		hashes["Vessel.Owner"] = calculateSimhash(e.Vessel.Owner)
		// Numerical fields: simple hash of string representation
		hashes["Vessel.Tonnage"] = hashString(fmt.Sprintf("%d", e.Vessel.Tonnage))
		hashes["Vessel.GrossRegisteredTonnage"] = hashString(fmt.Sprintf("%d", e.Vessel.GrossRegisteredTonnage))
	}

	// ContactInfo
	if len(e.PreparedFields.Contact.EmailAddresses) > 0 {
		hashes["Contact.EmailAddresses"] = calculateMinhash(normalizeNames(e.Contact.EmailAddresses), 128)
	}
	if len(e.PreparedFields.Contact.PhoneNumbers) > 0 {
		hashes["Contact.PhoneNumbers"] = calculateMinhash(e.PreparedFields.Contact.PhoneNumbers, 128)
	}
	if len(e.PreparedFields.Contact.FaxNumbers) > 0 {
		hashes["Contact.FaxNumbers"] = calculateMinhash(e.PreparedFields.Contact.FaxNumbers, 128)
	}
	if len(e.Contact.Websites) > 0 {
		hashes["Contact.Websites"] = calculateMinhash(normalizeNames(e.Contact.Websites), 128)
	}

	// Addresses
	if len(e.PreparedFields.Addresses) > 0 {
		for i, addr := range e.PreparedFields.Addresses {
			prefix := fmt.Sprintf("Addresses[%d]", i)
			hashes[prefix+".Line1"] = calculateSimhash(addr.Line1)
			if len(addr.Line1Fields) > 0 {
				hashes[prefix+".Line1Fields"] = calculateMinhash(addr.Line1Fields, 128)
			}
			hashes[prefix+".Line2"] = calculateSimhash(addr.Line2)
			if len(addr.Line2Fields) > 0 {
				hashes[prefix+".Line2Fields"] = calculateMinhash(addr.Line2Fields, 128)
			}
			hashes[prefix+".City"] = calculateSimhash(addr.City)
			if len(addr.CityFields) > 0 {
				hashes[prefix+".CityFields"] = calculateMinhash(addr.CityFields, 128)
			}
			hashes[prefix+".PostalCode"] = calculateSimhash(addr.PostalCode)
			hashes["Addresses["+strconv.Itoa(i)+"].State"] = calculateSimhash(addr.State)
			hashes[prefix+".Country"] = calculateSimhash(addr.Country)
		}
	}

	// CryptoAddresses
	if len(e.CryptoAddresses) > 0 {
		for i, crypto := range e.CryptoAddresses {
			prefix := fmt.Sprintf("CryptoAddresses[%d]", i)
			hashes[prefix+".Currency"] = hashString(crypto.Currency)
			hashes[prefix+".Address"] = calculateSimhash(crypto.Address)
		}
	}

	// Affiliations
	if len(e.Affiliations) > 0 {
		for i, aff := range e.Affiliations {
			prefix := fmt.Sprintf("Affiliations[%d]", i)
			hashes[prefix+".EntityName"] = calculateSimhash(aff.EntityName)
			hashes[prefix+".Type"] = hashString(aff.Type)
			hashes[prefix+".Details"] = calculateSimhash(aff.Details)
		}
	}

	// SanctionsInfo
	if e.SanctionsInfo != nil {
		if len(e.SanctionsInfo.Programs) > 0 {
			hashes["SanctionsInfo.Programs"] = calculateMinhash(normalizeNames(e.SanctionsInfo.Programs), 128)
		}
		hashes["SanctionsInfo.Secondary"] = hashString(fmt.Sprintf("%t", e.SanctionsInfo.Secondary))
		hashes["SanctionsInfo.Description"] = calculateSimhash(e.SanctionsInfo.Description)
	}

	// HistoricalInfo
	if len(e.HistoricalInfo) > 0 {
		for i, hist := range e.HistoricalInfo {
			prefix := fmt.Sprintf("HistoricalInfo[%d]", i)
			hashes[prefix+".Type"] = hashString(hist.Type)
			hashes[prefix+".Value"] = calculateSimhash(hist.Value)
			hashes[prefix+".Date"] = hashDate(&hist.Date)
		}
	}

	return hashes
}

// calculateSimhash computes a simhash for a string using the go-dedup/simhash library.
func calculateSimhash(s string) uint64 {
	if s == "" {
		return 0
	}
	fs := simhash.NewWordFeatureSet([]byte(prepare.LowerAndRemovePunctuation(s)))
	return simhash.Simhash(fs)
}

// calculateMinhash computes a minhash signature for a list of strings.
func calculateMinhash(list []string, k int) []uint64 {
	if len(list) == 0 {
		return nil
	}
	// Create a 64-bit hash function using Murmur3
	hashFunc := func(data []byte) uint64 {
		h := murmur3.New64()
		h.Write(data)
		return h.Sum64()
	}
	// Initialize BottomK with the hash function and k
	mh := minhash.NewBottomK(hashFunc, k)
	for _, item := range list {
		if item != "" {
			// Normalize the string before hashing
			normalized := prepare.LowerAndRemovePunctuation(item)
			mh.Push([]byte(normalized))
		}
	}
	return mh.Signature()
}

// hashString computes a simple hash for categorical fields or simple values.
func hashString(s string) uint64 {
	if s == "" {
		return 0
	}
	h := sha256.Sum256([]byte(s))
	return binary.BigEndian.Uint64(h[:8])
}

// hashDate computes a hash for a date by discretizing to year.
func hashDate(t *time.Time) uint64 {
	if t == nil {
		return 0
	}
	return uint64(t.Year())
}

const (
	DefaultBuckets = 256
)

// GenerateBuckets maps hashes to buckets for blocking, using a simple modulo operation.
func (h Hashes) GenerateBuckets(numBuckets int) map[string][]int {
	buckets := make(map[string][]int)
	for field, hash := range h {
		switch v := hash.(type) {
		case uint64:
			if v != 0 {
				bucket := int(v % uint64(numBuckets))
				buckets[field] = []int{bucket}
			}
		case []uint64:
			buckets[field] = make([]int, len(v))
			for i, val := range v {
				buckets[field][i] = int(val % uint64(numBuckets))
			}
		default:
			fmt.Printf("unexpected %v - %v (%T)\n", field, hash, hash)
		}
	}
	return buckets
}

// normalizeNames normalizes a slice of strings by lowercasing and removing punctuation.
func normalizeNames(altNames []string) []string {
	if len(altNames) == 0 {
		return nil
	}

	out := make([]string, len(altNames))
	for idx := range altNames {
		out[idx] = prepare.LowerAndRemovePunctuation(altNames[idx])
	}
	return out
}
