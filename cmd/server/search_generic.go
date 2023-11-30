// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"sync"
)

type Result[T any] struct {
	Data T

	match       float64
	matchedName string

	precomputedName string
	precomputedAlts []string
}

func (e Result[T]) MarshalJSON() ([]byte, error) {
	// Due to a problem with embedding type parameters we have to dig into
	// the parameterized type fields and include them in one object.
	//
	// Helpful Tips:
	// https://stackoverflow.com/a/64420452
	// https://github.com/golang/go/issues/41563

	v := reflect.ValueOf(e.Data)

	result := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i)
		value := v.Field(i)

		if key.IsExported() {
			result[key.Name] = value.Interface()
		}
	}

	result["match"] = e.match
	result["matchedName"] = e.matchedName

	return json.Marshal(result)
}

func topResults[T any](limit int, minMatch float64, name string, data []*Result[T]) []*Result[T] {
	if len(data) == 0 {
		return nil
	}

	name = precompute(name)
	nameTokens := strings.Fields(name)

	xs := newLargest(limit, minMatch)

	var wg sync.WaitGroup
	wg.Add(len(data))

	for i := range data {
		go func(i int) {
			defer wg.Done()

			it := &item{
				matched: data[i].precomputedName,
				value:   data[i],
				weight:  bestPairsJaroWinkler(nameTokens, data[i].precomputedName),
			}

			for _, alt := range data[i].precomputedAlts {
				if alt == "" {
					continue
				}

				score := bestPairsJaroWinkler(nameTokens, alt)
				if score > it.weight {
					it.matched = alt
					it.weight = score
				}
			}

			xs.add(it)
		}(i)
	}
	wg.Wait()

	out := make([]*Result[T], 0)
	for _, thisItem := range xs.items {
		if v := thisItem; v != nil {
			vv, ok := v.value.(*Result[T])
			if !ok {
				continue
			}
			res := &Result[T]{
				Data:            vv.Data,
				match:           v.weight,
				matchedName:     v.matched,
				precomputedName: vv.precomputedName,
				precomputedAlts: vv.precomputedAlts,
			}
			out = append(out, res)
		}
	}
	return out
}
