// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	moov "github.com/moov-io/watchman/client"

	"github.com/antihax/optional"
)

func getUIValues(ctx context.Context, api *moov.APIClient) error {
	req := &moov.GetUIValuesOpts{
		Limit: optional.NewInt32(10),
	}
	values, resp, err := api.WatchmanApi.GetUIValues(ctx, "sdnType", req)
	if err != nil {
		return fmt.Errorf("getUIValues: %v", err)
	}
	resp.Body.Close()

	if len(values) == 0 {
		return fmt.Errorf("no values found")
	}

	return nil
}
