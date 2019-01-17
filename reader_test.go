package ofac

import (
	"encoding/json"
	"os"
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

// TestInvalidAddressCSVFile validates validates an invalid Address CSV File returns an error
func TestInvalidAddressCSVFile(t *testing.T) {
	r := Reader{}

	r.FileName = "invalid/add.csv"
	if err := r.Read(); err != nil {
		if e, ok := err.(*os.PathError); ok {
			if !strings.Contains(e.Error(), "The system cannot find the path specified") {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidAlternateIDCSVFile validates validates an invalid Alternate ID CSV File returns an error
func TestInvalidAlternateIDCSVFile(t *testing.T) {
	r := Reader{}

	r.FileName = "invalid/alt.csv"
	if err := r.Read(); err != nil {
		if e, ok := err.(*os.PathError); ok {
			if !strings.Contains(e.Error(), "The system cannot find the path specified") {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}

	}
}

// TestInvalidSDNCSVFile validates an invalid SDN CSV File returns an error
func TestInvalidSDNCSVFile(t *testing.T) {
	r := Reader{}
	r.FileName = "invalid/sdn.csv"
	if err := r.Read(); err != nil {
		if e, ok := err.(*os.PathError); ok {
			if !strings.Contains(e.Error(), "The system cannot find the path specified") {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidSDNCommentsCSVFile validates an invalid SDNComments CSV File returns an error
func TestInvalidSDNCommentsCSVFile(t *testing.T) {
	r := Reader{}
	r.FileName = "invalid/sdn_comments.csv"
	if err := r.Read(); err != nil {
		if e, ok := err.(*os.PathError); ok {
			if !strings.Contains(e.Error(), "The system cannot find the path specified") {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddressCSVInvalidLineRead validates an error is returned when reading an OFAC Specially Designated National
// Address CSV File with an invalid line
func TestAddressCSVInvalidLineRead(t *testing.T) {
	r := Reader{}
	r.FileName = "test/testdata/invalidFiles/add.csv"
	if err := r.Read(); err != nil {
		if !strings.Contains(err.Error(), "wrong number of fields") {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAlternateIDCSVInvalidLineRead validates an error is returned when reading an OFAC Specially Designated National
// Alternate ID CSV File with an invalid line
func TestAlternateIDCSVInvalidLineRead(t *testing.T) {
	r := Reader{}
	r.FileName = "test/testdata/invalidFiles/alt.csv"
	if err := r.Read(); err != nil {
		if !strings.Contains(err.Error(), "wrong number of fields") {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSDNCSVInvalidLineRead validates an error is returned when reading an OFAC Specially Designated National
// CSV File with an invalid line
func TestSDNCSVInvalidLineRead(t *testing.T) {
	r := Reader{}
	r.FileName = "test/testdata/invalidFiles/sdn.csv"
	if err := r.Read(); err != nil {
		if !strings.Contains(err.Error(), "wrong number of fields") {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSDNCommentsCSVInvalidLineRead validates an error is returned when reading an OFAC Specially Designated National
// Comments CSV File with an invalid line
func TestSDNCommentsCSVInvalidLineRead(t *testing.T) {
	r := Reader{}
	r.FileName = "test/testdata/invalidFiles/sdn_comments.csv"
	if err := r.Read(); err != nil {
		if !strings.Contains(err.Error(), "wrong number of fields") {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseError validates parseError
func TestParseError(t *testing.T) {
	r := Reader{}
	r.FileName = "test/testdata/sdn_comments.csv"
	err := r.Read()

	if e := r.parseError(err); e != nil {
		t.Errorf("%T: %s", e, e)
	}
}
