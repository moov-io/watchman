// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"errors"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl_eu"
	"github.com/moov-io/watchman/pkg/csl_uk"
	"github.com/moov-io/watchman/pkg/csl_us"
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
	alt    *ofac.AlternateIdentity
	sdn    *ofac.SDN
	ssi    *csl_us.SSI
	uvl    *csl_us.UVL
	isn    *csl_us.ISN
	fse    *csl_us.FSE
	plc    *csl_us.PLC
	cap    *csl_us.CAP
	dtc    *csl_us.DTC
	cmic   *csl_us.CMIC
	ns_mbs *csl_us.NS_MBS

	eu_csl *csl_eu.CSLRecord

	uk_csl *csl_uk.CSLRecord

	uk_sanctionsList *csl_uk.SanctionsListRecord

	dp    *dpl.DPL
	el    *csl_us.EL
	addrs []*ofac.Address

	altNames []string
}

func SdnName(sdn *ofac.SDN, addrs []*ofac.Address) *Name {
	return &Name{
		Original:  sdn.SDNName,
		Processed: sdn.SDNName,
		sdn:       sdn,
		addrs:     addrs,
	}
}

func AltName(alt *ofac.AlternateIdentity) *Name {
	return &Name{
		Original:  alt.AlternateName,
		Processed: alt.AlternateName,
		alt:       alt,
	}
}

func DPName(dp *dpl.DPL) *Name {
	return &Name{
		Original:  dp.Name,
		Processed: dp.Name,
		dp:        dp,
	}
}

func CSLName(item interface{}) *Name {
	switch v := item.(type) {
	case *csl_us.EL:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			el:        v,
		}
	case *csl_us.MEU:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
		}
	case *csl_us.SSI:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			ssi:       v,
			altNames:  v.AlternateNames,
		}
	case *csl_us.UVL:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			uvl:       v,
		}
	case *csl_us.ISN:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			isn:       v,
			altNames:  v.AlternateNames,
		}
	case *csl_us.FSE:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			fse:       v,
		}
	case *csl_us.PLC:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			plc:       v,
			altNames:  v.AlternateNames,
		}
	case *csl_us.CAP:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			cap:       v,
			altNames:  v.AlternateNames,
		}
	case *csl_us.DTC:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			dtc:       v,
			altNames:  v.AlternateNames,
		}
	case *csl_us.CMIC:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			cmic:      v,
			altNames:  v.AlternateNames,
		}
	case *csl_us.NS_MBS:
		return &Name{
			Original:  v.Name,
			Processed: v.Name,
			ns_mbs:    v,
			altNames:  v.AlternateNames,
		}
	case *csl_eu.CSLRecord:
		if len(v.NameAliasWholeNames) >= 1 {
			var alts []string
			alts = append(alts, v.NameAliasWholeNames...)
			return &Name{
				Original:  v.NameAliasWholeNames[0],
				Processed: v.NameAliasWholeNames[0],
				eu_csl:    v,
				altNames:  alts,
			}
		}
	case *csl_uk.CSLRecord:
		if len(v.Names) >= 1 {
			var alts []string
			alts = append(alts, v.Names...)
			return &Name{
				Original:  v.Names[0],
				Processed: v.Names[0],
				uk_csl:    v,
				altNames:  alts,
			}
		}
	case *csl_uk.SanctionsListRecord:
		if len(v.Names) >= 1 {
			var alts []string
			alts = append(alts, v.Names...)
			return &Name{
				Original:         v.Names[0],
				Processed:        v.Names[0],
				uk_sanctionsList: v,
				altNames:         alts,
			}
		}

		return &Name{}
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

func NewPipeliner(logger log.Logger, debug bool) *Pipeliner {
	steps := []step{
		&reorderSDNStep{},
		&companyNameCleanupStep{},
		&stopwordsStep{},
		&normalizeStep{},
	}
	if debug {
		for i := range steps {
			steps[i] = &debugStep{logger: logger, step: steps[i]}
		}
	}
	return &Pipeliner{
		logger: logger,
		steps:  steps,
	}
}

type Pipeliner struct {
	logger log.Logger
	steps  []step
}

func (p *Pipeliner) Do(name *Name) error {
	if p == nil || p.steps == nil || p.logger == nil || name == nil {
		return errors.New("nil Pipeliner or Name")
	}
	for i := range p.steps {
		if name == nil {
			return p.logger.Error().LogErrorf("%T: nil Name", p.steps[i]).Err()
		}
		if err := p.steps[i].apply(name); err != nil {
			return p.logger.Error().LogErrorf("pipeline: %v", err).Err()
		}
	}
	return nil
}
