// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/moov-io/ofac"
	moov "github.com/moov-io/ofac/client"

	"github.com/antihax/optional"
)

var (
	defaultApiAddress = "https://api.moov.io"

	flagApiAddress = flag.String("address", defaultApiAddress, "Moov API address")
)

func main() {
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Printf("Starting moov/ofactest %s", ofac.Version)

	conf := moov.NewConfiguration()
	if v := *flagApiAddress; v == defaultApiAddress {
		conf.BasePath = v + "/v1/ofac"
	} else {
		conf.BasePath = v
	}
	conf.UserAgent = fmt.Sprintf("moov/ofactest:%s", ofac.Version)
	conf.HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 1 * time.Minute,
		},
	}

	log.Printf("[INFO] using %s for address", conf.BasePath)

	// Read OAuth token and set on conf
	accessToken := GetOAuthToken(OAuthTokenStorageFilepath)
	if accessToken == "" {
		log.Fatal("[FAILURE] no OAuth token found, run moov/apitest locally")
	} else {
		conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	// Setup OFAC API client
	api, ctx := moov.NewAPIClient(conf), context.TODO()

	// Ping OFAC
	if err := ping(ctx, api); err != nil {
		log.Fatal("[FAILURE] ping OFAC")
	} else {
		log.Println("[SUCCESS] ping")
	}

	// Check downloads
	if when, err := latestDownload(ctx, api); err != nil || when.IsZero() {
		log.Fatalf("[FAILURE] downloads: %v", err)
	} else {
		log.Printf("[SUCCESS] last download was: %v ago", time.Since(when).Truncate(1*time.Second))
	}

	// Search queries
	if err := searchSDNs(ctx, api); err != nil {
		log.Fatalf("[FAILURE] %v", err)
	} else {
		log.Println("[SUCCESS] search queries passed")
	}
}

func ping(ctx context.Context, api *moov.APIClient) error {
	resp, err := api.OFACApi.Ping(ctx)
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
	downloads, resp, err := api.OFACApi.GetLatestDownloads(ctx, &moov.GetLatestDownloadsOpts{
		Limit: optional.NewInt32(1),
	})
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
