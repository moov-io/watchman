// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"fmt"

	moov "github.com/moov-io/ofac/client"

	"github.com/antihax/optional"
)

func searchSDNs(ctx context.Context, api *moov.APIClient) error {
	// Name
	search, resp, err := api.OFACApi.SearchSDNs(ctx, &moov.SearchSDNsOpts{
		Name:  optional.NewString("alh"),
		Limit: optional.NewInt32(5),
	})
	if err != nil {
		return fmt.Errorf("searchSDNs: name search: %v", err)
	}
	defer resp.Body.Close()
	if len(search.SDNs) == 0 {
		return errors.New("searchSDNs: no SDNs found")
	}

	// Alt Name
	search, resp, err = api.OFACApi.SearchSDNs(ctx, &moov.SearchSDNsOpts{
		AltName: optional.NewString("alh"),
		Limit:   optional.NewInt32(5),
	})
	if err != nil {
		return fmt.Errorf("searchSDNs: AltName search: %v", err)
	}
	defer resp.Body.Close()
	if len(search.AltNames) == 0 {
		return errors.New("searchSDNs: no alt names found")
	}

	// Address
	search, resp, err = api.OFACApi.SearchSDNs(ctx, &moov.SearchSDNsOpts{
		Address: optional.NewString("St"),
		Limit:   optional.NewInt32(5),
	})
	if err != nil {
		return fmt.Errorf("searchSDNs: Address search: %v", err)
	}
	defer resp.Body.Close()
	if len(search.Addresses) == 0 {
		return errors.New("searchSDNs: no addresses found")
	}

	return nil
}
