// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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

func (s *searcher) spawnResearching(watchRepo watchRepository, webhookRepo webhookRepository, updates chan *downloadStats) {
	for {
		select {
		case <-updates:
			s.logger.Log("search", "async: starting re-search of watches")
			cursor := watchRepo.getWatchesCursor(watchResearchBatchSize)
			for {
				watches, _ := cursor.Next()
				if len(watches) == 0 {
					break
				}
				for i := range watches {
					customer, _ := getCustomerById(watches[i].customerId, s, nil)
					if customer == nil {
						// TODO(adam): remove watch?
						s.logger.Log("search", fmt.Sprintf("async: watch %s customer %v not found", watches[i].id, watches[i].customerId))
					}

					s.logger.Log("search", fmt.Sprintf("async: watch %s for customer %s found", watches[i].id, watches[i].customerId))

					now := time.Now()
					status, err := callWebhook(watches[i].id, customer, watches[i].webhook, watches[i].authToken)
					if err != nil {
						s.logger.Log("search", fmt.Errorf("async: problem writing watch (%s) webhook status: %v", watches[i].id, err))
					}
					if err := webhookRepo.recordWebhook(watches[i].id, now, status); err != nil {
						s.logger.Log("search", fmt.Errorf("async: problem writing watch (%s) webhook status: %v", watches[i].id, err))
					}
				}
			}

		}
	}
}
