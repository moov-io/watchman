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
	"strings"
	"sync"
	"time"
)

var (
	ofacFilenames = []string{
		"add.csv",          // Address
		"alt.csv",          // Alternate ID
		"sdn.csv",          // Specially Designated National
		"sdn_comments.csv", // Specially Designated National Comments
	}
	ofacURLTemplate = "https://www.treasury.gov/ofac/downloads/%s"

	dplFilenames = []string{
		"dpl.txt", // Denied Persons List (tab separated)
	}
	dplURLTemplate = "https://www.bis.doc.gov/dpl/%s"
)

func init() {
	if v := os.Getenv("OFAC_DOWNLOAD_TEMPLATE"); v != "" {
		ofacURLTemplate = v
	}
	if w := os.Getenv("DPL_DOWNLOAD_TEMPLATE"); w != "" {
		dplURLTemplate = w
	}
}

// Downloader will download and cache OFAC files in a temp directory.
//
// If HTTP is nil then http.DefaultClient will be used (which has NO timeouts).
//
// See: https://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/sdn_data.aspx
type Downloader struct {
	HTTP *http.Client
}

// GetFiles will download all OFAC related files and store them in a temporary directory
// returned and an error otherwise.
//
// Callers are expected to cleanup the temp directory.
func (dl *Downloader) GetFiles() (string, error) {
	if dl.HTTP == nil {
		dl.HTTP = http.DefaultClient
	}

	dir, err := ioutil.TempDir("", "ofac-and-dpl-downloader")
	if err != nil {
		return "", fmt.Errorf("OFAC: unable to make temp dir: %v", err)
	}

	// create a single list containing all filenames and source URLs
	namesAndSources := make(map[string]string)
	for _, fname := range ofacFilenames {
		namesAndSources[fname] = fmt.Sprintf(ofacURLTemplate, fname)
	}
	for _, fname := range dplFilenames {
		namesAndSources[fname] = fmt.Sprintf(dplURLTemplate, fname)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(namesAndSources))
	for name, source := range namesAndSources {
		go func(wg *sync.WaitGroup, filename, downloadURL string) {
			defer wg.Done()

			// Allow a couple retries for various sources (some are flakey)
			for i := 0; i < 3; i++ {
				resp, err := dl.HTTP.Get(downloadURL)
				if err != nil {
					time.Sleep(100 * time.Millisecond)
					continue // retry
				}
				// Copy resp.Body into a file in our temp dir
				fd, err := os.Create(filepath.Join(dir, filename))
				if err != nil {
					resp.Body.Close()
					return
				}
				io.Copy(fd, resp.Body) // copy file contents
				// close the open files
				fd.Close()
				resp.Body.Close()
				return // quit after successful download
			}
		}(&wg, name, source)
	}

	wg.Wait()

	// count files and error if the count isn't what we expected
	fds, err := ioutil.ReadDir(dir)
	if err != nil || len(fds) != len(namesAndSources) {
		matched, missing := compareNames(fds, namesAndSources)
		return "", fmt.Errorf("OFAC: problem downloading (matched=%v missing=%v): err=%v", matched, missing, err)
	}

	return dir, nil
}

func compareNames(found []os.FileInfo, expected map[string]string) (string, string) {
	var matched []string
	var missing []string

	for k := range expected {
		wasFound := false
		for i := range found {
			filename := filepath.Base(found[i].Name())
			if k == filename {
				wasFound = true
				matched = append(matched, filename)
				break
			}
		}
		if !wasFound {
			missing = append(missing, k)
		}
	}
	for i := range found {
		filename := filepath.Base(found[i].Name())
		if _, exists := expected[filename]; !exists {
			missing = append(missing, filename)
		}
	}

	return strings.Join(matched, ", "), strings.Join(missing, ", ")
}
