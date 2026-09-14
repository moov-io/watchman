package search

import (
	"runtime"
	"testing"

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
