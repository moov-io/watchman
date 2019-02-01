// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestWebhooks_validate(t *testing.T) {
	out, err := validateWebhook("")
	if err == nil {
		t.Error("expected error")
	}
	if out != "" {
		t.Errorf("got out=%q", out)
	}

	// happy path
	out, err = validateWebhook("https://ofac.example.com/callback")
	if err != nil {
		t.Error(err)
	}
	if out != "https://ofac.example.com/callback" {
		t.Errorf("got out=%q", out)
	}

	// HTTP endpoint
	out, err = validateWebhook("http://bad.example.com/callback")
	if err == nil {
		t.Error("expected error, but got none")
	}
	if out != "" {
		t.Errorf("out=%q", out)
	}
}
