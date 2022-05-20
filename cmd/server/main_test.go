// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleDownloadStats(t *testing.T) {
	updates := make(chan *DownloadStats)

	var received *DownloadStats
	go handleDownloadStats(updates, func(stats *DownloadStats) {
		received = stats
	})

	updates <- &DownloadStats{
		SDNs: 123,
	}

	require.NotNil(t, received)
	require.Equal(t, 123, received.SDNs)
}
