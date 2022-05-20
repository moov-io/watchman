// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

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
