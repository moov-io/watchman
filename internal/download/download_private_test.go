package download

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestConfig_findExtraLists(t *testing.T) {
	cases := []struct {
		name           string
		config, loaded []search.SourceList
		expected       string
	}{
		{
			name:     "Loaded exact",
			config:   []search.SourceList{search.SourceUSOFAC},
			loaded:   []search.SourceList{search.SourceUSOFAC},
			expected: "",
		},
		{
			name:     "Extra Configured",
			config:   []search.SourceList{search.SourceUSOFAC, search.SourceUSCSL},
			loaded:   []search.SourceList{search.SourceUSOFAC},
			expected: "us_csl",
		},
		{
			name:     "Extra Configured",
			config:   []search.SourceList{search.SourceUSOFAC, search.SourceList("us-csl"), search.SourceList("unknown")},
			loaded:   []search.SourceList{search.SourceUSOFAC},
			expected: "us-csl, unknown",
		},
		{
			name:     "Extra Loaded",
			config:   []search.SourceList{search.SourceUSOFAC},
			loaded:   []search.SourceList{search.SourceUSOFAC, search.SourceUSCSL},
			expected: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := findExtraLists(tc.config, tc.loaded)
			require.Equal(t, tc.expected, got)
		})
	}
}
