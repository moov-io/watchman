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
	"github.com/moov-io/watchman/pkg/sources/csl_uk"
)

func loadUKCSLRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-uk-csl-records")
	defer span.End()

	start := time.Now()

	// Download UK Sanctions List (ODS format)
	files, err := csl_uk.DownloadSanctionsList(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("UK CSL download: %w", err)
	}

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d UK CSL files found", len(files))
	}

	logger.Debug().Logf("finished UK CSL download: %v", time.Since(start))
	start = time.Now()

	// Parse ODS file and compute hash
	var records []csl_uk.SanctionsListRecord
	var hashData bytes.Buffer

	for filename, fd := range files {
		// Read file content for hashing
		content, err := io.ReadAll(fd)
		if closeErr := fd.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			return fmt.Errorf("reading UK CSL %s: %w", filename, err)
		}
		hashData.Write(content)

		// Parse from buffer
		parsed, _, err := csl_uk.ReadSanctionsListFile(io.NopCloser(bytes.NewReader(content)))
		if err != nil {
			return fmt.Errorf("parsing UK CSL %s: %w", filename, err)
		}
		records = append(records, parsed...)
	}

	// Calculate hash
	h := sha256.Sum256(hashData.Bytes())
	listHash := hex.EncodeToString(h[:])

	span.AddEvent("finished parsing")

	// Convert to search entities
	entities := csl_uk.ConvertSanctionsListData(records)

	logger.Debug().Logf("finished UK CSL preparation: %d entities in %v", len(entities), time.Since(start))
	span.AddEvent("finished UK CSL preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from UK CSL")
	}

	responseCh <- preparedList{
		ListName: search.SourceUKCSL,
		Entities: entities,
		Hash:     listHash,
	}

	return nil
}
