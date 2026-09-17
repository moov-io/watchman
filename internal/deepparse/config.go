package deepparse

import "time"

// Config controls the optional deepparse HTTP parser.
// Disabled by default; libpostal / usaddress remain the built-in parsers.
type Config struct {
	Enabled bool

	// BaseURL is the origin of a running deepparse FastAPI, for example
	// "http://localhost:8000". Overridable with DEEPPARSE_URL.
	BaseURL string

	// Model is the deepparse path segment (bpemb, bpemb-attention, fasttext, ...).
	// The published Docker image loads BPEmb models.
	Model string

	Timeout time.Duration
}
