// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/base/admin"
	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

func TestDebug__SDN(t *testing.T) {
	svc, err := admin.New(admin.Opts{
		Addr: ":0",
	})
	require.NoError(t, err)

	go svc.Listen()
	defer svc.Shutdown()

	svc.AddHandler(debugSDNPath, debugSDNHandler(log.NewNopLogger(), idSearcher))

	req, _ := http.NewRequest("GET", "http://"+svc.BindAddr()+"/debug/sdn/22790", nil)
	req.Header.Set("x-request-id", base.ID())

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bs, _ := io.ReadAll(resp.Body)
		t.Fatalf("bogus status code: %s: %s", resp.Status, string(bs))
	}

	var response struct {
		Debug struct {
			IndexedName     string `json:"indexedName"`
			ParsedRemarksID string `json:"parsedRemarksId"`
		} `json:"debug"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)

	require.Equal(t, "nicolas maduro moros", response.Debug.IndexedName)
	require.Equal(t, "5892464", response.Debug.ParsedRemarksID)
}
