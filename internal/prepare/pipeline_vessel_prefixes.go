// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import "strings"

// vesselNamePrefixes are leading tokens that remain after LowerAndRemovePunctuation
// turns "M/V", "M.T.", "N/M", and similar into spaced letters.
//
// Longer prefixes are listed first so "m v " matches before "mv " is irrelevant,
// and "msv " wins over "ms ".
var vesselNamePrefixes = []string{
	"msv ",
	"m v ",
	"m t ",
	"n m ",
	"t b ",
	"mv ",
	"mt ",
	"ss ",
	"ms ",
	"tb ",
	"nm ",
}

// StripVesselNamePrefixes removes leading ship-type markers from a prepared
// (already lowercased, punctuation-stripped) vessel name.
func StripVesselNamePrefixes(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	changed := true
	for changed {
		changed = false
		for _, prefix := range vesselNamePrefixes {
			if strings.HasPrefix(s, prefix) {
				s = strings.TrimSpace(s[len(prefix):])
				changed = true
				break
			}
		}
	}
	return s
}
