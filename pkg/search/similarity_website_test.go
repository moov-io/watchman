package search_test

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_Website(t *testing.T) {
	index := search.Entity[search.Value]{
		Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
		},
		Contact: search.ContactInfo{
			Websites: []string{"www.gicdf.org"},
		},
	}.Normalize()

	t.Run("https www url matches listed www host", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			},
			Contact: search.ContactInfo{
				Websites: []string{"https://www.gicdf.org/about"},
			},
		}.Normalize()

		withSite := search.Similarity(query, index)
		nameOnly := search.Similarity(search.Entity[search.Value]{
			Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			},
		}.Normalize(), index)

		require.Greater(t, withSite, nameOnly)
		require.Less(t, withSite, 1.0)
	})

	t.Run("bare host matches listed www host", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type: search.EntityBusiness,
			Contact: search.ContactInfo{
				Websites: []string{"gicdf.org"},
			},
		}.Normalize()

		got := search.Similarity(query, index)
		require.Greater(t, got, 0.0)
		require.Less(t, got, 1.0)
	})

	t.Run("different host does not match", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			},
			Contact: search.ContactInfo{
				Websites: []string{"unrelated.example"},
			},
		}.Normalize()

		nameOnly := search.Similarity(search.Entity[search.Value]{
			Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			Type: search.EntityBusiness,
			Business: &search.Business{
				Name: "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION",
			},
		}.Normalize(), index)

		got := search.Similarity(query, index)
		require.LessOrEqual(t, got, nameOnly)
	})
}
