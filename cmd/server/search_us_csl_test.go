// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"testing"
)

func TestSearcher_TopBISEntities(t *testing.T) {
	els := bisEntitySearcher.TopBISEntities(1, 0.00, "Khan")
	if len(els) == 0 {
		t.Fatal("empty ELs")
	}
	if els[0].Data.Name != "Mohammad Jan Khan Mangal" {
		t.Errorf("%#v", els[0].Data)
	}
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

func TestSearcher_TopSSIs(t *testing.T) {
	ssis := ssiSearcher.TopSSIs(1, 0.00, "ROSOBORONEKSPORT")
	if len(ssis) == 0 {
		t.Fatal("empty SSIs")
	}
	if ssis[0].Data.EntityID != "18782" {
		t.Errorf("%#v", ssis[0].Data)
	}
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
