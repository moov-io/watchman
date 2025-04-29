// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"context"
	"fmt"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	ofacFilenames = []string{
		"add.csv",          // Address
		"alt.csv",          // Alternate ID
		"sdn.csv",          // Specially Designated National
		"sdn_comments.csv", // Specially Designated National Comments
	}

	ofacURLTemplate = func() string {
		if v := os.Getenv("OFAC_DOWNLOAD_TEMPLATE"); v != "" {
			return v
		}
		return "https://sanctionslistservice.ofac.treas.gov/api/PublicationPreview/exports/%s"
	}()
)

func Download(ctx context.Context, logger log.Logger, initialDir string) (download.Files, error) {
	dl := download.New(logger, nil)

	addrs := make(map[string]string)
	for i := range ofacFilenames {
		addrs[ofacFilenames[i]] = fmt.Sprintf(ofacURLTemplate, ofacFilenames[i])
	}

	return dl.GetFiles(ctx, initialDir, addrs)
}
