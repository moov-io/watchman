// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/moov-io/ofac"
)

func main() {
	fmt.Printf("starting moov-io/ofac version %s\n", ofac.Version)

	if !inKubernetes() {
		es, err := ofac.NewElasticsearch()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
		defer es.Stop()

		fmt.Printf("ES is up and running! (Docker container ID: %s)\n", es.ID())
	}
}

func inKubernetes() bool {
	// kubernetes service account filepath (on default config). See https://stackoverflow.com/a/49045575
	_, err := os.Stat("/var/run/secrets/kubernetes.io")
	return err == nil // file exists
}
