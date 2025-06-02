// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/moov-io/watchman/pkg/download"
)

// SanctionsEntry represents a single row in the CSV
type SanctionsEntry struct {
	ID                     string
	Source                 string
	EntityNumber           string
	Type                   string
	Programs               string
	Name                   string
	Title                  string
	Addresses              string
	FederalRegisterNotice  string
	StartDate              string
	EndDate                string
	StandardOrder          string
	LicenseRequirement     string
	LicensePolicy          string
	CallSign               string
	VesselType             string
	GrossTonnage           string
	GrossRegisteredTonnage string
	VesselFlag             string
	VesselOwner            string
	Remarks                string
	SourceListURL          string
	AltNames               string
	Citizenships           string
	DatesOfBirth           string
	Nationalities          string
	PlacesOfBirth          string
	SourceInformationURL   string
	IDs                    string
}

// ListData holds the parsed CSV data and its hash
type ListData struct {
	SanctionsData []SanctionsEntry
	ListHash      string
}

// Read processes the input files and returns parsed CSV data
func Read(files download.Files) (*ListData, error) {
	for filename, contents := range files {
		switch strings.ToLower(filename) {
		case "consolidated.csv":
			return parseCSVContents(filename, contents)
		default:
			return nil, fmt.Errorf("unknown file %s", filename)
		}
	}
	return nil, errors.New("no files provided")
}

// parseCSVContents reads and parses the CSV contents
func parseCSVContents(filename string, contents io.ReadCloser) (*ListData, error) {
	var buf bytes.Buffer
	buftee := io.TeeReader(contents, &buf)
	defer contents.Close()

	// Read and parse CSV
	reader := csv.NewReader(buftee)
	reader.FieldsPerRecord = 29 // Based on the provided CSV structure
	reader.ReuseRecord = true   // Reuse the record slice to reduce memory allocations

	// Read header
	_, err := reader.Read() // Skip header row
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	var sanctionsData []SanctionsEntry
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse CSV: %w", err)
		}

		// Skip SDN records as per user instruction
		if strings.Contains(strings.ToLower(record[1]), "specially designated nationals") {
			continue
		}

		// Map CSV record to SanctionsEntry
		entry := SanctionsEntry{
			ID:                     record[0],
			Source:                 record[1],
			EntityNumber:           record[2],
			Type:                   record[3],
			Programs:               record[4],
			Name:                   record[5],
			Title:                  record[6],
			Addresses:              record[7],
			FederalRegisterNotice:  record[8],
			StartDate:              record[9],
			EndDate:                record[10],
			StandardOrder:          record[11],
			LicenseRequirement:     record[12],
			LicensePolicy:          record[13],
			CallSign:               record[14],
			VesselType:             record[15],
			GrossTonnage:           record[16],
			GrossRegisteredTonnage: record[17],
			VesselFlag:             record[18],
			VesselOwner:            record[19],
			Remarks:                record[20],
			SourceListURL:          record[21],
			AltNames:               record[22],
			Citizenships:           record[23],
			DatesOfBirth:           record[24],
			Nationalities:          record[25],
			PlacesOfBirth:          record[26],
			SourceInformationURL:   record[27],
			IDs:                    record[28],
		}
		sanctionsData = append(sanctionsData, entry)
	}

	// Compute hash of the entire file content
	listHash := sha256.Sum256(buf.Bytes())
	list := &ListData{
		SanctionsData: sanctionsData,
		ListHash:      hex.EncodeToString(listHash[:]),
	}

	return list, nil
}

// readCloserToReaderAt converts an io.ReadCloser to io.ReaderAt for compatibility
func readCloserToReaderAt(rc io.ReadCloser) (io.ReaderAt, int64, error) {
	defer rc.Close()

	tempFile, err := os.CreateTemp("", "temp_csv")
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
