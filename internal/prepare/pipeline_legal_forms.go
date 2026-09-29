// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import "strings"

// englishLegalFormPhrases maps spelled-out English legal forms onto the short
// token already used in list names (inc, ltd, llc, …). Longer phrases are listed
// first so "limited liability company" is not reduced to "ltd liability co".
//
// The input is already lowercased with punctuation turned into spaces.
var englishLegalFormPhrases = []struct{ from, to string }{
	{"limited liability company", "llc"},
	{"public limited company", "plc"},
	{"incorporated", "inc"},
	{"corporation", "corp"},
	{"company", "co"},
	{"limited", "ltd"},
}

// spacedLegalFormTokens are short forms that punctuation stripping can split
// into single letters (L.L.C. → "l l c", P.L.C. → "p l c").
var spacedLegalFormTokens = []string{"llc", "plc"}

// CanonicalizeEnglishLegalForms rewrites English legal-form phrases in a
// prepared business name. "Harrowfield Bearings Limited" and
// "HARROWFIELD BEARINGS LTD." both become "harrowfield bearings ltd".
func CanonicalizeEnglishLegalForms(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	padded := " " + s + " "
	for _, p := range englishLegalFormPhrases {
		padded = strings.ReplaceAll(padded, " "+p.from+" ", " "+p.to+" ")
	}
	for _, short := range spacedLegalFormTokens {
		spaced := " " + strings.Join(strings.Split(short, ""), " ") + " "
		padded = strings.ReplaceAll(padded, spaced, " "+short+" ")
	}
	return strings.TrimSpace(padded)
}
