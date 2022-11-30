// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"fmt"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	// Token is hardcoded on the EU site, but we offer an override.
	// https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en
	token               = strx.Or(os.Getenv("EU_CSL_TOKEN"), "dG9rZW4tMjAxNw")
	publicEUDownloadURL = fmt.Sprintf("https://webgate.ec.europa.eu/fsd/fsf/public/files/csvFullSanctionsList_1_1/content?token=%s", token)

	euDownloadURL = strx.Or(os.Getenv("EU_CSL_DOWNLOAD_URL"), publicEUDownloadURL)
)

func DownloadEU(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	euCSLNameAndSource := make(map[string]string)
	euCSLNameAndSource["eu_csl.csv"] = euDownloadURL

	file, err := dl.GetFiles(initialDir, euCSLNameAndSource)
	if len(file) == 0 || err != nil {
		return "", fmt.Errorf("eu csl download: %v", err)
	}
	return file[0], nil
}
