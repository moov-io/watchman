// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	publicUSDownloadURL = "https://api.trade.gov/static/consolidated_screening_list/%s"
	usDownloadURL       = strx.Or(os.Getenv("CSL_DOWNLOAD_TEMPLATE"), os.Getenv("US_CSL_DOWNLOAD_URL"), publicUSDownloadURL)
)

func Download(logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, download.HTTPClient)

	cslURL, err := buildDownloadURL(usDownloadURL)
	if err != nil {
		return nil, err
	}

	cslNameAndSource := make(map[string]string)
	cslNameAndSource["csl.csv"] = cslURL

	return dl.GetFiles(initialDir, cslNameAndSource)
}

func buildDownloadURL(urlStr string) (string, error) {
	cslURL, err := url.Parse(fmt.Sprintf(urlStr, "consolidated.csv"))
	if err != nil {
		return "", err
	}
	return cslURL.String(), nil
}
