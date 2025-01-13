// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package download

import (
	"context"
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
		Timeout: 45 * time.Second,
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

// GetFiles will initiate download of all provided files, return an io.ReadCloser to their content
//
// initialDir is an optional filepath to look for files in before attempting to download.
//
// Callers are expected to call the io.Closer interface method when they are done with the file
func (dl *Downloader) GetFiles(ctx context.Context, dir string, namesAndSources map[string]string) (map[string]io.ReadCloser, error) {
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
	// do not treat an nonexisting directory as error
	localFiles, _ := os.ReadDir(dir)

	var mu sync.Mutex
	out := make(map[string]io.ReadCloser)
	var wg sync.WaitGroup
	wg.Add(len(namesAndSources))

findfiles:
	for name, source := range namesAndSources {
		// Check if we have the file locally first
		for _, file := range localFiles {
			if strings.EqualFold(filepath.Base(file.Name()), name) {
				fn := filepath.Join(dir, file.Name())
				fd, err := os.Open(fn)
				if err != nil {
					dl.Logger.Error().LogErrorf("could not read file from %v initialDir: %v", fn, err)
					fd.Close()
					continue
				}
				mu.Lock()
				out[name] = fd
				mu.Unlock()
				// file is found, skip downloading
				wg.Done()
				continue findfiles
			}
		}

		// Download missing files
		go func(wg *sync.WaitGroup, filename, downloadURL string) {
			defer wg.Done()

			logger := dl.createLogger(filename, downloadURL)

			startTime := time.Now().In(time.UTC)
			content, err := dl.retryDownload(ctx, downloadURL)
			dur := time.Now().In(time.UTC).Sub(startTime)

			if err != nil {
				logger.Error().LogErrorf("FAILURE after %v to download: %v", dur, err)
				return
			}

			logger.Info().Logf("successful download after %v", dur)
			mu.Lock()
			out[filename] = content
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

func (dl *Downloader) retryDownload(ctx context.Context, downloadURL string) (io.ReadCloser, error) {
	// Allow a couple retries for various sources (some are flakey)
	for i := 0; i < 3; i++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
		if err != nil {
			return nil, dl.Logger.Error().LogErrorf("error building HTTP request: %v", err).Err()
		}
		req.Header.Set("User-Agent", fmt.Sprintf("moov-io/watchman:%v", watchman.Version))
		// in order to get passed europes 406 (Not Accepted)
		req.Header.Set("accept-language", "en-US,en;q=0.9")

		resp, err := dl.HTTP.Do(req)

		if err != nil {
			dl.Logger.Error().LogErrorf("err while doing client request: %v", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			resp.Body.Close()
			continue
		}
		return resp.Body, nil
	}
	return nil, errors.New("error max retries reached while trying to obtain file")
}
