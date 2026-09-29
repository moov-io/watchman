// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitExNames(t *testing.T) {
	cases := []struct {
		in      string
		cleaned string
		former  []string
	}{
		{"Ocean Pioneer (ex-Cape Diamond)", "Ocean Pioneer", []string{"Cape Diamond"}},
		{"Star Voyager (ex-Southern Glory)", "Star Voyager", []string{"Southern Glory"}},
		{"M/T Ondara Leigh (ex-Pelican Morrow)", "M/T Ondara Leigh", []string{"Pelican Morrow"}},
		{"Wheatland Glory (ex-Pacific Sower)", "Wheatland Glory", []string{"Pacific Sower"}},
		{"M.T. THONG NAWA 3 (Laem Chabang)", "M.T. THONG NAWA 3 (Laem Chabang)", nil},
		{"Solenne Harbour", "Solenne Harbour", nil},
		{"", "", nil},
		{"Foo (ex-Bar) (ex-Baz)", "Foo", []string{"Bar", "Baz"}},
		{"Foo (EX Cape Diamond)", "Foo", []string{"Cape Diamond"}},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			cleaned, former := SplitExNames(tc.in)
			require.Equal(t, tc.cleaned, cleaned)
			require.Equal(t, tc.former, former)
		})
	}
}
