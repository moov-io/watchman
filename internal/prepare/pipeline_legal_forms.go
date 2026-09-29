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
	{"l l c", "llc"},
	{"p l c", "plc"},
}

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
	return strings.TrimSpace(padded)
}
