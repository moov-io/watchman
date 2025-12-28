// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package opensanctions

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	// OpenSanctions PEP dataset in Senzing format
	// https://www.opensanctions.org/datasets/peps/
	publicDownloadURL = "https://data.opensanctions.org/datasets/latest/peps/senzing.json"
	downloadURL       = strx.Or(os.Getenv("OPENSANCTIONS_DOWNLOAD_URL"), publicDownloadURL)

	// API key for authenticated access
	// https://www.opensanctions.org/docs/api/
	apiKey = os.Getenv("OPENSANCTIONS_API_KEY")
)

func Download(ctx context.Context, logger log.Logger, initialDir string) (download.Files, error) {
	dl := download.New(logger, nil)

	targetURL := downloadURL
	if apiKey != "" {
		// Use API with authentication
		parsedURL, err := url.Parse(downloadURL)
		if err != nil {
			return nil, fmt.Errorf("unable to parse OpenSanctions download URL: %w", err)
		}
		q := parsedURL.Query()
		q.Set("api_key", apiKey)
		parsedURL.RawQuery = q.Encode()
		targetURL = parsedURL.String()

		logger.Info().Log("using OpenSanctions API with authentication")
	}

	addrs := make(map[string]string)
	addrs["peps_senzing.json"] = targetURL

	return dl.GetFiles(ctx, initialDir, addrs)
}
