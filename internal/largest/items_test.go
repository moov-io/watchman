// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package largest_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/largest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

// Helper to create a largest.Item with a given name and weight.
// The Value is a search.Entity[search.Value] with just a Name for illustration.
func makeItem(name string, w float64) largest.Item {
	return largest.Item{
		Value: search.Entity[search.Value]{
			Name: name,
			// Type: could be set if needed, but not required for this test
		},
		Weight: w,
	}
}

func TestItems_Basic(t *testing.T) {
	// We’ll track the top 3 items (capacity=3). Items below weight=2.0 are ignored (minMatch=2.0)
	xs := largest.NewItems(3, 2.0)

	// Initially, it should be empty
	require.Empty(t, xs.Items(), "expected no items at start")

	// 1) Below minMatch => should be ignored
	xs.Add(makeItem("A", 1.5))
	require.Empty(t, xs.Items(), "items below minMatch are ignored")

	// 2) Exactly minMatch => should be included
	xs.Add(makeItem("B", 2.0))
	got := xs.Items()
	require.Len(t, got, 1)
	require.Equal(t, "B", got[0].Value.Name)

	// 3) Insert item with weight=3 => should be included
	xs.Add(makeItem("C", 3.0))
	got = xs.Items()
	require.Len(t, got, 2)
	// Should be sorted descending: C(3.0), B(2.0)
	require.Equal(t, "C", got[0].Value.Name)
	require.Equal(t, "B", got[1].Value.Name)

	// 4) Insert item with weight=4 => included, still under capacity=3
	xs.Add(makeItem("D", 4.0))
	got = xs.Items()
	require.Len(t, got, 3)
	// Expect descending: D(4.0), C(3.0), B(2.0)
	require.Equal(t, []string{"D", "C", "B"}, []string{
		got[0].Value.Name,
		got[1].Value.Name,
		got[2].Value.Name,
	})

	// 5) Insert item with weight=5 => must remove smallest (B=2.0)
	xs.Add(makeItem("E", 5.0))
	got = xs.Items()
	require.Len(t, got, 3)
	// Expect descending: E(5.0), D(4.0), C(3.0)
	require.Equal(t, []string{"E", "D", "C"}, []string{
		got[0].Value.Name,
		got[1].Value.Name,
		got[2].Value.Name,
	})

	// 6) Insert something small but above minMatch => compare to smallest
	// Currently, smallest is C(3.0). We'll try F(2.5), which is less than 3.0 => ignore
	xs.Add(makeItem("F", 2.5))
	got = xs.Items()
	require.Len(t, got, 3, "should still have three items")
	require.Equal(t, []string{"E", "D", "C"}, []string{
		got[0].Value.Name,
		got[1].Value.Name,
		got[2].Value.Name,
	})
}

func TestItems_MinMatch(t *testing.T) {
	// Another quick test: if minMatch is 10.0, even high items below 10.0 get ignored
	xs := largest.NewItems(2, 10.0)

	xs.Add(makeItem("X", 9.9))
	xs.Add(makeItem("Y", 10.0))
	xs.Add(makeItem("Z", 12.0))

	// Expect only Y(10.0) and Z(12.0), ignoring X(9.9)
	got := xs.Items()
	require.Len(t, got, 2)
	// Descending: Z(12.0), Y(10.0)
	require.Equal(t, []string{"Z", "Y"}, []string{
		got[0].Value.Name,
		got[1].Value.Name,
	})
}
