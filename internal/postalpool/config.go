package postalpool

import (
	"time"
)

type Config struct {
	Enabled bool

	Instances      int
	StartingPort   int
	StartupTimeout time.Duration

	RequestTimeout time.Duration

	BinaryPath string
}
