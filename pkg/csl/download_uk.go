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
	// taken from https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents
	publicUKCSLDownloadURL = "https://ofsistorage.blob.core.windows.net/publishlive/2022format/ConList.csv"
	ukCSLDownloadURL       = strx.Or(os.Getenv("UK_CSL_DOWNLOAD_URL"), publicUKCSLDownloadURL)

	// https://www.gov.uk/government/publications/the-uk-sanctions-list
	publicUKSanctionsListURL = "https://assets.publishing.service.gov.uk/government/uploads/system/uploads/attachment_data/file/1121113/UK_Sanctions_List.ods"
	ukSanctionsListURL       = strx.Or(os.Getenv("UK_SANCTIONS_LIST_URL"), publicUKSanctionsListURL)
)

func DownloadUKCSL(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	ukCSLNameAndSource := make(map[string]string)
	ukCSLNameAndSource["ConList.csv"] = ukCSLDownloadURL

	file, err := dl.GetFiles(initialDir, ukCSLNameAndSource)
	if len(file) == 0 || err != nil {
		return "", fmt.Errorf("uk csl download: %v", err)
	}
	return file[0], nil
}

func DownloadUKSanctionsList(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	ukSanctionsNameAndSource := make(map[string]string)
	ukSanctionsNameAndSource["UK_Sanctions_List.ods"] = ukSanctionsListURL

	file, err := dl.GetFiles(initialDir, ukSanctionsNameAndSource)
	if len(file) == 0 || err != nil {
		return "", fmt.Errorf("uk download: %v", err)
	}
	return file[0], nil
}
