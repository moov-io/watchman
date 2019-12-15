package main

import (
	"testing"

	"github.com/moov-io/watchman/pkg/csl"

	"github.com/go-kit/kit/log"
)

// This file is for tests that depend on having a trade.gov API Key.
// These tests should be skipped if the TRADEGOV_API_KEY environment variable is not set.

func TestSearcher_refreshData_CSL(t *testing.T) {
	if testing.Short() || csl.ApiKey == "" {
		return
	}

	s := &searcher{
		logger: log.NewNopLogger(),
		pipe:   noLogPipeliner,
	}

	stats, err := s.refreshData("")
	if err != nil {
		t.Fatal(err)
	}
	if len(s.SSIs) == 0 || stats.SectoralSanctions == 0 {
		t.Errorf("empty SSIs=%d or stats.SectoralSanctions=%d", len(s.SSIs), stats.SectoralSanctions)
	}
	if len(s.BISEntities) == 0 || stats.BISEntities == 0 {
		t.Errorf("empty searcher.BISEntities=%d or stats.BISEntities=%d", len(s.BISEntities), stats.BISEntities)
	}
}
