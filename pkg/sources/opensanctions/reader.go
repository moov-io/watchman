// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package opensanctions

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/moov-io/watchman/pkg/sources/senzing"
	"github.com/moov-io/watchman/pkg/download"
	"github.com/moov-io/watchman/pkg/search"
)

// Results holds the parsed OpenSanctions PEP data and its hash
type Results struct {
	Entities []search.Entity[search.Value]
	ListHash string
}

// Read processes the downloaded Senzing JSON file and returns parsed entities
func Read(files download.Files) (*Results, error) {
	for filename, contents := range files {
		switch strings.ToLower(filename) {
		case "peps_senzing.json":
			return parseSenzingJSON(contents)
		default:
			return nil, fmt.Errorf("unknown file %s", filename)
		}
	}
	return nil, errors.New("no files provided")
}

// parseSenzingJSON reads and parses the Senzing JSON file
func parseSenzingJSON(contents io.ReadCloser) (*Results, error) {
	defer contents.Close()

	// Read all content for hashing
	var buf bytes.Buffer
	tee := io.TeeReader(contents, &buf)

	// Use the existing Senzing reader to parse entities
	entities, err := senzing.ReadEntities(tee, search.SourceOpenSanctionsPEP)
	if err != nil {
		return nil, fmt.Errorf("parsing senzing json: %w", err)
	}

	// Compute hash of the file content
	listHash := sha256.Sum256(buf.Bytes())

	return &Results{
		Entities: entities,
		ListHash: hex.EncodeToString(listHash[:]),
	}, nil
}
