package search

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func Benchmark_Search(b *testing.B) {
	svc := testService(b)
	ctx := context.Background()

	query := search.Entity[search.Value]{
		Name: "Xao Ling",
		Type: search.EntityPerson,
	}
	opts := SearchOpts{
		Limit:    20,
		MinMatch: 0.1,
	}

	b.ResetTimer()

	search := func(b *testing.B, groups int) {
		b.Helper()

		for i := 0; i < b.N; i++ {
			results, err := svc.Search(ctx, query, opts)
			require.NoError(b, err)
			require.Greater(b, len(results), 0)
		}
	}

	groups := []int{1, 3, 5, 10, 20, 25, 50, 100, 150, 200, 250}
	for _, g := range groups {
		b.Run(fmt.Sprintf("%d group", g), func(b *testing.B) {
			b.Setenv("SEARCH_GROUP_COUNT", strconv.Itoa(g))

			search(b, g)
		})
	}
}
