// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestHTTP__cleanMetricsPath(t *testing.T) {
	if v := cleanMetricsPath("/v1/ofac/companies/1234"); v != "v1-ofac-companies" {
		t.Errorf("got %q", v)
	}
	if v := cleanMetricsPath("/v1/ofac/ping"); v != "v1-ofac-ping" {
		t.Errorf("got %q", v)
	}
}
