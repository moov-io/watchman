// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package fincen_311

import (
	"context"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	publicFinCEN311URL = "https://www.fincen.gov/resources/statutes-and-regulations/311-and-9714-special-measures"
	fincen311URL       = strx.Or(os.Getenv("FINCEN_311_DOWNLOAD_URL"), publicFinCEN311URL)
)

// Download fetches the FinCEN 311/9714 Special Measures HTML page
func Download(ctx context.Context, logger log.Logger, initialDir string) (download.Files, error) {
	dl := download.New(logger, nil)

	addrs := make(map[string]string)
	addrs["fincen_311.html"] = fincen311URL

	return dl.GetFiles(ctx, initialDir, addrs)
}
