// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/base/admin"

	"github.com/go-kit/kit/log"
)

func TestDebug__SDN(t *testing.T) {
	svc := admin.NewServer(":0")
	go svc.Listen()
	defer svc.Shutdown()

	svc.AddHandler(debugSDNPath, debugSDNHandler(log.NewNopLogger(), idSearcher))

	req, _ := http.NewRequest("GET", "http://"+svc.BindAddr()+"/debug/sdn/22790", nil)
	req.Header.Set("x-request-id", base.ID())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bs, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("bogus status code: %s: %s", resp.Status, string(bs))
	}

	var response struct {
		Debug struct {
			IndexedName     string `json:"indexedName"`
			ParsedRemarksID string `json:"parsedRemarksId"`
		} `json:"debug"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Debug.IndexedName != "nicolas maduro moros" {
		t.Errorf("debug: indexed name: %v", response.Debug.IndexedName)
	}
	if response.Debug.ParsedRemarksID != "5892464" {
		t.Errorf("got parsed ID: %s", response.Debug.ParsedRemarksID)
	}
}
