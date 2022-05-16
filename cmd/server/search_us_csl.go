// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"github.com/moov-io/watchman/pkg/csl"
)

func precomputeSSIs(ssis []*csl.SSI, pipe *pipeliner) []*Result[csl.SSI] {
	out := make([]*Result[csl.SSI], len(ssis))
	for i, ssi := range ssis {
		nn := ssiName(ssi)
		if err := pipe.Do(nn); err != nil {
			pipe.logger.LogErrorf("problem pipelining SSI: %v", err)
			continue
		}

		var altNames []string
		for i := range ssi.AlternateNames {
			altNN := &Name{Processed: ssi.AlternateNames[i]}
			if err := pipe.Do(altNN); err != nil {
				pipe.logger.LogErrorf("problem pipelining alt: %v", err)
				continue
			}
			altNames = append(altNames, altNN.Processed)
		}
		ssi.AlternateNames = altNames

		out[i] = &Result[csl.SSI]{
			Data:            *ssi,
			precomputedName: nn.Processed,
			precomputedAlts: altNames,
		}
	}
	return out
}

// TopSSIs searches Sectoral Sanctions records by Name and Alias
func (s *searcher) TopSSIs(limit int, minMatch float64, name string) []*Result[csl.SSI] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start() // TODO(adam): This used to be on a pre-record gate, so this may have different perf metrics
	defer s.Gate.Done()

	return topResults[csl.SSI](limit, minMatch, name, s.SSIs)
}

func precomputeBISEntities(els []*csl.EL, pipe *pipeliner) []*Result[csl.EL] {
	out := make([]*Result[csl.EL], len(els))
	for i, el := range els {
		nn := bisEntityName(el)
		if err := pipe.Do(nn); err != nil {
			pipe.logger.LogErrorf("problem pipelining EL: %v", err)
			continue
		}

		var altNames []string
		for i := range el.AlternateNames {
			altNN := &Name{Processed: el.AlternateNames[i]}
			if err := pipe.Do(altNN); err != nil {
				pipe.logger.LogErrorf("problem pipelining alt: %v", err)
				continue
			}
			altNames = append(altNames, altNN.Processed)
		}
		el.AlternateNames = altNames

		out[i] = &Result[csl.EL]{
			Data:            *el,
			precomputedName: nn.Processed,
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
