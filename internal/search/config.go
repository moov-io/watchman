package search

import (
	"cmp"
	"os"
	"runtime"
	"strconv"
)

type Config struct {
	SearchGroups SearchGroups
}

type SearchGroups struct {
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

	return Config{
		SearchGroups: SearchGroups{
			Default: cpus * 2,
			Min:     cpus,
			Max:     cpus * 4,
		},
	}
}
