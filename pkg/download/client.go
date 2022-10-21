// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman"
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

	// Check the initial directory for files we don't need to download
	var dir string
	if initialDir != "" {
		dir = initialDir // empty, but use it as a directory
	}
	// Create a temporary directory for downloads if needed
	if dir == "" {
		temp, err := os.MkdirTemp("", "downloader")
		if err != nil {
			return nil, fmt.Errorf("downloader: unable to make temp dir: %v", err)
		}
		dir = temp
	}

	localFiles, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("readdir %s: %v", dir, err)
	}

	var mu sync.Mutex
	var out []string

	var wg sync.WaitGroup
	wg.Add(len(namesAndSources))
	for name, source := range namesAndSources {
		// Check if we have the file locally first
		found := false
		for i := range localFiles {
			if strings.EqualFold(filepath.Base(localFiles[i].Name()), name) {
				found = true
				mu.Lock()
				out = append(out, filepath.Join(dir, localFiles[i].Name()))
				mu.Unlock()
				break
			}
		}
		// Skip downloading this file since we found it
		if found {
			wg.Done()
			continue
		}

		// Download missing files
		go func(wg *sync.WaitGroup, filename, downloadURL string) {
			defer wg.Done()

			logger := dl.createLogger(filename, downloadURL)

			startTime := time.Now().In(time.UTC)
			err := dl.retryDownload(dir, filename, downloadURL)
			dur := time.Now().In(time.UTC).Sub(startTime)

			if err != nil {
				logger.Error().LogErrorf("FAILURE after %v to download: %v", dur, err)
			} else {
				logger.Error().LogErrorf("successful download after %v", dur)
			}

			mu.Lock()
			out = append(out, filepath.Join(dir, filename))
			mu.Unlock()
		}(&wg, name, source)
	}
	wg.Wait()

	return out, nil
}

func (dl *Downloader) createLogger(filename, downloadURL string) log.Logger {
	var host string
	u, _ := url.Parse(downloadURL)
	if u != nil {
		host = u.Host
	}
	return dl.Logger.With(log.Fields{
		"host":     log.String(host),
		"filename": log.String(filename),
	})
}

func (dl *Downloader) retryDownload(dir, filename, downloadURL string) error {
	// Allow a couple retries for various sources (some are flakey)
	for i := 0; i < 3; i++ {
		req, err := http.NewRequest("GET", downloadURL, nil)
		if err != nil {
			return dl.Logger.Error().LogErrorf("error building HTTP request: %v", err).Err()
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
			return fmt.Errorf("attempt %d failed to create file: %v", i, err)
		}

		io.Copy(fd, resp.Body) // copy file contents

		// close the open files
		fd.Close()
		resp.Body.Close()
		return nil // quit after successful download
	}
	return nil
}
