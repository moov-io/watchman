// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/moov-io/ofac"
	moov "github.com/moov-io/ofac/client"
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

	// Setup OFAC API client
	api := moov.NewAPIClient(conf)

	// Ping OFAC
	resp, err := api.OFACApi.Ping(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("[FAILURE] ping error (stats code: %d): %v", resp.StatusCode, err)
	} else {
		log.Println("[SUCCESS] ping OFAC")
	}
}
