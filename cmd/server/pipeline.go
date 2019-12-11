// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/go-kit/kit/log"
)

type Name struct {
	Original  string
	Processed string

	// optional metadata of where a name came from
	sdn *ofac.SDN
	ssi *csl.SSI
	dp  *dpl.DPL
	el  *csl.EL

	addrs []*ofac.Address
}

func sdnName(sdn *ofac.SDN) *Name {
	return &Name{
		Original:  sdn.SDNName,
		Processed: sdn.SDNName,
		sdn:       sdn,
	}
}

func ssiName(ssi *csl.SSI) *Name {
	return &Name{
		Original:  ssi.Name,
		Processed: ssi.Name,
		ssi:       ssi,
	}
}

func dpName(dp *dpl.DPL) *Name {
	return &Name{
		Original:  dp.Name,
		Processed: dp.Name,
		dp:        dp,
	}
}

func bisEntityName(el *csl.EL) *Name {
	return &Name{
		Original:  el.Name,
		Processed: el.Name,
		el:        el,
	}
}

type step interface {
	apply(*Name) error
}

type debugStep struct {
	step

	logger log.Logger
}

func (ds *debugStep) apply(in *Name) error {
	if err := ds.step.apply(in); err != nil {
		// TODO(adam): PII log, we can't have this
		ds.logger.Log("pipeline", fmt.Sprintf("%T encountered error: %v", ds.step, err), "original", in.Original)
		return err
	}
	ds.logger.Log("pipeline", fmt.Sprintf("%T", ds.step), "result", in.Processed, "original", in.Original) // TODO(adam): PII log
	return nil
}

func newPiepliner(logger log.Logger) *pipeliner {
	return &pipeliner{
		logger: logger,
		steps: []step{
			&debugStep{logger: logger, step: &reorderSDNStep{}},
			&debugStep{logger: logger, step: &stopwordsStep{}},
			&debugStep{logger: logger, step: &normalizeStep{}},
		},
	}
}

type pipeliner struct {
	logger log.Logger
	steps  []step
}

func (p *pipeliner) Do(name *Name) error {
	for i := range p.steps {
		if name == nil {
			return fmt.Errorf("%T: nil Name", p.steps[i])
		}
		if err := p.steps[i].apply(name); err != nil {
			return fmt.Errorf("pipeline: %v", err)
		}
	}
	return nil
}
