// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"context"
	"io"

	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	//publicUNSanctionsListURL is the URL to download the UN CSL from. It can be overridden by setting the UN_SANCTIONS_LIST_URL env var.
	publicUNSanctionsListURL = "https://scsanctions.un.org/resources/xml/en/consolidated.xml"

	unSanctionsListURL = strx.Or(os.Getenv("UN_SANCTIONS_LIST_URL"), publicUNSanctionsListURL)
)

func DownloadSanctionsList_UN(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, nil)

	logger.Info().Logf("downloading UN sanctions list from %s", unSanctionsListURL)

	unSanctionsNameAndSource := map[string]string{
		"UN_Sanctions_List.xml": unSanctionsListURL,
	}

	return dl.GetFiles(ctx, initialDir, unSanctionsNameAndSource)
}
