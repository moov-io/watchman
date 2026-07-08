// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package us_tel

// DetailsURL returns the US State Terrorist Exclusion List page.
//
//	The US does not provide a unique URL for each entity,
//
// so we return the main
func DetailsURL(sourceID string) string {
	return "https://www.state.gov/terrorist-exclusion-list"
}
