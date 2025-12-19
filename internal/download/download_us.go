package download

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/csl_us"
	"github.com/moov-io/watchman/pkg/sources/ofac"
	"github.com/moov-io/watchman/pkg/sources/us_non_sdn"

	"github.com/moov-io/base/log"
)

func loadOFACRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-ofac-records")
	defer span.End()

	start := time.Now()
	files, err := ofac.Download(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("OFAC download: %v", err)
	}
	defer files.Close()

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d OFAC files found", len(files))
	}

	logger.Debug().Logf("finished OFAC download: %v", time.Since(start))
	start = time.Now()

	res, err := ofac.Read(files)
	if err != nil {
		return fmt.Errorf("parsing OFAC: %w", err)
	}
	span.AddEvent("finished parsing")

	entities := ofac.GroupIntoEntities(res.SDNs, res.Addresses, res.SDNComments, res.AlternateIdentities)

	logger.Debug().Logf("finished OFAC preparation: %v", time.Since(start))
	span.AddEvent("finished OFAC preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from US OFAC")
	}

	responseCh <- preparedList{
		ListName: search.SourceUSOFAC,
		Entities: entities,
		Hash:     res.ListHash,
	}
	return nil
}

func loadUSNonSDNRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-us-non-sdn-records")
	defer span.End()

	start := time.Now()
	files, err := us_non_sdn.Download(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("US Non-SDN download: %v", err)
	}
	defer files.Close()

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d US Non-SDN files found", len(files))
	}

	logger.Debug().Logf("finished US Non-SDN download: %v", time.Since(start))
	start = time.Now()

	// The US Non-SDN list downloads OFAC-compatible files
	res, err := ofac.Read(files)
	if err != nil {
		return fmt.Errorf("parsing US Non-SDN: %w", err)
	}
	span.AddEvent("finished parsing")

	entities := ofac.GroupIntoEntities(res.SDNs, res.Addresses, res.SDNComments, res.AlternateIdentities,
		ofac.WithSourceList(search.SourceUSNonSDN),
	)

	logger.Debug().Logf("finished US Non-SDN preparation: %v", time.Since(start))
	span.AddEvent("finished US Non-SDN preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from US Non-SDN")
	}

	responseCh <- preparedList{
		ListName: search.SourceUSNonSDN,
		Entities: entities,
		Hash:     res.ListHash,
	}
	return nil
}

func loadCSLUSRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-us-csl-records")
	defer span.End()

	start := time.Now()
	files, err := csl_us.Download(ctx, logger, initialDataDirectory(conf))
	if err != nil {
		return fmt.Errorf("US CSL download: %w", err)
	}
	defer files.Close()

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d US CSL files found", len(files))
	}

	logger.Debug().Logf("finished US CSL download: %v", time.Since(start))
	start = time.Now()

	res, err := csl_us.Read(files)
	if err != nil {
		return fmt.Errorf("parsing US CSL: %w", err)
	}
	span.AddEvent("finished parsing")

	entities := csl_us.ConvertSanctionsData(res)

	logger.Debug().Logf("finished US CSL preparation: %v", time.Since(start))
	span.AddEvent("finished US CSL preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from US CSL")
	}

	responseCh <- preparedList{
		ListName: search.SourceUSCSL,
		Entities: entities,
		Hash:     res.ListHash,
	}

	return nil
}
