// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"
)

// Name represents an individual or entity name to be processed for search.
type Name struct {
	// Original is the initial value and MUST not be changed by any pipeline step.
	Original string

	// Processed is the mutable value that each pipeline step can optionally
	// replace and is read as the input to each step.
	Processed string

	// optional metadata of where a name came from
	alt   *ofac.AlternateIdentity
	sdn   *ofac.SDN
	ssi   *csl.SSI
	dp    *dpl.DPL
	el    *csl.EL
	addrs []*ofac.Address

	altNames []string
}

func sdnName(sdn *ofac.SDN, addrs []*ofac.Address) *Name {
	return &Name{
		Original:  sdn.SDNName,
		Processed: sdn.SDNName,
		sdn:       sdn,
		addrs:     addrs,
	}
}

func altName(alt *ofac.AlternateIdentity) *Name {
	return &Name{
		Original:  alt.AlternateName,
		Processed: alt.AlternateName,
		alt:       alt,
	}
}

func dpName(dp *dpl.DPL) *Name {
	return &Name{
		Original:  dp.Name,
		Processed: dp.Name,
		dp:        dp,
	}
}

func cslName(item interface{}) *Name {
	switch v := item.(type) {
	case *csl.EL:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			el:        v,
		}
	case *csl.MEU:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
		}
	case *csl.SSI:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			ssi:       v,
			altNames:  v.AlternateNames,
		}
	}
	return &Name{}
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
		return ds.logger.Info().With(log.Fields{
			"original": log.String(in.Original),
		}).LogErrorf("%T encountered error: %v", ds.step, err).Err()
	}
	ds.logger.Info().With(log.Fields{
		"original": log.String(in.Original),
		"result":   log.String(in.Processed),
	}).Logf("%T", ds.step)
	return nil
}

func newPipeliner(logger log.Logger) *pipeliner {
	return &pipeliner{
		logger: logger,
		steps: []step{
			&debugStep{logger: logger, step: &reorderSDNStep{}},
			&debugStep{logger: logger, step: &companyNameCleanupStep{}},
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
	if p == nil || p.steps == nil || p.logger == nil || name == nil {
		return errors.New("nil pipeliner or Name")
	}
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
