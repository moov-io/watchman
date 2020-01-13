// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestLargest(t *testing.T) {
	xs := newLargest(10)

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
