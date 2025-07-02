package linksim

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSortKeys(t *testing.T) {
	cases := []struct {
		entity   search.Entity[search.Value]
		expected []string
	}{
		{
			entity: john,
			expected: []string{
				"TYPE:0230",
				"NAME:0190",
				"GOVID:C0173|T0190|X0146",
				"ADDR:C0143|S0021|P0007|Y0023|L0201,0028,0173",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.entity.Name, func(t *testing.T) {
			hashes := GenerateHashes(tc.entity)
			buckets := hashes.GenerateBuckets(DefaultBuckets)
			keys := SortKeys(buckets)

			require.ElementsMatch(t, tc.expected, keys)
		})
	}
}
