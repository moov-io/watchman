// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package moovoauth

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// TokenFilepath returns a local filesystem path for the Moov API OAuth token.
// This is written by the apitest binary (or inside its docker image) when executed.
func TokenFilepath() string {
	filename := "access-token"

	// If we're ran as user 'moov' write to current dir
	if u, _ := user.Current(); u != nil {
		if u.Username == "moov" || u.Name == "moov" {
			dir, _ := os.Getwd()
			return filepath.Join(dir, filename)
		}
	}

	switch runtime.GOOS {
	case "darwin", "linux":
		return filepath.Join(os.Getenv("HOME"), ".moov", "api", "oauth", filename)
	case "windows":
		// return filepath.Join(os.Getenv("USERPROFILE"), filename) // TODO(adam): better path
	}
	return ""
}
