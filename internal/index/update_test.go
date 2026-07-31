package index

import (
	"context"
	"fmt"
	"testing"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestLists_Update_CorpusOwnsEntities(t *testing.T) {
	lists := NewLists(nil).(*lists)

	gen := func(tag string, n int) []search.Entity[search.Value] {
		out := make([]search.Entity[search.Value], n)
		for i := range out {
			out[i] = search.Entity[search.Value]{
				Name:     fmt.Sprintf("%s-%d", tag, i),
				Type:     search.EntityPerson,
				Source:   search.SourceUSOFAC,
				SourceID: fmt.Sprintf("%s-%d", tag, i),
			}.Normalize()
		}
		return out
	}

	first := gen("first", 50)
	lists.Update(download.Stats{
		Entities: first,
		Lists:    map[string]int{string(search.SourceUSOFAC): len(first)},
	})

	// Metadata must not hold a second root to the entity slice.
	require.Nil(t, lists.latestStats.Entities)
	require.Nil(t, lists.latestStats.TFIDFIndex)
	require.NotNil(t, lists.corpus)
	require.Len(t, lists.corpus.entities, 50)

	got, err := lists.GetEntities(context.Background(), search.SourceUSOFAC)
	require.NoError(t, err)
	require.Len(t, got, 50)
	require.Equal(t, "first-0", got[0].SourceID)

	// Second generation replaces the first; search still works.
	second := gen("second", 30)
	lists.Update(download.Stats{
		Entities: second,
		Lists:    map[string]int{string(search.SourceUSOFAC): len(second)},
	})
	require.Nil(t, lists.latestStats.Entities)
	require.Len(t, lists.corpus.entities, 30)

	got, err = lists.GetEntities(context.Background(), search.SourceUSOFAC)
	require.NoError(t, err)
	require.Len(t, got, 30)
	require.Equal(t, "second-0", got[0].SourceID)

	// Candidate selection uses the new corpus only.
	cands, err := lists.SelectCandidates(context.Background(), search.Entity[search.Value]{
		Name:   "second-1",
		Type:   search.EntityPerson,
		Source: search.SourceUSOFAC,
	}.Normalize())
	require.NoError(t, err)
	require.NotEmpty(t, cands)

	stats := lists.LatestStats()
	require.Nil(t, stats.Entities)
	require.Equal(t, 30, stats.Lists[string(search.SourceUSOFAC)])
}

func TestLists_Update_Empty(t *testing.T) {
	lists := NewLists(nil)
	lists.Update(download.Stats{
		Lists: map[string]int{},
	})
	got, err := lists.GetEntities(context.Background(), "")
	require.NoError(t, err)
	require.Empty(t, got)
}
