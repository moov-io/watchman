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

	stats := download.Stats{
		Entities: entities,
		Lists: map[string]int{
			string(search.SourceUSOFAC): 2,
			string(search.SourceEUCSL):  1,
		},
	}

	idx := NewLists(nil)
	idx.Update(stats)
	ctx := context.Background()

	t.Run("GetEntities by source", func(t *testing.T) {
		got, err := idx.GetEntities(ctx, search.SourceUSOFAC)
		require.NoError(t, err)
		require.Len(t, got, 2)
	})

	t.Run("type partition via candidates", func(t *testing.T) {
		query := mustNorm(search.Entity[search.Value]{
			Name:   "Shipping Limited",
			Type:   search.EntityBusiness,
			Source: search.SourceUSOFAC,
		})
		cands, err := idx.SelectCandidates(ctx, query)
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
		cands, err := idx.SelectCandidates(ctx, query)
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
		cands, err := idx.SelectCandidates(ctx, query)
		require.NoError(t, err)
		// Full person partition for US OFAC
		require.Len(t, cands, 1)
		require.Equal(t, "1", cands[0].SourceID)
	})

	t.Run("GetEntities empty partition does not leak other sources", func(t *testing.T) {
		// Source is registered in Lists but has no entities in the corpus
		idx.Update(download.Stats{
			Entities: entities,
			Lists: map[string]int{
				string(search.SourceUSOFAC): 2,
				string(search.SourceEUCSL):  1,
				string(search.SourceUKCSL):  0, // listed but empty
			},
		})
		got, err := idx.GetEntities(ctx, search.SourceUKCSL)
		require.NoError(t, err)
		require.Empty(t, got, "empty source partition must not fall back to all entities")
	})

	t.Run("name tokens deduped per entity", func(t *testing.T) {
		// Rebuild with an entity that repeats a token across primary and alt names
		dup := mustNorm(search.Entity[search.Value]{
			Name:     "Smith Trading Smith",
			Type:     search.EntityBusiness,
			Source:   search.SourceUSOFAC,
			SourceID: "dup",
			Business: &search.Business{
				Name:     "Smith Trading Smith",
				AltNames: []string{"Smith Holdings"},
			},
		})
		idx.Update(download.Stats{
			Entities: []search.Entity[search.Value]{dup},
			Lists:    map[string]int{string(search.SourceUSOFAC): 1},
		})

		cands, err := idx.SelectCandidates(ctx, mustNorm(search.Entity[search.Value]{
			Name:   "Smith",
			Type:   search.EntityBusiness,
			Source: search.SourceUSOFAC,
		}))
		require.NoError(t, err)
		require.Len(t, cands, 1)

		impl := idx.(*lists)
		impl.mu.RLock()
		postings := impl.corpus.nameTokens["smith"]
		impl.mu.RUnlock()
		require.Equal(t, []int{0}, postings, "entity index should appear once per token")
	})
}

func mustNorm(e search.Entity[search.Value]) search.Entity[search.Value] {
	return e.Normalize()
}
