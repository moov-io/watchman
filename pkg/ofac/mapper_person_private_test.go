package ofac

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAltNames(t *testing.T) {
	tests := []struct {
		remarks  []string
		expected []string
	}{
		{
			remarks:  []string{"a.k.a. 'SMITH, John'"},
			expected: []string{"SMITH, John"},
		},
		{
			remarks:  []string{"a.k.a. 'SMITH, John'; a.k.a. 'DOE, Jane'"},
			expected: []string{"SMITH, John", "DOE, Jane"},
		},
		{
			remarks:  []string{"Some other remark", "a.k.a. 'SMITH, John'"},
			expected: []string{"SMITH, John"},
		},
		{
			remarks:  []string{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		result := parseAltNames(tt.remarks)
		require.Equal(t, tt.expected, result)
	}
}
