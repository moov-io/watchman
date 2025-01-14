// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/moov-io/watchman/pkg/csl_us/gen/ENHANCED_XML"
)

func Read(files map[string]io.ReadCloser) (*ENHANCED_XML.SanctionsData, error) {
	for filename, contents := range files {
		switch strings.ToLower(filename) {
		case "cons_enhanced.zip":
			return parseZipContents(filename, contents)
		default:
			return nil, fmt.Errorf("unknown file %s", filename)
		}
	}
	return nil, errors.New("no files provided")
}

func parseZipContents(filename string, contents io.ReadCloser) (*ENHANCED_XML.SanctionsData, error) {
	readerAt, size, err := readCloserToReaderAt(contents)
	if err != nil {
		return nil, fmt.Errorf("preparing %s for zip read: %w", filename, err)
	}

	r, err := zip.NewReader(readerAt, size)
	if err != nil {
		return nil, fmt.Errorf("opening %s as zip: %w", filename, err)
	}

	needle := "CONS_ENHANCED.XML"

	var file *zip.File
	for _, f := range r.File {
		if strings.EqualFold(f.Name, needle) {
			file = f
			break
		}
	}
	if file == nil {
		return nil, fmt.Errorf("%s not found", needle)
	}

	fr, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open %s inside zip: %w", needle, err)
	}
	defer fr.Close()

	var doc ENHANCED_XML.SanctionsData
	decoder := xml.NewDecoder(fr)

	err = decoder.Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return &doc, nil

}

func readCloserToReaderAt(rc io.ReadCloser) (io.ReaderAt, int64, error) {
	defer rc.Close()

	tempFile, err := os.CreateTemp("", "temp_zip")
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create temp file: %w", err)
	}

	size, err := io.Copy(tempFile, rc)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, 0, fmt.Errorf("failed to write to temp file: %w", err)
	}

	if err := tempFile.Sync(); err != nil {
		return nil, 0, fmt.Errorf("tempfile.sync failed: %w", err)
	}

	if _, err := tempFile.Seek(0, io.SeekStart); err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, 0, fmt.Errorf("failed to seek temp file: %w", err)
	}

	return tempFile, size, nil
}

// func ToEntity(src ENHANCED_XML.EntitiesEntity) search.Entity[ENHANCED_XML.EntitiesEntity]
