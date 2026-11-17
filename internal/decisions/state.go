package decisions

import (
	"fmt"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

const (
	maxAltNames     = 8
	maxIDs          = 8
	maxAddresses    = 4
	maxCrypto       = 4
	maxAffiliations = 4
	maxHistorical   = 4
)

type screeningState struct {
	Query      screeningEntity   `json:"query"`
	Candidates []screeningEntity `json:"candidates"`
}

type screeningEntity struct {
	Index        *int          `json:"index,omitempty"`
	Name         string        `json:"name,omitempty"`
	Type         string        `json:"entityType,omitempty"`
	Source       string        `json:"sourceList,omitempty"`
	SourceID     string        `json:"sourceID,omitempty"`
	Match        float64       `json:"match,omitempty"`
	AltNames     []string      `json:"altNames,omitempty"`
	BirthDate    string        `json:"birthDate,omitempty"`
	Created      string        `json:"created,omitempty"`
	Built        string        `json:"built,omitempty"`
	IDs          []screeningID `json:"ids,omitempty"`
	IMONumber    string        `json:"imoNumber,omitempty"`
	MMSI         string        `json:"mmsi,omitempty"`
	CallSign     string        `json:"callSign,omitempty"`
	ICAOCode     string        `json:"icaoCode,omitempty"`
	SerialNumber string        `json:"serialNumber,omitempty"`
	SwiftBICs    []string      `json:"swiftBICs,omitempty"`
	Addresses    []string      `json:"addresses,omitempty"`
	Crypto       []string      `json:"cryptoAddresses,omitempty"`
	Affiliations []string      `json:"affiliations,omitempty"`
	FormerNames  []string      `json:"formerNames,omitempty"`
}

type screeningID struct {
	Type       string `json:"type"`
	Country    string `json:"country,omitempty"`
	Identifier string `json:"identifier"`
}

func buildState(query search.Entity[search.Value], results []search.SearchedEntity[search.Value]) screeningState {
	st := screeningState{
		Query:      compactEntity(query, -1, 0),
		Candidates: make([]screeningEntity, 0, len(results)),
	}
	for i, result := range results {
		st.Candidates = append(st.Candidates, compactEntity(result.Entity, i, result.Match))
	}
	return st
}

func compactEntity(entity search.Entity[search.Value], index int, match float64) screeningEntity {
	out := screeningEntity{
		Name:      entity.Name,
		Type:      string(entity.Type),
		Source:    string(entity.Source),
		SourceID:  entity.SourceID,
		SwiftBICs: capStrings(entity.SwiftBICs, maxIDs),
	}
	if index >= 0 {
		idx := index
		out.Index = &idx
		out.Match = match
	}

	switch {
	case entity.Person != nil:
		out.AltNames = capStrings(entity.Person.AltNames, maxAltNames)
		out.BirthDate = formatDate(entity.Person.BirthDate)
		out.IDs = compactIDs(entity.Person.GovernmentIDs)
	case entity.Business != nil:
		out.AltNames = capStrings(entity.Business.AltNames, maxAltNames)
		out.Created = formatDate(entity.Business.Created)
		out.IDs = compactIDs(entity.Business.GovernmentIDs)
	case entity.Organization != nil:
		out.AltNames = capStrings(entity.Organization.AltNames, maxAltNames)
		out.Created = formatDate(entity.Organization.Created)
		out.IDs = compactIDs(entity.Organization.GovernmentIDs)
	case entity.Vessel != nil:
		out.AltNames = capStrings(entity.Vessel.AltNames, maxAltNames)
		out.IMONumber = entity.Vessel.IMONumber
		out.MMSI = entity.Vessel.MMSI
		out.CallSign = entity.Vessel.CallSign
		out.Built = formatDate(entity.Vessel.Built)
	case entity.Aircraft != nil:
		out.AltNames = capStrings(entity.Aircraft.AltNames, maxAltNames)
		out.SerialNumber = entity.Aircraft.SerialNumber
		out.ICAOCode = entity.Aircraft.ICAOCode
		out.Built = formatDate(entity.Aircraft.Built)
	}

	if len(entity.Addresses) > 0 {
		n := min(len(entity.Addresses), maxAddresses)
		out.Addresses = make([]string, 0, n)
		for i := 0; i < n; i++ {
			if formatted := entity.Addresses[i].Format(); formatted != "" {
				out.Addresses = append(out.Addresses, formatted)
			}
		}
	}

	if len(entity.CryptoAddresses) > 0 {
		n := min(len(entity.CryptoAddresses), maxCrypto)
		out.Crypto = make([]string, 0, n)
		for i := 0; i < n; i++ {
			addr := entity.CryptoAddresses[i]
			if addr.Currency == "" && addr.Address == "" {
				continue
			}
			out.Crypto = append(out.Crypto, addr.Currency+":"+addr.Address)
		}
	}

	if len(entity.Affiliations) > 0 {
		n := min(len(entity.Affiliations), maxAffiliations)
		out.Affiliations = make([]string, 0, n)
		for i := 0; i < n; i++ {
			aff := entity.Affiliations[i]
			if aff.EntityName == "" {
				continue
			}
			if aff.Type != "" {
				out.Affiliations = append(out.Affiliations, aff.Type+": "+aff.EntityName)
				continue
			}
			out.Affiliations = append(out.Affiliations, aff.EntityName)
		}
	}

	if len(entity.HistoricalInfo) > 0 {
		n := min(len(entity.HistoricalInfo), maxHistorical)
		out.FormerNames = make([]string, 0, n)
		for i := 0; i < n; i++ {
			item := entity.HistoricalInfo[i]
			if item.Value == "" {
				continue
			}
			out.FormerNames = append(out.FormerNames, item.Value)
		}
	}

	return out
}

func compactIDs(ids []search.GovernmentID) []screeningID {
	if len(ids) == 0 {
		return nil
	}
	n := min(len(ids), maxIDs)
	out := make([]screeningID, 0, n)
	for i := 0; i < n; i++ {
		id := ids[i]
		if id.Identifier == "" {
			continue
		}
		out = append(out, screeningID{
			Type:       string(id.Type),
			Country:    id.Country,
			Identifier: id.Identifier,
		})
	}
	return out
}

func capStrings(in []string, max int) []string {
	if len(in) == 0 {
		return nil
	}
	if len(in) > max {
		return in[:max]
	}
	return in
}

func formatDate(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}

func sameEntityKey(i int) string { return fmt.Sprintf("same_entity_%d", i) }
func actionKey(i int) string     { return fmt.Sprintf("action_%d", i) }
func qualityKey(i int) string    { return fmt.Sprintf("quality_%d", i) }
