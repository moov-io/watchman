// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

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

func TestEUDownload(t *testing.T) {
	if testing.Short() {
		return
	}

	files, err := DownloadEU(context.Background(), log.NewNopLogger(), "")
	require.NoError(t, err)
	require.Len(t, files, 1)

	for filename := range files {
		if !strings.EqualFold("eu_csl.csv", filepath.Base(filename)) {
			t.Errorf("unknown file %s", filename)
		}
	}
}

func TestEUDownload_initialDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "initial-dir")
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
	mk(t, "eu_csl.csv", "file=eu_csl.csv")

	file, err := DownloadEU(context.Background(), log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(file) == 0 {
		t.Fatal("no EU CSL file")
	}

	for fn, fd := range file {
		if strings.EqualFold("eu_csl.csv", filepath.Base(fn)) {
			bs, err := io.ReadAll(fd)
			if err != nil {
				t.Fatal(err)
			}
			if v := string(bs); v != "file=eu_csl.csv" {
				t.Errorf("eu_csl.csv: %v", v)
			}
		} else {
			t.Fatalf("unknown file: %v", file)
		}
	}
}
