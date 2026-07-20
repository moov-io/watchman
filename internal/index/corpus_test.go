package index

import (
	"context"
	"testing"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestCorpus_PartitionAndCandidates(t *testing.T) {
	entities := []search.Entity[search.Value]{
		mustNorm(search.Entity[search.Value]{
			Name:     "John Smith",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "1",
			Person:   &search.Person{Name: "John Smith"},
		}),
		mustNorm(search.Entity[search.Value]{
			Name:     "Acme Shipping Limited",
			Type:     search.EntityBusiness,
			Source:   search.SourceUSOFAC,
			SourceID: "2",
			Business: &search.Business{Name: "Acme Shipping Limited"},
		}),
		mustNorm(search.Entity[search.Value]{
			Name:     "Other Corp",
			Type:     search.EntityBusiness,
			Source:   search.SourceEUCSL,
			SourceID: "3",
			Business: &search.Business{Name: "Other Corp"},
			CryptoAddresses: []search.CryptoAddress{
				{Currency: "XBT", Address: "abc123"},
			},
		}),
	}

	lists := NewLists(nil)
	lists.Update(download.Stats{
		Entities: entities,
		Lists: map[string]int{
			string(search.SourceUSOFAC): 2,
			string(search.SourceEUCSL):  1,
		},
	})

	ctx := context.Background()

	t.Run("GetEntities by source", func(t *testing.T) {
		got, err := lists.GetEntities(ctx, search.SourceUSOFAC)
		require.NoError(t, err)
		require.Len(t, got, 2)
	})

	t.Run("type partition via candidates", func(t *testing.T) {
		query := mustNorm(search.Entity[search.Value]{
			Name:   "Shipping Limited",
			Type:   search.EntityBusiness,
			Source: search.SourceUSOFAC,
		})
		cands, err := lists.SelectCandidates(ctx, query)
		require.NoError(t, err)
		// Should not include the person or EU entity
		require.NotEmpty(t, cands)
		for _, c := range cands {
			require.Equal(t, search.EntityBusiness, c.Type)
			require.Equal(t, search.SourceUSOFAC, c.Source)
		}
		// Token "shipping" should hit entity 2
		require.True(t, len(cands) <= 2)
	})

	t.Run("crypto exact candidate", func(t *testing.T) {
		query := mustNorm(search.Entity[search.Value]{
			Type: search.EntityBusiness,
			CryptoAddresses: []search.CryptoAddress{
				{Currency: "XBT", Address: "abc123"},
			},
		})
		// empty source → all sources partition
		cands, err := lists.SelectCandidates(ctx, query)
		require.NoError(t, err)
		require.Len(t, cands, 1)
		require.Equal(t, "3", cands[0].SourceID)
	})

	t.Run("typo falls back to partition", func(t *testing.T) {
		query := mustNorm(search.Entity[search.Value]{
			Name:   "Zzznotatoken",
			Type:   search.EntityPerson,
			Source: search.SourceUSOFAC,
		})
		cands, err := lists.SelectCandidates(ctx, query)
		require.NoError(t, err)
		// Full person partition for US OFAC
		require.Len(t, cands, 1)
		require.Equal(t, "1", cands[0].SourceID)
	})
}

func mustNorm(e search.Entity[search.Value]) search.Entity[search.Value] {
	return e.Normalize()
}
