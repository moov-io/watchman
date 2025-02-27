// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"

	"github.com/antchfx/htmlquery"
)

var (
	// taken from https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents
	publicCSLDownloadURL = "https://ofsistorage.blob.core.windows.net/publishlive/2022format/ConList.csv"
	ukCSLDownloadURL     = strx.Or(os.Getenv("UK_CSL_DOWNLOAD_URL"), publicCSLDownloadURL)
)

func DownloadCSL(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, download.HTTPClient)

	ukCSLNameAndSource := make(map[string]string)
	ukCSLNameAndSource["ConList.csv"] = ukCSLDownloadURL

	return dl.GetFiles(ctx, initialDir, ukCSLNameAndSource)
}

func DownloadSanctionsList(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, download.HTTPClient)

	ukSanctionsNameAndSource := make(map[string]string)

	latestURL, err := fetchLatestUKSanctionsListURL(ctx, logger, initialDir)
	if err != nil {
		return nil, err
	}
	logger.Info().Logf("downloading UK sanctions from %s", latestURL)

	ukSanctionsNameAndSource["UK_Sanctions_List.ods"] = latestURL

	return dl.GetFiles(ctx, initialDir, ukSanctionsNameAndSource)
}

var (
	defaultUKSanctionsListHTML = strx.Or(os.Getenv("UK_CSL_HTML_INDEX_URL"), "https://www.gov.uk/government/publications/the-uk-sanctions-list")
)

func fetchLatestUKSanctionsListURL(ctx context.Context, logger log.Logger, initialDir string) (string, error) {
	fromEnv := strings.TrimSpace(os.Getenv("UK_SANCTIONS_LIST_URL"))
	if fromEnv != "" {
		return fromEnv, nil
	}

	// Fetch the HTML page and look for the latest link
	ukSanctionsNameAndSource := make(map[string]string)
	ukSanctionsNameAndSource["UK_Sanctions_List.ods"] = defaultUKSanctionsListHTML

	dl := download.New(logger, download.HTTPClient)

	pages, err := dl.GetFiles(ctx, initialDir, ukSanctionsNameAndSource)
	if err != nil {
		return "", fmt.Errorf("getting UK Sanctions html index: %w", err)
	}

	indexContents, exists := pages["UK_Sanctions_List.ods"]
	if !exists {
		return "", fmt.Errorf("UK sanctions index page %s not found", defaultUKSanctionsListHTML)
	}

	index, err := htmlquery.Parse(indexContents)
	if err != nil {
		return "", fmt.Errorf("parsing UK sanctions index page: %w", err)
	}

	links, err := htmlquery.QueryAll(index, `//a[contains(@class, 'govuk-link') and contains(@href, '.ods')]`)
	if err != nil {
		return "", fmt.Errorf("html xpath failed: %w", err)
	}

	for _, link := range links {
		for _, attr := range link.Attr {
			if attr.Key == "href" && strings.HasSuffix(attr.Val, ".ods") {
				return attr.Val, nil
			}
		}
	}

	return "", nil
}
