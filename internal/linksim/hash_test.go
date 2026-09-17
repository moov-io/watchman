// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package linksim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDigest(t *testing.T) {
	require.Empty(t, digest("name-soundex", ""))
	require.Empty(t, digest("name-soundex", "   "))

	a := digest("gov-id", "1234567890")
	require.Len(t, a, 8)
	require.Equal(t, a, digest("gov-id", "1234567890"))
	require.NotEqual(t, a, digest("gov-country", "1234567890"), "field name must domain-separate hashes")
	require.NotEqual(t, a, digest("gov-id", "1234567891"))
}
