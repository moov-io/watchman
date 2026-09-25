// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

var nativeIDPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)^ofac-(\d+)$`),
	regexp.MustCompile(`(?i)^gb-hmt-(\d+)$`),
	regexp.MustCompile(`(?i)^eu-fsf-eu-(\d+)`),
	regexp.MustCompile(`(?i)^eu-tb-logical-(\d+)$`),
	regexp.MustCompile(`(?i)^unsc-(\d+)$`),
	regexp.MustCompile(`(?i)^ch-seco-(\d+)$`),
}

type searchConfig struct {
	evalConfig
	address  string
	searchN  int
	minMatch float64
	timeout  time.Duration
}

func scoreSearch(ctx context.Context, pairs []Pair, cfg searchConfig) ([]pairResult, error) {
	if cfg.timeout <= 0 {
		cfg.timeout = 30 * time.Second
	}
	client := search.NewClient(&http.Client{Timeout: cfg.timeout}, cfg.address)
	out := make([]pairResult, len(pairs))
	workers := cfg.workers
	if workers < 1 {
		workers = 1
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobs := make(chan int)
	var (
		wg       sync.WaitGroup
		errOnce  sync.Once
		firstErr error
	)
	setErr := func(err error) {
		errOnce.Do(func() {
			firstErr = err
			cancel()
		})
	}

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conv := convertOpts{maxAltNames: cfg.maxAltNames, nameOnly: cfg.nameOnly}
			opts := search.SearchOpts{
				Limit:     cfg.searchN,
				MinMatch:  cfg.minMatch,
				Algorithm: cfg.algorithm,
			}
			for i := range jobs {
				if ctx.Err() != nil {
					return
				}
				p := pairs[i]
				query, _ := convertPair(p, conv)
				query.SourceID = ""

				resp, err := client.SearchByEntity(ctx, query, opts)
				if err != nil {
					setErr(fmt.Errorf("search pair %d (%s): %w", i, p.Left.ID, err))
					return
				}
				score, found, rank, hit := matchRight(p.Right, resp.Entities)
				res := makeResult(i, p, score, cfg.threshold)
				res.SearchFound = found
				res.SearchRank = rank
				if found {
					res.SearchSource = string(hit.Source)
					res.SearchSourceID = hit.SourceID
				}
				out[i] = res
			}
		}()
	}

send:
	for i := range pairs {
		select {
		case <-ctx.Done():
			break send
		case jobs <- i:
		}
	}
	close(jobs)
	wg.Wait()
	return out, firstErr
}

func matchRight(right FTMEntity, hits []search.SearchedEntity[search.Value]) (score float64, found bool, rank int, hit search.SearchedEntity[search.Value]) {
	want := rightIDs(right)
	caption := strings.TrimSpace(right.Caption)
	for i, h := range hits {
		if idMatch(h, want) || nameMatch(h.Name, caption) {
			return h.Match, true, i + 1, h
		}
	}
	return 0, false, 0, search.SearchedEntity[search.Value]{}
}

func rightIDs(e FTMEntity) map[string]struct{} {
	out := make(map[string]struct{})
	add := func(id string) {
		id = strings.TrimSpace(id)
		if id == "" {
			return
		}
		out[strings.ToLower(id)] = struct{}{}
		for _, re := range nativeIDPatterns {
			m := re.FindStringSubmatch(id)
			if len(m) == 2 {
				out[strings.ToLower(m[1])] = struct{}{}
			}
		}
	}
	add(e.ID)
	for _, r := range e.Referents {
		add(r)
	}
	return out
}

func idMatch(h search.SearchedEntity[search.Value], want map[string]struct{}) bool {
	if h.SourceID == "" {
		return false
	}
	_, ok := want[strings.ToLower(h.SourceID)]
	return ok
}

func nameMatch(got, want string) bool {
	got = strings.TrimSpace(got)
	want = strings.TrimSpace(want)
	if got == "" || want == "" {
		return false
	}
	return strings.EqualFold(got, want)
}
