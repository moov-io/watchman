// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	publicUSDownloadURL = "https://data.trade.gov/downloadable_consolidated_screening_list/v1/%s"
	usDownloadURL       = strx.Or(os.Getenv("US_CSL_DOWNLOAD_TEMPLATE"), os.Getenv("US_CSL_DOWNLOAD_URL"), publicUSDownloadURL)
)

func Download(ctx context.Context, logger log.Logger, initialDir string) (download.Files, error) {
	dl := download.New(logger, nil)

	where, err := url.Parse(fmt.Sprintf(usDownloadURL, "consolidated.csv"))
	if err != nil {
		return nil, fmt.Errorf("building url: %w", err)
	}

	addrs := make(map[string]string)
	addrs["consolidated.csv"] = where.String()

	return dl.GetFiles(ctx, initialDir, addrs)
}
