// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/moov-io/base"
	moov "github.com/moov-io/ofac/client"

	"github.com/antihax/optional"
)

func addCompanyWatch(ctx context.Context, api *moov.APIClient, id string, webhook string) error {
	// add watch
	req := moov.WatchRequest{
		AuthToken: base.ID(),
		Webhook:   webhook,
	}
	opts := &moov.AddOFACCompanyWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
		XUserID:    optional.NewString(*flagUserID),
	}
	watch, resp, err := api.OFACApi.AddOFACCompanyWatch(ctx, id, req, opts)
	if err != nil {
		return fmt.Errorf("addCompanyWatch: %v", err)
	}
	defer resp.Body.Close()

	// remove watch
	resp, err = api.OFACApi.RemoveOFACCompanyWatch(ctx, id, watch.WatchID, &moov.RemoveOFACCompanyWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
		XUserID:    optional.NewString(*flagUserID),
	})
	if err != nil {
		return fmt.Errorf("addCompanyWatch: remove: %v", err)
	}
	resp.Body.Close()
	return nil
}

func addCustomerWatch(ctx context.Context, api *moov.APIClient, id string, webhook string) error {
	// add watch
	req := moov.WatchRequest{
		AuthToken: base.ID(),
		Webhook:   webhook,
	}
	opts := &moov.AddOFACCustomerWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
		XUserID:    optional.NewString(*flagUserID),
	}
	watch, resp, err := api.OFACApi.AddOFACCustomerWatch(ctx, id, req, opts)
	if err != nil {
		return fmt.Errorf("addCustomerWatch: add: %v", err)
	}
	resp.Body.Close()

	// remove watch
	resp, err = api.OFACApi.RemoveOFACCustomerWatch(ctx, id, watch.WatchID, &moov.RemoveOFACCustomerWatchOpts{
		XRequestID: optional.NewString(*flagRequestID),
		XUserID:    optional.NewString(*flagUserID),
	})
	if err != nil {
		return fmt.Errorf("addCustomerWatch: remove: %v", err)
	}
	resp.Body.Close()
	return nil
}
