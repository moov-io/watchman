// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/cslustest"

	"github.com/stretchr/testify/require"
)

func BenchmarkFindEntity(b *testing.B) {
	for b.Loop() {
		entity := cslustest.FindEntity(b, "233a4d725770c81fb561ffe3842c14010c2201971d6be62eca1e613b")
		require.NotEmpty(b, entity.Name)
	}
}
