package search

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"
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

	search := func(b *testing.B) {
		b.Helper()

		for i := 0; i < b.N; i++ {
			results, err := svc.Search(ctx, query, opts)
			require.NoError(b, err)
			require.Greater(b, len(results), 0)
		}
	}

	b.Run("fixed group size", func(b *testing.B) {
		groups := []int{1, 3, 5, 10, 20, 25, 50, 100, 150, 200, 250}
		for _, g := range groups {
			b.Run(fmt.Sprintf("%d", g), func(b *testing.B) {
				b.Setenv("SEARCH_GROUP_COUNT", strconv.Itoa(g))

				search(b)
			})
		}
	})

	b.Run("dynamic group size", func(b *testing.B) {
		b.Setenv("SEARCH_GROUP_COUNT", "") // clear

		search(b)
	})
}

func Benchmark_SearchParallel(b *testing.B) {
	svc := testService(b)
	ctx := context.Background()

	opts := SearchOpts{
		Limit:    5,
		MinMatch: 0.80,
	}

	b.Run("individual", func(b *testing.B) {
		query := ofactest.FindEntity(b, "29702")

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				results, err := svc.Search(ctx, query, opts)
				require.NoError(b, err)
				require.Greater(b, len(results), 0)
			}
		})
	})

	b.Run("business", func(b *testing.B) {
		query := ofactest.FindEntity(b, "44525")

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				results, err := svc.Search(ctx, query, opts)
				require.NoError(b, err)
				require.Greater(b, len(results), 0)
			}
		})
	})

	b.Run("vessel", func(b *testing.B) {
		query := ofactest.FindEntity(b, "50972")

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				results, err := svc.Search(ctx, query, opts)
				require.NoError(b, err)
				require.Greater(b, len(results), 0)
			}
		})
	})

	b.Run("aircraft", func(b *testing.B) {
		query := ofactest.FindEntity(b, "48727")

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				results, err := svc.Search(ctx, query, opts)
				require.NoError(b, err)
				require.Greater(b, len(results), 0)
			}
		})
	})
}
