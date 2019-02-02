// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	watchResearchBatchSize = 100
)

func init() {
	if v := os.Getenv("WEBHOOK_BATCH_SIZE"); v != "" {
		d, _ := strconv.Atoi(v)
		if d > 0 {
			watchResearchBatchSize = d
		}
	}
}

func (s *searcher) spawnResearching(watchRepo watchRepository, updates chan *downloadStats) {
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
				for i := range watches { // .id, .customerId, .webhook
					customer := getCustomerById(watches[i].customerId, s)
					if customer == nil {
						// TODO(adam): remove watch?
						s.logger.Log("search", fmt.Sprintf("async: customer %v not found for watchId=%q", watches[i].customerId, watches[i].id))
					}

					s.logger.Log("search", fmt.Sprintf("async: watch for customer %s found", watches[i].customerId))

					if err := callWebhook(watches[i].id, customer, watches[i].webhook); err != nil {
						s.logger.Log("search", fmt.Sprintf("async: %v", err))
					}
				}
			}

		}
	}
}
