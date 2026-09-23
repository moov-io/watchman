// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package index

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectSorted(t *testing.T) {
	require.Empty(t, intersectSorted(nil, []int{1}))
	require.Empty(t, intersectSorted([]int{1}, nil))
	require.Equal(t, []int{2, 4}, intersectSorted([]int{1, 2, 3, 4}, []int{2, 4, 6}))
	require.Equal(t, []int{5}, intersectSorted([]int{5}, []int{1, 3, 5, 7, 9}))
	require.Empty(t, intersectSorted([]int{1, 2}, []int{3, 4}))
}

func TestUnionSorted(t *testing.T) {
	require.Nil(t, unionSorted(nil))
	require.Equal(t, []int{1, 2, 3, 5}, unionSorted([][]int{{1, 3, 5}, {1, 2}}))
}
