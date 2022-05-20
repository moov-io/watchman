// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleDownloadStats(t *testing.T) {
	updates := make(chan *DownloadStats)

	var wg sync.WaitGroup // make the race detector happy

	var received *DownloadStats
	go handleDownloadStats(updates, func(stats *DownloadStats) {
		received = stats
		wg.Done()
	})

	wg.Add(1)
	updates <- &DownloadStats{
		SDNs: 123,
	}
	wg.Wait()

	require.NotNil(t, received)
	require.Equal(t, 123, received.SDNs)
}
