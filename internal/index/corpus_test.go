package index

import (
	"context"
	"fmt"
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

	t.Run("empty type within known source does not scan full corpus", func(t *testing.T) {
		// US OFAC has persons and businesses but no aircraft
		query := mustNorm(search.Entity[search.Value]{
			Name:   "Anything",
			Type:   search.EntityAircraft,
			Source: search.SourceUSOFAC,
		})
		cands, err := idx.SelectCandidates(ctx, query)
		require.NoError(t, err)
		require.Empty(t, cands, "empty type partition must not fall back to scoring all entities")
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

	t.Run("empty source/type does not duplicate all-partition entries", func(t *testing.T) {
		bare := mustNorm(search.Entity[search.Value]{
			Name:     "Bare Entity",
			SourceID: "bare",
			// Source and Type intentionally empty
		})
		idx.Update(download.Stats{
			Entities: []search.Entity[search.Value]{bare},
			Lists:    map[string]int{"": 1},
		})
		impl := idx.(*lists)
		impl.mu.RLock()
		all := impl.corpus.bySourceType[""][""]
		impl.mu.RUnlock()
		require.Equal(t, []int{0}, all, "empty source/type must append entity once to all-partition")
	})

	t.Run("crypto without name tokens does not expand to full partition", func(t *testing.T) {
		idx.Update(stats) // restore multi-entity corpus
		query := mustNorm(search.Entity[search.Value]{
			// No name — only crypto. Must not expand to full business partition.
			Type: search.EntityBusiness,
			CryptoAddresses: []search.CryptoAddress{
				{Currency: "XBT", Address: "abc123"},
			},
		})
		require.Empty(t, query.PreparedFields.NameFields)
		cands, err := idx.SelectCandidates(ctx, query)
		require.NoError(t, err)
		require.Len(t, cands, 1)
		require.Equal(t, "3", cands[0].SourceID)
	})

	t.Run("multi-token query intersects from the rarest token", func(t *testing.T) {
		johnSmith := mustNorm(search.Entity[search.Value]{
			Name:     "John Smith",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "js",
			Person:   &search.Person{Name: "John Smith"},
		})
		johnDoe := mustNorm(search.Entity[search.Value]{
			Name:     "John Doe",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "jd",
			Person:   &search.Person{Name: "John Doe"},
		})
		janeSmith := mustNorm(search.Entity[search.Value]{
			Name:     "Jane Smith",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "jas",
			Person:   &search.Person{Name: "Jane Smith"},
		})
		idx.Update(download.Stats{
			Entities: []search.Entity[search.Value]{johnSmith, johnDoe, janeSmith},
			Lists:    map[string]int{string(search.SourceUSOFAC): 3},
		})

		cands, err := idx.SelectCandidates(ctx, mustNorm(search.Entity[search.Value]{
			Name:   "John Smith",
			Type:   search.EntityPerson,
			Source: search.SourceUSOFAC,
		}))
		require.NoError(t, err)
		require.Len(t, cands, 1)
		require.Equal(t, "js", cands[0].SourceID)
	})

	t.Run("misspelled extra token does not drop the matching token", func(t *testing.T) {
		johnSmith := mustNorm(search.Entity[search.Value]{
			Name:     "John Smith",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "js",
			Person:   &search.Person{Name: "John Smith"},
		})
		idx.Update(download.Stats{
			Entities: []search.Entity[search.Value]{johnSmith},
			Lists:    map[string]int{string(search.SourceUSOFAC): 1},
		})

		cands, err := idx.SelectCandidates(ctx, mustNorm(search.Entity[search.Value]{
			Name:   "John Zzznotatoken",
			Type:   search.EntityPerson,
			Source: search.SourceUSOFAC,
		}))
		require.NoError(t, err)
		require.Len(t, cands, 1)
		require.Equal(t, "js", cands[0].SourceID)
	})

	t.Run("disjoint token hits fall back to union", func(t *testing.T) {
		johnDoe := mustNorm(search.Entity[search.Value]{
			Name:     "John Doe",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "jd",
			Person:   &search.Person{Name: "John Doe"},
		})
		janeSmith := mustNorm(search.Entity[search.Value]{
			Name:     "Jane Smith",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "jas",
			Person:   &search.Person{Name: "Jane Smith"},
		})
		idx.Update(download.Stats{
			Entities: []search.Entity[search.Value]{johnDoe, janeSmith},
			Lists:    map[string]int{string(search.SourceUSOFAC): 2},
		})

		cands, err := idx.SelectCandidates(ctx, mustNorm(search.Entity[search.Value]{
			Name:   "John Smith",
			Type:   search.EntityPerson,
			Source: search.SourceUSOFAC,
		}))
		require.NoError(t, err)
		require.Len(t, cands, 2)
		ids := []string{cands[0].SourceID, cands[1].SourceID}
		require.ElementsMatch(t, []string{"jd", "jas"}, ids)
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

func TestCorpus_BlockingKeys(t *testing.T) {
	john := mustNorm(search.Entity[search.Value]{
		Name:     "John Smith",
		Type:     search.EntityPerson,
		Source:   search.SourceUSOFAC,
		SourceID: "j1",
		Person: &search.Person{
			GovernmentIDs: []search.GovernmentID{
				{Type: search.GovernmentIDPassport, Country: "US", Identifier: "1234567890"},
			},
		},
		Addresses: []search.Address{
			{Line1: "541 First St", City: "Anytown", State: "CA", PostalCode: "90210", Country: "US"},
		},
	})
	jane := mustNorm(search.Entity[search.Value]{
		Name:     "Jane Doe",
		Type:     search.EntityPerson,
		Source:   search.SourceUSOFAC,
		SourceID: "j2",
		Person: &search.Person{
			GovernmentIDs: []search.GovernmentID{
				{Type: search.GovernmentIDPassport, Country: "GB", Identifier: "999888777"},
			},
		},
		Addresses: []search.Address{
			{Line1: "10 Downing Street", City: "London", PostalCode: "SW1A 2AA", Country: "GB"},
		},
	})

	idx := NewLists(nil)
	idx.Update(download.Stats{
		Entities: []search.Entity[search.Value]{john, jane},
		Lists:    map[string]int{string(search.SourceUSOFAC): 2},
	})
	ctx := context.Background()

	t.Run("government ID query does not scan the other person", func(t *testing.T) {
		query := mustNorm(search.Entity[search.Value]{
			Type:   search.EntityPerson,
			Source: search.SourceUSOFAC,
			Person: &search.Person{
				GovernmentIDs: []search.GovernmentID{
					{Type: search.GovernmentIDPassport, Country: "US", Identifier: "1234567890"},
				},
			},
		})
		require.Empty(t, query.PreparedFields.NameFields)
		cands, err := idx.SelectCandidates(ctx, query)
		require.NoError(t, err)
		require.Len(t, cands, 1)
		require.Equal(t, "j1", cands[0].SourceID)
	})

	t.Run("address-only query stays in the matching country block", func(t *testing.T) {
		query := mustNorm(search.Entity[search.Value]{
			Type:   search.EntityPerson,
			Source: search.SourceUSOFAC,
			Addresses: []search.Address{
				{City: "Anytown", State: "CA", PostalCode: "90210", Country: "US"},
			},
		})
		require.Empty(t, query.PreparedFields.NameFields)
		cands, err := idx.SelectCandidates(ctx, query)
		require.NoError(t, err)
		require.Len(t, cands, 1)
		require.Equal(t, "j1", cands[0].SourceID)
	})

	t.Run("unknown government ID falls back to the partition", func(t *testing.T) {
		query := mustNorm(search.Entity[search.Value]{
			Type:   search.EntityPerson,
			Source: search.SourceUSOFAC,
			Person: &search.Person{
				GovernmentIDs: []search.GovernmentID{
					{Type: search.GovernmentIDPassport, Country: "US", Identifier: "0000000000"},
				},
			},
		})
		cands, err := idx.SelectCandidates(ctx, query)
		require.NoError(t, err)
		require.Len(t, cands, 2, "no blocking-key hits must not drop recall")
	})
}

func mustNorm(e search.Entity[search.Value]) search.Entity[search.Value] {
	return e.Normalize()
}

func BenchmarkSelectCandidates(b *testing.B) {
	entities := make([]search.Entity[search.Value], 0, 4400)
	for i := 0; i < 4000; i++ {
		entities = append(entities, mustNorm(search.Entity[search.Value]{
			Name:     "Acme Company Limited",
			Type:     search.EntityBusiness,
			Source:   search.SourceUSOFAC,
			SourceID: fmt.Sprintf("c%d", i),
			Business: &search.Business{Name: "Acme Company Limited"},
		}))
	}
	for i := 0; i < 200; i++ {
		entities = append(entities, mustNorm(search.Entity[search.Value]{
			Name:     "Ocean Shipping Limited",
			Type:     search.EntityBusiness,
			Source:   search.SourceUSOFAC,
			SourceID: fmt.Sprintf("s%d", i),
			Business: &search.Business{Name: "Ocean Shipping Limited"},
		}))
	}
	for i := 0; i < 200; i++ {
		entities = append(entities, mustNorm(search.Entity[search.Value]{
			Name:     "Ocean Freight Group",
			Type:     search.EntityBusiness,
			Source:   search.SourceUSOFAC,
			SourceID: fmt.Sprintf("f%d", i),
			Business: &search.Business{Name: "Ocean Freight Group"},
		}))
	}

	idx := NewLists(nil)
	idx.Update(download.Stats{
		Entities: entities,
		Lists:    map[string]int{string(search.SourceUSOFAC): len(entities)},
	})
	query := mustNorm(search.Entity[search.Value]{
		Name:   "Ocean Shipping Limited",
		Type:   search.EntityBusiness,
		Source: search.SourceUSOFAC,
	})
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cands, err := idx.SelectCandidates(ctx, query)
		if err != nil {
			b.Fatal(err)
		}
		if len(cands) == 0 {
			b.Fatal("expected candidates")
		}
	}
}
