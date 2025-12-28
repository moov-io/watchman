// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package download

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/opensanctions"
)

func loadOpenSanctionsPEPRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-opensanctions-pep-records")
	defer span.End()

	start := time.Now()
	files, err := opensanctions.Download(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("OpenSanctions PEP download: %w", err)
	}
	defer files.Close()

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d OpenSanctions PEP files found", len(files))
	}

	logger.Debug().Logf("finished OpenSanctions PEP download: %v", time.Since(start))
	start = time.Now()

	res, err := opensanctions.Read(files)
	if err != nil {
		return fmt.Errorf("parsing OpenSanctions PEP: %w", err)
	}
	span.AddEvent("finished parsing")

	logger.Debug().Logf("finished OpenSanctions PEP preparation: %v", time.Since(start))
	span.AddEvent("finished OpenSanctions PEP preparation")

	if len(res.Entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from OpenSanctions PEP")
	}

	responseCh <- preparedList{
		ListName: search.SourceOpenSanctionsPEP,
		Entities: res.Entities,
		Hash:     res.ListHash,
	}
	return nil
}
