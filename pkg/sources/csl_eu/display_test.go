// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/assert"
)

func TestDetailsURL_WithPublicationURL(t *testing.T) {
	entity := search.Entity[search.Value]{
		Source:   search.SourceEUCSL,
		SourceID: "123",
		SourceData: CSLRecord{
			EntityPublicationURL: "http://example.com/entity/123",
		},
	}

	url := DetailsURL(entity)
	assert.Equal(t, "http://example.com/entity/123", url)
}

func TestDetailsURL_WithoutPublicationURL(t *testing.T) {
	entity := search.Entity[search.Value]{
		Source:   search.SourceEUCSL,
		SourceID: "456",
		SourceData: CSLRecord{
			EntityPublicationURL: "",
		},
	}

	url := DetailsURL(entity)
	assert.Equal(t, defaultEUCSLURL, url)
}

func TestDetailsURL_WithNonCSLRecord(t *testing.T) {
	entity := search.Entity[search.Value]{
		Source:     search.SourceEUCSL,
		SourceID:   "789",
		SourceData: "some other type",
	}

	url := DetailsURL(entity)
	assert.Equal(t, defaultEUCSLURL, url)
}
