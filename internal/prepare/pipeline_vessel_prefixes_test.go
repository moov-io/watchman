// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripVesselNamePrefixes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"mv solenne harbour", "solenne harbour"},
		{"m t ardrevan pearl", "ardrevan pearl"},
		{"mt ardrevan pearl", "ardrevan pearl"},
		{"ss maltrevor queen", "maltrevor queen"},
		{"m v pelgarra glory", "pelgarra glory"},
		{"solenne harbour", "solenne harbour"},
		{"", ""},
		{"mv", "mv"},
		{"msv sahil e sindh", "sahil e sindh"},
		{"n m bahia trader", "bahia trader"},
		{"t b chao le 8", "chao le 8"},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, StripVesselNamePrefixes(tc.in))
		})
	}
}
