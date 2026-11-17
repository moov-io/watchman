package decisions

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
			name: "enabled from env",
			env: map[string]string{
				"DECISIONS_ENABLED": "true",
			},
			conf:     Config{},
			expected: Config{Enabled: true},
		},
		{
			name: "enabled can be turned off by env",
			env: map[string]string{
				"DECISIONS_ENABLED": "false",
			},
			conf:     Config{Enabled: true},
			expected: Config{Enabled: false},
		},
		{
			name: "strings numbers and duration from env",
			env: map[string]string{
				"DECISIONS_API_KEY":      "secret",
				"DECISIONS_BASE_URL":     "http://127.0.0.1:11434/v1",
				"DECISIONS_MODEL":        "tev1",
				"DECISIONS_TIMEOUT":      "2s",
				"DECISIONS_MAX_ENTITIES": "4",
				"DECISIONS_MIN_MATCH":    "0.5",
			},
			conf: Config{},
			expected: Config{
				APIKey:      "secret",
				BaseURL:     "http://127.0.0.1:11434/v1",
				Model:       "tev1",
				Timeout:     2 * time.Second,
				MaxEntities: 4,
				MinMatch:    0.5,
			},
		},
		{
			name: "invalid values are ignored",
			env: map[string]string{
				"DECISIONS_TIMEOUT":      "soon",
				"DECISIONS_MAX_ENTITIES": "0",
				"DECISIONS_MIN_MATCH":    "1.5",
			},
			conf: Config{
				Timeout:     time.Second,
				MaxEntities: 3,
				MinMatch:    0.2,
			},
			expected: Config{
				Timeout:     time.Second,
				MaxEntities: 3,
				MinMatch:    0.2,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ClearTestEnv(t.Setenv)
			for key, value := range tc.env {
				t.Setenv(key, value)
			}

			conf := tc.conf
			conf.LoadFromEnv()
			require.Equal(t, tc.expected, conf)
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	t.Run("disabled skips validation", func(t *testing.T) {
		require.NoError(t, Config{}.Validate())
	})

	t.Run("enabled requires base URL and model", func(t *testing.T) {
		require.ErrorIs(t, Config{Enabled: true, Model: "nimble"}.Validate(), ErrBaseURLRequired)
		require.ErrorIs(t, Config{Enabled: true, BaseURL: "http://localhost:11434/v1"}.Validate(), ErrModelRequired)
	})

	t.Run("enabled valid", func(t *testing.T) {
		require.NoError(t, DefaultConfig().Validate())
		conf := DefaultConfig()
		conf.Enabled = true
		require.NoError(t, conf.Validate())
	})

	t.Run("bounds", func(t *testing.T) {
		conf := DefaultConfig()
		conf.Enabled = true
		conf.Timeout = -time.Second
		require.ErrorIs(t, conf.Validate(), ErrInvalidTimeout)

		conf = DefaultConfig()
		conf.Enabled = true
		conf.MaxEntities = -1
		require.ErrorIs(t, conf.Validate(), ErrInvalidMaxEntities)

		conf = DefaultConfig()
		conf.Enabled = true
		conf.MinMatch = 1.2
		require.ErrorIs(t, conf.Validate(), ErrInvalidMinMatch)
	})
}

func TestConfig_MaxEntitiesCap(t *testing.T) {
	require.Equal(t, 10, Config{}.maxEntitiesOrDefault())
	require.Equal(t, 4, Config{MaxEntities: 4}.maxEntitiesOrDefault())
	require.Equal(t, hardMaxEntities, Config{MaxEntities: 100}.maxEntitiesOrDefault())
	require.Equal(t, 21, hardMaxEntities)
	require.LessOrEqual(t, hardMaxEntities*questionsPerHit, ollamaMaxQuestions)
}

func TestSystemOneURL(t *testing.T) {
	require.Equal(t, "http://localhost:11434/v1/systemone", systemOneURL("http://localhost:11434/v1"))
	require.Equal(t, "http://localhost:11434/v1/systemone", systemOneURL("http://localhost:11434"))
	require.Equal(t, "https://api.typesafe.ai/v1/systemone", systemOneURL("https://api.typesafe.ai/v1/"))
	require.Equal(t, "https://api.typesafe.ai/v1/systemone", systemOneURL("https://api.typesafe.ai/v1/systemone"))
}
