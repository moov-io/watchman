// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package download

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/fincen_311"
)

func loadFinCEN311Records(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-fincen-311-records")
	defer span.End()

	start := time.Now()
	files, err := fincen_311.Download(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("FinCEN 311 download: %v", err)
	}
	defer files.Close()

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d FinCEN 311 files found", len(files))
	}

	logger.Debug().Logf("finished FinCEN 311 download: %v", time.Since(start))
	start = time.Now()

	res, err := fincen_311.Read(files)
	if err != nil {
		return fmt.Errorf("parsing FinCEN 311: %w", err)
	}
	span.AddEvent("finished parsing")

	entities := fincen_311.ConvertSpecialMeasures(res)

	logger.Debug().Logf("finished FinCEN 311 preparation: %v", time.Since(start))
	span.AddEvent("finished FinCEN 311 preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from FinCEN 311")
	}

	responseCh <- preparedList{
		ListName: search.SourceUSFinCEN311,
		Entities: entities,
		Hash:     res.ListHash,
	}
	return nil
}
