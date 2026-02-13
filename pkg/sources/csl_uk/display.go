// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

// DetailsURL returns the UK government sanctions list page.
// The UK sanctions list doesn't support direct entity linking,
// so we return the main search page.
func DetailsURL(sourceID string) string {
	return "https://www.gov.uk/government/publications/the-uk-sanctions-list"
}
