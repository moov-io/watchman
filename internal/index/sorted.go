// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package index

import "slices"

// intersectSorted returns the sorted intersection of two sorted int slices.
// Neither input is modified. A nil or empty input yields nil.
func intersectSorted(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}
	if len(a) > len(b) {
		a, b = b, a
	}

	// Binary-search the smaller list into the larger when that is cheaper than
	// a linear merge (tiny posting lists against a large partition).
	if len(a) < 32 || uint64(len(a))*uint64(64) < uint64(len(b)) {
		out := make([]int, 0, len(a))
		for _, v := range a {
			if _, found := slices.BinarySearch(b, v); found {
				out = append(out, v)
			}
		}
		return out
	}

	out := make([]int, 0, len(a))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] == b[j]:
			out = append(out, a[i])
			i++
			j++
		case a[i] < b[j]:
			i++
		default:
			j++
		}
	}
	return out
}

// unionSorted returns the sorted unique union of sorted int slices.
func unionSorted(lists [][]int) []int {
	n := 0
	for _, l := range lists {
		n += len(l)
	}
	if n == 0 {
		return nil
	}
	out := make([]int, 0, n)
	for _, l := range lists {
		out = append(out, l...)
	}
	slices.Sort(out)
	return slices.Compact(out)
}
