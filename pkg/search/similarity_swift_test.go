package search_test

import (
	"encoding/json"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_SwiftBIC(t *testing.T) {
	index := search.Entity[search.Value]{
		Name: "ONERBANK ZAO",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "ONERBANK ZAO",
		},
		SwiftBICs: []string{"HNRBBY2X"},
	}.Normalize()

	t.Run("exact BIC forces 1.0 despite different name", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: "Honor Bank",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "Honor Bank",
			},
			SwiftBICs: []string{"HNRBBY2X"},
		}.Normalize()

		got := search.Similarity(query, index)
		require.InDelta(t, 1.0, got, 0.001)
	})

	t.Run("11-character XXX form matches listed BIC8", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type:      search.EntityBusiness,
			SwiftBICs: []string{"HNRBBY2XXXX"},
		}.Normalize()

		got := search.Similarity(query, index)
		require.InDelta(t, 1.0, got, 0.001)
	})

	t.Run("BIC-only query without type scores 1.0", func(t *testing.T) {
		query := search.Entity[search.Value]{
			SwiftBICs: []string{"hnrbby2x"},
		}.Normalize()

		got := search.Similarity(query, index)
		require.InDelta(t, 1.0, got, 0.001)
	})

	t.Run("organization query vs business index scores 1.0", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type:      search.EntityOrganization,
			SwiftBICs: []string{"HNRBBY2X"},
			Organization: &search.Organization{
				Name: "Honor Bank",
			},
		}.Normalize()

		got := search.Similarity(query, index)
		require.InDelta(t, 1.0, got, 0.001)
	})

	t.Run("different BIC does not force 1.0", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: "ONERBANK ZAO",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "ONERBANK ZAO",
			},
			SwiftBICs: []string{"RUDVRUMM"},
		}.Normalize()

		got := search.Similarity(query, index)
		require.Less(t, got, 1.0)
	})
}

func TestExactOverride_SwiftBICForcesOneDespiteName(t *testing.T) {
	query := search.Entity[search.Value]{
		Name:      "Wrong Name LLC",
		Type:      search.EntityOrganization,
		SwiftBICs: []string{"RUDVRUMM"},
		Organization: &search.Organization{
			Name: "Wrong Name LLC",
		},
	}.Normalize()
	index := search.Entity[search.Value]{
		Name:      "SME BANK",
		Type:      search.EntityOrganization,
		SwiftBICs: []string{"RUDVRUMM"},
		Organization: &search.Organization{
			Name: "SME BANK",
		},
	}.Normalize()

	got := search.Similarity(query, index)
	require.InDelta(t, 1.0, got, 0.001)
}

func TestIdentifierConflict_SwiftBIC(t *testing.T) {
	sameName := search.Entity[search.Value]{
		Name: "Same Bank",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "Same Bank",
		},
		SwiftBICs: []string{"HNRBBY2X"},
	}.Normalize()
	other := search.Entity[search.Value]{
		Name: "Same Bank",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "Same Bank",
		},
		SwiftBICs: []string{"RUDVRUMM"},
	}.Normalize()

	withConflict := search.Similarity(sameName, other)
	matching := other
	matching.SwiftBICs = []string{"HNRBBY2X"}
	matching = matching.Normalize()
	withoutConflict := search.Similarity(sameName, matching)

	require.Less(t, withConflict, withoutConflict)
	require.InDelta(t, 1.0, withoutConflict, 0.001)
}

func TestEntity_SwiftBICJSON(t *testing.T) {
	with := search.Entity[search.Value]{
		Name:      "ONERBANK ZAO",
		Type:      search.EntityBusiness,
		SwiftBICs: []string{"HNRBBY2X"},
	}
	raw, err := json.Marshal(with)
	require.NoError(t, err)
	require.Contains(t, string(raw), `"swiftBICs":["HNRBBY2X"]`)

	empty := search.Entity[search.Value]{
		Name: "No BIC",
		Type: search.EntityBusiness,
	}
	raw, err = json.Marshal(empty)
	require.NoError(t, err)
	require.NotContains(t, string(raw), "swiftBICs")
}
