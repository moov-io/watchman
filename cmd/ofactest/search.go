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

	// Load SDN
	if err := loadSDN(ctx, api, search.SDNs[0].EntityID); err != nil {
		return err
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

	// Load AltNames
	if err := loadAltNames(ctx, api, search.AltNames[0].EntityID); err != nil {
		return err
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

	// Load Address
	if err := loadAddresses(ctx, api, search.Addresses[0].EntityID); err != nil {
		return err
	}

	return nil
}

func loadSDN(ctx context.Context, api *moov.APIClient, id string) error {
	sdn, resp, err := api.OFACApi.GetSDN(ctx, id, nil)
	if err != nil {
		return fmt.Errorf("loadSDN: %v", err)
	}
	defer resp.Body.Close()
	if sdn.EntityID != id {
		return fmt.Errorf("loadSDN: wrong SDN: expected %s but got %s", id, sdn.EntityID)
	}
	return nil
}

func loadAddresses(ctx context.Context, api *moov.APIClient, id string) error {
	addr, resp, err := api.OFACApi.GetSDNAddresses(ctx, id, nil)
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

func loadAltNames(ctx context.Context, api *moov.APIClient, id string) error {
	alt, resp, err := api.OFACApi.GetSDNAltNames(ctx, id, nil)
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

func loadCustomer(ctx context.Context, api *moov.APIClient, id string) error {
	cust, resp, err := api.OFACApi.GetCustomer(ctx, id, nil)
	if err != nil {
		return fmt.Errorf("loadCustomer: %v", err)
	}
	defer resp.Body.Close()
	if cust.Id != id {
		return fmt.Errorf("loadCustomer: wrong Customer: expected %s but got %s", id, cust.Id)
	}
	return nil
}
