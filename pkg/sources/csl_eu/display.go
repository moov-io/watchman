// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

import (
	"github.com/moov-io/watchman/pkg/search"
)

const (
	defaultEUCSLURL = "https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions"
)

// DetailsURL returns the EU sanctions publication URL for an entity.
// Extracts EntityPublicationURL from SourceData if available, otherwise returns default.
func DetailsURL(entity search.Entity[search.Value]) string {
	if m, ok := entity.SourceData.(map[string]interface{}); ok {
		if url, ok := m["entityPublicationUrl"].(string); ok && url != "" {
			return url
		}
	}

	if record, ok := entity.SourceData.(CSLRecord); ok {
		if record.EntityPublicationURL != "" {
			return record.EntityPublicationURL
		}
	}

	return defaultEUCSLURL
}
