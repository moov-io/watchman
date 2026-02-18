// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package download

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/csl_eu"
)

func loadEUCSLRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-eu-csl-records")
	defer span.End()

	start := time.Now()

	// Download EU CSL CSV file
	files, err := csl_eu.DownloadEU(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("EU CSL download: %w", err)
	}

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d EU CSL files found", len(files))
	}

	logger.Debug().Logf("finished EU CSL download: %v", time.Since(start))
	start = time.Now()

	// Parse CSV file and compute hash
	var records []csl_eu.CSLRecord
	var hashData bytes.Buffer

	for filename, fd := range files {
		// Read file content for hashing
		content, err := io.ReadAll(fd)
		if closeErr := fd.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			return fmt.Errorf("reading EU CSL %s: %w", filename, err)
		}
		hashData.Write(content)

		// Parse from buffer
		parsed, _, err := csl_eu.ParseEU(io.NopCloser(bytes.NewReader(content)))
		if err != nil {
			return fmt.Errorf("parsing EU CSL %s: %w", filename, err)
		}
		records = append(records, parsed...)
	}

	// Calculate hash
	h := sha256.Sum256(hashData.Bytes())
	listHash := hex.EncodeToString(h[:])

	span.AddEvent("finished parsing")

	// Convert to search entities
	entities := csl_eu.ConvertEUCSLData(records)

	logger.Debug().Logf("finished EU CSL preparation: %d entities in %v", len(entities), time.Since(start))
	span.AddEvent("finished EU CSL preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from EU CSL")
	}

	responseCh <- preparedList{
		ListName: search.SourceEUCSL,
		Entities: entities,
		Hash:     listHash,
	}

	return nil
}
