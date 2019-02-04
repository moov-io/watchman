// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestWebhook_batchSize(t *testing.T) {
	if d := readWebhookBatchSize(""); d != watchResearchBatchSize {
		t.Errorf("expected watchResearchBatchSize default, but got %d", d)
	}

	if d := readWebhookBatchSize("42"); d != 42 {
		t.Errorf("expected watchResearchBatchSize default, but got %d", d)
	}
}
