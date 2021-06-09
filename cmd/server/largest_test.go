// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestLargest(t *testing.T) {
	xs := newLargest(10, 0.0)

	min := 10000.0
	for i := 0; i < 1000; i++ {
		it := &item{
			value:  i,
			weight: float64(rand.Intn(1000)),
		}
		xs.add(it)
		min = math.Min(min, it.weight)
	}

	// Check we didn't overflow capacity
	if len(xs.items) != xs.capacity {
		t.Errorf("len(xs.items)=%d != xs.capacity=%d", len(xs.items), xs.capacity)
	}

	for i := range xs.items {
		if i+1 > len(xs.items)-1 {
			continue // don't hit index out of bounds
		}

		if xs.items[i].weight < xs.items[i+1].weight {
			t.Errorf("xs.items[%d].weight=%.2f < xs.items[%d].weight=%.2f", i, xs.items[i].weight, i+1, xs.items[i+1].weight)
		}
	}
}

// TestLargest_MaxOrdering will test the ordering of 1.0 values to see
// if they hold their insert ordering.
func TestLargest_MaxOrdering(t *testing.T) {
	xs := newLargest(10, 0.0)

	xs.add(&item{value: "A", weight: 0.99})
	xs.add(&item{value: "B", weight: 1.0})
	xs.add(&item{value: "C", weight: 1.0})
	xs.add(&item{value: "D", weight: 1.0})
	xs.add(&item{value: "E", weight: 0.97})

	if n := len(xs.items); n != 10 {
		t.Fatalf("found %d items: %#v", n, xs.items)
	}

	if s, ok := xs.items[0].value.(string); !ok || s != "B" {
		t.Errorf("xs.items[0]=%#v", xs.items[0])
	}
	if s, ok := xs.items[1].value.(string); !ok || s != "C" {
		t.Errorf("xs.items[1]=%#v", xs.items[1])
	}
	if s, ok := xs.items[2].value.(string); !ok || s != "D" {
		t.Errorf("xs.items[2]=%#v", xs.items[2])
	}
	if s, ok := xs.items[3].value.(string); !ok || s != "A" {
		t.Errorf("xs.items[3]=%#v", xs.items[3])
	}
	if s, ok := xs.items[4].value.(string); !ok || s != "E" {
		t.Errorf("xs.items[4]=%#v", xs.items[4])
	}
	for i := 5; i < 10; i++ {
		if xs.items[i] != nil {
			t.Errorf("#%d was non-nil: %#v", i, xs.items[i])
		}
	}
}

func TestLargest__MinMatch(t *testing.T) {
	xs := newLargest(2, 0.96)

	xs.add(&item{value: "A", weight: 0.94})
	xs.add(&item{value: "B", weight: 1.0})
	xs.add(&item{value: "C", weight: 0.95})
	xs.add(&item{value: "D", weight: 0.09})

	require.Equal(t, "B", xs.items[0].value)
	require.Nil(t, xs.items[1])
}
