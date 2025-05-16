package fshelp_test

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/internal/fshelp"

	"github.com/stretchr/testify/require"
)

func TestFindPkgDir(t *testing.T) {
	pkg, err := fshelp.FindPkgDir()
	require.NoError(t, err)

	_, leaf := filepath.Split(pkg)
	require.Equal(t, "pkg", leaf)

	t.Run("not found", func(t *testing.T) {
		dir := t.TempDir()
		t.Chdir(dir)

		pkg, err := fshelp.FindPkgDir()
		require.ErrorContains(t, err, "no pkg ancestor found")
		require.Empty(t, pkg)
	})
}
