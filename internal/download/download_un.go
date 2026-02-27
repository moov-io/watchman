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
	"github.com/moov-io/watchman/pkg/sources/csl_un"
)

func loadUNCSLRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-un-csl-records")
	defer span.End()

	start := time.Now()

	// Download UN Sanctions List (CSV format)
	files, err := csl_un.DownloadSanctionsList_UN(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("UN CSL download: %w", err)
	}

	span.AddEvent("finished downloading UN LISt")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d UN CSL files found", len(files))
	}

	logger.Debug().Logf("finished UNCSL download: %v", time.Since(start))
	start = time.Now()

	var entities []search.Entity[search.Value]
	var hashData bytes.Buffer

	for filename, fd := range files {
		// Create a TeeReader so we can calculate the hash while streaming the XML
		hasher := sha256.New()
		tee := io.TeeReader(fd, hasher)

		// Use the optimized Reader to process records
		reader := csl_un.NewReader(tee)
		err := reader.Read(
			func(p csl_un.UNIndividual) {
				entities = append(entities, p.ToEntity())

			},
			func(e csl_un.UNEntity) {
				entities = append(entities, e.ToEntity())
			},
		)

		if closeErr := fd.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			return fmt.Errorf("parsing UN list %s: %w", filename, err)
		}

		// Collect hash from the hasher
		hashData.Write(hasher.Sum(nil))
	}

	// Final list hash calculation
	h := sha256.Sum256(hashData.Bytes())
	listHash := hex.EncodeToString(h[:])

	span.AddEvent("finished parsing")

	logger.Debug().Logf("finished UN CSL preparation: %d entities in %v", len(entities), time.Since(start))

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from UN list")
	}

	// Send to the central preparation channel
	responseCh <- preparedList{
		ListName: search.SourceUNCSL, // Ensure this constant matches your search package
		Entities: entities,
		Hash:     listHash,
	}

	return nil
}
