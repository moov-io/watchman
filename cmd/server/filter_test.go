// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"testing"

	"github.com/moov-io/ofac"
)

func TestFilter__buildFilterRequest(t *testing.T) {
	u, _ := url.Parse("/search?q=jane+doe&sdnType=individual&program=SDGT")
	req := buildFilterRequest(u)
	if req.empty() {
		t.Error("filterRequest is not empty")
	}
	if req.sdnType != "individual" {
		t.Errorf("req.sdnType=%s", req.sdnType)
	}
	if req.program != "SDGT" {
		t.Errorf("req.program=%s", req.program)
	}

	// just the sdnType filter
	u, _ = url.Parse("/search?q=jane+doe&sdnType=individual")
	req = buildFilterRequest(u)
	if req.empty() {
		t.Error("filterRequest is not empty")
	}
	if req.sdnType != "individual" {
		t.Errorf("req.sdnType=%s", req.sdnType)
	}
	if req.program != "" {
		t.Errorf("req.program=%s", req.program)
	}

	// empty request
	u, _ = url.Parse("/search?q=jane+doe")
	req = buildFilterRequest(u)
	if !req.empty() {
		t.Error("filterRequest is empty")
	}
	if req.sdnType != "" || req.program != "" {
		t.Errorf("req.sdnType=%s req.program=%s", req.sdnType, req.program)
	}
}

var (
	filterableSDNs = []SDN{
		{
			SDN: &ofac.SDN{
				EntityID: "12",
				SDNName:  "Jane Doe",
				SDNType:  "individual",
				Program:  "other",
			},
		},
		{
			SDN: &ofac.SDN{
				EntityID: "13",
				SDNName:  "EP-1111",
				SDNType:  "aircraft",
				Program:  "SDGT",
			},
		},
	}
	missingSDNType = []SDN{
		{
			SDN: &ofac.SDN{
				EntityID: "14",
				SDNName:  "missing sdnType",
				Program:  "SDGT",
			},
		},
	}
	missingProgram = []SDN{
		{
			SDN: &ofac.SDN{
				EntityID: "15",
				SDNName:  "missing program",
				SDNType:  "individual",
			},
		},
	}
)

func TestFilter__sdnType(t *testing.T) {
	sdns := filterSDNs(filterableSDNs, filterRequest{sdnType: "individual"})
	if len(sdns) != 1 {
		t.Errorf("got: %#v", sdns)
	}
	if sdns[0].EntityID != "12" {
		t.Errorf("sdns[0].EntityID=%s", sdns[0].EntityID)
	}

	sdns = filterSDNs(filterableSDNs, filterRequest{})
	if len(sdns) != 2 {
		t.Errorf("got %#v", sdns)
	}

	sdns = filterSDNs(filterableSDNs, filterRequest{sdnType: "other"})
	if len(sdns) != 0 {
		t.Errorf("got %#v", sdns)
	}
}

func TestFilter__program(t *testing.T) {
	sdns := filterSDNs(filterableSDNs, filterRequest{program: "SDGT"})
	if len(sdns) != 1 {
		t.Errorf("got: %#v", sdns)
	}
	if sdns[0].EntityID != "13" {
		t.Errorf("sdns[0].EntityID=%s", sdns[0].EntityID)
	}

	sdns = filterSDNs(filterableSDNs, filterRequest{})
	if len(sdns) != 2 {
		t.Errorf("got %#v", sdns)
	}

	sdns = filterSDNs(filterableSDNs, filterRequest{program: "unknown"})
	if len(sdns) != 0 {
		t.Errorf("got %#v", sdns)
	}
}

func TestFilter__multiple(t *testing.T) {
	sdns := filterSDNs(filterableSDNs, filterRequest{sdnType: "aircraft", program: "SDGT"})
	if len(sdns) != 1 {
		t.Errorf("got: %#v", sdns)
	}
	if sdns[0].EntityID != "13" {
		t.Errorf("sdns[0].EntityID=%s", sdns[0].EntityID)
	}

	sdns = filterSDNs(filterableSDNs, filterRequest{})
	if len(sdns) != 2 {
		t.Errorf("got %#v", sdns)
	}

	sdns = filterSDNs(filterableSDNs, filterRequest{sdnType: "other", program: "unknown"})
	if len(sdns) != 0 {
		t.Errorf("got %#v", sdns)
	}
}

func TestFilter__missing(t *testing.T) {
	if len(missingSDNType) != 1 {
		t.Fatalf("%#v", missingSDNType)
	}
	sdns := filterSDNs(append(filterableSDNs, missingSDNType...), filterRequest{sdnType: "individual"})
	if len(sdns) != 1 || sdns[0].EntityID != "12" {
		t.Errorf("sdns=%#v", sdns)
	}

	if len(missingProgram) != 1 {
		t.Fatalf("%#v", missingProgram)
	}
	sdns = filterSDNs(append(filterableSDNs, missingProgram...), filterRequest{program: "SDGT"})
	if len(sdns) != 1 || sdns[0].EntityID != "13" {
		t.Errorf("sdns=%#v", sdns)
	}
}
