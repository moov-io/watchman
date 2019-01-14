// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package k8s

import (
	"os"
)

// Inside returns true if ran from inside a Kubernetes cluster.
func Inside() bool {
	// kubernetes service account filepath (on default config).
	// See https://stackoverflow.com/a/49045575
	path := "/var/run/secrets/kubernetes.io"
	if v := os.Getenv("KUBERNETES_SERVICE_ACCOUNT_FILEPATH"); v != "" {
		path = v
	}
	_, err := os.Stat(path)
	return err == nil // file exists
}
