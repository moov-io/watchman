// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"strings"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/ofac"
)

type cryptoAddressSearchResult struct {
	OFAC []SDNWithDigitalCurrencyAddress `json:"ofac"`
}

func searchByCryptoAddress(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cryptoAddress := strings.TrimSpace(r.URL.Query().Get("address"))
		cryptoName := strings.TrimSpace(r.URL.Query().Get("name"))
		if cryptoAddress == "" {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}

		limit := extractSearchLimit(r)

		// Find SDNs with a crypto address that exactly matches
		resp := cryptoAddressSearchResult{
			OFAC: searcher.FindSDNCryptoAddresses(limit, cryptoName, cryptoAddress),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

type SDNWithDigitalCurrencyAddress struct {
	SDN *ofac.SDN `json:"sdn"`

	DigitalCurrencyAddresses []ofac.DigitalCurrencyAddress `json:"digitalCurrencyAddresses"`
}

func (s *searcher) FindSDNCryptoAddresses(limit int, name, needle string) []SDNWithDigitalCurrencyAddress {
	s.RLock()
	defer s.RUnlock()

	var out []SDNWithDigitalCurrencyAddress
	for i := range s.SDNComments {
		addresses := s.SDNComments[i].DigitalCurrencyAddresses
		for j := range addresses {
			// Skip addresses of a different coin
			if name != "" && addresses[j].Currency != name {
				continue
			}
			if addresses[j].Address == needle {
				// Find SDN
				sdn := s.findSDNWithoutLock(s.SDNComments[i].EntityID)
				if sdn != nil {
					out = append(out, SDNWithDigitalCurrencyAddress{
						SDN:                      sdn.SDN,
						DigitalCurrencyAddresses: addresses,
					})
				}
			}
		}
	}
	return out
}
