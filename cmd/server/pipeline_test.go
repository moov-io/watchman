// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-kit/kit/log"
)

var (
	noopPipeliner = &pipeliner{
		logger: log.NewNopLogger(),
		steps:  []step{},
	}
)
