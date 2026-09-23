// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package index

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQwertyVariants(t *testing.T) {
	vars := qwertyVariants("9263643")
	require.Equal(t, "9263643", vars[0])
	require.True(t, slices.Contains(vars, "9263642"), "3 is adjacent to 2 on the number row")
	require.True(t, slices.Contains(vars, "9263463"), "adjacent transposition")
	require.False(t, slices.Contains(vars, "9263649"), "3 is not adjacent to 9")
}

func TestQwertyVariantsEmpty(t *testing.T) {
	require.Nil(t, qwertyVariants(""))
}
