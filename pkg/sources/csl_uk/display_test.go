// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetailsURL(t *testing.T) {
	// Should return the UK government sanctions page regardless of sourceID
	url := DetailsURL("UK-123")
	require.Equal(t, "https://www.gov.uk/government/publications/the-uk-sanctions-list", url)

	// Should work with empty sourceID
	url = DetailsURL("")
	require.Equal(t, "https://www.gov.uk/government/publications/the-uk-sanctions-list", url)
}
