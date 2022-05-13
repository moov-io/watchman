// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/ofac"
)

var (
	noopPipeliner = &pipeliner{
		logger: log.NewNopLogger(),
		steps:  []step{},
	}

	noLogPipeliner = newPipeliner(log.NewNopLogger())
)

func TestPipelineNoop(t *testing.T) {
	if err := noopPipeliner.Do(&Name{}); err != nil {
		t.Fatal(err)
	}
}

func TestFullPipeline(t *testing.T) {
	individual := func(in string) *Name {
		return &Name{
			Processed: in,
			sdn: &ofac.SDN{
				SDNType: "individual",
			},
		}
	}
	company := func(in string) *Name {
		return &Name{
			Processed: in,
			sdn: &ofac.SDN{
				SDNType: "", // blank refers to a company
			},
		}
	}

	cases := []struct {
		in       *Name
		expected string
	}{
		// input edge cases
		{individual(""), ""},
		{individual(" "), ""},
		{individual("  "), ""},
		{company(""), ""},
		{company(" "), ""},
		{company("  "), ""},

		// Re-order individual names
		{individual("MADURO MOROS, Nicolas"), "nicolas maduro moros"},

		// Remove Company Suffixes
		{company("YAKIMA OIL TRADING, LLP"), "yakima oil trading"},                                                      // SDN 20259
		{company("MKS INTERNATIONAL CO. LTD."), "mks international"},                                                    // SDN 21553
		{company("SHANGHAI NORTH TRANSWAY INTERNATIONAL TRADING CO."), "shanghai north transway international trading"}, // SDN 22246

		// Remove stopwords
		{company("INVERSIONES LA QUINTA Y CIA. LTDA."), "inversiones la quinta y cia"},

		// Normalize ("-" -> " ")
		{company("ANGLO-CARIBBEAN CO., LTD."), "anglo caribbean"},
	}
	for i := range cases {
		if err := noLogPipeliner.Do(cases[i].in); err != nil {
			t.Error(err)
		} else {
			if cases[i].in.Processed != cases[i].expected {
				t.Errorf("%d# output=%q expected=%q", i, cases[i].in.Processed, cases[i].expected)
			}
		}
	}
}
