package search

import (
	"runtime"
	"testing"

	"github.com/moov-io/watchman/internal/embeddings"
	"github.com/moov-io/watchman/internal/index"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestConfig_getMaxInFlight(t *testing.T) {
	cases := []struct {
		name     string
		conf     Config
		envValue string
		expected int
	}{
		{
			name:     "empty config and env",
			conf:     Config{},
			expected: runtime.GOMAXPROCS(0),
		},
		{
			name:     "from config field only",
			conf:     Config{MaxInFlight: 3},
			expected: 3,
		},
		{
			name:     "from env var only",
			conf:     Config{},
			envValue: "5",
			expected: 5,
		},
		{
			name:     "env var overrides config field",
			conf:     Config{MaxInFlight: 3},
			envValue: "5",
			expected: 5,
		},
		{
			name:     "env var with whitespace",
			conf:     Config{},
			envValue: " 7 ",
			expected: 7,
		},
		{
			name:     "unparsable env var keeps config field",
			conf:     Config{MaxInFlight: 3},
			envValue: "abc",
			expected: 3,
		},
		{
			name:     "zero env var keeps config field",
			conf:     Config{MaxInFlight: 3},
			envValue: "0",
			expected: 3,
		},
		{
			name:     "negative env var and config fall back to GOMAXPROCS",
			conf:     Config{MaxInFlight: -2},
			envValue: "-1",
			expected: runtime.GOMAXPROCS(0),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("SEARCH_MAX_IN_FLIGHT", tc.envValue)

			require.Equal(t, tc.expected, getMaxInFlight(tc.conf))
		})
	}
}

func TestNewService_MaxInFlightFromEnv(t *testing.T) {
	t.Setenv("SEARCH_MAX_IN_FLIGHT", "4")

	conf := Config{
		Goroutines:  DefaultConfig().Goroutines,
		MaxInFlight: 2,
	}
	svc, err := NewService(log.NewTestLogger(), conf, nil, index.NewLists(nil))
	require.NoError(t, err)

	require.Equal(t, 4, cap(svc.(*service).searchSem))
}

func TestNewService_EmbeddingsEnabledFromEnv(t *testing.T) {
	newService := func(t *testing.T, yamlEnabled bool) *service {
		t.Helper()

		emb := embeddings.DefaultConfig()
		emb.Enabled = yamlEnabled
		emb.Provider.Name = "mock"
		emb.Provider.BaseURL = "http://mock"
		emb.Provider.Model = "mock"
		emb.Provider.Dimension = 16

		conf := Config{
			Goroutines: DefaultConfig().Goroutines,
			Embeddings: emb,
		}
		svc, err := NewService(log.NewTestLogger(), conf, nil, index.NewLists(nil))
		require.NoError(t, err)

		return svc.(*service)
	}

	t.Run("env enables", func(t *testing.T) {
		t.Setenv("EMBEDDINGS_ENABLED", "true")

		svc := newService(t, false)
		require.NotNil(t, svc.embeddings)
	})

	t.Run("env disables", func(t *testing.T) {
		t.Setenv("EMBEDDINGS_ENABLED", "false")

		svc := newService(t, true)
		require.Nil(t, svc.embeddings)
	})

	t.Run("no env keeps yaml", func(t *testing.T) {
		t.Setenv("EMBEDDINGS_ENABLED", "")

		require.Nil(t, newService(t, false).embeddings)
		require.NotNil(t, newService(t, true).embeddings)
	})
}
