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

func TestUKCSL(t *testing.T) {
	t.Skip("UK CSL is currently broken, looks like they require API access now")

	if testing.Short() {
		t.Skip("ignorning network test")
	}

	logger := log.NewNopLogger()
	dir, err := os.MkdirTemp("", "ukcsl")
	require.NoError(t, err)

	file, err := DownloadUK(logger, dir)
	require.NoError(t, err)

	ukcslRecords, _, err := ReadUKFile(file)
	require.NoError(t, err)

	if len(ukcslRecords) == 0 {
		t.Error("parsed zero UK CSL records")
	}
}
