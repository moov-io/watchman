package main

import (
	"testing"

	"github.com/go-kit/kit/log"
)

func TestSearcher_refreshData_CSL(t *testing.T) {
	if testing.Short() {
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
