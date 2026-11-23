package search_test

import (
	"bytes"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func queryAutoMart() search.Entity[search.Value] {
	return search.Entity[search.Value]{
		Name: "GINO'S AUTO MART INC",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name:     "GINO'S AUTO MART INC",
			AltNames: []string{"GINO'S AUTO MART"},
			GovernmentIDs: []search.GovernmentID{
				{Type: search.GovernmentIDTax, Country: "United States", Identifier: "119998877"},
			},
		},
		Contact: search.ContactInfo{
			EmailAddresses: []string{"parts@example.com"},
			PhoneNumbers:   []string{"1 3125550100"},
			Websites:       []string{"https://example.com/autos"},
		},
		Addresses: []search.Address{{
			Line1:      "2400 W Cedar St",
			City:       "Chicago",
			PostalCode: "60608",
			State:      "IL",
			Country:    "United States",
		}},
	}.Normalize()
}

func listedBusiness() search.Entity[search.Value] {
	return search.Entity[search.Value]{
		Name:     "Ginos Auto Mart Inc",
		Type:     search.EntityBusiness,
		Source:   "test-list",
		SourceID: "TEST-1",
		Business: &search.Business{
			Name: "Ginos Auto Mart Inc",
		},
		Contact: search.ContactInfo{
			PhoneNumbers: []string{"3125550100"},
		},
		Addresses: []search.Address{{
			Line1:      "900 N Walnut Ave",
			City:       "Chicago",
			PostalCode: "60612",
			State:      "IL",
			Country:    "US",
		}},
	}.Normalize()
}

func debugScore(t *testing.T, query, index search.Entity[search.Value]) (search.SimilarityScore, string) {
	t.Helper()
	var buf bytes.Buffer
	got := search.DebugSimilarityWithOpts(&buf, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	return got, buf.String()
}

func pieceByType(t *testing.T, got search.SimilarityScore, pieceType string) search.ScorePiece {
	t.Helper()
	for _, p := range got.Pieces {
		if p.PieceType == pieceType {
			return p
		}
	}
	t.Fatalf("missing %s piece", pieceType)
	return search.ScorePiece{}
}

func TestSimilarity_PossessiveBusinessExactPhone(t *testing.T) {
	query := queryAutoMart()
	index := listedBusiness()

	require.Equal(t, []string{"ginos", "auto", "mart", "inc"}, query.PreparedFields.NameFields)
	require.Equal(t, []string{"ginos", "auto", "mart", "inc"}, index.PreparedFields.NameFields)
	require.Equal(t, []string{"13125550100"}, query.PreparedFields.Contact.PhoneNumbers)
	require.Equal(t, []string{"13125550100"}, index.PreparedFields.Contact.PhoneNumbers)

	got, debug := debugScore(t, query, index)
	require.GreaterOrEqual(t, got.FinalScore, 0.85, "name + exact phone should clear the screening line with room to spare")
	require.NotContains(t, debug, "fields.required < 2")

	name := pieceByType(t, got, "name")
	require.True(t, name.Exact)
	require.InDelta(t, 1.0, name.Score, 0.0001)
}

func TestSimilarity_ExactPhoneSkipsMinRequired(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "GINO'S AUTO MART INC",
		Type: search.EntityBusiness,
		Contact: search.ContactInfo{
			PhoneNumbers: []string{"1 3125550100"},
		},
	}.Normalize()
	index := listedBusiness()

	got, debug := debugScore(t, query, index)
	require.GreaterOrEqual(t, got.FinalScore, 0.85)
	require.NotContains(t, debug, "fields.required < 2")
	// Phone is evidence, not a unique identity key, so name-only still applies
	// when the query has no address.
	require.Contains(t, debug, "reduced name-only match penalty")
}

func TestSimilarity_NameOnlyStillPenalized(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "GINO'S AUTO MART INC",
		Type: search.EntityBusiness,
	}.Normalize()
	index := listedBusiness()

	got, debug := debugScore(t, query, index)
	phone := search.Similarity(search.Entity[search.Value]{
		Name:    query.Name,
		Type:    search.EntityBusiness,
		Contact: search.ContactInfo{PhoneNumbers: []string{"1 3125550100"}},
	}.Normalize(), index)
	require.InDelta(t, 0.812, got.FinalScore, 0.001)
	require.Greater(t, phone, got.FinalScore)
	require.Contains(t, debug, "fields.required < 2")
	require.Contains(t, debug, "reduced name-only match penalty")
}

func listedWebsite() search.Entity[search.Value] {
	return search.Entity[search.Value]{
		Name:     "Ginos Auto Mart Inc",
		Type:     search.EntityBusiness,
		Source:   "test-list",
		SourceID: "TEST-2",
		Business: &search.Business{
			Name: "Ginos Auto Mart Inc",
		},
		Contact: search.ContactInfo{
			Websites: []string{"example.com"},
		},
	}.Normalize()
}

func TestSimilarity_WebsiteOnlyCountsOnce(t *testing.T) {
	query := search.Entity[search.Value]{
		Type: search.EntityBusiness,
		Contact: search.ContactInfo{
			Websites: []string{"https://example.com/autos"},
		},
	}.Normalize()
	index := listedWebsite()

	got, debug := debugScore(t, query, index)
	contact := pieceByType(t, got, "contact-exact")
	require.GreaterOrEqual(t, contact.FieldsCompared, 2)
	require.True(t, contact.Exact)
	require.Contains(t, debug, "fields.required < 2")
	require.Less(t, got.FinalScore, 1.0)
}

func TestSimilarity_NameAndWebsiteSkipsMinRequired(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "GINO'S AUTO MART INC",
		Type: search.EntityBusiness,
		Contact: search.ContactInfo{
			Websites: []string{"https://example.com/autos"},
		},
	}.Normalize()
	index := listedWebsite()

	got, debug := debugScore(t, query, index)
	require.GreaterOrEqual(t, got.FinalScore, 0.85)
	require.Less(t, got.FinalScore, 1.0)
	require.NotContains(t, debug, "fields.required < 2")
}
