// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
)

func TestDownloader(t *testing.T) {
	if testing.Short() {
		return
	}

	files, err := Download(log.NewNopLogger(), "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(filepath.Dir(files[0]))

	if len(files) != 4 {
		t.Errorf("OFAC: found %d files", len(files))
	}
	for i := range files {
		name := filepath.Base(files[i])
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

	files, err := Download(log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	for i := range files {
		switch filepath.Base(files[i]) {
		case "sdn.txt":
			bs, err := os.ReadFile(files[i])
			if err != nil {
				t.Fatal(err)
			}
			if v := string(bs); v != "file=sdn.csv" {
				t.Errorf("sdn.csv: %v", v)
			}
		}
	}
}
