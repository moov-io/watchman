// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIssue326(t *testing.T) {
	india := precompute("Huawei Technologies India Private Limited")
	investment := precompute("Huawei Technologies Investment Co. Ltd.")

	// Cuba
	score := jaroWinkler(precompute("Huawei Cuba"), precompute("Huawei"))
	assert.Equal(t, 0.7444444444444445, score)

	// India
	score = jaroWinkler(india, precompute("Huawei"))
	assert.Equal(t, 0.4846031746031746, score)
	score = jaroWinkler(india, precompute("Huawei Technologies"))
	assert.Equal(t, 0.6084415584415584, score)

	// Investment
	score = jaroWinkler(investment, precompute("Huawei"))
	assert.Equal(t, 0.3788888888888889, score)
	score = jaroWinkler(investment, precompute("Huawei Technologies"))
	assert.Equal(t, 0.5419191919191919, score)
}
