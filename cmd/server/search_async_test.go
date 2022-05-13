// Copyright 2022 The Moov Authors
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

func TestSearchAsync__renderBody(t *testing.T) {
	w := watch{
		id:           "12345",
		customerID:   "306",
		customerName: "BANCO NACIONAL DE CUBA",
		webhook:      "https://example.com",
		authToken:    "secret-value",
	}

	companyRepo := createTestCompanyRepository(t)
	defer companyRepo.close()
	customerRepo := createTestCustomerRepository(t)
	defer customerRepo.close()

	body, err := customerSearcher.renderBody(w, companyRepo, customerRepo)
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Fatal("nil Body")
	}

	w = watch{
		id:          "54321",
		companyID:   "21206",
		companyName: "AL-HISN",
		webhook:     "https://moov.io",
		authToken:   "hidden",
	}
	body, err = companySearcher.renderBody(w, companyRepo, customerRepo)
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Fatal("nil Body")
	}
}

func TestSearchAsync__renderBodyName(t *testing.T) {
	w := watch{
		id:           "12345",
		customerName: "BANCO NACIONAL DE CUBA",
		webhook:      "https://example.com",
		authToken:    "secret-value",
	}

	companyRepo := createTestCompanyRepository(t)
	defer companyRepo.close()
	customerRepo := createTestCustomerRepository(t)
	defer customerRepo.close()

	body, err := customerSearcher.renderBody(w, companyRepo, customerRepo)
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Fatal("nil Body")
	}

	w = watch{
		id:          "54321",
		companyName: "AL-HISN",
		webhook:     "https://moov.io",
		authToken:   "hidden",
	}
	body, err = companySearcher.renderBody(w, companyRepo, customerRepo)
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Fatal("nil Body")
	}
}

func TestSearchAsync_getCompanyBody(t *testing.T) {
	repo := createTestCompanyRepository(t)
	defer repo.close()

	body, err := getCompanyBody(companySearcher, "watchID", "21206", 1.0, repo)
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
	if (1.0 - company.Match) > 0.001 {
		t.Errorf("unexpected company.Match=%.2f", company.Match)
	}

	// Company not found
	body, err = getCompanyBody(companySearcher, "watchID", "", 0.0, repo)
	if err == nil || body != nil {
		t.Fatal("expected error and no body")
	}
}

func TestSearchAsync_getCustomerBody(t *testing.T) {
	repo := createTestCustomerRepository(t)
	defer repo.close()

	body, err := getCustomerBody(customerSearcher, "watchID", "306", 0.91, repo)
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
	if (0.91 - customer.Match) > 0.001 {
		t.Errorf("unexpected customer.Match=%.2f", customer.Match)
	}

	// Customer not found
	body, err = getCustomerBody(customerSearcher, "watchID", "", 0.0, repo)
	if err == nil || body != nil {
		t.Fatal("expected error and no body")
	}
}
