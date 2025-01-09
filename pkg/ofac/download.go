// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	ofacFilenames = []string{
		"ADD.CSV",          // Address
		"ALT.CSV",          // Alternate ID
		"SDN.CSV",          // Specially Designated National
		"SDN_COMMENTS.CSV", // Specially Designated National Comments
	}

	ofacURLTemplate = func() string {
		if v := os.Getenv("OFAC_DOWNLOAD_TEMPLATE"); v != "" {
			return v
		}
		return "https://sanctionslistservice.ofac.treas.gov/api/PublicationPreview/exports/%s"
	}()
)

func Download(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, download.HTTPClient)

	addrs := make(map[string]string)
	for i := range ofacFilenames {
		addrs[ofacFilenames[i]] = fmt.Sprintf(ofacURLTemplate, ofacFilenames[i])
	}

	return dl.GetFiles(ctx, initialDir, addrs)
}
