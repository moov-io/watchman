// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"strings"
)

type filterRequest struct {
	sdnType     string
	ofacProgram string
}

func (req filterRequest) empty() bool {
	return req.sdnType == "" && req.ofacProgram == ""
}

func buildFilterRequest(u *url.URL) filterRequest {
	return filterRequest{
		sdnType:     u.Query().Get("sdnType"),
		ofacProgram: u.Query().Get("ofacProgram"),
	}
}

func filterSDNs(sdns []SDN, req filterRequest) []SDN {
	if req.empty() {
		// short-circuit and return if we have no filters
		return sdns
	}

	var out []SDN
	for i := range sdns {
		// by default exclude the result (as at least one filter is non-empty)
		keep := false

		// Look at all our filters
		// If the filter is non-empty AND matches the SDN's field then keep it
		//
		// NOTE: If we add more filters don't forget to also add them in values.go
		if req.sdnType != "" {
			if sdns[i].SDNType != "" {
				if strings.EqualFold(sdns[i].SDNType, req.sdnType) {
					keep = true
				}
			} else {
				// 'entity' is a special case value for ?sdnType in that it refers to a company or organization
				// and not an individual, however OFAC's data files do not contain this value and we must infer
				// it instead.
				if sdns[i].SDNType == "" && strings.EqualFold(req.sdnType, "entity") {
					keep = true
				} else {
					continue // skip this SDN as the filter didn't match
				}
			}
		}
		if req.ofacProgram != "" {
			for j := range sdns[i].Programs {
				if strings.EqualFold(sdns[i].Programs[j], req.ofacProgram) {
					keep = true
				}
			}
		}

		if keep {
			out = append(out, sdns[i])
		}
	}
	return out
}
