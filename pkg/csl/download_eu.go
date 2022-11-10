// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"fmt"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

// TODO: where can I get a new token from?
const uri = "https://webgate.ec.europa.eu/fsd/fsf/public/files/csvFullSanctionsList_1_1/content?token=dG9rZW4tMjAxNw"

func DownloadEU(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	euCSLNameAndSource := make(map[string]string)
	euCSLNameAndSource["eu_csl.csv"] = uri

	file, err := dl.GetFiles(initialDir, euCSLNameAndSource)
	if len(file) == 0 || err != nil {
		return "", fmt.Errorf("eu csl download: %v", err)
	}
	return file[0], nil
}
