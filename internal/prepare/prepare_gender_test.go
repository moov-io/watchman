package prepare

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestPrepare_NormalizeGender(t *testing.T) {
	cases := []struct {
		input    string
		expected search.Gender
	}{
		{"M", search.GenderMale},
		{"Male", search.GenderMale},
		{"guy", search.GenderMale},
		{"F", search.GenderFemale},
		{"FEMALE", search.GenderFemale},
		{"Girl", search.GenderFemale},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got := NormalizeGender(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
