package norm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard name",
			input:    "AEROCARIBBEAN AIRLINES",
			expected: "aerocaribbean airlines",
		},
		{
			name:     "name with punctuation",
			input:    "ANGLO-CARIBBEAN CO., LTD.",
			expected: "anglo caribbean co ltd",
		},
		{
			name:     "extra whitespace",
			input:    "  BANCO   NACIONAL  DE   CUBA  ",
			expected: "banco nacional de cuba",
		},
		{
			name:     "mixed case with special chars",
			input:    "Banco.Nacional_de@Cuba",
			expected: "banco nacional de cuba",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only special chars",
			input:    ".,!@#$%^&*()",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Name(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveInsignificantTerms(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "all significant terms",
			input:    "banco nacional cuba",
			expected: []string{"banco", "nacional", "cuba"},
		},
		{
			name:     "with noise terms",
			input:    "the banco of nacional and cuba",
			expected: []string{"banco", "nacional", "cuba"},
		},
		{
			name:     "with short terms",
			input:    "al banco de nacional",
			expected: []string{"banco", "nacional"},
		},
		{
			name:     "only noise terms",
			input:    "the of and in at",
			expected: nil,
		},
		{
			name:     "empty input",
			input:    "",
			expected: nil,
		},
		{
			name:     "mixed case terms",
			input:    "THE Banco OF Nacional",
			expected: []string{"Banco", "Nacional"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveInsignificantTerms(strings.Fields(tt.input))
			require.Equal(t, tt.expected, result)
		})
	}
}
