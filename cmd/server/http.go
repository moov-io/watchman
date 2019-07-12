// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/idempotent/lru"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	routeHistogram = prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
		Name: "http_response_duration_seconds",
		Help: "Histogram representing the http response durations",
	}, []string{"route"})

	inmemIdempotentRecorder = lru.New()
)

func wrapResponseWriter(logger log.Logger, w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	route := fmt.Sprintf("%s-%s", strings.ToLower(r.Method), cleanMetricsPath(r.URL.Path))
	return moovhttp.Wrap(logger, routeHistogram.With("route", route), w, r)
}

var baseIDRegex = regexp.MustCompile(`([a-f0-9]{40})`)

// cleanMetricsPath takes a URL path and formats it for Prometheus metrics
//
// This method replaces /'s with -'s and clean out OFAC ID's (which are numeric).
// This method also strips out moov/base.ID() values from URL path slugs.
func cleanMetricsPath(path string) string {
	parts := strings.Split(path, "/")
	var out []string
	for i := range parts {
		if n, _ := strconv.Atoi(parts[i]); n > 0 || parts[i] == "" {
			continue // numeric ID
		}
		if baseIDRegex.MatchString(parts[i]) {
			continue // assume it's a moov/base.ID() value
		}
		out = append(out, parts[i])
	}
	return strings.Join(out, "-")
}
