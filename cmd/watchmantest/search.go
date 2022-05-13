// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"fmt"

	moov "github.com/moov-io/watchman/client"

	"github.com/antihax/optional"
)

// searchByName will attempt sanctions searches for the provided name and then load the SDN metadata
// associated to the company/organization or individual.
func searchByName(ctx context.Context, api *moov.APIClient, name string) (*moov.OfacSdn, error) {
	opts := &moov.SearchOpts{
		Limit:      optional.NewInt32(2),
		Name:       optional.NewString(name),
		XRequestID: optional.NewString(*flagRequestID),
	}

	search, resp, err := api.WatchmanApi.Search(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("searchByName: %v", err)
	}
	defer resp.Body.Close()

	if len(search.SDNs) == 0 {
		return nil, fmt.Errorf("searchByName: found no SDNs for %q", name)
	}

	// Find Customer or company
	if search.SDNs[0].SdnType == moov.SDNTYPE_INDIVIDUAL {
		if err := getCustomer(ctx, api, search.SDNs[0].EntityID); err != nil {
			return nil, err
		}
	} else {
		if err := getCompany(ctx, api, search.SDNs[0].EntityID); err != nil {
			return nil, err
		}
	}
	return &search.SDNs[0], nil
}

// searchByAltName will attempt sanctions searches and retrieval of all alt names associated to the first result
// for the provided altName and error if none are found.
func searchByAltName(ctx context.Context, api *moov.APIClient, alt string) error {
	opts := &moov.SearchOpts{
		AltName:    optional.NewString(alt),
		Limit:      optional.NewInt32(2),
		XRequestID: optional.NewString(*flagRequestID),
	}

	search, resp, err := api.WatchmanApi.Search(ctx, opts)
	if err != nil {
		return fmt.Errorf("searchByAltName: %v", err)
	}
	defer resp.Body.Close()

	if len(search.AltNames) == 0 {
		return fmt.Errorf("searchByAltName: found no AltNames for %q", alt)
	}
	return getSDNAltNames(ctx, api, search.AltNames[0].EntityID)
}

// searchByAddress will attempt sanctions searches and retrieval of all addresses associated to the first result
// for the provided address and error if none are found.
func searchByAddress(ctx context.Context, api *moov.APIClient, address string) error {
	opts := &moov.SearchOpts{
		Address:    optional.NewString(address),
		Limit:      optional.NewInt32(2),
		XRequestID: optional.NewString(*flagRequestID),
	}

	search, resp, err := api.WatchmanApi.Search(ctx, opts)
	if err != nil {
		return fmt.Errorf("searchByAddress: %v", err)
	}
	defer resp.Body.Close()

	if len(search.Addresses) == 0 {
		return fmt.Errorf("searchByAddress: found no Addresses for %q", address)
	}
	return getSDNAddresses(ctx, api, search.Addresses[0].EntityID)
}

func getSDNAddresses(ctx context.Context, api *moov.APIClient, id string) error {
	addr, resp, err := api.WatchmanApi.GetSDNAddresses(ctx, id, nil)
	if err != nil {
		return fmt.Errorf("loadAddresses: %v", err)
	}
	defer resp.Body.Close()
	if len(addr) == 0 {
		return errors.New("loadAddresses: no Addresses found")
	}
	if addr[0].EntityID != id {
		return fmt.Errorf("loadAddresses: wrong Address: expected %s but got %s", id, addr[0].EntityID)
	}
	return nil
}

func getSDNAltNames(ctx context.Context, api *moov.APIClient, id string) error {
	alt, resp, err := api.WatchmanApi.GetSDNAltNames(ctx, id, nil)
	if err != nil {
		return fmt.Errorf("loadAltNames: %v", err)
	}
	defer resp.Body.Close()
	if len(alt) == 0 {
		return errors.New("loadAltNames: no AltNames found")
	}
	if alt[0].EntityID != id {
		return fmt.Errorf("loadAltNames: wrong AltName: expected %s but got %s", id, alt[0].EntityID)
	}
	return nil
}

func getCustomer(ctx context.Context, api *moov.APIClient, id string) error {
	cust, resp, err := api.WatchmanApi.GetOfacCustomer(ctx, id, nil)
	if err != nil {
		return fmt.Errorf("loadCustomer: %v", err)
	}
	defer resp.Body.Close()
	if cust.ID != id {
		return fmt.Errorf("loadCustomer: wrong Customer: expected %s but got %s", id, cust.ID)
	}
	return nil
}

func getCompany(ctx context.Context, api *moov.APIClient, id string) error {
	company, resp, err := api.WatchmanApi.GetOfacCompany(ctx, id, nil)
	if err != nil {
		return fmt.Errorf("loadCompany: %v", err)
	}
	defer resp.Body.Close()
	if company.ID != id {
		return fmt.Errorf("loadCompany: wrong Company: expected %s but got %s", id, company.ID)
	}
	return nil
}
