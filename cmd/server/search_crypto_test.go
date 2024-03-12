// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

var (
	cryptoSearcher = newSearcher(log.NewNopLogger(), noLogPipeliner, 1)
)

func init() {
	// Set SDN Comments
	fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", "sdn_comments.csv"))
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	ofacResults, err := ofac.Read(map[string]io.ReadCloser{"sdn_comments.csv": fd})
	if err != nil {
		panic(fmt.Sprintf("ERROR reading sdn_comments.csv: %v", err))
	}

	cryptoSearcher.SDNComments = ofacResults.SDNComments
	cryptoSearcher.SDNs = precomputeSDNs([]*ofac.SDN{
		{
			EntityID: "39796", // matches TestSearchCrypto
			SDNName:  "Person A",
			SDNType:  "individual",
			Title:    "Guy or Girl doing crypto stuff",
		},
	}, nil, noLogPipeliner)
}

func TestSearchCryptoSetup(t *testing.T) {
	require.Len(t, cryptoSearcher.SDNComments, 13)
	require.Len(t, cryptoSearcher.SDNs, 1)
}

type expectedCryptoAddressSearchResult struct {
	OFAC []SDNWithDigitalCurrencyAddress `json:"ofac"`
}

func TestSearchCrypto(t *testing.T) {
	router := mux.NewRouter()
	addSearchRoutes(log.NewNopLogger(), router, cryptoSearcher)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/crypto?address=0x242654336ca2205714071898f67E254EB49ACdCe", nil)
	router.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	var response expectedCryptoAddressSearchResult
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	require.Len(t, response.OFAC, 1)
	require.Equal(t, "39796", response.OFAC[0].SDN.EntityID)

	// Now with cryptocurrency name specified
	req = httptest.NewRequest("GET", "/crypto?name=ETH&address=0x242654336ca2205714071898f67E254EB49ACdCe", nil)
	router.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	require.Len(t, response.OFAC, 1)
	require.Equal(t, "39796", response.OFAC[0].SDN.EntityID)

	// With wrong cryptocurrency name
	req = httptest.NewRequest("GET", "/crypto?name=QRR&address=0x242654336ca2205714071898f67E254EB49ACdCe", nil)
	router.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	require.Len(t, response.OFAC, 0)
}
