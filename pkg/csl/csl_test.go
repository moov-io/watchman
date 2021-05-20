// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"io/ioutil"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/require"
)

func TestCSL(t *testing.T) {
	t.Skip("CSL is currently broken, looks like they require API access now")

	if testing.Short() {
		t.Skip("ignorning network test")
	}

	logger := log.NewNopLogger()
	dir, err := ioutil.TempDir("", "csl")
	require.NoError(t, err)

	file, err := Download(logger, dir)
	require.NoError(t, err)

	cslRecords, err := Read(file)
	require.NoError(t, err)

	if len(cslRecords.SSIs) == 0 {
		t.Error("parsed zero CSL SSI records")
	}
	if len(cslRecords.ELs) == 0 {
		t.Error("parsed zero CSL EL records")
	}
}
