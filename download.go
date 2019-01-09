// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	ofacFilenames = []string{
		"add.csv",          // Address
		"alt.csv",          // Alternate ID
		"sdn.csv",          // Specially Designated National
		"sdn_comments.csv", // Specially Designated National Comments
	}
	ofacURLTemplate = "https://www.treasury.gov/ofac/downloads/%s"
)

func init() {
	if v := os.Getenv("OFAC_DOWNLOAD_TEMPLATE"); v != "" {
		ofacURLTemplate = v
	}
}

// ofacDownloader will download and cache OFAC files in a temp directory.
//
// If HTTP is nil then http.DefaultClient will be used (which has NO timeouts).
//
// See: https://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/sdn_data.aspx
type ofacDownloader struct {
	HTTP *http.Client
}

// getFiles will download all OFAC related files and store them in a temporary directory
// returned and an error otherwise.
//
// Callers are expected to cleanup the temp directory.
func (dl *ofacDownloader) getFiles() (string, error) {
	if dl.HTTP == nil {
		dl.HTTP = http.DefaultClient
	}

	dir, err := ioutil.TempDir("", "ofac-downloader")
	if err != nil {
		return "", fmt.Errorf("OFAC: unable to make temp dir: %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ofacFilenames))

	for i := range ofacFilenames {
		name := ofacFilenames[i]

		go func(wg *sync.WaitGroup, filename string) {
			defer wg.Done()

			resp, err := dl.HTTP.Get(fmt.Sprintf(ofacURLTemplate, filename))
			if err != nil {
				return
			}
			defer resp.Body.Close()

			// Copy resp.Body into a file in our temp dir
			fd, err := os.Create(filepath.Join(dir, filename))
			if err != nil {
				return
			}
			defer fd.Close()

			io.Copy(fd, resp.Body) // copy contents
		}(&wg, name)
	}

	wg.Wait()

	// count files and error if the count isn't what we expected
	fds, err := ioutil.ReadDir(dir)
	if err != nil || len(fds) != len(ofacFilenames) {
		return "", fmt.Errorf("OFAC: problem downloading (found=%d, expected=%d): err=%v", len(fds), len(ofacFilenames), err)
	}

	return dir, nil
}
