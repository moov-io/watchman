// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"reflect"

	"github.com/moov-io/watchman/pkg/csl"
)

func precomputeCSLEntities[T any](items []*T, pipe *pipeliner) []*Result[T] {
	out := make([]*Result[T], len(items))

	for i, item := range items {
		name := cslName(item)
		if err := pipe.Do(name); err != nil {
			pipe.logger.LogErrorf("problem pipelining %T: %v", item, err)
			continue
		}

		var altNames []string

		elm := reflect.ValueOf(item).Elem()
		for i := 0; i < elm.NumField(); i++ {
			name := elm.Type().Field(i).Name
			_type := elm.Type().Field(i).Type.String()

			if name == "AlternateNames" && _type == "[]string" {
				alts, ok := elm.Field(i).Interface().([]string)
				if !ok {
					continue
				}
				for j := range alts {
					alt := &Name{Processed: alts[j]}
					pipe.Do(alt)
					altNames = append(altNames, alt.Processed)
				}
			}
		}

		out[i] = &Result[T]{
			Data:            *item,
			precomputedName: name.Processed,
			precomputedAlts: altNames,
		}
	}

	return out
}

// TopBISEntities searches BIS Entity List records by name and alias
func (s *searcher) TopBISEntities(limit int, minMatch float64, name string) []*Result[csl.EL] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start() // TODO(adam): This used to be on a pre-record gate, so this may have different perf metrics
	defer s.Gate.Done()

	return topResults[csl.EL](limit, minMatch, name, s.BISEntities)
}

// TopMEUs searches Military End User records by name and alias
func (s *searcher) TopMEUs(limit int, minMatch float64, name string) []*Result[csl.MEU] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.MEU](limit, minMatch, name, s.MilitaryEndUsers)
}

// TopSSIs searches Sectoral Sanctions records by Name and Alias
func (s *searcher) TopSSIs(limit int, minMatch float64, name string) []*Result[csl.SSI] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.SSI](limit, minMatch, name, s.SSIs)
}
