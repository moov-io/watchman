// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	OAuthTokenStorageFilepath = filepath.Join(os.Getenv("HOME"), ".moov/api/oauth/access-token") // TODO(adam): windows, os.UserHomeDir in Go 1.12
)

// GetOAuthToken will read a Moov API OAuth token from disk and return it
func GetOAuthToken(path string) string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(bytes.TrimSpace(bs))
}
