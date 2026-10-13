// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestToEntity_NonproliferationSanctions(t *testing.T) {
	cases := []struct {
		name   string
		source string
	}{
		{
			name:   "legacy ISN source code",
			source: "Nonproliferation Sanctions (ISN) - State Department",
		},
		{
			name:   "ACN source code",
			source: "Nonproliferation Sanctions (ACN) - State Department",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			entity := ToEntity(SanctionsEntry{
				ID:       "test-id",
				Source:   tc.source,
				Name:     "Sinobright Import and Export Company",
				Programs: "INKSNA",
			})

			require.Equal(t, search.EntityBusiness, entity.Type)
			require.Equal(t, search.SourceUSCSL, entity.Source)
			require.Equal(t, "test-id", entity.SourceID)
			require.NotNil(t, entity.Business)
			require.Nil(t, entity.Person)
			require.Equal(t, "Sinobright Import and Export Company", entity.Name)

			us, ok := entity.SourceData.(SanctionsEntry)
			require.True(t, ok)
			require.Equal(t, tc.source, us.Source)
		})
	}
}

func TestIsNonproliferationSanctions(t *testing.T) {
	cases := []struct {
		source string
		want   bool
	}{
		{source: "Nonproliferation Sanctions (ISN) - State Department", want: true},
		{source: "Nonproliferation Sanctions (ACN) - State Department", want: true},
		{source: "(ACN) - State Department", want: true},
		{source: "(ISN) - State Department", want: true},
		{source: "Denied Persons List (DPL) - Bureau of Industry and Security", want: false},
		{source: "Entity List (EL) - Bureau of Industry and Security", want: false},
		{source: "ITAR Debarred (DTC) - State Department", want: false},
	}

	for _, tc := range cases {
		t.Run(tc.source, func(t *testing.T) {
			require.Equal(t, tc.want, isNonproliferationSanctions(tc.source))
		})
	}
}

func TestGetAllNamesFromCSV(t *testing.T) {
	// Create a sample SanctionsEntry with alternate names
	entry := SanctionsEntry{
		ID:       "9640",
		Name:     "Mohammed ABU TEIR",
		Type:     "Individual",
		AltNames: "Mohammed Mahmud ABU TAIR;Mohammad Mahmoud ABOU TAYR",
	}

	got := getAllNames(entry, "individual")
	want := []string{
		"Mohammed Mahmud ABU TAIR",
		"Mohammad Mahmoud ABOU TAYR",
	}

	require.Len(t, got, len(want), "expected %d alternate names, got %d", len(want), len(got))
	for i, name := range got {
		require.Equal(t, want[i], name, "alternate name at index %d", i)
	}
}

func TestMapFeaturesFromCSV(t *testing.T) {
	// Create a sample SanctionsEntry with birth date and place of birth
	entry := SanctionsEntry{
		ID:           "9640",
		Name:         "Mohammed ABU TEIR",
		Type:         "Individual",
		DatesOfBirth: "1951",
		IDs:          "Gender, Male",
	}

	person := mapPerson(entry)

	// Check birth date year
	if person.BirthDate == nil {
		t.Fatal("BirthDate is nil, want non-nil")
	}
	require.Equal(t, 1951, person.BirthDate.Year(), "expected BirthDate year to be 1951")

	// Check gender
	require.Equal(t, search.GenderMale, person.Gender, "expected Gender to be Male")

	// Note: PlaceOfBirth is not supported in search.Person; commented-out test is ignored
}

func TestMapSanctionsProgramsFromCSV(t *testing.T) {
	// Create a sample SanctionsEntry with programs
	entry := SanctionsEntry{
		ID:       "9640",
		Name:     "Mohammed ABU TEIR",
		Type:     "Individual",
		Programs: "NS-PLC",
	}

	got := mapSanctionsInfo(entry)
	require.Len(t, got.Programs, 1, "expected 1 program")
	require.Equal(t, "NS-PLC", got.Programs[0], "expected program to be NS-PLC")
}

