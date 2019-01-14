// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"fmt"
	"os/exec"
	"runtime"
	"testing"

	"github.com/moov-io/base/docker"
)

func TestElasticsearch(t *testing.T) {
	if !docker.Enabled() {
		t.Skip(fmt.Sprintf("docker wasn't found on %s", runtime.GOOS))
	}

	es, err := NewElasticsearch()
	if err != nil {
		t.Fatalf("%#v", err)
	}
	defer func() {
		if err := es.Stop(); err != nil {
			t.Fatalf("%#v", err)
		}
	}()

	// Check we can ping
	if err := <-es.ping(); err != nil {
		t.Fatalf("%#v", err)
	}
}
