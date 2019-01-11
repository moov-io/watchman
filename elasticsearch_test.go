// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"fmt"
	"os/exec"
	"runtime"
	"testing"
)

func isDockerEnabled() bool {
	bin, err := exec.LookPath("docker")
	return bin != "" && err == nil // 'docker' was found on PATH
}

func TestElasticsearch(t *testing.T) {
	if !isDockerEnabled() {
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
