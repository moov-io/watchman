// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"sync"
	"testing"

	"github.com/go-kit/kit/log"
)

var (
	testLiveSearcher = &searcher{
		logger: log.NewNopLogger(),
		pipe:   noLogPipeliner,
	}
	testSearcherStats *downloadStats
	testSearcherOnce  sync.Once
)

func createTestSearcher(t *testing.T) *searcher {
	if testing.Short() {
		t.Skip("-short enabled")
	}

	testSearcherOnce.Do(func() {
		stats, err := testLiveSearcher.refreshData("")
		if err != nil {
			t.Fatal(err)
		}
		testSearcherStats = stats
	})

	return testLiveSearcher
}

func createBenchmarkSearcher(b *testing.B) *searcher {
	testSearcherOnce.Do(func() {
		stats, err := testLiveSearcher.refreshData("")
		if err != nil {
			b.Fatal(err)
		}
		testSearcherStats = stats
	})
	return testLiveSearcher
}
