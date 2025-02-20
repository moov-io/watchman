package postalpool

import (
	"net"
	"net/http"
	"time"
)

type Config struct {
	Enabled bool

	Instances      int
	StartingPort   int
	StartupTimeout time.Duration

	RequestTimeout time.Duration

	Dialer    *net.Dialer
	Transport *http.Transport

	BinaryPath       string
	CGOSelfInstances int
}
