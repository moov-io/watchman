// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/moov-io/base/docker"
)

func TestElasticsearch(t *testing.T) {
	if !docker.Enabled() {
		t.Skip(fmt.Sprintf("docker wasn't found on %s", runtime.GOOS))
	}
	if testing.Short() {
		t.Skip("skipping expensive Elasticsearch test in short mode")
	}

	es, err := NewElasticsearch(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := es.Stop(nil); err != nil {
			t.Fatal(err)
		}
	}()

	// Check we can ping
	if err := <-es.ping(); err != nil {
		t.Fatal(err)
	}
	if err := es.Ping(); err != nil {
		t.Fatal(err)
	}

	// Grab Docker container ID
	if id := es.ID(); id == "" {
		t.Fatal("no docker container ID found")
	}
}
