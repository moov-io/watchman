// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"fmt"
	"net/url"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

const uri = "https://webgate.ec.europa.eu/fsd/fsf/public/files/csvFullSanctionsList_1_1/content?token=dG9rZW4tMjAxNw"

var (
	euCSLDownloadTemplate = func() string {
		if w := os.Getenv("EU_CSL_DOWNLOAD_TEMPLATE"); w != "" {
			return w
		}
		return uri
	}()
)

func DownloadEU(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	euCSLURL, err := buildEUDownloadURL(euCSLDownloadTemplate)
	if err != nil {
		return "", err
	}

	fmt.Println("euCSLURL: ", euCSLURL)

	euCSLNameAndSource := make(map[string]string)
	euCSLNameAndSource["eu_csl.csv"] = euCSLURL

	file, err := dl.GetFiles(initialDir, euCSLNameAndSource)
	if len(file) == 0 || err != nil {
		return "", fmt.Errorf("eu csl download: %v", err)
	}
	return file[0], nil
}

func buildEUDownloadURL(urlStr string) (string, error) {
	euCSLURL, err := url.Parse(fmt.Sprintf(urlStr, "eu_csl_processed.csv"))
	if err != nil {
		return "", err
	}
	return euCSLURL.String(), nil
}
