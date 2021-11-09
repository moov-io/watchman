// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package dpl

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

// Read parses DPL records from a TXT file and populates the associated arrays.
//
// For more details on the raw DPL files see https://moov-io.github.io/watchman/file-structure.html
func Read(path string) ([]*DPL, error) {
	// open txt file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// create a new csv.Reader and set the delim char to txtDelim char
	reader := csv.NewReader(f)
	reader.Comma = '\t'

	// Loop through all lines we can
	var out []*DPL
	for {
		line, err := reader.Read()
		if err != nil {
			if err != nil {
				// reached the last line
				if errors.Is(err, io.EOF) {
					break
				}
				// malformed row
				if errors.Is(err, csv.ErrFieldCount) ||
					errors.Is(err, csv.ErrBareQuote) ||
					errors.Is(err, csv.ErrQuote) {
					continue
				}
				return nil, err
			}
		}

		if len(line) < 12 || (len(line) >= 2 && line[1] == "Street_Address") {
			continue // skip malformed headers
		}

		deniedPerson := &DPL{
			Name:           line[0],
			StreetAddress:  line[1],
			City:           line[2],
			State:          line[3],
			Country:        line[4],
			PostalCode:     line[5],
			EffectiveDate:  line[6],
			ExpirationDate: line[7],
			StandardOrder:  line[8],
			LastUpdate:     line[9],
			Action:         line[10],
			FRCitation:     line[11],
		}
		out = append(out, deniedPerson)
	}
	return out, nil
}
