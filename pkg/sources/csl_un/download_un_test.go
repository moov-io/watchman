// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
)

func TestUNSanctionsListDownload_initialDir(t *testing.T) {
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
	mk(t, "UN_Sanctions_List.xml", "file=UN_Sanctions_List.xml")

	file, err := DownloadSanctionsList_UN(context.Background(), log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(file) == 0 {
		t.Fatal("no UN Sanctions List file")
	}

	for fn, fd := range file {
		if strings.EqualFold("UN_Sanctions_List.xml", filepath.Base(fn)) {
			_, err := io.ReadAll(fd)
			if err != nil {
				t.Fatal(err)
			}
		} else {
			t.Fatalf("unknown file: %v", file)
		}
	}
}

// Test downloading from a supplied file:// URL via environment override.
func TestUNSanctionsListDownload_fromURL(t *testing.T) {
	// prepare temporary file with known content
	f, err := os.CreateTemp("", "unlist-*.xml")
	if err != nil {
		t.Fatal(err)
	}
	content := "hello-un-list"
	if _, err := f.WriteString(content); err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(f.Name())

	// override both the environment and the package variable
	orig := unSanctionsListURL
	unSanctionsListURL = "file://" + f.Name()
	defer func() { unSanctionsListURL = orig }()

	os.Setenv("UN_SANCTIONS_LIST_URL", "file://"+f.Name())
	defer os.Unsetenv("UN_SANCTIONS_LIST_URL")

	dir, err := os.MkdirTemp("", "download-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	files, err := DownloadSanctionsList_UN(context.Background(), log.NewNopLogger(), dir)
	if err != nil {
		t.Fatalf("expected download success, got %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("expected exactly 1 file, got %d", len(files))
	}

	for _, rc := range files {
		buf, err := io.ReadAll(rc)
		if err != nil {
			t.Fatal(err)
		}
		if string(buf) != content {
			t.Errorf("downloaded content mismatch: %q", string(buf))
		}
	}
}

// When the URL is invalid, we should get an error
func TestUNSanctionsListDownload_badURL(t *testing.T) {
	orig := unSanctionsListURL
	defer func() { unSanctionsListURL = orig }()

	unSanctionsListURL = "file:///nonexistent/path"

	files, err := DownloadSanctionsList_UN(context.Background(), log.NewNopLogger(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected no files for bad URL, got %d", len(files))
	}
}
