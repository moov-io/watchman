// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestOFACDownloader(t *testing.T) {
	dl := ofacDownloader{}
	dir, err := dl.getFiles()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// check file count
	fds, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(fds) != len(ofacFilenames) {
		t.Errorf("OAFC: expected %d files but found %d", len(fds), len(ofacFilenames))
	}
	for i := range fds {
		name := fds[i].Name()
		switch name {
		case "add.csv", "alt.csv", "sdn.csv", "sdn_comments.csv":
			continue
		default:
			t.Errorf("unknown file %s", name)
		}
	}
}
