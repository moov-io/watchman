// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package opensanctions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetailsURL(t *testing.T) {
	tests := []struct {
		entityID string
		expected string
	}{
		{
			entityID: "Q123456",
			expected: "https://www.opensanctions.org/entities/Q123456/",
		},
		{
			entityID: "NK-abc-123",
			expected: "https://www.opensanctions.org/entities/NK-abc-123/",
		},
	}

	for _, tc := range tests {
		t.Run(tc.entityID, func(t *testing.T) {
			url := DetailsURL(tc.entityID)
			require.Equal(t, tc.expected, url)
		})
	}
}
