// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/stretchr/testify/require"
)

func TestDownloadStats(t *testing.T) {
	when := time.Date(2022, time.May, 21, 9, 4, 0, 0, time.UTC)
	bs, err := json.Marshal(&DownloadStats{
		SDNs: 1,
		Errors: []error{
			errors.New("bad thing"),
		},
		RefreshedAt: when,
	})
	require.NoError(t, err)

	var wrapper struct {
		SDNs      int
		Errors    []string
		Timestamp time.Time
	}
	err = json.NewDecoder(bytes.NewReader(bs)).Decode(&wrapper)
	require.NoError(t, err)

	require.Equal(t, 1, wrapper.SDNs)
	require.Len(t, wrapper.Errors, 1)
	require.Equal(t, when, wrapper.Timestamp)
}

func TestSearcher__refreshInterval(t *testing.T) {
	if v := getDataRefreshInterval(log.NewNopLogger(), ""); v.String() != "12h0m0s" {
		t.Errorf("Got %v", v)
	}
	if v := getDataRefreshInterval(log.NewNopLogger(), "60s"); v.String() != "1m0s" {
		t.Errorf("Got %v", v)
	}
	if v := getDataRefreshInterval(log.NewNopLogger(), "off"); v != 0*time.Second {
		t.Errorf("got %v", v)
	}

	// cover another branch
	s := newSearcher(log.NewNopLogger(), noLogPipeliner, 1)
	s.periodicDataRefresh(0*time.Second, nil)
}

func TestSearcher__refreshData(t *testing.T) {
	s := createTestSearcher(t) // TODO(adam): initial setup
	stats := testSearcherStats

	if len(s.Addresses) == 0 || stats.Addresses == 0 {
		t.Errorf("empty Addresses=%d stats.Addresses=%d", len(s.Addresses), stats.Addresses)
	}
	if len(s.Alts) == 0 || stats.Alts == 0 {
		t.Errorf("empty Alts=%d or stats.Alts=%d", len(s.Alts), stats.Alts)
	}
	if len(s.SDNs) == 0 || stats.SDNs == 0 {
		t.Errorf("empty SDNs=%d or stats.SDNs=%d", len(s.SDNs), stats.SDNs)
	}
	if len(s.DPs) == 0 || stats.DeniedPersons == 0 {
		t.Errorf("empty DPs=%d or stats.DeniedPersons=%d", len(s.DPs), stats.DeniedPersons)
	}
	if len(s.SSIs) == 0 || stats.SectoralSanctions == 0 {
		t.Errorf("empty SSIs=%d or stats.SectoralSanctions=%d", len(s.SSIs), stats.SectoralSanctions)
	}
	if len(s.BISEntities) == 0 || stats.BISEntities == 0 {
		t.Errorf("empty searcher.BISEntities=%d or stats.BISEntities=%d", len(s.BISEntities), stats.BISEntities)
	}
}

func TestDownload__lastRefresh(t *testing.T) {
	start := time.Now()
	time.Sleep(5 * time.Millisecond) // force start to be before our calls

	if when := lastRefresh(""); when.Before(start) {
		t.Errorf("expected time.Now()")
	}

	// make a temp dir (initially with nothing in it)
	dir, err := os.MkdirTemp("", "lastRefresh")
	if err != nil {
		t.Fatal(err)
	}

	if when := lastRefresh(dir); !when.IsZero() {
		t.Errorf("expected zero time: %v", t)
	}

	// add a file and get it's mtime
	path := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(path, []byte("hello, world"), 0600); err != nil {
		t.Fatal(err)
	}
	if info, err := os.Stat(path); err != nil {
		t.Fatal(err)
	} else {
		if when := lastRefresh(dir); !when.Equal(info.ModTime()) {
			t.Errorf("t=%v", when)
		}
	}
}

func TestDownload__OFAC_Spillover(t *testing.T) {
	logger := log.NewTestLogger()
	initialDir := filepath.Join("..", "..", "test", "testdata", "static")

	res, err := ofacRecords(logger, initialDir)
	require.NoError(t, err)

	var sdn *ofac.SDN
	idx := slices.IndexFunc(res.SDNs, func(s *ofac.SDN) bool {
		return s.EntityID == "12300"
	})
	if idx >= 0 {
		sdn = res.SDNs[idx]
	}
	require.NotNil(t, sdn)

	//nolint:misspell
	expected := `DOB 13 May 1965; alt. DOB 13 Apr 1968; alt. DOB 07 Jul 1964; POB Medellin, Colombia; alt. POB Marinilla, Antioquia, Colombia; alt. POB Ciudad Victoria, Tamaulipas, Mexico; Cedula No. 7548733 (Colombia); alt. Cedula No. 70163752 (Colombia); alt. Cedula No. 172489729-1 (Ecuador); Passport AL720622 (Colombia); R.F.C. CIVJ650513LJA (Mexico); alt. R.F.C. OUSV-640707 (Mexico); C.U.R.P. CIVJ650513HNEFLR06 (Mexico); alt. C.U.R.P. OUVS640707HTSSLR07 (Mexico); Matricula Mercantil No 181301-1 Cali (Colombia); alt. Matricula Mercantil No 405885 Bogota (Colombia); Linked To: BIO FORESTAL S.A.S.; Linked To: CUBI CAFE CLICK CUBE MEXICO, S.A. DE C.V.; Linked To: DOLPHIN DIVE SCHOOL S.A.; Linked To: GANADERIA LA SORGUITA S.A.S.; Linked To: GESTORES DEL ECUADOR GESTORUM S.A.; Linked To: INVERPUNTO DEL VALLE S.A.; Linked To: INVERSIONES CIFUENTES Y CIA. S. EN C.; Linked To: LE CLAUDE, S.A. DE C.V.; Linked To: OPERADORA NUEVA GRANADA, S.A. DE C.V.; Linked To: PARQUES TEMATICOS S.A.S.; Linked To: PROMO RAIZ S.A.S.; Linked To: RED MUNDIAL INMOBILIARIA, S.A. DE C.V.; Linked To: FUNDACION PARA EL BIENESTAR Y EL PORVENIR; Linked To: C.I. METALURGIA EXTRACTIVA DE COLOMBIA S.A.S.; Linked To: GRUPO MUNDO MARINO, S.A.; Linked To: C.I. DISERCOM S.A.S.; Linked To: C.I. OKCOFFEE COLOMBIA S.A.S.; Linked To: C.I. OKCOFFEE INTERNATIONAL S.A.S.; Linked To: FUNDACION OKCOFFEE COLOMBIA; Linked To: CUBICAFE S.A.S.; Linked To: HOTELES Y BIENES S.A.; Linked To: FUNDACION SALVA LA SELVA; Linked To: LINEA AEREA PUEBLOS AMAZONICOS S.A.S.; Linked To: DESARROLLO MINERO RESPONSABLE C.I. S.A.S.; Linked To: R D I S.A.`
	require.Equal(t, expected, sdn.Remarks)
}
