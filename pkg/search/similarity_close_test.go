package search

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCompareDates(t *testing.T) {
	cases := []struct {
		name         string
		query, index time.Time
		expected     float64
	}{
		{
			name:     "exact match",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			expected: 1.0,
		},
		{
			name:     "completely different",
			query:    time.Date(1929, time.May, 21, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.0,
		},
		{
			name:     "adjacent year",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1971, time.March, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.3,
		},
		{
			name:     "adjacent month",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.April, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.575,
		},
		{
			name:     "adjacent day",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.March, 11, 0, 0, 0, 0, time.UTC),
			expected: 0.925,
		},
		// TODO(adam): more tests
		// only year, only month, only year + month
		// examples in func comment
		// test ranges for OFAC (<= 5 years, closer is higher weight, decayed)
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := compareDates(&tc.query, &tc.index)
			require.InDelta(t, got, tc.expected, 0.001)
		})
	}
}

func TestCompareDates_EdgeCases(t *testing.T) {
	when := time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		name         string
		query, index *time.Time
		expected     float64
	}{
		{
			name:     "nil query",
			query:    nil,
			index:    &when,
			expected: 0.0,
		},
		{
			name:     "nil index",
			query:    &when,
			index:    nil,
			expected: 0.0,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := compareDates(tc.query, tc.index)
			require.InDelta(t, got, tc.expected, 0.001)
		})
	}
}
