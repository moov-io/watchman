// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"context"
	"io"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	// UK Sanctions List CSV from https://www.gov.uk/government/publications/the-uk-sanctions-list
	publicUKSanctionsListURL = "https://sanctionslist.fcdo.gov.uk/docs/UK-Sanctions-List.csv"
	ukSanctionsListURL       = strx.Or(os.Getenv("UK_SANCTIONS_LIST_URL"), publicUKSanctionsListURL)
)

func DownloadSanctionsList(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, nil)

	logger.Info().Logf("downloading UK sanctions list from %s", ukSanctionsListURL)

	ukSanctionsNameAndSource := map[string]string{
		"UK_Sanctions_List.csv": ukSanctionsListURL,
	}

	return dl.GetFiles(ctx, initialDir, ukSanctionsNameAndSource)
}
