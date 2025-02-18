package norm_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/norm"

	"github.com/stretchr/testify/require"
)

func TestPhoneNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid US number",
			input:    "123-456-7890",
			expected: "1234567890",
		},
		{
			name:     "valid international number",
			input:    "+44 20 7123 4567",
			expected: "442071234567",
		},
		{
			name:     "valid number with parentheses",
			input:    "(555) 123-4567",
			expected: "5551234567",
		},
		{
			name:     "number with country code without plus",
			input:    "44 20 7123 4567",
			expected: "442071234567",
		},
		{
			name:     "US number with country code",
			input:    "1-555-123-4567",
			expected: "15551234567",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "too short",
			input:    "123456",
			expected: "123456",
		},
		{
			name:     "too long",
			input:    "123456789012345678",
			expected: "123456789012345678",
		},
		{
			name:     "invalid characters",
			input:    "123-ABC-7890",
			expected: "123ABC7890",
		},
		{
			name:     "multiple plus signs",
			input:    "+1+2345678901",
			expected: "12345678901",
		},
		{
			name:     "plus sign in middle",
			input:    "123+4567890",
			expected: "1234567890",
		},
		{
			name:     "valid Indian number",
			input:    "+91 98765 43210",
			expected: "919876543210",
		},
		{
			name:     "valid Chinese number",
			input:    "+86 123 4567 8901",
			expected: "8612345678901",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := norm.PhoneNumber(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
