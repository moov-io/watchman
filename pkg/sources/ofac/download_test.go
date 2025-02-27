// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestDownloader(t *testing.T) {
	if testing.Short() {
		return
	}

	files, err := Download(context.Background(), log.NewNopLogger(), "")
	require.NoError(t, err)
	require.Len(t, files, 4)

	for fn := range files {
		name := strings.ToLower(filepath.Base(fn))
		switch name {
		case "add.csv", "alt.csv", "sdn.csv", "sdn_comments.csv":
			continue
		default:
			t.Errorf("unknown file %s", name)
		}
	}
}

func TestDownloader__initialDir(t *testing.T) {
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
	mk(t, "sdn.csv", "file=sdn.csv")
	mk(t, "dpl.txt", "file=dpl.txt")

	files, err := Download(context.Background(), log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	for fn := range files {
		switch filepath.Base(fn) {
		case "sdn.txt":
			bs, err := os.ReadFile(fn)
			if err != nil {
				t.Fatal(err)
			}
			if v := string(bs); v != "file=sdn.csv" {
				t.Errorf("sdn.csv: %v", v)
			}
		}
	}
}
