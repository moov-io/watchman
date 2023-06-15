// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"os"
	"path/filepath"
	"reflect"
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

func TestCleanPrgmsList(t *testing.T) {
	tests := []struct {
		prgms    string
		expected string
	}{
		{"SDGT] ", "SDGT"},
		{" SDGT] [IFSR", "SDGT; IFSR"},
		{"SDNTK] [FTO] [SDGT", "SDNTK; FTO; SDGT"},
		{"SDNTK] [FTO] [SDGT; IFSR]", "SDNTK; FTO; SDGT; IFSR"},
		{"[IFSR] [SDNTK] [FTO] [SDGT", "IFSR; SDNTK; FTO; SDGT"},
	}

	for _, test := range tests {
		if actual := cleanPrgmsList(test.prgms); actual != test.expected {
			t.Errorf("Wanted %q, got %q", test.expected, actual)
		}
	}
}

func TestSplitPrograms(t *testing.T) {
	if items := splitPrograms("SGDT"); !reflect.DeepEqual(items, []string{"SGDT"}) {
		t.Errorf("items=%v", items)
	}
	if items := splitPrograms("IRAN; SGDT"); !reflect.DeepEqual(items, []string{"IRAN", "SGDT"}) {
		t.Errorf("items=%v", items)
	}
	if items := splitPrograms("IFSR; SDNTK; FTO; SDGT"); !reflect.DeepEqual(items, []string{"IFSR", "SDNTK", "FTO", "SDGT"}) {
		t.Errorf("items=%v", items)
	}
}

func TestSDNComments(t *testing.T) {
	fd, err := os.CreateTemp("", "sdn-csv")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fd.WriteString(`28264,"hone Number 8613314257947; alt. Phone Number 8618004121000; Identification Number 210302198701102136 (China); a.k.a. "blackjack1987"; a.k.a. "khaleesi"; Linked To: LAZARUS GROUP."`); err != nil {
		t.Fatal(err)
	}

	// read with lazy quotes enabled
	if res, err := csvSDNCommentsFile(fd.Name()); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else {
		if len(res.SDNComments) != 1 {
			t.Errorf("SDNComments=%#v", res.SDNComments)
		}
		for i := range res.SDNComments {
			t.Logf("\ncomment #%d\n entity=%s\n remarks=%v", i, res.SDNComments[i].EntityID, res.SDNComments[i].RemarksExtended)
		}
	}
}
