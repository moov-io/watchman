// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import "strings"

// partyLabelPrefixes are ASCII leaders copied from trade documents (SWIFT,
// customs, letters of credit). Longer prefixes are listed first.
var partyLabelPrefixes = []string{
	"ordering customer:",
	"shipper:",
	"beneficiary:",
	"field 59:",
	"messrs.",
	"messrs ",
}

// StripLeadingPartyLabels removes a leading role marker from a raw name.
// "SHIPPER: TRANSPORTES REFRIGERADOS DEL SURESTE SA DE CV" becomes
// "TRANSPORTES REFRIGERADOS DEL SURESTE SA DE CV".
func StripLeadingPartyLabels(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	lower := strings.ToLower(s)
	for _, prefix := range partyLabelPrefixes {
		if strings.HasPrefix(lower, prefix) {
			return strings.TrimSpace(s[len(prefix):])
		}
	}
	return s
}
