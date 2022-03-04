// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package dpl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	dplDownloadTemplate = func() string {
		if w := os.Getenv("DPL_DOWNLOAD_TEMPLATE"); w != "" {
			return w
		}
		return "https://www.bis.doc.gov/dpl/%s" // Denied Persons List (tab separated)
	}()
)

// Download returns an array of absolute filepaths for files downloaded
func Download(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	addrs := make(map[string]string)
	addrs["dpl.txt"] = fmt.Sprintf(dplDownloadTemplate, "dpl.txt")

	files, err := dl.GetFiles(initialDir, addrs)
	if len(files) == 0 || err != nil {
		return "", fmt.Errorf("dpl download: %v", err)
	}
	for i := range files {
		if filepath.Base(files[i]) == "dpl.txt" {
			return files[i], nil
		}
	}
	return "", nil
}
