// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/watchman/pkg/ofac"
)

func TestPipeline__companyNameCleanupStep(t *testing.T) {
	nn := &Name{
		Processed: "SAI ADVISORS INC.",
		sdn: &ofac.SDN{
			SDNType: "",
		},
	}

	step := &companyNameCleanupStep{}
	if err := step.apply(nn); err != nil {
		t.Fatal(err)
	}

	if nn.Processed != "SAI ADVISORS" {
		t.Errorf("nn.Processed=%s", nn.Processed)
	}
}

func TestRemoveCompanyTitles(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"SIS D.O.O.", "SIS"},
		{"SAI ADVISORS INC.", "SAI ADVISORS"},                                                                  // SDN 24428
		{"COBALT REFINERY CO. INC.", "COBALT REFINERY"},                                                        // SDN 3748
		{"AL BARAKA EXCHANGE LLC", "AL BARAKA EXCHANGE"},                                                       // SDN 6953
		{"RUNNING BROOK, LLC (USA)", "RUNNING BROOK, (USA)"},                                                   // SDN 11589
		{"YAKIMA OIL TRADING, LLP", "YAKIMA OIL TRADING,"},                                                     // SDN 20259
		{"MKS INTERNATIONAL CO. LTD.", "MKS INTERNATIONAL"},                                                    // SDN 21553
		{"SHANGHAI NORTH TRANSWAY INTERNATIONAL TRADING CO.", "SHANGHAI NORTH TRANSWAY INTERNATIONAL TRADING"}, // SDN 22246
		{"DANDONG ZHICHENG METALLIC MATERIAL CO., LTD.", "DANDONG ZHICHENG METALLIC MATERIAL."},                // SDN 22603
		{"ADVANCED ELECTRONICS DEVELOPMENT, LTD", "ADVANCED ELECTRONICS DEVELOPMENT"},                          // SDN 8310
		{"AMD CO. LTD AGENCY", "AMD AGENCY"},                                                                   // SDN 8340
		{"REYNOLDS AND WILSON, LTD.", "REYNOLDS AND WILSON."},                                                  // SDN 8397
		{"AEROCOMERCIAL ALAS DE COLOMBIA LTDA.", "AEROCOMERCIAL ALAS DE COLOMBIA"},                             // SDN 8732,
		{"DIMABE LTDA.", "DIMABE"},                                                                             // SDN 8877
		{"ASCOTEC STEEL TRADING GMBH", "ASCOTEC STEEL TRADING"},                                                // SDN 11613
		{"TROPIC TOURS GMBH", "TROPIC TOURS"},                                                                  // SDN 2110,
		{"MC OVERSEAS TRADING COMPANY SA DE CV", "MC OVERSEAS TRADING COMPANY"},                                // SDN 10252
		{"SIRJANCO TRADING L.L.C.", "SIRJANCO TRADING"},                                                        // SDN 15985

		// Issue 483
		{"11420 CORP.", "11420 CORP."},
		{"11,420.2-1 CORP.", "11,420.2-1 CORP."},

		// Controls
		{"TADBIR ECONOMIC DEVELOPMENT GROUP", "TADBIR ECONOMIC DEVELOPMENT GROUP"}, // SDN 16006
		{"DI LAURO, Marco", "DI LAURO, Marco"},                                     // SDN 16128
		{"PETRO ROYAL FZE", "PETRO ROYAL FZE"},                                     // SDN 16136
	}
	for i := range cases {
		if ans := removeCompanyTitles(cases[i].input); cases[i].expected != ans {
			t.Errorf("#%d input=%q expected=%q got=%q", i, cases[i].input, cases[i].expected, ans)
		}
	}
}
