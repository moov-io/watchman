package download

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestConfig_getIgnoredDownloadErrors(t *testing.T) {
	cases := []struct {
		name     string
		conf     Config
		envValue string
		expected []search.SourceList
	}{
		{
			name:     "empty config",
			conf:     Config{},
			expected: []search.SourceList{},
		},
		{
			name: "from config field only",
			conf: Config{
				IgnoredDownloadErrors: []search.SourceList{search.SourceUSOFAC, search.SourceEUCSL},
			},
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC}, // sorted
		},
		{
			name: "from config field normalized to lowercase and trimmed",
			conf: Config{
				IgnoredDownloadErrors: []search.SourceList{" US_OFAC ", "EU_CSL", "\teu_csl\n"},
			},
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
		{
			name:     "from env var only",
			conf:     Config{},
			envValue: "us_ofac,eu_csl",
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC}, // sorted
		},
		{
			name: "merged and deduped",
			conf: Config{
				IgnoredDownloadErrors: []search.SourceList{"US_OFAC "},
			},
			envValue: " eu_csl ,US_OFAC",
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
		{
			name:     "blanks and spaces ignored",
			conf:     Config{},
			envValue: " us_ofac , , eu_csl , ",
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
		{
			name:     "normalized to lowercase",
			conf:     Config{},
			envValue: "US_OFAC,EU_CSL",
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envValue != "" {
				t.Setenv("IGNORED_DOWNLOAD_ERRORS", tc.envValue)
			}
			got := getIgnoredDownloadErrors(tc.conf)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestConfig_getIncludedLists(t *testing.T) {
	cases := []struct {
		name     string
		conf     Config
		envValue string
		expected []search.SourceList
	}{
		{
			name:     "empty config",
			conf:     Config{},
			expected: []search.SourceList{},
		},
		{
			name: "normalizes IncludedLists from config",
			conf: Config{
				IncludedLists: []search.SourceList{" US_OFAC ", "EU_CSL"},
			},
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
		{
			name:     "normalizes INCLUDED_LISTS from env",
			conf:     Config{},
			envValue: " us_ofac ,EU_CSL ",
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
		{
			name: "merged and deduped across config and env",
			conf: Config{
				IncludedLists: []search.SourceList{"US_OFAC"},
			},
			envValue: "eu_csl, us_ofac ",
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
		{
			name: "normalizes custom Senzing list names",
			conf: Config{
				Senzing: []SenzingList{
					{SourceList: " MyCustomList ", Location: "file:///dev/null"},
				},
			},
			expected: []search.SourceList{"mycustomlist"},
		},
		{
			name:     "blanks and spaces ignored in env",
			conf:     Config{},
			envValue: " , us_ofac , , eu_csl , ",
			expected: []search.SourceList{search.SourceEUCSL, search.SourceUSOFAC},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envValue != "" {
				t.Setenv("INCLUDED_LISTS", tc.envValue)
			}
			got := getIncludedLists(tc.conf)
			require.Equal(t, tc.expected, got)
		})
	}
}

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
