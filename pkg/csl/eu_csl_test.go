// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"os"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestEUCSL(t *testing.T) {
	t.Skip("CSL is currently broken, looks like they require API access now")

	if testing.Short() {
		t.Skip("ignorning network test")
	}

	logger := log.NewNopLogger()
	dir, err := os.MkdirTemp("", "eucsl")
	require.NoError(t, err)

	file, err := DownloadEU(logger, dir)
	require.NoError(t, err)

	eucslRecords, _, err := ReadEUFile(file)
	require.NoError(t, err)

	if len(eucslRecords) == 0 {
		t.Error("parsed zero EU CSL records")
	}
}
