// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"strings"
)

// original list: inc, incorporated, llc, llp, co, ltd, limited, sa de cv, corporation, corp, ltda,
//                open joint stock company, pty ltd, public limited company, ag, cjsc, plc, as, aps,
//                oy, sa, gmbh, se, pvt ltd, sp zoo, ooo, sl, pjsc, jsc, bv, pt, tbk

var (
	companySuffixReplacer = strings.NewReplacer(
		" CO.", "",
		" D.O.O.", "",
		" INC.", "",
		" GMBH", "",
		" LLC", "",
		" L.L.C.", "",
		" LLP", "",
		" LTD.", "",
		" LTD ", " ",
		", LTD", "",
		" LTDA.", "",
		" SA DE CV", "",
	)
)

func RemoveCompanyTitles(in string) string {
	return companySuffixReplacer.Replace(in)
}

// companySuffixTokens are legal-form words after LowerAndRemovePunctuation.
// Candidate selection treats them as optional so a query of
// "Ocean Shipping Limited" still matches a DBA of "Ocean Shipping".
var companySuffixTokens = map[string]struct{}{
	"ag":           {},
	"aps":          {},
	"bv":           {},
	"cjsc":         {},
	"co":           {},
	"company":      {},
	"corp":         {},
	"corporation":  {},
	"gmbh":         {},
	"inc":          {},
	"incorporated": {},
	"jsc":          {},
	"limited":      {},
	"llc":          {},
	"llp":          {},
	"ltd":          {},
	"ltda":         {},
	"nv":           {},
	"ooo":          {},
	"oy":           {},
	"pjsc":         {},
	"plc":          {},
	"pt":           {},
	"pty":          {},
	"pvt":          {},
	"sa":           {},
	"se":           {},
	"sl":           {},
}

// IsCompanySuffixToken reports whether tok is a legal-entity suffix
// (limited, llc, gmbh, …) after name normalization.
func IsCompanySuffixToken(tok string) bool {
	_, ok := companySuffixTokens[tok]
	return ok
}
