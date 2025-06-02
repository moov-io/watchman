// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"

	"github.com/stretchr/testify/require"
)

func BenchmarkFindEntity(b *testing.B) {
	for b.Loop() {
		entity := ofactest.FindEntity(b, "33151")
		require.NotEmpty(b, entity.Name)
	}
}
