// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
)

type fd struct {
	name string
}

func (fd *fd) Name() string       { return fd.name }
func (fd *fd) Size() int64        { return 1 }
func (fd *fd) Mode() os.FileMode  { return 0600 }
func (fd *fd) ModTime() time.Time { return time.Now() }
func (fd *fd) IsDir() bool        { return false }
func (fd *fd) Sys() interface{}   { return nil }

func TestDownloader__compareNames(t *testing.T) {
	var fds = []os.FileInfo{
		&fd{name: "sdn.csv"},
		&fd{name: "dpl.txt"},
	}
	expected := map[string]string{
		"sdn.csv": "https://example.com",
		"dpl.txt": "https://example.com",
	}

	// first case, all matched
	matched, missing := compareNames(fds, expected)
	if (matched != "sdn.csv, dpl.txt" && matched != "dpl.txt, sdn.csv") || missing != "" {
		t.Errorf("matched=%q missing=%q", matched, missing)
	}
	matched, missing = compareNames(nil, expected)
	if matched != "" || (missing != "sdn.csv, dpl.txt" && missing != "dpl.txt, sdn.csv") {
		t.Errorf("matched=%q missing=%q", matched, missing)
	}
	matched, missing = compareNames([]os.FileInfo{&fd{name: "dpl.txt"}}, expected)
	if matched != "dpl.txt" || missing != "sdn.csv" {
		t.Errorf("matched=%q missing=%q", matched, missing)
	}
	matched, missing = compareNames(fds, make(map[string]string))
	if matched != "" || (missing != "sdn.csv, dpl.txt" && missing != "dpl.txt, sdn.csv") {
		t.Errorf("matched=%q missing=%q", matched, missing)
	}
}

func TestDownloader(t *testing.T) {
	if testing.Short() {
		return
	}

	dl := Downloader{
		Logger: log.NewNopLogger(),
	}
	dir, err := dl.GetFiles("")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// check file count
	fds, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	numFiles := len(ofacFilenames) + len(dplFilenames)
	if len(fds) != numFiles {
		t.Errorf("OAFC: expected %d files but found %d", len(fds), numFiles)
	}
	for i := range fds {
		name := fds[i].Name()
		switch name {
		case "add.csv", "alt.csv", "sdn.csv", "sdn_comments.csv", "dpl.txt":
			continue
		default:
			t.Errorf("unknown file %s", name)
		}
	}

	// nil Downloader
	var dl2 *Downloader
	if _, err := dl2.GetFiles(""); err == nil {
		t.Error("expected error")
	}
}

func TestDownloader__initialDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "iniital-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	mk := func(t *testing.T, name string, body string) {
		if err := ioutil.WriteFile(filepath.Join(dir, name), []byte(body), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// create each file
	mk(t, "add.csv", "file=add.csv")
	mk(t, "alt.csv", "file=alt.csv")
	mk(t, "sdn.csv", "file=sdn.csv")
	mk(t, "sdn_comments.csv", "file=sdn_comments.csv")
	mk(t, "dpl.txt", "file=dpl.txt")

	dl := Downloader{
		// Logger: log.NewNopLogger(), // nil so GetFiles sets this
	}
	out, err := dl.GetFiles(dir)
	if err != nil {
		t.Fatal(err)
	}

	// read a couple files
	bs, err := ioutil.ReadFile(filepath.Join(out, "add.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if v := string(bs); v != "file=add.csv" {
		t.Error("unexpected contents in add.csv")
	}
	// read another file
	bs, err = ioutil.ReadFile(filepath.Join(out, "dpl.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if v := string(bs); v != "file=dpl.txt" {
		t.Error("unexpected contents in dpl.txt")
	}

	// use an invalid initial directory to get an error
	out, err = dl.GetFiles(filepath.Join("this", "path", "doesn't", "exist"))
	if err == nil {
		t.Error("expected error")
	}
	if len(out) != 0 {
		t.Errorf("got %d files", len(out))
	}
}
