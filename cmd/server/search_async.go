// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/base/log"
)

var (
	watchResearchBatchSize = 100
)

func init() {
	watchResearchBatchSize = readWebhookBatchSize(os.Getenv("WEBHOOK_BATCH_SIZE"))
}

func readWebhookBatchSize(str string) int {
	if str == "" {
		return watchResearchBatchSize
	}
	d, _ := strconv.Atoi(str)
	if d > 0 {
		return d
	}
	return watchResearchBatchSize
}

// spawnResearching will block and select on updates for when to re-inspect all watches setup.
// Since watches are used to post list data via webhooks they are used as catalysts in other systems.
func (s *searcher) spawnResearching(logger log.Logger, companyRepo companyRepository, custRepo customerRepository, watchRepo watchRepository, webhookRepo webhookRepository) {
	s.logger.Log("async: starting re-search of watches")
	cursor := watchRepo.getWatchesCursor(logger, watchResearchBatchSize)
	for {
		watches, _ := cursor.Next()
		if len(watches) == 0 {
			break
		}
		for i := range watches {
			body, err := s.renderBody(watches[i], companyRepo, custRepo)
			if err != nil {
				s.logger.Logf("async: watch %s: %v", watches[i].id, err)
				continue
			}
			if body == nil {
				s.logger.Logf("async: no body rendered for watchID=%s - skipping", watches[i].id)
				continue
			}

			// Send HTTP webhook
			now := time.Now()
			status, err := callWebhook(body, watches[i].webhook, watches[i].authToken)
			if err != nil {
				s.logger.Logf("async: problem writing watch (%s) webhook status: %v", watches[i].id, err)
			}
			if err := webhookRepo.recordWebhook(watches[i].id, now, status); err != nil {
				s.logger.Logf("async: problem writing watch (%s) webhook status: %v", watches[i].id, err)
			}
		}
	}
	s.logger.Log("async: finished re-search of watches")
}

func (s *searcher) renderBody(w watch, companyRepo companyRepository, custRepo customerRepository) (*bytes.Buffer, error) {
	keeper := keepSDN(filterRequest{})

	// Perform a query (ID watches) or search (name watches) and encode the model in JSON for calling the webhook.
	switch {
	case w.customerID != "":
		s.logger.Logf("async: watch %s for customer %s found", w.id, w.customerID)
		return getCustomerBody(s, w.id, w.customerID, 1.0, custRepo)

	case w.customerName != "":
		s.logger.Logf("async: name watch '%s' for customer %s found", w.customerName, w.id)
		sdns := s.TopSDNs(5, 0.00, w.customerName, keeper)
		for j := range sdns {
			if strings.EqualFold(sdns[j].SDNType, "individual") {
				return getCustomerBody(s, w.id, sdns[j].EntityID, sdns[j].match, custRepo)
			}
		}

	case w.companyID != "":
		s.logger.Logf("async: watch %s for company %s found", w.id, w.companyID)
		return getCompanyBody(s, w.id, w.companyID, 1.0, companyRepo)

	case w.companyName != "":
		s.logger.Logf("async: name watch '%s' for company %s found", w.companyName, w.id)
		sdns := s.TopSDNs(5, 0.00, w.companyName, keeper)
		for j := range sdns {
			if !strings.EqualFold(sdns[j].SDNType, "individual") {
				return getCompanyBody(s, w.id, sdns[j].EntityID, sdns[j].match, companyRepo)
			}
		}
	}
	return nil, nil
}

// getCustomerBody returns the JSON encoded form of a given customer by their EntityID
func getCustomerBody(s *searcher, watchID string, customerID string, match float64, repo customerRepository) (*bytes.Buffer, error) {
	customer, _ := getCustomerByID(customerID, s, repo)
	if customer == nil {
		return nil, fmt.Errorf("async: watch %s customer %v not found", watchID, customerID)
	}
	customer.Match = match

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(customer); err != nil {
		return nil, fmt.Errorf("problem creating JSON for customer watch %s: %v", watchID, err)
	}
	return &buf, nil
}

// getCompanyBody returns the JSON encoded form of a given customer by their EntityID
func getCompanyBody(s *searcher, watchID string, companyID string, match float64, repo companyRepository) (*bytes.Buffer, error) {
	company, _ := getCompanyByID(companyID, s, repo)
	if company == nil {
		return nil, fmt.Errorf("async: watch %s company %v not found", watchID, companyID)
	}
	company.Match = match

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(company); err != nil {
		return nil, fmt.Errorf("problem creating JSON for company watch %s: %v", watchID, err)
	}
	return &buf, nil
}
