// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/internal/stringscore"

	"github.com/stretchr/testify/assert"
)

func TestIssue326(t *testing.T) {
	india := prepare.LowerAndRemovePunctuation("Huawei Technologies India Private Limited")
	investment := prepare.LowerAndRemovePunctuation("Huawei Technologies Investment Co. Ltd.")

	// Cuba
	score := stringscore.JaroWinkler(prepare.LowerAndRemovePunctuation("Huawei Cuba"), prepare.LowerAndRemovePunctuation("Huawei"))
	assert.Equal(t, 0.7444444444444445, score)

	// India
	score = stringscore.JaroWinkler(india, prepare.LowerAndRemovePunctuation("Huawei"))
	assert.Equal(t, 0.4846031746031746, score)
	score = stringscore.JaroWinkler(india, prepare.LowerAndRemovePunctuation("Huawei Technologies"))
	assert.Equal(t, 0.6084415584415584, score)

	// Investment
	score = stringscore.JaroWinkler(investment, prepare.LowerAndRemovePunctuation("Huawei"))
	assert.Equal(t, 0.3788888888888889, score)
	score = stringscore.JaroWinkler(investment, prepare.LowerAndRemovePunctuation("Huawei Technologies"))
	assert.Equal(t, 0.5419191919191919, score)
}
