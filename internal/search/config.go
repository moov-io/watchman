package search

import (
	"cmp"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/moov-io/watchman/internal/embeddings"
)

type Config struct {
	Goroutines Goroutines
	Embeddings embeddings.Config

	// MaxInFlight bounds concurrent full-list searches (admission control).
	// Zero means GOMAXPROCS. Override with SEARCH_MAX_IN_FLIGHT.
	MaxInFlight int
}

type Goroutines struct {
	Default int
	Min     int
	Max     int
}

func DefaultConfig() Config {
	cpus := runtime.NumCPU()

	if v := os.Getenv("GOMAXPROCS"); v != "" {
		n, _ := strconv.ParseInt(v, 10, 8)
		cpus = cmp.Or(cpus, int(n))
	}

	inFlight := runtime.GOMAXPROCS(0)
	if v := strings.TrimSpace(os.Getenv("SEARCH_MAX_IN_FLIGHT")); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			inFlight = n
		}
	}

	return Config{
		Goroutines: Goroutines{
			Default: cpus * 2,
			Min:     cpus,
			Max:     cpus * 4,
		},
		Embeddings:  embeddings.DefaultConfig(),
		MaxInFlight: inFlight,
	}
}

// getMaxInFlight returns how many large searches may run at once (admission control).
// The SEARCH_MAX_IN_FLIGHT environment variable takes precedence over the
// YAML-configured MaxInFlight field when it holds a positive integer.
func getMaxInFlight(conf Config) int {
	if v := strings.TrimSpace(os.Getenv("SEARCH_MAX_IN_FLIGHT")); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	if conf.MaxInFlight > 0 {
		return conf.MaxInFlight
	}
	return max(runtime.GOMAXPROCS(0), 1)
}
