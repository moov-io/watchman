// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// batchsearch is a cli tool used for testing batches of names against Moov's sanctions service.
//
// With no arguments the contaier runs tests against the production API, but we strongly ask you
// run batchsearch against local instances of Watchman.
//
//  $ go install ./cmd/batchsearch
//  $ batchsearch -allowed-file users.txt -blocked-file criminals.txt -threshold 0.99 -sdn-type individual -v
//  2019/10/09 17:36:16.160025 main.go:61: Starting moov/batchsearch v0.12.0-dev
//  2019/10/09 17:36:16.160055 main.go:64: [INFO] using http://localhost:8084 for address
//  2019/10/09 17:36:16.161818 main.go:73: [SUCCESS] ping
//  2019/10/09 17:36:16.174108 main.go:156: [INFO] didn't block 'Husein HAZEM'
//  2019/10/09 17:36:16.212986 main.go:148: [INFO] blocked 'Nicolas Ernesto MADURO GUERRA'
//  2019/10/09 17:36:16.213423 main.go:146: [ERROR] 'Maria Alexandra PERDOMO' wasn't blocked (match=0.96)
//  exit status 1
//
// batchsearch is not a stable tool. Please contact Moov developers if you intend to use this tool,
// otherwise we might change the tool (or remove it) without notice.
package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/moov-io/watchman"
	moov "github.com/moov-io/watchman/client"
	"github.com/moov-io/watchman/cmd/internal"

	"github.com/antihax/optional"
	"go4.org/syncutil"
)

var (
	flagApiAddress = flag.String("address", internal.DefaultApiAddress, "Moov API address")
	flagLocal      = flag.Bool("local", false, "Use local HTTP addresses")

	flagThreshold = flag.Float64("threshold", 0.99, "Minimum match percentage required for blocking")

	flagAllowedFile = flag.String("allowed-file", filepath.Join("test", "testdata", "allowed.txt"), "Filepath to file with names expected to be allowed")
	flagBlockedFile = flag.String("blocked-file", filepath.Join("test", "testdata", "blocked.txt"), "Filepath to file with names expected to be blocked")

	flagSdnType = flag.String("sdn-type", "individual", "sdnType query param")

	flagRequestID = flag.String("request-id", "", "Override what is set for the X-Request-ID HTTP header")

	flagVerbose = flag.Bool("v", false, "Enable detailed logging")

	flagWorkers = flag.Int("workers", runtime.NumCPU(), "How many tasks to run concurrently")
)

func main() {
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Printf("Starting moov/batchsearch %s", watchman.Version)

	conf := internal.Config(*flagApiAddress, *flagLocal)
	log.Printf("[INFO] using %s for address", conf.BasePath)

	// Setup API client
	api, ctx := moov.NewAPIClient(conf), context.TODO()

	// Ping
	if err := ping(ctx, api); err != nil {
		log.Fatalf("[FAILURE] ping Sanctions Search: %v", err)
	} else {
		log.Println("[SUCCESS] ping")
	}

	// Perform checks over the two incoming files
	if path := *flagAllowedFile; path != "" {
		names, err := readNames(path)
		if err != nil {
			log.Fatalf("[FAILURE] %v", err)
		}
		if n := checkNames(BlockUnexpected, names, *flagThreshold, api); n == Failure {
			os.Exit(int(n))
		}
	}
	if path := *flagBlockedFile; path != "" {
		names, err := readNames(path)
		if err != nil {
			log.Fatalf("[FAILURE] %v", err)
		}
		if n := checkNames(BlockExpected, names, *flagThreshold, api); n == Failure {
			os.Exit(int(n))
		}
	}
	log.Println("[SUCCESS] all tests passed")
}

func ping(ctx context.Context, api *moov.APIClient) error {
	resp, err := api.WatchmanApi.Ping(ctx)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("ping error (stats code: %d): %v", resp.StatusCode, err)
	}
	return nil
}

type action int

var (
	BlockExpected   action = 1
	BlockUnexpected action = 2
)

var (
	Success int64 = 0
	Failure int64 = 1
)

func checkNames(desc action, names []string, threshold float64, api *moov.APIClient) int64 {
	var wg sync.WaitGroup
	wg.Add(len(names))

	var exitCode int64 // must be protected with atomic calls
	markFailure := func() {
		atomic.CompareAndSwapInt64(&exitCode, Success, Failure) // set Failure as exit code
	}

	workers := syncutil.NewGate(*flagWorkers)

	for i := range names {
		workers.Start()
		go func(name string) {
			defer workers.Done()
			defer wg.Done()

			if match, err := searchByName(api, name); err != nil {
				markFailure()
				log.Printf("[FATAL] problem searching for '%s': %v", name, err)
			} else {
				switch desc {
				case BlockExpected:
					if match < threshold {
						markFailure()
						log.Printf("[ERROR] '%s' wasn't blocked (match=%.2f)", name, match)
					} else {
						log.Printf("[INFO] blocked '%s'", name)
					}
				case BlockUnexpected:
					if match > threshold {
						markFailure()
						log.Printf("[ERROR] '%s' was blocked (match=%.2f)", name, match)
					} else {
						if *flagVerbose {
							log.Printf("[INFO] didn't block '%s'", name)
						}
					}
				}
			}
		}(names[i])
	}

	wg.Wait() // block until all requests are finished

	return exitCode
}

func readNames(path string) ([]string, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("problem reading %s: %v", path, err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var names []string
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(name, "//") || strings.HasPrefix(name, "#") {
			continue
		}
		names = append(names, name)
	}
	return names, nil
}

func searchByName(api *moov.APIClient, name string) (float64, error) {
	opts := &moov.SearchOpts{
		Limit:      optional.NewInt32(1),
		Name:       optional.NewString(name),
		SdnType:    optional.NewInterface(*flagSdnType),
		XRequestID: optional.NewString(*flagRequestID),
	}

	ctx, cancelFunc := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancelFunc()

	search, resp, err := api.WatchmanApi.Search(ctx, opts)
	if err != nil {
		return 0.0, fmt.Errorf("searchByName: %v", err)
	}
	defer resp.Body.Close()

	if len(search.SDNs) == 0 {
		return 0.0, errors.New("no SDNs returned")
	}
	return float64(search.SDNs[0].Match), nil
}
