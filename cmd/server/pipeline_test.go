// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/go-kit/kit/log"
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
	sdn := func(in string) *Name {
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
		// Remove Company Suffixes
		{sdn("YAKIMA OIL TRADING, LLP"), "yakima oil trading"},                                                      // SDN 20259
		{sdn("MKS INTERNATIONAL CO. LTD."), "mks international"},                                                    // SDN 21553
		{sdn("SHANGHAI NORTH TRANSWAY INTERNATIONAL TRADING CO."), "shanghai north transway international trading"}, // SDN 22246

		// Remove stopwords
		{sdn("INVERSIONES LA QUINTA Y CIA. LTDA."), "inversiones la quinta y cia"},
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
