// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/moov-io/base/log"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func BenchmarkSearchHandler(b *testing.B) {
	searcher := createTestSearcher(b) // Uses live data
	b.ResetTimer()

	router := mux.NewRouter()
	addSearchRoutes(log.NewNopLogger(), router, searcher)

	g := &errgroup.Group{}
	g.SetLimit(10)

	for i := 0; i < b.N; i++ {
		g.Go(func() error {
			name := fake.Person().Name()

			v := make(url.Values, 0)
			v.Set("name", name)
			v.Set("limit", "10")
			v.Set("minMatch", "0.70")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/search?%s", v.Encode()), nil)
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				return fmt.Errorf("unexpected status: %v", w.Code)
			}
			return nil
		})
	}
	require.NoError(b, g.Wait())
}
