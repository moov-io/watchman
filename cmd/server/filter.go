// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"strings"
)

type filterRequest struct {
	sdnType string
	program string
}

func (req filterRequest) empty() bool {
	return req.sdnType == "" && req.program == ""
}

func buildFilterRequest(u *url.URL) filterRequest {
	return filterRequest{
		sdnType: u.Query().Get("sdnType"),
		program: u.Query().Get("program"),
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
		// TODO(adam): If we add more filters don't forget to also add them in values.go
		if req.sdnType != "" {
			if sdns[i].SDNType != "" && strings.EqualFold(sdns[i].SDNType, req.sdnType) {
				keep = true
			} else {
				continue // skip this SDN as the filter didn't match
			}
		}
		if req.program != "" {
			if sdns[i].Program != "" && strings.EqualFold(sdns[i].Program, req.program) {
				keep = true
			} else {
				continue
			}
		}

		if keep {
			out = append(out, sdns[i])
		}
	}
	return out
}
