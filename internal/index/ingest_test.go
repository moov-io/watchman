package index

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

type countingRepo struct {
	ingest.MockRepository
	listAll   atomic.Int32
	checksums atomic.Int32
}

func (r *countingRepo) ListAll(ctx context.Context) ([]search.Entity[search.Value], error) {
	r.listAll.Add(1)
	return r.MockRepository.ListAll(ctx)
}

func (r *countingRepo) Checksums(ctx context.Context) ([]ingest.SourceChecksum, error) {
	r.checksums.Add(1)
	return r.MockRepository.Checksums(ctx)
}

func testIngestEntity(id int) search.Entity[search.Value] {
	return search.Entity[search.Value]{
		Name:     fmt.Sprintf("Ingest Person %05d", id),
		Type:     search.EntityPerson,
		SourceID: fmt.Sprintf("%05d", id),
	}.Normalize()
}

func TestLists_IngestCorpus(t *testing.T) {
	repo := &countingRepo{}
	lists := NewLists(repo)
	ctx := context.Background()

	const n = 1500
	ents := make([]search.Entity[search.Value], n)
	for i := 0; i < n; i++ {
		ents[i] = testIngestEntity(i)
	}
	require.NoError(t, repo.Upsert(ctx, "fincen-person", ents))

	query := search.Entity[search.Value]{
		Name:   "Ingest Person 01499",
		Type:   search.EntityPerson,
		Source: "fincen-person",
	}.Normalize()

	cands, err := lists.SelectCandidates(ctx, query)
	require.NoError(t, err)
	require.Greater(t, cands.Len(), 0)

	found := false
	for i := 0; i < cands.Len(); i++ {
		if cands.At(i).SourceID == "01499" {
			found = true
			break
		}
	}
	require.True(t, found, "row 1499 must be a candidate")

	got, err := lists.GetEntities(ctx, "fincen-person")
	require.NoError(t, err)
	require.Len(t, got, n)

	firstListAll := repo.listAll.Load()
	require.Equal(t, int32(1), firstListAll)

	cands, err = lists.SelectCandidates(ctx, query)
	require.NoError(t, err)
	require.Greater(t, cands.Len(), 0)
	require.Equal(t, firstListAll, repo.listAll.Load(), "checksum hit must not rescan entities")
	require.Greater(t, repo.checksums.Load(), int32(1))

	stats := lists.LatestStats()
	require.Equal(t, n, stats.Lists["fincen-person"])
	require.Len(t, stats.ListHashes["fincen-person"], 64)
}

func TestLists_IngestSurvivesDownloadRefresh(t *testing.T) {
	repo := &countingRepo{}
	lists := NewLists(repo)
	ctx := context.Background()

	require.NoError(t, repo.Upsert(ctx, "fincen-person", []search.Entity[search.Value]{testIngestEntity(1)}))
	require.NoError(t, lists.RefreshIngest(ctx))

	ofac := search.Entity[search.Value]{
		Name:     "Nicolas MADURO MOROS",
		Type:     search.EntityPerson,
		Source:   search.SourceUSOFAC,
		SourceID: "22790",
	}.Normalize()
	lists.Update(download.Stats{
		Entities: []search.Entity[search.Value]{ofac},
		Lists:    map[string]int{string(search.SourceUSOFAC): 1},
	})

	got, err := lists.GetEntities(ctx, "fincen-person")
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, "00001", got[0].SourceID)

	cands, err := lists.SelectCandidates(ctx, search.Entity[search.Value]{
		Name:   "Ingest Person 00001",
		Type:   search.EntityPerson,
		Source: "fincen-person",
	}.Normalize())
	require.NoError(t, err)
	require.Greater(t, cands.Len(), 0)

	all, err := lists.SelectCandidates(ctx, search.Entity[search.Value]{
		Name: "Person",
		Type: search.EntityPerson,
	}.Normalize())
	require.NoError(t, err)
	require.GreaterOrEqual(t, all.Len(), 2)
}

func TestLists_RefreshIngest_Empty(t *testing.T) {
	lists := NewLists(nil)
	require.NoError(t, lists.RefreshIngest(context.Background()))
}

func TestLists_UnknownIngestSource(t *testing.T) {
	repo := &countingRepo{}
	lists := NewLists(repo)
	ctx := context.Background()

	got, err := lists.GetEntities(ctx, "missing-file")
	require.NoError(t, err)
	require.Empty(t, got)

	cands, err := lists.SelectCandidates(ctx, search.Entity[search.Value]{
		Name:   "nobody",
		Type:   search.EntityPerson,
		Source: "missing-file",
	}.Normalize())
	require.NoError(t, err)
	require.Equal(t, 0, cands.Len())
}
