// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package largest

import (
	"slices"
	"sync"
)

// Item represents an arbitrary value with an associated weight
type Item struct {
	Value  interface{}
	Weight float64
}

// NewItems returns a structure which can be used to track items with the highest weights
func NewItems(capacity int, minMatch float64) *Items {
	return &Items{
		items:    make([]*Item, capacity),
		capacity: capacity,
		minMatch: minMatch,
	}
}

// Items keeps track of a set of items with the lowest weights. This is used to
// find the largest weighted values out of a much larger set.
type Items struct {
	items    []*Item
	capacity int
	minMatch float64
	mu       sync.Mutex
}

func (xs *Items) Add(it *Item) {
	if it.Weight < xs.minMatch {
		return // skip item as it's below our threshold
	}

	xs.mu.Lock()
	defer xs.mu.Unlock()

	for i := range xs.items {
		if xs.items[i] == nil {
			xs.items[i] = it // insert if we found empty slot
			break
		}
		if xs.items[i].Weight < it.Weight {
			xs.items = slices.Insert(xs.items, i, it)
			break
		}
	}
	if len(xs.items) > xs.capacity {
		xs.items = xs.items[:xs.capacity]
	}
}

func (xs *Items) Items() []*Item {
	xs.mu.Lock()
	defer xs.mu.Unlock()

	out := make([]*Item, len(xs.items))
	copy(out, xs.items)
	return out
}
