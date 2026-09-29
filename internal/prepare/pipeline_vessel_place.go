// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import "strings"

// StripTrailingVesselPlace removes a trailing port, flag, or place from a raw
// vessel name: a final parenthetical that is not "(ex-…)", then a final comma
// clause. "MV SIAM ORCHID 7, Bangkok" becomes "MV SIAM ORCHID 7".
func StripTrailingVesselPlace(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	for {
		trimmed, ok := stripTrailingParen(s)
		if !ok {
			break
		}
		s = trimmed
	}

	if i := strings.LastIndex(s, ","); i > 0 {
		head := strings.TrimSpace(s[:i])
		if head != "" {
			s = head
		}
	}
	return s
}

func stripTrailingParen(s string) (string, bool) {
	s = strings.TrimSpace(s)
	if !strings.HasSuffix(s, ")") {
		return s, false
	}
	start := strings.LastIndex(s, "(")
	if start < 0 {
		return s, false
	}
	inner := strings.TrimSpace(s[start+1 : len(s)-1])
	if strings.HasPrefix(strings.ToLower(inner), "ex") {
		return s, false
	}
	return strings.TrimSpace(s[:start]), true
}
