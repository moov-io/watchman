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
			expected: 0.26, // Year > 5 (0.2), month distant (0.3), day distant (0.3): (0.4*0.2) + (0.3*0.3) + (0.3*0.3)
		},
		{
			name:     "adjacent year",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1971, time.March, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.96, // Year diff=1 (0.9), month same (1.0), day same (1.0): (0.4*0.9) + (0.3*1.0) + (0.3*1.0)
		},
		{
			name:     "adjacent month",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.April, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.97, // Year same (1.0), month diff=1 (0.9), day same (1.0): (0.4*1.0) + (0.3*0.9) + (0.3*1.0)
		},
		{
			name:     "adjacent day",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.March, 11, 0, 0, 0, 0, time.UTC),
			expected: 0.98, // Year same (1.0), month same (1.0), day diff=1 (0.95): (0.4*1.0) + (0.3*1.0) + (0.3*0.95)
		},
		{
			name:     "same year and month, similar day (1 vs 11)",
			query:    time.Date(1970, time.March, 1, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.March, 11, 0, 0, 0, 0, time.UTC),
			expected: 0.91, // Year same (1.0), month same (1.0), day similar (0.7): (0.4*1.0) + (0.3*1.0) + (0.3*0.7)
		},
		{
			name:     "same year and day, month 1 vs 11",
			query:    time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.November, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.91, // Year same (1.0), month 1 vs 11 (0.7), day same (1.0): (0.4*1.0) + (0.3*0.7) + (0.3*1.0)
		},
		{
			name:     "same month and day, 5 years apart",
			query:    time.Date(1970, time.March, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1975, time.March, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.8, // Year diff=5 (0.5), month same (1.0), day same (1.0): (0.4*0.5) + (0.3*1.0) + (0.3*1.0)
		},
		{
			name:     "only year matches",
			query:    time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.July, 15, 0, 0, 0, 0, time.UTC),
			expected: 0.7, // Year same (1.0), month distant (0.3), day distant (0.3): (0.4*1.0) + (0.3*0.3) + (0.3*0.3)
		},
		{
			name:     "only month matches",
			query:    time.Date(1960, time.March, 1, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.March, 15, 0, 0, 0, 0, time.UTC),
			expected: 0.59, // Year > 5 (0.2), month same (1.0), day distant (0.3): (0.4*0.2) + (0.3*1.0) + (0.3*0.3)
		},
		{
			name:     "only day matches",
			query:    time.Date(1960, time.January, 10, 0, 0, 0, 0, time.UTC),
			index:    time.Date(1970, time.June, 10, 0, 0, 0, 0, time.UTC),
			expected: 0.47,
		},
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
