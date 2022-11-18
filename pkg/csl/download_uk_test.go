// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
)

func TestUKDownload(t *testing.T) {
	if testing.Short() {
		return
	}

	file, err := DownloadUK(log.NewNopLogger(), "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("file in test: ", file)
	if file == "" {
		t.Fatal("no UK CSL file")
	}
	defer os.RemoveAll(filepath.Dir(file))

	if !strings.EqualFold("ConList.csv", filepath.Base(file)) {
		t.Errorf("unknown file %s", file)
	}
}

func TestUKDownload_initialDir(t *testing.T) {
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
	mk(t, "ConList.csv", "file=ConList.csv")

	file, err := DownloadUK(log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	if file == "" {
		t.Fatal("no UK CSL file")
	}

	if strings.EqualFold("ConList.csv", filepath.Base(file)) {
		bs, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		if v := string(bs); v != "file=ConList.csv" {
			t.Errorf("ConList.csv: %v", v)
		}
	} else {
		t.Fatalf("unknown file: %v", file)
	}
}
