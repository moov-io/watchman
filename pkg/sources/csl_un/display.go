// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

// DetailsURL returns the UN government sanctions list page.
//
//	The UN does not provide a unique URL for each entity,
//
// so we return the main
func DetailsURL(sourceID string) string {
	return "https://www.un.org/securitycouncil/sanctions/information"
}
