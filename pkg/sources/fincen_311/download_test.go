// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package fincen_311

import (
	"context"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	files, err := Download(context.Background(), log.NewNopLogger(), "")
	require.NoError(t, err)
	require.Len(t, files, 1)

	file, found := files["fincen_311.html"]
	require.True(t, found)
	require.NotNil(t, file)
	require.NoError(t, file.Close())
}

func TestFullPipeline(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// Download real page
	files, err := Download(context.Background(), log.NewNopLogger(), "")
	require.NoError(t, err)
	require.Len(t, files, 1)

	// Parse HTML
	data, err := Read(files)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotEmpty(t, data.ListHash)

	// Should have at least 10 entries (there are ~40+ on the real page)
	require.GreaterOrEqual(t, len(data.SpecialMeasures), 10,
		"expected at least 10 special measures entries")

	// Convert to entities
	entities := ConvertSpecialMeasures(data)
	require.Len(t, entities, len(data.SpecialMeasures))

	// Verify some known entities exist
	var foundIran, foundBank bool
	for _, e := range entities {
		nameLower := strings.ToLower(e.Name)
		if strings.Contains(nameLower, "iran") {
			foundIran = true
		}
		if strings.Contains(nameLower, "bank") {
			foundBank = true
		}
	}

	require.True(t, foundIran, "expected to find Iran in the list")
	require.True(t, foundBank, "expected to find at least one bank in the list")

	// Log some stats for debugging
	t.Logf("Parsed %d special measures entries", len(data.SpecialMeasures))

	var institutions, jurisdictions, transactions int
	for _, sm := range data.SpecialMeasures {
		switch sm.EntityType {
		case SMTypeFinancialInstitution:
			institutions++
		case SMTypeJurisdiction:
			jurisdictions++
		case SMTypeTransactionClass:
			transactions++
		}
	}
	t.Logf("Types: %d institutions, %d jurisdictions, %d transaction classes",
		institutions, jurisdictions, transactions)
}
