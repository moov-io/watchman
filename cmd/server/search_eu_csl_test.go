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

func TestSearch__EU_CSL(t *testing.T) {
	w := httptest.NewRecorder()
	// misspelled on purpose to also check jaro winkler is picking up the right records
	req := httptest.NewRequest("GET", "/search/eu-csl?name=Saddam%20Hussien", nil)

	router := mux.NewRouter()
	addSearchRoutes(log.NewNopLogger(), router, eu_cslSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"match":0.92419`)
	require.Contains(t, w.Body.String(), `"matchedName":"saddam hussein al tikriti"`)

	var wrapper struct {
		EUConsolidatedSanctionsList []csl.EUCSLRecord `json:"euConsolidatedSanctionsList"`
	}
	err := json.NewDecoder(w.Body).Decode(&wrapper)
	require.NoError(t, err)

	require.Len(t, wrapper.EUConsolidatedSanctionsList, 1)
	prolif := wrapper.EUConsolidatedSanctionsList[0]
	require.Equal(t, 13, prolif.EntityLogicalID)
}
