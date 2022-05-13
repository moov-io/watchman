// Copyright 2022 The Moov Authors
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

func filterSDNs(sdns []*SDN, req filterRequest) []*SDN {
	if req.empty() {
		// short-circuit and return if we have no filters
		return sdns
	}

	keeper := keepSDN(req)

	var out []*SDN
	for i := range sdns {
		if keeper(sdns[i]) {
			out = append(out, sdns[i])
		}
	}
	return out
}

func keepSDN(req filterRequest) func(*SDN) bool {
	return func(sdn *SDN) bool {
		if req.empty() {
			return true // short-circuit if we have no filters
		}
		// by default exclude the result (as at least one filter is non-empty)
		keep := false

		// Look at all our filters
		// If the filter is non-empty AND matches the SDN's field then keep it
		//
		// NOTE: If we add more filters don't forget to also add them in values.go
		if req.sdnType != "" {
			if sdn.SDNType != "" {
				if strings.EqualFold(sdn.SDNType, req.sdnType) {
					keep = true
				}
			} else {
				// 'entity' is a special case value for ?sdnType in that it refers to a company or organization
				// and not an individual, however OFAC's data files do not contain this value and we must infer
				// it instead.
				if sdn.SDNType == "" && strings.EqualFold(req.sdnType, "entity") {
					keep = true
				} else {
					return false // skip this SDN as the filter didn't match
				}
			}
		}
		if req.ofacProgram != "" {
			for j := range sdn.Programs {
				if strings.EqualFold(sdn.Programs[j], req.ofacProgram) {
					keep = true
				}
			}
		}
		return keep
	}
}
