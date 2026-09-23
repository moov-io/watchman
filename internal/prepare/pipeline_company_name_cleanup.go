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
