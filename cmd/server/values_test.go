// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/base/log"

	"github.com/gorilla/mux"
)

func TestValues__accumulator(t *testing.T) {
	acc := newAccumulator(2)
	acc.add("v2") // out of alphanumeric order
	acc.add("")   // empty value, ignored
	acc.add("v1")
	acc.add("v1") // duplicate

	xs := acc.getValues()
	if len(xs) != 2 {
		t.Errorf("got values: %v", xs)
	}
	if xs[0] != "v1" || xs[1] != "v2" {
		t.Errorf("values: %v", xs)
	}

	// add another past the limit and expect it to be excluded
	acc.add("v3")
	if len(xs) != 2 {
		t.Errorf("got values: %v", xs)
	}
}

func TestValues__getValues(t *testing.T) {
	router := mux.NewRouter()
	addValuesRoutes(log.NewNopLogger(), router, sdnSearcher)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ui/values/sdnType", nil)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var values []string
	if err := json.NewDecoder(w.Body).Decode(&values); err != nil {
		t.Error(err)
	}
	if len(values) != 2 {
		t.Errorf("values: %v", values)
	}
	for i := range values {
		switch values[i] {
		case "individual", "entity":
			continue
		default:
			t.Errorf("values[%d]=%s", i, values[i])
		}
	}
}

func TestValues__getValuesLimit(t *testing.T) {
	router := mux.NewRouter()
	addValuesRoutes(log.NewNopLogger(), router, sdnSearcher)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ui/values/ofacProgram?limit=1", nil)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var values []string
	if err := json.NewDecoder(w.Body).Decode(&values); err != nil {
		t.Error(err)
	}
	if len(values) != 1 {
		t.Errorf("values: %v", values)
	}
}

func TestValues__getValuesErr(t *testing.T) {
	router := mux.NewRouter()
	addValuesRoutes(log.NewNopLogger(), router, sdnSearcher)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ui/values/other", nil)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %s", w.Code, w.Body.String())
	}
}
