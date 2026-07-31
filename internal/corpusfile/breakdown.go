package corpusfile

import (
	"fmt"
	"sort"

	"github.com/moov-io/watchman/pkg/search"
)

// FieldBreakdown estimates gob-encoded contribution of entity field groups.
// Used to decide what to move cold next; not a precise RSS accounting.
type FieldBreakdown struct {
	Entities int
	Groups   []FieldGroupSize
	Total    int64
}

// FieldGroupSize is one logical field group.
type FieldGroupSize struct {
	Name  string
	Bytes int64
	Pct   float64
}

func (b FieldBreakdown) String() string {
	s := fmt.Sprintf("entities=%d total=%.2fMB\n", b.Entities, float64(b.Total)/(1<<20))
	for _, g := range b.Groups {
		s += fmt.Sprintf("  %-22s %8.2f MB  %5.1f%%\n", g.Name, float64(g.Bytes)/(1<<20), g.Pct)
	}
	return s
}

// MeasureFieldBreakdown gob-encodes each field group in isolation across the corpus.
func MeasureFieldBreakdown(entities []search.Entity[search.Value]) (FieldBreakdown, error) {
	type bucket struct {
		name string
		fn   func(e search.Entity[search.Value]) any
	}
	buckets := []bucket{
		{"SourceData", func(e search.Entity[search.Value]) any { return e.SourceData }},
		{"Name+IDs", func(e search.Entity[search.Value]) any {
			return struct {
				Name, SourceID string
				Type           search.EntityType
				Source         search.SourceList
			}{e.Name, e.SourceID, e.Type, e.Source}
		}},
		{"PreparedFields", func(e search.Entity[search.Value]) any { return e.PreparedFields }},
		{"Person", func(e search.Entity[search.Value]) any { return e.Person }},
		{"Business", func(e search.Entity[search.Value]) any { return e.Business }},
		{"Organization", func(e search.Entity[search.Value]) any { return e.Organization }},
		{"Aircraft", func(e search.Entity[search.Value]) any { return e.Aircraft }},
		{"Vessel", func(e search.Entity[search.Value]) any { return e.Vessel }},
		{"Addresses", func(e search.Entity[search.Value]) any { return e.Addresses }},
		{"Contact", func(e search.Entity[search.Value]) any { return e.Contact }},
		{"CryptoAddresses", func(e search.Entity[search.Value]) any { return e.CryptoAddresses }},
		{"Affiliations", func(e search.Entity[search.Value]) any { return e.Affiliations }},
		{"SanctionsInfo", func(e search.Entity[search.Value]) any { return e.SanctionsInfo }},
		{"HistoricalInfo", func(e search.Entity[search.Value]) any { return e.HistoricalInfo }},
	}

	out := FieldBreakdown{Entities: len(entities)}
	for _, b := range buckets {
		var total int64
		for i := range entities {
			n, err := gobSize(b.fn(entities[i]))
			if err != nil {
				return out, fmt.Errorf("%s[%d]: %w", b.name, i, err)
			}
			total += n
		}
		out.Groups = append(out.Groups, FieldGroupSize{Name: b.name, Bytes: total})
		out.Total += total
	}

	for i := range out.Groups {
		if out.Total > 0 {
			out.Groups[i].Pct = 100 * float64(out.Groups[i].Bytes) / float64(out.Total)
		}
	}
	sort.Slice(out.Groups, func(i, j int) bool {
		return out.Groups[i].Bytes > out.Groups[j].Bytes
	})
	return out, nil
}
