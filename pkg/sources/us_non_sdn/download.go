// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package us_non_sdn

import (
	"context"
	"fmt"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	filenames = []string{
		"CONS_PRIM.CSV",     // Names
		"CONS_ADD.CSV",      // Addresses
		"CONS_ALT.CSV",      // Alternate Names
		"CONS_COMMENTS.CSV", // Comments
	}

	urlTemplate = func() string {
		if v := os.Getenv("US_NON_SDN_DOWNLOAD_TEMPLATE"); v != "" {
			return v
		}
		return "https://sanctionslistservice.ofac.treas.gov/api/PublicationPreview/exports/%s"
	}()
)

func Download(ctx context.Context, logger log.Logger, initialDir string) (download.Files, error) {
	dl := download.New(logger, nil)

	addrs := make(map[string]string)
	for _, name := range filenames {
		addrs[name] = fmt.Sprintf(urlTemplate, name)
	}

	return dl.GetFiles(ctx, initialDir, addrs)
}
