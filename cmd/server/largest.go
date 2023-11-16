// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import "sync"

// item represents an arbitrary value with an associated weight
type item struct {
	matched string
	value   interface{}
	weight  float64
}

// newLargest returns a `largest` instance which can be used to track items with the highest weights
func newLargest(capacity int, minMatch float64) *largest {
	return &largest{
		items:    make([]*item, capacity),
		capacity: capacity,
		minMatch: minMatch,
	}
}

// largest keeps track of a set of items with the lowest weights. This is used to
// find the largest weighted values out of a much larger set.
type largest struct {
	items    []*item
	capacity int
	minMatch float64
	mu       sync.Mutex
}

func (xs *largest) add(it *item) {
	xs.mu.Lock()
	defer xs.mu.Unlock()

	if it.weight < xs.minMatch {
		return // skip item as it's below our threshold
	}

	for i := range xs.items {
		if xs.items[i] == nil {
			xs.items[i] = it // insert if we found empty slot
			break
		}
		if xs.items[i].weight < it.weight {
			// insert at i, slide other items over
			xs.items = append(xs.items, nil)
			copy(xs.items[i+1:], xs.items[i:])
			xs.items[i] = it
			break
		}
	}
	if len(xs.items) > xs.capacity {
		xs.items = xs.items[:xs.capacity]
	}
}
