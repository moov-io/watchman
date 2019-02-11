// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"

	"github.com/moov-io/api/pkg/moovoauth"
)

// getOAuthToken will read a Moov API OAuth token from disk and return it
//
// See github.com/moov-io/api's cmd/apitest for OAuth token writing to disk
func getOAuthToken() string {
	path := moovoauth.TokenFilepath()
	if path == "" {
		return ""
	}

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return string(bytes.TrimSpace(bs))
}
