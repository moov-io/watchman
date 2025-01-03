// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package largest

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func randomWeight() float64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))
	return float64(n.Int64()) / 100.0
}

func TestLargest(t *testing.T) {
	xs := NewItems(10, 0.0)

	min := 10000.0
	for i := 0; i < 1000; i++ {
		it := &Item{
			Value:  i,
			Weight: randomWeight(),
		}
		xs.Add(it)
		min = math.Min(min, it.Weight)
	}

	// Check we didn't overflow
	items := xs.Items()
	require.Equal(t, len(items), 10)

	for i := range items {
		if i+1 > len(items)-1 {
			continue // don't hit index out of bounds
		}
		if items[i].Weight < 0.0001 {
			t.Fatalf("weight of %.2f is too low", items[i].Weight)
		}
		if items[i].Weight < items[i+1].Weight {
			t.Errorf("items[%d].Weight=%.2f < items[%d].Weight=%.2f", i, items[i].Weight, i+1, items[i+1].Weight)
		}
	}
}

// TestLargest_MaxOrdering will test the ordering of 1.0 values to see
// if they hold their insert ordering.
func TestLargest_MaxOrdering(t *testing.T) {
	xs := NewItems(10, 0.0)

	xs.Add(&Item{Value: "A", Weight: 0.99})
	xs.Add(&Item{Value: "B", Weight: 1.0})
	xs.Add(&Item{Value: "C", Weight: 1.0})
	xs.Add(&Item{Value: "D", Weight: 1.0})
	xs.Add(&Item{Value: "E", Weight: 0.97})

	if n := len(xs.items); n != 10 {
		t.Fatalf("found %d items: %#v", n, xs.items)
	}

	if s, ok := xs.items[0].Value.(string); !ok || s != "B" {
		t.Errorf("xs.items[0]=%#v", xs.items[0])
	}
	if s, ok := xs.items[1].Value.(string); !ok || s != "C" {
		t.Errorf("xs.items[1]=%#v", xs.items[1])
	}
	if s, ok := xs.items[2].Value.(string); !ok || s != "D" {
		t.Errorf("xs.items[2]=%#v", xs.items[2])
	}
	if s, ok := xs.items[3].Value.(string); !ok || s != "A" {
		t.Errorf("xs.items[3]=%#v", xs.items[3])
	}
	if s, ok := xs.items[4].Value.(string); !ok || s != "E" {
		t.Errorf("xs.items[4]=%#v", xs.items[4])
	}
	for i := 5; i < 10; i++ {
		if xs.items[i] != nil {
			t.Errorf("#%d was non-nil: %#v", i, xs.items[i])
		}
	}
}

func TestLargest__MinMatch(t *testing.T) {
	xs := NewItems(2, 0.96)

	xs.Add(&Item{Value: "A", Weight: 0.94})
	xs.Add(&Item{Value: "B", Weight: 1.0})
	xs.Add(&Item{Value: "C", Weight: 0.95})
	xs.Add(&Item{Value: "D", Weight: 0.09})

	require.Equal(t, "B", xs.items[0].Value)
	require.Nil(t, xs.items[1])
}

func BenchmarkLargest(b *testing.B) {
	size := b.N * 500_000

	scores := make([]float64, size)
	for i := 0; i < b.N; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			b.Fatal(err)
		}
		scores[i] = float64(n.Int64()) / 100.0
	}

	limit := 20
	matches := []float64{0.1, 0.25, 0.5, 0.75, 0.9, 0.99}
	for i := range matches {
		b.Run(fmt.Sprintf("%.2f%%", matches[i]*100), func(b *testing.B) {
			// accumulate scores
			xs := NewItems(limit, matches[i])

			g := &errgroup.Group{}
			for i := range scores {
				score := scores[i]
				g.Go(func() error {
					xs.Add(&Item{
						Value:  ofac.SDN{},
						Weight: score,
					})
					return nil
				})
			}
			require.NoError(b, g.Wait())
			require.Len(b, xs.items, limit)
			require.Equal(b, limit, cap(xs.items))
		})
	}
}
