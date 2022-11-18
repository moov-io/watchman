// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"fmt"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

// taken from https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents
const (
	ukuri = "https://ofsistorage.blob.core.windows.net/publishlive/2022format/ConList.csv"
)

func DownloadUK(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	ukCSLNameAndSource := make(map[string]string)
	ukCSLNameAndSource["ConList.csv"] = ukuri

	file, err := dl.GetFiles(initialDir, ukCSLNameAndSource)
	if len(file) == 0 || err != nil {
		return "", fmt.Errorf("uk csl download: %v", err)
	}
	return file[0], nil
}
