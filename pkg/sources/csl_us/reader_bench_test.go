package csl_us

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/fshelp"

	"github.com/stretchr/testify/require"
)

func BenchmarkRead(b *testing.B) {
	logger := log.NewTestLogger()

	pkg, err := fshelp.FindPkgDir()
	require.NoError(b, err)

	initialDir := filepath.Join(pkg, "pkg", "sources", "csl_us", "testdata")

	files, err := Download(b.Context(), logger, initialDir)
	require.NoError(b, err)

	for b.Loop() {
		data, err := Read(files)
		require.NoError(b, err)

		if n := len(data.SanctionsData); n != 5512 {
			b.Fatalf("read %d records", n)
		}
	}
}
