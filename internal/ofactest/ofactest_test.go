package ofactest_test

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

const (
	downloadRealRecords = false
)

func TestOFACTest_Sample(t *testing.T) {
	logger := log.NewTestLogger()

	var dl download.Downloader
	var err error

	if downloadRealRecords {
		// Use a real downloader (to pull from OFAC website)
		dl, err = download.NewDownloader(logger, download.Config{
			IncludedLists: []search.SourceList{
				search.SourceUSOFAC,
			},
		})
	} else {
		// Use a mock downloader with OFAC files from ./pkg/sources/ofac/testdata/
		dl = ofactest.GetDownloader(t)
	}
	require.NoError(t, err)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)

	gather := func(entity search.Entity[search.Value]) string {
		if len(entity.Contact.PhoneNumbers) > 0 {
			return entity.Contact.PhoneNumbers[0]
		}
		return ""
	}

	var sample []search.Entity[search.Value]
	for idx := range stats.Entities {
		// Logic to include Entity
		if gather(stats.Entities[idx]) == "" {
			continue
		}

		// Entity has passed our checks, so sample it
		if shouldSample(0.01) {
			sample = append(sample, stats.Entities[idx])
		}
	}

	// Show entities
	for idx := range sample {
		t.Logf("%v - %q", sample[idx].SourceID, gather(sample[idx]))
	}
	t.Logf("sampled %d records", len(sample))
}

var (
	max = big.NewInt(100)
)

func shouldSample(chance float64) bool {
	if chance < 1.0 {
		chance *= 100.0
	}

	n, _ := rand.Int(rand.Reader, max)

	return float64(n.Int64()) < chance
}
