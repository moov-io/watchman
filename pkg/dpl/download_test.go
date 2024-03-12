// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package dpl

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
)

func TestDownloader(t *testing.T) {
	if testing.Short() {
		return
	}

	file, err := Download(log.NewNopLogger(), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(file) == 0 {
		t.Fatal("no DPL file")
	}
	for filename := range file {
		if !strings.EqualFold("dpl.txt", filepath.Base(filename)) {
			t.Errorf("unknown file %s", file)
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

	file, err := Download(log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(file) == 0 {
		t.Fatal("no DPL file")
	}
	for fn, fd := range file {
		if strings.EqualFold("dpl.txt", filepath.Base(fn)) {
			bs, err := io.ReadAll(fd)
			if err != nil {
				t.Fatal(err)
			}
			if v := string(bs); v != "file=dpl.txt" {
				t.Errorf("dpl.txt: %v", v)
			}
		} else {
			t.Fatalf("unknown file: %v", file)
		}
	}
}
