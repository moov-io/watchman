// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/go-kit/kit/log"
)

var (
	noopPipeliner = &pipeliner{
		logger: log.NewNopLogger(),
		steps:  []step{},
	}

	noLogPipeliner = newPipeliner(log.NewNopLogger())
)

func TestPipelineNoop(t *testing.T) {
	if err := noopPipeliner.Do(&Name{}); err != nil {
		t.Fatal(err)
	}
}
