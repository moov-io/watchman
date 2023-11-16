// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestSearch_UK_CSL(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search/uk-csl?name=%27ABD%20AL-NASIR", nil)

	router := mux.NewRouter()
	addSearchRoutes(log.NewNopLogger(), router, uk_cslSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"match":1`)
	require.Contains(t, w.Body.String(), `"matchedName":"'abd al nasir"`)

	var wrapper struct {
		UKCSL []csl.UKCSLRecord `json:"ukConsolidatedSanctionsList"`
	}
	err := json.NewDecoder(w.Body).Decode(&wrapper)
	require.NoError(t, err)

	require.Greater(t, len(wrapper.UKCSL), 0)

	require.Equal(t, int(13720), wrapper.UKCSL[0].GroupID)
}

func TestSearch_UK_SanctionsList(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search/uk-csl?name=HAJI%20KHAIRULLAH%20HAJI%20SATTAR%20MONEY%20EXCHANGE", nil)

	router := mux.NewRouter()
	addSearchRoutes(log.NewNopLogger(), router, uk_sanctionsListSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"match":1`)
	require.Contains(t, w.Body.String(), `"matchedName":"haji khairullah haji sattar money exchange"`)

	var wrapper struct {
		UKSanctionsList []csl.UKSanctionsListRecord `json:"ukSanctionsList"`
	}
	err := json.NewDecoder(w.Body).Decode(&wrapper)
	require.NoError(t, err)

	require.Greater(t, len(wrapper.UKSanctionsList), 0)

	require.Equal(t, "AFG0001", wrapper.UKSanctionsList[0].UniqueID)
}
