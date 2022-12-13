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

func TestUKCSLDownload(t *testing.T) {
	if testing.Short() {
		return
	}

	file, err := DownloadUKCSL(log.NewNopLogger(), "")
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

func TestUKCSLDownload_initialDir(t *testing.T) {
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

	file, err := DownloadUKCSL(log.NewNopLogger(), dir)
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

func TestUKSanctionsListDownload(t *testing.T) {
	if testing.Short() {
		return
	}

	file, err := DownloadUKSanctionsList(log.NewNopLogger(), "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("file in test: ", file)
	if file == "" {
		t.Fatal("no UK Sanctions List file")
	}
	defer os.RemoveAll(filepath.Dir(file))

	if !strings.EqualFold("UK_Sanctions_List.ods", filepath.Base(file)) {
		t.Errorf("unknown file %s", file)
	}
}

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
	mk(t, "UK_Sanctions_List.ods", "file=UK_Sanctions_List.ods")

	file, err := DownloadUKSanctionsList(log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	if file == "" {
		t.Fatal("no UK Sanctions List file")
	}

	if strings.EqualFold("UK_Sanctions_List.ods", filepath.Base(file)) {
		_, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		// if v := string(bs); v != "file=UK_Sanctions_List.ods" {
		// 	t.Errorf("UK_Sanctions_List.ods: %v", v)
		// }
	} else {
		t.Fatalf("unknown file: %v", file)
	}
}
