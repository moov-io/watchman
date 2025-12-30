// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package fincen_311

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/moov-io/watchman/pkg/download"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// Read parses the downloaded HTML file and extracts special measures
func Read(files download.Files) (*ListData, error) {
	for filename, contents := range files {
		if strings.Contains(strings.ToLower(filename), "fincen_311") {
			return parseHTMLContents(contents)
		}
	}
	return nil, errors.New("no FinCEN 311 file found")
}

func parseHTMLContents(contents io.ReadCloser) (*ListData, error) {
	var buf bytes.Buffer
	buftee := io.TeeReader(contents, &buf)
	defer contents.Close()

	doc, err := htmlquery.Parse(buftee)
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %w", err)
	}

	measures, err := extractSpecialMeasures(doc)
	if err != nil {
		return nil, fmt.Errorf("extracting special measures: %w", err)
	}

	// Compute hash for change detection
	listHash := sha256.Sum256(buf.Bytes())

	return &ListData{
		SpecialMeasures: measures,
		ListHash:        hex.EncodeToString(listHash[:]),
	}, nil
}

func extractSpecialMeasures(doc *html.Node) ([]SpecialMeasure, error) {
	var measures []SpecialMeasure

	// Try multiple XPath strategies for robustness
	xpaths := []string{
		`//table//tbody//tr`,
		`//table//tr[position()>1]`,
		`//div[contains(@class, 'table')]//tr`,
		`//article//table//tr`,
	}

	var rows []*html.Node
	var err error
	for _, xpath := range xpaths {
		rows, err = htmlquery.QueryAll(doc, xpath)
		if err == nil && len(rows) > 0 {
			break
		}
	}

	if len(rows) == 0 {
		return nil, errors.New("no table rows found - HTML structure may have changed")
	}

	for _, row := range rows {
		measure, err := parseTableRow(row)
		if err != nil {
			// Skip malformed rows but continue processing
			continue
		}
		// Skip empty names and header rows
		if measure.EntityName != "" &&
			!strings.EqualFold(measure.EntityName, "Entity Name") &&
			!strings.EqualFold(measure.EntityName, "Name") {
			measures = append(measures, measure)
		}
	}

	if len(measures) == 0 {
		return nil, errors.New("no valid entities parsed from HTML")
	}

	return measures, nil
}

func parseTableRow(row *html.Node) (SpecialMeasure, error) {
	var measure SpecialMeasure

	cells, _ := htmlquery.QueryAll(row, `.//td`)
	if len(cells) < 2 {
		return measure, errors.New("insufficient cells in row")
	}

	// Column 0: Entity Name
	measure.EntityName = cleanText(htmlquery.InnerText(cells[0]))

	// Column 1: Finding (may contain link and date)
	if len(cells) > 1 {
		measure.FindingURL, measure.FindingDate = extractLinkAndDate(cells[1])
	}

	// Column 2: NPRM
	if len(cells) > 2 {
		measure.NPRMURL, measure.NPRMDate = extractLinkAndDate(cells[2])
	}

	// Column 3: Final Rule
	if len(cells) > 3 {
		measure.FinalRuleURL, measure.FinalRuleDate = extractLinkAndDate(cells[3])
	}

	// Column 4: Rescinded (if exists)
	if len(cells) > 4 {
		measure.RescindedURL, measure.RescindedDate = extractLinkAndDate(cells[4])
		cellText := strings.ToLower(htmlquery.InnerText(cells[4]))
		measure.IsRescinded = measure.RescindedURL != "" ||
			strings.Contains(cellText, "rescind") ||
			(measure.RescindedDate != "" && measure.RescindedDate != "---")
	}

	// Classify entity type based on name
	measure.EntityType = classifyEntityType(measure.EntityName)

	return measure, nil
}

func extractLinkAndDate(cell *html.Node) (url, date string) {
	// Look for anchor tags
	links, _ := htmlquery.QueryAll(cell, `.//a`)
	for _, link := range links {
		for _, attr := range link.Attr {
			if attr.Key == "href" && url == "" {
				url = attr.Val
				// Make relative URLs absolute
				if strings.HasPrefix(url, "/") {
					url = "https://www.fincen.gov" + url
				}
				break
			}
		}
	}

	// Extract date text
	text := cleanText(htmlquery.InnerText(cell))
	if text != "" && text != "---" && text != "-" {
		date = text
	}

	return url, date
}

func classifyEntityType(name string) SMType {
	lower := strings.ToLower(name)

	// Financial Institution indicators
	bankKeywords := []string{
		"bank", "credit union", "financial institution",
		"fbme", "ablv", "banca privada", "asiauniversalbank",
		"bitzlato", "suex", "garantex", "casa de cambio",
		"pm2btc", "liberty reserve", "delta asia",
		"casa de bolsa", "institución de banca",
	}
	for _, keyword := range bankKeywords {
		if strings.Contains(lower, keyword) {
			return SMTypeFinancialInstitution
		}
	}

	// Jurisdiction indicators (country/region names)
	jurisdictionKeywords := []string{
		"iran", "korea", "burma", "myanmar",
		"lebanon", "venezuela", "republic of", "nauru",
		"ukraine", "people's republic", "democratic republic",
		"islamic republic",
	}
	for _, keyword := range jurisdictionKeywords {
		if strings.Contains(lower, keyword) {
			return SMTypeJurisdiction
		}
	}

	// Transaction class indicators
	transactionKeywords := []string{
		"virtual currency", "mixing", "convertible",
		"transaction", "class of", "gambling",
	}
	for _, keyword := range transactionKeywords {
		if strings.Contains(lower, keyword) {
			return SMTypeTransactionClass
		}
	}

	// Default to financial institution for unknown types
	return SMTypeFinancialInstitution
}

func cleanText(s string) string {
	// Remove extra whitespace and newlines
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	// Collapse multiple spaces
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}
