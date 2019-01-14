// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package docker

import (
	"os/exec"
)

// Enabled returns true if Docker is available when called.
func Enabled() bool {
	bin, err := exec.LookPath("docker")
	return bin != "" && err == nil // 'docker' was found on PATH
}
