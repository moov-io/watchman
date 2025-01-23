package search

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_buildQueryParameters(t *testing.T) {
	cases := []struct {
		entity   Entity[Value]
		opts     SearchOpts
		expected map[string][]string
	}{
		{
			entity: Entity[Value]{
				Name: "john doe",
				Type: EntityPerson,
				Person: &Person{
					AltNames:  []string{"jon doe", "johnny doe"},
					Gender:    "male",
					BirthDate: ptr(time.Date(1998, time.April, 12, 10, 30, 0, 0, time.UTC)),
				},
			},
			opts: SearchOpts{
				Limit:    3,
				MinMatch: 0.9,
			},
			expected: map[string][]string{
				"name":      []string{"john doe"},
				"type":      []string{"person"},
				"altNames":  []string{"jon doe", "johnny doe"},
				"gender":    []string{"male"},
				"birthDate": []string{"1998-04-12"},
				"limit":     []string{"3"},
				"minMatch":  []string{"0.90"},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.entity.Name, func(t *testing.T) {
			q := make(url.Values)
			got := buildQueryParameters(q, tc.entity, tc.opts)

			require.Equal(t, len(tc.expected), len(got)) // same key size

			for expectedKey, expectedValues := range tc.expected {
				gotValues, found := got[expectedKey]
				require.True(t, found, fmt.Sprintf("looking for %s", expectedKey))
				require.ElementsMatch(t, expectedValues, gotValues)
			}
		})
	}
}

func ptr[T any](in T) *T {
	return &in
}
