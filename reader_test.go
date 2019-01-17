// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestAddressCSVFileRead validates reading an OFAC Address CSV File
func TestAddressCSVFileRead(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/add.csv"
	if err := r.Read(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAlternateIDCSVFileRead validates reading an OFAC Alternate ID CSV File
func TestAlternateIDCSVFileRead(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/alt.csv"
	if err := r.Read(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSDNCSVFileRead validates reading an OFAC Specially Designated National CSV File
func TestSDNCSVFileRead(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/sdn.csv"
	if err := r.Read(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSDNCommentsCSVFileRead validates reading an OFAC Specially Designated National Comments CSV File
func TestSDNCommentsCSVFileRead(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/sdn_comments.csv"
	if err := r.Read(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestInvalidFileExtension validates the file extension is csv
func TestInvalidFileExtension(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/add.csb"
	err := r.Read()

	if !strings.Contains(err.Error(), "file type") {
		t.Errorf("%T: %s", err, err)
	}
}

// TestInvalidFileName validates the file name is valid
func TestInvalidFileName(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/xyz.csv"
	err := r.Read()

	if !strings.Contains(err.Error(), "file name") {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSDNCSVFileHit validates reading an OFAC Specially Designated National CSV File Hit
func TestSDNCSVFileHit(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/sdn.csv"
	err := r.Read()

	if err != nil {
		t.Errorf("%T: %s", err, err)
	}

	hit := false
	for _, sdn := range r.SDNArray {
		if sdn.SDNName == "HAWATMA, Nayif" {
			hit = true
		}
	}
	if !hit {
		t.Errorf("%s", "the check missed a specially designated name")
	}
}

// TestSDNCSVFileJSON tests reading an OFAC Specially Designated National CSV File and formatting to JSON
func TestSDNCSVFileJSON(t *testing.T) {
	r := Reader{}

	r.FileName = "test/testdata/sdn.csv"
	if err := r.Read(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if _, err := json.Marshal(r.SDNArray); err != nil {
		t.Errorf("%T: %s", err, err)
	}

}
