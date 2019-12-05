// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/moov-io/watchman/pkg/download"
	"net/url"
	"os"
	"strings"
)

var (
	cslDownloadTemplate = func() string {
		if w := os.Getenv("CSL_DOWNLOAD_TEMPLATE"); w != "" {
			return w
		}
		return "https://api.trade.gov/consolidated_screening_list/%s"
	}()
)

func Download(logger log.Logger, initialDir string) (string, error) {
	dl := download.New(logger, download.HTTPClient)

	cslURL, err := buildDownloadURL(cslDownloadTemplate)
	if err != nil {
		return "", err
	}

	cslNameAndSource := make(map[string]string)
	cslNameAndSource["csl.csv"] = cslURL

	file, err := dl.GetFiles(initialDir, cslNameAndSource)
	if len(file) == 0 || err != nil {
		return "", fmt.Errorf("csl download: %v", err)
	}
	return file[0], nil
}

func buildDownloadURL(urlStr string) (string, error) {
	cslURL, err := url.Parse(fmt.Sprintf(urlStr, "search.csv"))
	if err != nil {
		return "", err
	}

	if strings.EqualFold(cslURL.Host, "api.trade.gov") { // only require api key if source is api.trade.gov
		q := cslURL.Query()
		keyOverride := os.Getenv("TRADEGOV_API_KEY")
		if keyOverride == "" && q.Get("api_key") == "" {
			return "", fmt.Errorf("csl download: missing api key")
		}

		if q.Get("api_key") == "" { // download template did not include api key
			q.Set("api_key", keyOverride)
		}

		cslURL.RawQuery = q.Encode()
	}

	return cslURL.String(), nil
}
