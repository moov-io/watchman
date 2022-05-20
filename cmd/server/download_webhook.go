// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"

	"github.com/moov-io/base/log"
)

func callDownloadWebook(logger log.Logger, stats *DownloadStats) {
	webhookURL := strings.TrimSpace(os.Getenv("DOWNLOAD_WEBHOOK_URL"))
	webhookAuthToken := strings.TrimSpace(os.Getenv("DOWNLOAD_WEBHOOK_AUTH_TOKEN"))

	if webhookURL == "" {
		return
	}
	logger.Info().Log("sending stats to download webhook url")

	var body bytes.Buffer
	json.NewEncoder(&body).Encode(stats)

	statusCode, err := callWebhook(&body, webhookURL, webhookAuthToken)
	if err != nil {
		logger.Error().LogErrorf("calling download webhook: %v", err)
	} else {
		logger.Info().Logf("http status code %d from download webhook", statusCode)
	}
}
