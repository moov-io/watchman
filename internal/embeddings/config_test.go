package embeddings

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfig_LoadFromEnv(t *testing.T) {
	cases := []struct {
		name     string
		env      map[string]string
		conf     Config
		expected Config
	}{
		{
			name:     "no env keeps config",
			conf:     DefaultConfig(),
			expected: DefaultConfig(),
		},
		{
			name: "enabled and cross script only from env",
			env: map[string]string{
				"EMBEDDINGS_ENABLED":           "true",
				"EMBEDDINGS_CROSS_SCRIPT_ONLY": "no",
			},
			conf: Config{CrossScriptOnly: true},
			expected: Config{
				Enabled:         true,
				CrossScriptOnly: false,
			},
		},
		{
			name: "enabled can be turned off by env",
			env: map[string]string{
				"EMBEDDINGS_ENABLED": "false",
			},
			conf:     Config{Enabled: true},
			expected: Config{Enabled: false},
		},
		{
			name: "numbers and duration from env",
			env: map[string]string{
				"EMBEDDINGS_CACHE_SIZE":          "5",
				"EMBEDDINGS_BATCH_SIZE":          "8",
				"EMBEDDINGS_INDEX_BUILD_TIMEOUT": "2m",
				"EMBEDDINGS_DIMENSION":           "4",
			},
			conf: Config{},
			expected: Config{
				Provider:          ProviderConfig{Dimension: 4},
				Cache:             CacheConfig{Size: 5},
				BatchSize:         8,
				IndexBuildTimeout: 2 * time.Minute,
			},
		},
		{
			name: "invalid values are ignored",
			env: map[string]string{
				"EMBEDDINGS_CACHE_SIZE":          "abc",
				"EMBEDDINGS_BATCH_SIZE":          "0",
				"EMBEDDINGS_INDEX_BUILD_TIMEOUT": "soon",
				"EMBEDDINGS_DIMENSION":           "-1",
			},
			conf: Config{
				Provider:          ProviderConfig{Dimension: 3},
				Cache:             CacheConfig{Size: 10},
				BatchSize:         32,
				IndexBuildTimeout: time.Minute,
			},
			expected: Config{
				Provider:          ProviderConfig{Dimension: 3},
				Cache:             CacheConfig{Size: 10},
				BatchSize:         32,
				IndexBuildTimeout: time.Minute,
			},
		},
		{
			name: "strings from env",
			env: map[string]string{
				"EMBEDDINGS_API_KEY":  "key",
				"EMBEDDINGS_BASE_URL": "http://localhost:11434/v1",
				"EMBEDDINGS_MODEL":    "m",
			},
			conf: Config{Provider: ProviderConfig{Model: "yaml"}},
			expected: Config{Provider: ProviderConfig{
				APIKey:  "key",
				BaseURL: "http://localhost:11434/v1",
				Model:   "m",
			}},
		},
	}

	names := []string{
		"EMBEDDINGS_ENABLED", "EMBEDDINGS_API_KEY", "EMBEDDINGS_BASE_URL", "EMBEDDINGS_MODEL", "EMBEDDINGS_DIMENSION",
		"EMBEDDINGS_CACHE_SIZE", "EMBEDDINGS_CROSS_SCRIPT_ONLY",
		"EMBEDDINGS_BATCH_SIZE", "EMBEDDINGS_INDEX_BUILD_TIMEOUT",
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, name := range names {
				t.Setenv(name, tc.env[name])
			}

			conf := tc.conf
			conf.LoadFromEnv()

			require.Equal(t, tc.expected, conf)
		})
	}
}
