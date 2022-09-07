// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestSearch__US_CSL(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search/us-csl?name=Khan&limit=1", nil)

	router := mux.NewRouter()
	addSearchRoutes(log.NewNopLogger(), router, isnSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"match":0.6333`)

	var wrapper struct {
		NonProliferationSanctions []csl.ISN `json:"nonProliferationSanctions"`
	}
	err := json.NewDecoder(w.Body).Decode(&wrapper)
	require.NoError(t, err)

	require.Len(t, wrapper.NonProliferationSanctions, 1)
	prolif := wrapper.NonProliferationSanctions[0]
	require.Equal(t, "2d2db09c686e4829d0ef1b0b04145eec3d42cd88", prolif.EntityID)
}

func TestSearcher_TopBISEntities(t *testing.T) {
	els := bisEntitySearcher.TopBISEntities(1, 0.00, "Khan")
	if len(els) == 0 {
		t.Fatal("empty ELs")
	}
	if els[0].Data.Name != "Mohammad Jan Khan Mangal" {
		t.Errorf("%#v", els[0].Data)
	}

	// Verify AlternateNames are passed through
	require.Len(t, els[0].Data.AlternateNames, 1)
	require.Equal(t, "Air I", els[0].Data.AlternateNames[0])
}

func TestSearcher_TopBISEntities_AltName(t *testing.T) {
	els := bisEntitySearcher.TopBISEntities(1, 0.00, "Luqman Sehreci.")
	if len(els) == 0 {
		t.Fatal("empty ELs")
	}
	if els[0].Data.Name != "Luqman Yasin Yunus Shgragi" {
		t.Errorf("%#v", els[0].Data)
	}
	if math.Abs(1.0-els[0].match) > 0.001 {
		t.Errorf("Expected match=1.0 for alt names: %f - %#v", els[0].match, els[0].Data)
	}
}

func TestSearcher_TopMEUs(t *testing.T) {
	meus := meuSearcher.TopMEUs(1, 0.00, "China Gas")
	require.Len(t, meus, 1)

	require.Equal(t, "d54346ef81802673c1b1daeb2ca8bd5d13755abd", meus[0].Data.EntityID)
	require.Equal(t, "0.70597", fmt.Sprintf("%.5f", meus[0].match))
}

func TestSearcher_TopSSIs(t *testing.T) {
	ssis := ssiSearcher.TopSSIs(1, 0.00, "ROSOBORONEKSPORT")
	if len(ssis) == 0 {
		t.Fatal("empty SSIs")
	}
	if ssis[0].Data.EntityID != "18782" {
		t.Errorf("%#v", ssis[0].Data)
	}

	// Verify AlternateNames are passed through
	require.Len(t, ssis[0].Data.AlternateNames, 6)
	require.Equal(t, "RUSSIAN DEFENSE EXPORT ROSOBORONEXPORT", ssis[0].Data.AlternateNames[0])
}

func TestSearcher_TopSSIs_limit(t *testing.T) {
	ssis := ssiSearcher.TopSSIs(2, 0.00, "SPECIALIZED DEPOSITORY")
	if len(ssis) != 2 {
		t.Fatalf("Expected 2 results, found %d", len(ssis))
	}
	if ssis[0].Data.EntityID != "18736" {
		t.Errorf("%#v", ssis[0].Data)
	}
}

func TestSearcher_TopSSIs_reportAltNameWeight(t *testing.T) {
	ssis := ssiSearcher.TopSSIs(1, 0.00, "KENKYUSHO")
	if len(ssis) == 0 {
		t.Fatal("empty SSIs")
	}
	if ssis[0].Data.EntityID != "18782" {
		t.Errorf("%f - %#v", ssis[0].match, ssis[0].Data)
	}
	if math.Abs(1.0-ssis[0].match) > 0.001 {
		t.Errorf("Expected match=1.0 for alt names: %f - %#v", ssis[0].match, ssis[0].Data)
	}
}

func TestSearcher_TopISNs(t *testing.T) {
	isns := isnSearcher.TopISNs(1, 0.00, "Abdul Qadeer K")
	require.Len(t, isns, 1)

	isn := isns[0]
	require.Equal(t, "2d2db09c686e4829d0ef1b0b04145eec3d42cd88", isn.Data.EntityID)
	require.Equal(t, "0.92", fmt.Sprintf("%.2f", isn.match))
}

func TestSearcher_TopUVLs(t *testing.T) {
	uvls := uvlSearcher.TopUVLs(1, 0.00, "Atlas Sanatgaran")
	require.Len(t, uvls, 1)

	uvl := uvls[0]
	require.Equal(t, "f15fa805ff4ac5e09026f5e78011a1bb6b26dec2", uvl.Data.EntityID)
	require.Equal(t, "1", strconv.Itoa(int(uvl.match)))
}

func TestSearcher_TopFSEs(t *testing.T) {
	fses := fseSearcher.TopFSEs(1, 0.00, "BEKTAS, Halis")
	require.Len(t, fses, 1)

	fse := fses[0]
	require.Equal(t, "17526", fse.Data.EntityID)
	require.Equal(t, "1", strconv.Itoa(int(fse.match)))
}

func TestSearcher_TopPLCs(t *testing.T) {
	plcs := plcSearcher.TopPLCs(1, 0.00, "SALAMEH, Salem")
	require.Len(t, plcs, 1)

	plc := plcs[0]
	require.Equal(t, "9702", plc.Data.EntityID)
	require.Equal(t, "1", strconv.Itoa(int(plc.match)))
}

func TestSearcher_TopCAPs(t *testing.T) {
	caps := capSearcher.TopCAPs(1, 0.00, "BM BANK PUBLIC JOINT STOCK COMPANY")
	require.Len(t, caps, 1)

	cap := caps[0]
	require.Equal(t, "20002", cap.Data.EntityID)
	require.Equal(t, "1", strconv.Itoa(int(cap.match)))
}

func TestSearcher_TopDTCs(t *testing.T) {
	dtcs := dtcSearcher.TopDTCs(1, 0.00, "Yasmin Ahmed")
	require.Len(t, dtcs, 1)

	dtc := dtcs[0]
	require.Equal(t, "d44d88d0265d93927b9ff1c13bbbb7c7db64142c", dtc.Data.EntityID)
	require.Equal(t, "1", strconv.Itoa(int(dtc.match)))
}

func TestSearcher_TopCMICs(t *testing.T) {
	cmics := cmicSearcher.TopCMICs(1, 0.00, "PROVEN HONOUR CAPITAL LIMITED")
	require.Len(t, cmics, 1)

	cmic := cmics[0]
	require.Equal(t, "32091", cmic.Data.EntityID)
	require.Equal(t, "1", strconv.Itoa(int(cmic.match)))
}

func TestSearcher_TopNSMBSs(t *testing.T) {
	ns_mbss := ns_mbsSearcher.TopNS_MBS(1, 0.00, "GAZPROMBANK JOINT STOCK COMPANY")
	require.Len(t, ns_mbss, 1)

	ns_mbs := ns_mbss[0]
	require.Equal(t, "17016", ns_mbs.Data.EntityID)
	require.Equal(t, "1", strconv.Itoa(int(ns_mbs.match)))
}
