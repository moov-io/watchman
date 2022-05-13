// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type companyNameCleanupStep struct {
}

func (s *companyNameCleanupStep) apply(in *Name) error {
	switch {
	case in.sdn != nil && in.sdn.SDNType == "":
		in.Processed = removeCompanyTitles(in.Processed)

	case in.ssi != nil && in.ssi.Type == "":
		in.Processed = removeCompanyTitles(in.Processed)
	}
	return nil
}

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

func removeCompanyTitles(in string) string {
	return companySuffixReplacer.Replace(in)
}
