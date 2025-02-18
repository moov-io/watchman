package ofactest_test

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestOFACTest_Sample(t *testing.T) {
	dl := ofactest.GetDownloader(t)

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
		if shouldSample(1.0) {
			sample = append(sample, stats.Entities[idx])
		}
	}

	// Show entities
	for idx := range sample {
		t.Logf("%v - %q", sample[idx].SourceID, gather(sample[idx]))
	}
}

var (
	max = big.NewInt(100)
)

func shouldSample(chance float64) bool {
	n, _ := rand.Int(rand.Reader, max)

	return float64(n.Int64()/100.0) < chance
}
