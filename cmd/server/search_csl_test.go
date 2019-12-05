package main

import (
	"github.com/go-kit/kit/log"
	"os"
	"testing"
)

// This file is for tests that depend on having a trade.gov API Key.
// These tests should be skipped if the TRADEGOV_API_KEY environment variable is not set.

func TestSearcher_refreshData_CSL(t *testing.T) {
	if os.Getenv("TRADEGOV_API_KEY") == "" {
		return
	}
	s := &searcher{logger: log.NewNopLogger()}
	stats, err := s.refreshData("")
	if err != nil {
		t.Fatal(err)
	}
	if len(s.SSIs) == 0 || stats.SectoralSanctions == 0 {
		t.Errorf("empty SSIs=%d or stats.SectoralSanctions=%d", len(s.SSIs), stats.SectoralSanctions)
	}
	if len(s.ELs) == 0 || stats.EntityList == 0 {
		t.Errorf("empty ELs=%d or stats.EntityList=%d", len(s.ELs), stats.EntityList)
	}
}
