// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ru_fa

import (
	"context"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

func Download(ctx context.Context, logger log.Logger, initialDir string) (download.Files, error) {
	dl := download.New(logger, download.HTTPClient)

	addrs := make(map[string]string)
		addrs[htmlFilename] = urlToDownload

	return dl.GetFiles(ctx, initialDir, addrs)
}
