// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package dpl

import (
	"fmt"
	"io"
	"os"

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
func Download(logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, download.HTTPClient)

	addrs := make(map[string]string)
	addrs["dpl.txt"] = fmt.Sprintf(dplDownloadTemplate, "dpl.txt")

	return dl.GetFiles(initialDir, addrs)
}
