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
	"github.com/moov-io/watchman/pkg/sources/us_tel"
)

func loadUSTelRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-us-tel-records")
	defer span.End()

	start := time.Now()

	files, err := us_tel.DownloadUsTel(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("US TEL download error: %w", err)
	}

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d US TEL files found", len(files))
	}

	logger.Debug().Logf("finished US TEL download: %v", time.Since(start))
	start = time.Now()

	var entities []search.Entity[search.Value]
	var hashData bytes.Buffer

	for filename, fd := range files {
		content, err := io.ReadAll(fd)
		if closeErr := fd.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			return fmt.Errorf("reading US TEL %s: %w", filename, err)
		}
		hashData.Write(content)

		// Parse from buffer
		reader := us_tel.NewReader(bytes.NewReader(content))
		err = reader.Read(func(record us_tel.TELRecord) {
			entities = append(entities, record.ToEntity())
		})
		if err != nil {
			return fmt.Errorf("parsing US TEL %s: %w", filename, err)
		}
	}

	h := sha256.Sum256(hashData.Bytes())
	listHash := hex.EncodeToString(h[:])

	span.AddEvent("finished parsing")

	logger.Debug().Logf("finished US TEL preparation: %d entities in %v", len(entities), time.Since(start))
	span.AddEvent("finished US TEL preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from US TEL")
	}

	responseCh <- preparedList{
		ListName: search.SourceUSTEL,
		Entities: entities,
		Hash:     listHash,
	}

	return nil
}
