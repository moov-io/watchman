// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package us_tel

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"
)


func TestTELRecord_ToEntity_LegalEntity(t *testing.T) {
	r := TELRecord{
		ID:      "NK-245arpDkLm74auKv2h9w3m",
		Caption: "Japanese Red Army (JRA)",
		Schema:  "LegalEntity",
		Referents: []string{
			"us-dos-terr-9905ea8331b9d183638112392c6d35b1b4bafada",
			"us-dos-fto-078566d5306ffe94820da46e73a2311ec39f495c",
		},
		Datasets: []string{"us_state_terrorist_exclusion"},
		Origin:   []string{"inferred"},
		FirstSeen:  "2024-06-10T20:32:02",
		LastSeen:   "2026-06-09T20:32:01",
		LastChange: "2024-06-16T20:32:02",
		Properties: TELProperties{
			Program: []string{
				"Section 411 of the USA PATRIOT ACT of 2001 (8 U.S.C. § 1182) Terrorist Exclusion List (TEL)",
			},
			Topics: []string{"sanction"},
			Name: []string{
				"Japanese Red Army (JRA)",
				"JRA",
			},
		},
		Target: true,
	}

	ent := r.ToEntity()

	expName := "Japanese Red Army (JRA)"
	if ent.Name != expName {
		t.Errorf("expected Name %q, got %q", expName, ent.Name)
	}

	if ent.SourceID != r.ID {
		t.Errorf("expected SourceID %q, got %q", r.ID, ent.SourceID)
	}

	if ent.Type != search.EntityBusiness {
		t.Errorf("expected EntityBusiness, got %v", ent.Type)
	}

	if ent.Business == nil {
		t.Fatal("Business pointer should not be nil")
	}

	if ent.Business.Name != expName {
		t.Errorf("business name mismatch: expected %q got %q", expName, ent.Business.Name)
	}

	expAlts := []string{"JRA"}
	if len(ent.Business.AltNames) != len(expAlts) {
		t.Fatalf("expected %d alt names, got %d", len(expAlts), len(ent.Business.AltNames))
	}

	for i, v := range expAlts {
		if ent.Business.AltNames[i] != v {
			t.Errorf("alt name[%d] expected %q got %q", i, v, ent.Business.AltNames[i])
		}
	}
}

func TestTELRecord_ToEntity_UsesCaptionWhenNameMissing(t *testing.T) {
	r := TELRecord{
		ID:      "NK-46PWVzWZRu6PZ6YxW4emNg",
		Caption: "Darkazanli Company",
		Schema:  "LegalEntity",
		Properties: TELProperties{
			Program: []string{
				"Section 411 of the USA PATRIOT ACT of 2001 (8 U.S.C. § 1182) Terrorist Exclusion List (TEL)",
			},
			Topics: []string{"sanction"},
			Name:   []string{},
		},
		Target: true,
	}

	ent := r.ToEntity()

	if ent.Name != r.Caption {
		t.Errorf("expected Name from Caption %q, got %q", r.Caption, ent.Name)
	}

	if ent.Business == nil {
		t.Fatal("Business pointer should not be nil")
	}

	if ent.Business.Name != r.Caption {
		t.Errorf("expected Business.Name from Caption %q, got %q", r.Caption, ent.Business.Name)
	}
}

func TestTELRecord_ToEntity_EmptyAliasesIgnored(t *testing.T) {
	r := TELRecord{
		ID:      "NK-AEU6mpsqUQKpqbR9nmQv23",
		Caption: "Continuity Army Council",
		Schema:  "LegalEntity",
		Properties: TELProperties{
			Name: []string{
				"Continuity Army Council",
				"",
				"Continuity Irish Republican Army (CIRA)",
			},
			Topics: []string{"sanction"},
		},
		Target: true,
	}

	ent := r.ToEntity()

	expAlts := []string{"Continuity Irish Republican Army (CIRA)"}

	if ent.Business == nil {
		t.Fatal("Business pointer should not be nil")
	}

	if len(ent.Business.AltNames) != len(expAlts) {
		t.Fatalf("expected %d alt names, got %d", len(expAlts), len(ent.Business.AltNames))
	}

	for i, v := range expAlts {
		if ent.Business.AltNames[i] != v {
			t.Errorf("alt name[%d] expected %q got %q", i, v, ent.Business.AltNames[i])
		}
	}
}