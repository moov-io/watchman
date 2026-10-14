package csl_us

import (
	"testing"

	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

func BenchmarkRead(b *testing.B) {
	logger := log.NewTestLogger()

	for b.Loop() {
		b.StopTimer()
		// testdata/ is next to this package (FindPkgDir already ends in /pkg).
		files, err := Download(b.Context(), logger, "testdata")
		require.NoError(b, err)
		b.StartTimer()

		data, err := Read(files)
		require.NoError(b, err)
		require.NotEmpty(b, data.SanctionsData)
	}
}
