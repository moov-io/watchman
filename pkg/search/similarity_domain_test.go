package search_test

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_Domain(t *testing.T) {
	index := search.Entity[search.Value]{
		Name: "SUEX OTC, S.R.O.",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "SUEX OTC, S.R.O.",
		},
		Contact: search.ContactInfo{
			Websites: []string{"suex.io"},
		},
	}.Normalize()

	t.Run("subdomain query matches listed eTLD+1", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type: search.EntityBusiness,
			Contact: search.ContactInfo{
				Domains: []string{"pay.suex.io"},
			},
		}.Normalize()

		got := search.Similarity(query, index)
		require.Greater(t, got, 0.0)
		require.Less(t, got, 1.0)
	})

	t.Run("listed host query matches listed eTLD+1", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type: search.EntityBusiness,
			Contact: search.ContactInfo{
				Domains: []string{"suex.io"},
			},
		}.Normalize()

		got := search.Similarity(query, index)
		require.Greater(t, got, 0.0)
		require.Less(t, got, 1.0)
	})

	t.Run("unrelated domain does not match", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: "SUEX OTC, S.R.O.",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "SUEX OTC, S.R.O.",
			},
			Contact: search.ContactInfo{
				Domains: []string{"unrelated.example"},
			},
		}.Normalize()

		nameOnly := search.Similarity(search.Entity[search.Value]{
			Name: "SUEX OTC, S.R.O.",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "SUEX OTC, S.R.O.",
			},
		}.Normalize(), index)

		got := search.Similarity(query, index)
		require.LessOrEqual(t, got, nameOnly)
	})
}

func TestSimilarity_DomainFromEmail(t *testing.T) {
	index := search.Entity[search.Value]{
		Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
		},
		Contact: search.ContactInfo{
			EmailAddresses: []string{"info@gicdf.org"},
		},
	}.Normalize()

	query := search.Entity[search.Value]{
		Type: search.EntityBusiness,
		Contact: search.ContactInfo{
			Domains: []string{"mail.gicdf.org"},
		},
	}.Normalize()

	got := search.Similarity(query, index)
	require.Greater(t, got, 0.0)
	require.Less(t, got, 1.0)
}
