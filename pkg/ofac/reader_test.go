// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"path/filepath"
	"testing"
)

// TestOFAC__read validates reading an OFAC Address CSV File
func TestOFAC__read(t *testing.T) {
	res, err := Read(filepath.Join("..", "..", "test", "testdata", "add.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Addresses) == 0 {
		t.Errorf("found no Addresses")
	}

	res, err = Read(filepath.Join("..", "..", "test", "testdata", "alt.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if len(res.AlternateIdentities) == 0 {
		t.Errorf("found no AlternateIdentities")
	}

	res, err = Read(filepath.Join("..", "..", "test", "testdata", "sdn.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if len(res.SDNs) == 0 {
		t.Errorf("found no SDNs")
	}

	res, err = Read(filepath.Join("..", "..", "test", "testdata", "sdn_comments.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if len(res.SDNComments) == 0 {
		t.Errorf("found no SDN comments")
	}
}

func TestReplaceNull(t *testing.T) {
	ans := replaceNull(nil)
	if ans != nil {
		t.Errorf("Got %v", ans)
	}
	ans = replaceNull([]string{" -0-"})
	if len(ans) != 1 || ans[0] != "" {
		t.Errorf("Got %v", ans)
	}
	ans = replaceNull([]string{"foo", " -0-"})
	if len(ans) != 2 || ans[0] != "foo" || ans[1] != "" {
		t.Errorf("Got %v", ans)
	}
}
