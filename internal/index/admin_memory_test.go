package index

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMemoryHandler(t *testing.T) {
	lists := NewLists(nil)
	lists.Update(download.Stats{
		Entities: []search.Entity[search.Value]{
			{Name: "A", Type: search.EntityPerson, Source: search.SourceUSOFAC, SourceID: "1"},
			{Name: "B", Type: search.EntityBusiness, Source: search.SourceUSOFAC, SourceID: "2"},
		},
		Lists:     map[string]int{string(search.SourceUSOFAC): 2},
		StartedAt: time.Now().Add(-time.Minute),
		EndedAt:   time.Now(),
		Version:   "test",
	})

	req := httptest.NewRequest(http.MethodGet, "/debug/memory", nil)
	rr := httptest.NewRecorder()
	memoryHandler(lists)(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Contains(t, rr.Header().Get("Content-Type"), "application/json")

	var report MemoryReport
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &report))
	require.Equal(t, 2, report.EntityCount)
	require.Equal(t, map[string]int{string(search.SourceUSOFAC): 2}, report.Lists)
	require.NotEmpty(t, report.PprofHints.Heap)
	require.Greater(t, report.Sys, uint64(0))
}

func TestEntityCount(t *testing.T) {
	lists := NewLists(nil)
	require.Equal(t, 0, EntityCount(lists))

	lists.Update(download.Stats{
		Entities: []search.Entity[search.Value]{{Name: "x", Source: search.SourceUSOFAC}},
		Lists:    map[string]int{string(search.SourceUSOFAC): 1},
	})
	require.Equal(t, 1, EntityCount(lists))
}
