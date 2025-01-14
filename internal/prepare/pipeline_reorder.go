// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"fmt"
	"regexp"
	"strings"
)

// This pattern looks for:
//
//	a comma, optional whitespace, then
//	1+ Unicode letters and diacritics (\p{L}\p{M}), plus allowed punctuation/apostrophes/hyphens/spaces
//	until the end of the string.
//
// Examples that should match (and reorder):
//
//	"AL-ZAYDI, Shibl Muhsin 'Ubayd" --> "Shibl Muhsin 'Ubayd AL-ZAYDI"
//	"MADURO MOROS, Nicolas"         --> "Nicolas MADURO MOROS"
var (
	surnamePrecedes = regexp.MustCompile(`,(?:\s+)?([\p{L}\p{M}'’\-\.\s]+)$`)
)

func ReorderSDNNames(names []string, sdnType string) []string {
	if len(names) == 0 {
		return nil
	}

	out := make([]string, len(names))
	for idx := range names {
		out[idx] = ReorderSDNName(names[idx], sdnType)
	}
	return out
}

// ReorderSDNName will take a given SDN name and, if it matches "Surname, FirstName(s)",
// reorder it to "FirstName(s) Surname" (only for type == "individual").
func ReorderSDNName(name, sdnType string) string {
	// Only reorder for individuals
	if !strings.EqualFold(sdnType, "individual") {
		return name
	}

	// Try matching the pattern
	match := surnamePrecedes.FindStringSubmatch(name)
	if len(match) < 2 {
		// No match => no reordering
		return name
	}

	// match[1] is the part after the comma (the "first/given names" portion).
	givenNames := strings.TrimSpace(match[1])

	// match[0] is the entire substring matching the pattern, including the comma;
	// remove it from the original to isolate the "surname" part.
	surname := strings.TrimSuffix(name, match[0])
	surname = strings.TrimSpace(surname)

	// Rebuild as "GivenName(s) Surname"
	out := fmt.Sprintf("%s %s", givenNames, surname)
	return strings.TrimSpace(out)
}
