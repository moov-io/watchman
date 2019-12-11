// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"strings"
)

// original list: inc, incorporated, llc, llp, co, ltd, limited, sa de cv, corporation, corp, ltda,
//                open joint stock company, pty ltd, public limited company, ag, cjsc, plc, as, aps,
//                oy, sa, gmbh, se, pvt ltd, sp zoo, ooo, sl, pjsc, jsc, bv, pt, tbk

var (
	companySuffixReplacer = strings.NewReplacer(
		" CO.", "",
		" INC.", "",
		" GMBH", "",
		" LLC", "",
		" LLP", "",
		" LTD.", "",
		" LTD ", " ",
		", LTD", "",
		" LTDA.", "",
	)
)

func removeCompanyTitles(in string) string {
	return companySuffixReplacer.Replace(in)
}
