// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"encoding/xml"
	"io"
)

// Reader handles the streaming of UN Consolidated List records.
type Reader struct {
	scanner io.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{scanner: r}
}

// Read calls the provided functions for each Individual and Entity found in the XML.
func (r *Reader) Read(onIndividual func(UNIndividual), onEntity func(UNEntity)) error {
	decoder := xml.NewDecoder(r.scanner)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "INDIVIDUAL" {
				var p UNIndividual
				if err := decoder.DecodeElement(&p, &se); err == nil {
					onIndividual(p)
				}
			} else if se.Name.Local == "ENTITY" {
				var e UNEntity
				if err := decoder.DecodeElement(&e, &se); err == nil {
					onEntity(e)
				}
			}
		}
	}
	return nil
}
