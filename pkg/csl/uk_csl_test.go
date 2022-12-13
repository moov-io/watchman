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
	t.Skip("skipping UK CSL download")

	if testing.Short() {
		t.Skip("ignorning network test")
	}

	logger := log.NewNopLogger()
	dir, err := os.MkdirTemp("", "ukcsl")
	require.NoError(t, err)

	file, err := DownloadUKCSL(logger, dir)
	require.NoError(t, err)

	ukcslRecords, _, err := ReadUKCSLFile(file)
	require.NoError(t, err)

	if len(ukcslRecords) == 0 {
		t.Error("parsed zero UK CSL records")
	}
}

func TestUKSanctionsList(t *testing.T) {
	t.Skip("skipping UK Sanctions List download")

	if testing.Short() {
		t.Skip("ignorning network test")
	}

	logger := log.NewNopLogger()
	dir, err := os.MkdirTemp("", "ukSanctionsList")
	require.NoError(t, err)

	file, err := DownloadUKSanctionsList(logger, dir)
	require.NoError(t, err)

	ukSanctionsListRecords, _, err := ReadUKSanctionsListFile(file)
	require.NoError(t, err)

	if len(ukSanctionsListRecords) == 0 {
		t.Error("parsed zero UK Sanctions List records")
	}
}
