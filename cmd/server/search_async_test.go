// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"testing"
)

func TestSearchAsync_batchSize(t *testing.T) {
	if d := readWebhookBatchSize(""); d != watchResearchBatchSize {
		t.Errorf("expected watchResearchBatchSize default, but got %d", d)
	}

	if d := readWebhookBatchSize("42"); d != 42 {
		t.Errorf("expected watchResearchBatchSize default, but got %d", d)
	}
}

func TestSearchAsync_getCompanyBody(t *testing.T) {
	repo := createTestCompanyRepository(t)
	defer repo.close()

	body, err := getCompanyBody(companySearcher, "watchId", "21206", repo)
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Error("empty body")
	}

	var company Company
	if err := json.NewDecoder(body).Decode(&company); err != nil {
		t.Error(err)
	}
	if company.ID == "" {
		t.Errorf("empty company: %#v", company)
	}

	// Company not found
	body, err = getCompanyBody(companySearcher, "watchId", "", repo)
	if err == nil || body != nil {
		t.Fatal("expected error and no body")
	}
}

func TestSearchAsync_getCustomerBody(t *testing.T) {
	repo := createTestCustomerRepository(t)
	defer repo.close()

	body, err := getCustomerBody(customerSearcher, "watchId", "306", repo)
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Error("empty body")
	}

	var customer Customer
	if err := json.NewDecoder(body).Decode(&customer); err != nil {
		t.Error(err)
	}
	if customer.ID == "" {
		t.Errorf("empty customer: %#v", customer)
	}

	// Customer not found
	body, err = getCustomerBody(customerSearcher, "watchId", "", repo)
	if err == nil || body != nil {
		t.Fatal("expected error and no body")
	}
}
