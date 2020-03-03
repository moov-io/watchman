// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package download

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/watchman"

	"github.com/go-kit/kit/log"
)

var (
	HTTPClient = &http.Client{
		Timeout: 15 * time.Second,
	}
)

func New(logger log.Logger, httpClient *http.Client) *Downloader {
	return &Downloader{
		HTTP:   httpClient,
		Logger: logger,
	}
}

// Downloader will download and cache DPL files in a temp directory.
//
// If HTTP is nil then http.DefaultClient will be used (which has NO timeouts).
//
// See: https://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/sdn_data.aspx
type Downloader struct {
	HTTP   *http.Client
	Logger log.Logger
}

// GetFiles will download all provided files, return their filepaths, and store them in a
// temporary directory and an error otherwise.
//
// initialDir is an optional filepath to look for files in before attempting to download.
//
// Callers are expected to cleanup the temp directory.
func (dl *Downloader) GetFiles(initialDir string, namesAndSources map[string]string) ([]string, error) {
	if dl == nil {
		return nil, errors.New("nil Downloader")
	}
	if dl.HTTP == nil {
		dl.HTTP = http.DefaultClient
	}
	if dl.Logger == nil {
		dl.Logger = log.NewNopLogger()
	}

	// Create a temporary directory for downloads
	dir, err := ioutil.TempDir("", "downloader")
	if err != nil {
		return nil, fmt.Errorf("downloader: unable to make temp dir: %v", err)
	}

	// Check the initial directory for files we don't need to download
	if initialDir == "" {
		initialDir = dir // empty, but use it as a directory
	}
	localFiles, err := ioutil.ReadDir(initialDir)
	if err != nil {
		return nil, fmt.Errorf("readdir %s: %v", initialDir, err)
	}

	var wg sync.WaitGroup
	wg.Add(len(namesAndSources))
	for name, source := range namesAndSources {
		go func(wg *sync.WaitGroup, filename, downloadURL string) {
			defer wg.Done()

			// Check if we have the file locally first
			for i := range localFiles {
				if strings.EqualFold(filepath.Base(localFiles[i].Name()), filename) {
					in, err := os.Open(filepath.Join(initialDir, localFiles[i].Name()))
					if err != nil {
						dl.Logger.Log("download", fmt.Errorf("problem opening local file %s: %v", localFiles[i].Name(), err))
						return
					}
					// Copy the local file to our output directory
					out, err := os.Create(filepath.Join(dir, filename))
					if err != nil {
						dl.Logger.Log("download", fmt.Errorf("problem creating out file (from local) %s: %v", filename, err))
						in.Close()
						return
					}
					if n, err := io.Copy(out, in); err != nil { // copy file contents
						dl.Logger.Log("download", fmt.Errorf("copied (n=%d) from local file %s: %v", n, filename, err))
						return
					}

					in.Close()
					out.Close()

					return // quit as we've copied instead of downloading
				}
			}

			// Allow a couple retries for various sources (some are flakey)
			for i := 0; i < 3; i++ {
				req, err := http.NewRequest("GET", downloadURL, nil)
				if err != nil {
					dl.Logger.Log("download", fmt.Sprintf("error building HTTP request: %v", err))
					return
				}
				req.Header.Set("User-Agent", fmt.Sprintf("moov-io/watchman:%v", watchman.Version))

				resp, err := dl.HTTP.Do(req)
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
		return nil, fmt.Errorf("DPL: problem downloading (matched=%v missing=%v): err=%v", matched, missing, err)
	}
	var out []string
	for i := range fds {
		out = append(out, filepath.Join(dir, filepath.Base(fds[i].Name())))
	}
	return out, nil
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
