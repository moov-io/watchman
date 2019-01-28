// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"path/filepath"

	"github.com/moov-io/ofac"
)

func (s *searcher) refreshData() error {
	if s.logger != nil {
		s.logger.Log("download", "Starting refresh of OFAC data")
	}

	// Download files
	dir, err := (&ofac.Downloader{}).GetFiles()
	if err != nil {
		return fmt.Errorf("ERROR: downloading OFAC data: %v", err)
	}

	// Parse each OFAC file
	r := &ofac.Reader{}
	r.FileName = filepath.Join(dir, "add.csv")
	if err := r.Read(); err != nil {
		return fmt.Errorf("ERROR: reading add.csv: %v", err)
	}
	r.FileName = filepath.Join(dir, "alt.csv")
	if err := r.Read(); err != nil {
		return fmt.Errorf("ERROR: reading alt.csv: %v", err)
	}
	r.FileName = filepath.Join(dir, "sdn.csv")
	if err := r.Read(); err != nil {
		return fmt.Errorf("ERROR: reading sdn.csv: %v", err)
	}

	// Precompute new data
	sdns := precomputeSDNs(r.SDNs)
	adds := precomputeAddresses(r.Addresses)
	alts := precomputeAlts(r.AlternateIdentities)

	// Set new records after precomputation (to minimize lock contention)
	s.Lock()
	s.SDNs = sdns
	s.Addresses = adds
	s.Alts = alts
	s.Unlock()

	if s.logger != nil {
		s.logger.Log("download", "Finished refresh of OFAC data")
	}

	return nil
}