func TestToEntity_SwiftBIC(t *testing.T) {
	t.Run("type-first SWIFT/BIC", func(t *testing.T) {
		got := ToEntity(SanctionsEntry{
			ID:   "18283",
			Name: "SME BANK",
			Type: "Entity",
			IDs:  "Public Registration Number, 1027739108649; SWIFT/BIC, RUDVRUMM; Website, mspbank.ru",
		})
		require.Equal(t, []string{"RUDVRUMM"}, got.SwiftBICs)
		require.Equal(t, search.EntityOrganization, got.Type)
	})

	t.Run("value-first SWIFT/BIC", func(t *testing.T) {
		got := ToEntity(SanctionsEntry{
			ID:   "15268",
			Name: "BANK OF KUNLUN CO LTD",
			Type: "Entity",
			IDs:  "CKLBCNBJ, SWIFT/BIC",
		})
		require.Equal(t, []string{"CKLBCNBJ"}, got.SwiftBICs)
	})

	t.Run("Gazprombank GAZPRUMM", func(t *testing.T) {
		got := ToEntity(SanctionsEntry{
			ID:   "17016",
			Name: "GAZPROMBANK OAO",
			Type: "Entity",
			IDs:  "09807684, Government Gazette Number; 7744001497, Tax ID No.; GAZPRUMM, SWIFT/BIC",
		})
		require.Equal(t, []string{"GAZPRUMM"}, got.SwiftBICs)
	})
}

func TestToEntity_SwiftBICFromTestdata(t *testing.T) {
	ctx := context.Background()
	logger := log.NewTestLogger()

	wd, err := os.Getwd()
	require.NoError(t, err)

	files, err := Download(ctx, logger, filepath.Join(wd, "testdata"))
	require.NoError(t, err)

	data, err := Read(files)
	require.NoError(t, err)

	want := map[string][]string{
		"15268": {"CKLBCNBJ"},
		"18283": {"RUDVRUMM"},
		"17016": {"GAZPRUMM"},
	}

	found := 0
	withBIC := 0
	for _, row := range data.SanctionsData {
		e := ToEntity(row)
		if len(e.SwiftBICs) > 0 {
			withBIC++
		}
		expect, ok := want[e.SourceID]
		if !ok {
			continue
		}
		require.Equal(t, expect, e.SwiftBICs, e.SourceID)
		require.Equal(t, expect, e.PreparedFields.SwiftBICs, e.SourceID)
		found++
	}
	require.Equal(t, len(want), found)
	require.Greater(t, withBIC, 20)
}

func TestMapIDTypeFromCSV(t *testing.T) {
	// Create a sample SanctionsEntry with IDs
	entry := SanctionsEntry{
		ID:   "26182",
		Name: "Evren KAYAKIRAN",
		Type: "Individual",
		IDs:  "Passport, U00242309, TR",
	}

	ids := mapGovernmentIDs(entry)
	require.Len(t, ids, 1, "expected 1 government ID")

	want := search.GovernmentID{
		Name:       "Passport",
		Type:       search.GovernmentIDPassport,
		Country:    "Turkey",
		Identifier: "U00242309",
	}
	require.Equal(t, want, ids[0], "expected government ID to match")
}

func TestReaderWithFile(t *testing.T) {
	ctx := context.Background()
	logger := log.NewTestLogger()

	// Download the gzipped CSV file
	wd, err := os.Getwd()
	require.NoError(t, err)

	files, err := Download(ctx, logger, filepath.Join(wd, "testdata"))
	require.NoError(t, err, "failed to download testdata/consolidated.csv")
	require.Len(t, files, 1, "expected 1 file")

	// Read and parse the CSV
	data, err := Read(files)
	require.NoError(t, err, "failed to read CSV")
	require.NotEmpty(t, data.ListHash, "expected non-empty ListHash")

	// Check the number of entities (20 total, 1 SDN filtered out = 19)
	entities := data.SanctionsData
	require.Len(t, entities, 5512, fmt.Sprintf("found %d entities after filtering out SDN record", len(entities)))

	// Verify a sample entity
	found := false
	for _, entity := range entities {
		if entity.ID == "1b5979e8fd3920a1a0985bf5088dad930a6a5562d1513f9e4b015e42" {
			require.Equal(t, "Atlas Sanatgaran", entity.Name, "expected name to match")
			require.Equal(t, "Unverified List (UVL) - Bureau of Industry and Security", entity.Source, "expected source to match")
			require.Equal(t, "Komitas 26/114, Yerevan, Armenia, AM", entity.Addresses, "expected addresses to match")
			found = true
			break
		}
	}
	require.True(t, found, "expected to find Atlas Sanatgaran entity")
}
