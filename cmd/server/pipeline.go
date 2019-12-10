// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/ofac"
)

type Name struct {
	Original  string
	Processed string

	// optional metadata of where a name came from
	sdn *ofac.SDN
	ssi *csl.SSI
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
	ds.logger.Log("pipeline", fmt.Sprintf("%T result", ds.step), "result", in.Processed, "original", in.Original) // TODO(adam): PII log
	return nil
}

func pipeline(logger log.Logger, name *Name) error {
	steps := []step{
		&debugStep{logger: logger, step: &reorderSDNStep{}},
	}
	for i := range steps {
		if err := steps[i].apply(name); err != nil {
			return fmt.Errorf("pipeline: %v", err)
		}
	}
	return nil
}
