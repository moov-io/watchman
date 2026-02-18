// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
)

func TestUKSanctionsListDownload_initialDir(t *testing.T) {
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
	mk(t, "UK_Sanctions_List.csv", "file=UK_Sanctions_List.csv")

	file, err := DownloadSanctionsList(context.Background(), log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(file) == 0 {
		t.Fatal("no UK Sanctions List file")
	}

	for fn, fd := range file {
		if strings.EqualFold("UK_Sanctions_List.csv", filepath.Base(fn)) {
			_, err := io.ReadAll(fd)
			if err != nil {
				t.Fatal(err)
			}
		} else {
			t.Fatalf("unknown file: %v", file)
		}
	}
}
