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

func TestEUDownload(t *testing.T) {
	if testing.Short() {
		return
	}

	file, err := DownloadEU(log.NewNopLogger(), "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("file in test: ", file)
	if file == "" {
		t.Fatal("no EU CSL file")
	}
	defer os.RemoveAll(filepath.Dir(file))

	if !strings.EqualFold("eu_csl.csv", filepath.Base(file)) {
		t.Errorf("unknown file %s", file)
	}
}

func TestEUDownload_initialDir(t *testing.T) {
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
	mk(t, "eu_csl.csv", "file=eu_csl.csv")

	file, err := DownloadEU(log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	if file == "" {
		t.Fatal("no EU CSL file")
	}

	if strings.EqualFold("eu_csl.csv", filepath.Base(file)) {
		bs, err := os.ReadFile(file)
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
