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
	assert.Equal(t, score, 0.8055555555555556)

	// India
	score = jaroWinkler(india, precompute("Huawei"))
	assert.Equal(t, score, 0.5592063492063492)
	score = jaroWinkler(india, precompute("Huawei Technologies"))
	assert.Equal(t, score, 0.6903174603174603)

	// Investment
	score = jaroWinkler(investment, precompute("Huawei"))
	assert.Equal(t, score, 0.3788888888888889)
	score = jaroWinkler(investment, precompute("Huawei Technologies"))
	assert.Equal(t, score, 0.7377777777777779)
}
