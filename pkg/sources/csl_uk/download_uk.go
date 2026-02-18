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
	// Consolidated List CSV from https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets
	publicCSLDownloadURL = "https://ofsistorage.blob.core.windows.net/publishlive/2022format/ConList.csv"
	ukCSLDownloadURL     = strx.Or(os.Getenv("UK_CSL_DOWNLOAD_URL"), publicCSLDownloadURL)

	// UK Sanctions List CSV from https://sanctionslist.fcdo.gov.uk/
	publicUKSanctionsListURL = "https://sanctionslist.fcdo.gov.uk/docs/UK-Sanctions-List.csv"
	ukSanctionsListURL       = strx.Or(os.Getenv("UK_SANCTIONS_LIST_URL"), publicUKSanctionsListURL)
)

func DownloadCSL(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, nil)

	ukCSLNameAndSource := make(map[string]string)
	ukCSLNameAndSource["ConList.csv"] = ukCSLDownloadURL

	return dl.GetFiles(ctx, initialDir, ukCSLNameAndSource)
}

func DownloadSanctionsList(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, nil)

	logger.Info().Logf("downloading UK sanctions list from %s", ukSanctionsListURL)

	ukSanctionsNameAndSource := map[string]string{
		"UK_Sanctions_List.csv": ukSanctionsListURL,
	}

	return dl.GetFiles(ctx, initialDir, ukSanctionsNameAndSource)
}
