// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// watchmantest is a cli tool used for testing the Moov Sanction Search service.
//
// With no arguments the contaier runs tests against the production API.
// This tool requires an OAuth token provided by github.com/moov-io/api written
// to the local disk, but running apitest first will write this token.
//
// This tool can be used to query with custom searches:
//  $ go install ./cmd/watchmantest
//  $ watchmantest -local moh
//  2019/02/14 23:37:44.432334 main.go:44: Starting moov/watchmantest v0.4.1-dev
//  2019/02/14 23:37:44.432366 main.go:60: [INFO] using http://localhost:8084 for address
//  2019/02/14 23:37:44.434534 main.go:76: [SUCCESS] ping
//  2019/02/14 23:37:44.435204 main.go:83: [SUCCESS] last download was: 3h45m58s ago
//  2019/02/14 23:37:44.440230 main.go:96: [SUCCESS] name search passed, query="moh"
//  2019/02/14 23:37:44.441506 main.go:104: [SUCCESS] added customer=24032 watch
//  2019/02/14 23:37:44.445473 main.go:118: [SUCCESS] alt name search passed
//  2019/02/14 23:37:44.449367 main.go:123: [SUCCESS] address search passed
//
// watchmantest is not a stable tool. Please contact Moov developers if you intend to use this tool,
// otherwise we might change the tool (or remove it) without notice.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/moov-io/watchman"
	moov "github.com/moov-io/watchman/client"
	"github.com/moov-io/watchman/cmd/internal"

	"github.com/antihax/optional"
)

var (
	flagApiAddress = flag.String("address", internal.DefaultApiAddress, "Moov API address")
	flagLocal      = flag.Bool("local", false, "Use local HTTP addresses")
	flagWebhook    = flag.String("webhook", "https://moov.io/watchman", "Secure HTTP address for webhooks")

	flagRequestID = flag.String("request-id", "", "Override what is set for the X-Request-ID HTTP header")
)

func main() {
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Printf("Starting moov/watchmantest %s", watchman.Version)

	conf := internal.Config(*flagApiAddress, *flagLocal)
	log.Printf("[INFO] using %s for address", conf.BasePath)

	// Read OAuth token and set on conf
	if v := os.Getenv("OAUTH_TOKEN"); v != "" {
		conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", v))
	} else {
		if local := *flagLocal; !local {
			log.Fatal("[FAILURE] no OAuth token provided (try adding -local for http://localhost requests)")
		}
	}

	// Setup API client
	api, ctx := moov.NewAPIClient(conf), context.TODO()

	// Ping
	if err := ping(ctx, api); err != nil {
		log.Fatal("[FAILURE] ping Sanction Search")
	} else {
		log.Println("[SUCCESS] ping")
	}

	// Check downloads
	if when, err := latestDownload(ctx, api); err != nil || when.IsZero() {
		log.Fatalf("[FAILURE] downloads: %v", err)
	} else {
		log.Printf("[SUCCESS] last download was: %v ago", time.Since(when).Truncate(1*time.Second))
	}

	query := "alh" // string that matches a lot of records
	if v := flag.Arg(0); v != "" {
		query = v
	}

	// Search queries
	sdn, err := searchByName(ctx, api, query)
	if err != nil {
		log.Fatalf("[FAILURE] problem searching SDNs: %v", err)
	} else {
		log.Printf("[SUCCESS] name search passed, query=%q", query)
	}

	// Add watch on the SDN
	if sdn.SdnType == moov.SDNTYPE_INDIVIDUAL {
		if err := addCustomerWatch(ctx, api, sdn.EntityID, *flagWebhook); err != nil {
			log.Fatalf("[FAILURE] problem adding customer watch: %v", err)
		} else {
			log.Printf("[SUCCESS] added customer=%s watch", sdn.EntityID)
		}
	} else {
		if err := addCompanyWatch(ctx, api, sdn.EntityID, *flagWebhook); err != nil {
			log.Fatalf("[FAILURE] problem adding company watch: %v", err)
		} else {
			log.Printf("[SUCCESS] added company=%s watch", sdn.EntityID)
		}
	}

	// Load alt names and addresses
	if err := searchByAltName(ctx, api, query); err != nil {
		log.Fatalf("[FAILURE] problem searching Alt Names: %v", err)
	} else {
		log.Println("[SUCCESS] alt name search passed")
	}
	if err := searchByAddress(ctx, api, "St"); err != nil {
		log.Fatalf("[FAILURE] problem searching addresses: %v", err)
	} else {
		log.Println("[SUCCESS] address search passed")
	}

	// Lookup UI values
	if err := getUIValues(ctx, api); err != nil {
		log.Fatalf("[FAILURE] problem looking up UI values: %v", err)
	} else {
		log.Println("[SUCCESS] UI values lookup passed")
	}
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

func latestDownload(ctx context.Context, api *moov.APIClient) (time.Time, error) {
	opts := &moov.GetLatestDownloadsOpts{
		Limit:      optional.NewInt32(1),
		XRequestID: optional.NewString(*flagRequestID),
	}
	downloads, resp, err := api.WatchmanApi.GetLatestDownloads(ctx, opts)
	if err != nil {
		return time.Time{}, err
	}
	resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return time.Time{}, fmt.Errorf("download error (stats code: %d): %v", resp.StatusCode, err)
	}
	if len(downloads) == 0 {
		return time.Time{}, errors.New("empty downloads response")
	}
	return downloads[0].Timestamp, nil
}
