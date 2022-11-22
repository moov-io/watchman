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
	req := httptest.NewRequest("GET", "/search/uk-csl?name=organization", nil)

	router := mux.NewRouter()
	addSearchRoutes(log.NewNopLogger(), router, isnSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"match":0.6333`)

	var wrapper struct {
		UKCSL []csl.UKCSLRecord `json:"ukConsolidatedSanctionsList"`
	}
	err := json.NewDecoder(w.Body).Decode(&wrapper)
	require.NoError(t, err)

	// require.Equal(t, "2d2db09c686e4829d0ef1b0b04145eec3d42cd88", prolif.EntityID)
}
