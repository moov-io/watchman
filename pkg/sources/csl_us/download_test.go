// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	if testing.Short() {
		return
	}

	files, err := Download(context.Background(), log.NewNopLogger(), "")
	require.NoError(t, err)
	require.Len(t, files, 1)

	file, found := files["CONS_ENHANCED.ZIP"]
	require.True(t, found)
	require.NotNil(t, file)
	require.NoError(t, file.Close())
}

func TestDownload_initialDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "iniital-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	mk := func(t *testing.T, name string, body string) {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(body), 0600); err != nil {
			t.Fatalf("writing %s: %v", path, err)
		}
	}

	// create each file
	mk(t, "CONS_ENHANCED.ZIP", "file=cons_enhanced.zip")

	files, err := Download(context.Background(), log.NewNopLogger(), dir)
	require.NoError(t, err)
	require.Len(t, files, 1)

	for filename, fd := range files {
		if strings.EqualFold("CONS_ENHANCED.ZIP", filepath.Base(filename)) {
			bs, err := io.ReadAll(fd)
			require.NoError(t, err)

			require.Equal(t, "file=cons_enhanced.zip", string(bs))
		} else {
			t.Fatalf("unknown file: %v", filename)
		}
	}
}
