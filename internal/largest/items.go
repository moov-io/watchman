// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package largest

import (
	"sort"
	"sync"

	"github.com/moov-io/watchman/pkg/search"
)

type Item struct {
	Value  search.Entity[search.Value]
	Weight float64
}

// Items keeps track of a set of Items (with a Value and Weight),
// up to a fixed capacity, only retaining the highest weights (>= minMatch).
type Items struct {
	mu       sync.Mutex
	items    []Item
	capacity int
	minMatch float64
}

// NewItems returns a structure which tracks the top-weighted Items,
// subject to minMatch and a fixed capacity.
func NewItems(capacity int, minMatch float64) *Items {
	if minMatch <= 0.001 {
		minMatch = 0.01
	}
	return &Items{
		items:    make([]Item, 0, capacity),
		capacity: capacity,
		minMatch: minMatch,
	}
}

// Add inserts an Item if it meets the minMatch threshold,
// ensuring we only keep the top N items by Weight.
func (xs *Items) Add(it Item) {
	if it.Weight < xs.minMatch {
		// Skip if below minMatch threshold
		return
	}

	xs.mu.Lock()
	defer xs.mu.Unlock()

	// If there's room, just insert in the correct spot
	if len(xs.items) < xs.capacity {
		xs.insertDescending(it)
		return
	}

	// We are at capacity, so compare the new item to the smallest in our list
	// (in descending order, the smallest is the last element).
	if it.Weight <= xs.items[len(xs.items)-1].Weight {
		// New item is not heavier than our smallest stored item
		return
	}

	// Remove the smallest item at the end...
	xs.items = xs.items[:len(xs.items)-1]

	// ...and insert the new one
	xs.insertDescending(it)
}

// insertDescending inserts an Item so that xs.items remains
// sorted by Weight in descending order (index 0 is highest).
func (xs *Items) insertDescending(it Item) {
	// Find the position using binary search
	// We want the first spot where items[i].Weight < it.Weight.
	i := sort.Search(len(xs.items), func(i int) bool {
		return xs.items[i].Weight < it.Weight
	})
	// Extend the slice by 1
	xs.items = append(xs.items, it)
	// Shift everything after i right by 1
	copy(xs.items[i+1:], xs.items[i:])
	// Place the new item at position i
	xs.items[i] = it
}

// All returns a copy of all items in descending Weight order.
func (xs *Items) Items() []Item {
	xs.mu.Lock()
	defer xs.mu.Unlock()

	out := make([]Item, len(xs.items))
	copy(out, xs.items)
	return out
}
