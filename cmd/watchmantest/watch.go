// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/moov-io/base"
	moov "github.com/moov-io/watchman/client"

	"github.com/antihax/optional"
)

func addCompanyWatch(ctx context.Context, api *moov.APIClient, id string, webhook string) error {
	// add watch
	req := moov.OfacWatchRequest{
		AuthToken: base.ID(),
		Webhook:   webhook,
	}
	opts := &moov.AddOfacCompanyWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
	}
	watch, resp, err := api.WatchmanApi.AddOfacCompanyWatch(ctx, id, req, opts)
	if err != nil {
		return fmt.Errorf("addCompanyWatch: %v", err)
	}
	defer resp.Body.Close()

	// remove watch
	resp, err = api.WatchmanApi.RemoveOfacCompanyWatch(ctx, id, watch.WatchID, &moov.RemoveOfacCompanyWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
	})
	if err != nil {
		return fmt.Errorf("addCompanyWatch: remove: %v", err)
	}
	resp.Body.Close()
	return nil
}

func addCustomerWatch(ctx context.Context, api *moov.APIClient, id string, webhook string) error {
	// add watch
	req := moov.OfacWatchRequest{
		AuthToken: base.ID(),
		Webhook:   webhook,
	}
	opts := &moov.AddOfacCustomerWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
	}
	watch, resp, err := api.WatchmanApi.AddOfacCustomerWatch(ctx, id, req, opts)
	if err != nil {
		return fmt.Errorf("addCustomerWatch: add: %v", err)
	}
	resp.Body.Close()

	// remove watch
	resp, err = api.WatchmanApi.RemoveOfacCustomerWatch(ctx, id, watch.WatchID, &moov.RemoveOfacCustomerWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
	})
	if err != nil {
		return fmt.Errorf("addCustomerWatch: remove: %v", err)
	}
	resp.Body.Close()
	return nil
}
