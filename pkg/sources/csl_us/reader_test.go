// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"context"
	"fmt"
	"testing"

	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	ctx := context.Background()
	logger := log.NewTestLogger()

	// Download the gzipped CSV file
	files, err := Download(ctx, logger, "testdata")
	require.NoError(t, err)
	require.Len(t, files, 1)

	// Read and parse the CSV
	data, err := Read(files)
	require.NoError(t, err)
	require.NotEmpty(t, data.ListHash)

	// Check the number of entities (20 total, 1 SDN filtered out = 19)
	entities := data.SanctionsData
	require.Len(t, entities, 5512, fmt.Sprintf("found %d entities after filtering out SDN record", len(entities)))

	// Verify a sample entity
	for _, entity := range entities {
		if entity.ID == "1b5979e8fd3920a1a0985bf5088dad930a6a5562d1513f9e4b015e42" {
			require.Equal(t, "Atlas Sanatgaran", entity.Name)
			require.Equal(t, "Unverified List (UVL) - Bureau of Industry and Security", entity.Source)
			require.Equal(t, "Komitas 26/114, Yerevan, Armenia, AM", entity.Addresses)
			break
		}
	}
}
