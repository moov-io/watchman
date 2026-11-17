package decisions

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/base/strx"
)

const (
	defaultBaseURL     = "http://localhost:11434/v1"
	defaultModel       = "nimble"
	defaultTimeout     = 5 * time.Second
	defaultMaxEntities = 10
	questionsPerHit    = 3
	// Ollama System One caps questions at 64 (3 per hit → 21 entities).
	ollamaMaxQuestions = 64
	hardMaxEntities    = ollamaMaxQuestions / questionsPerHit
)

// Config controls the optional Jev-style decision model used to annotate
// search hits. Disabled by default. When enabled, Watchman POSTs the query
// and top entities to POST {BaseURL}/systemone (Ollama 0.35+ or a hosted
// TypeSafe-compatible API) and attaches typed answers on each hit.
type Config struct {
	Enabled bool `json:"enabled"`

	// BaseURL is the OpenAI-style /v1 root.
	// Examples:
	//   - Ollama: "http://localhost:11434/v1"
	//   - TypeSafe hosted: "https://api.typesafe.ai/v1"
	BaseURL string `json:"baseURL"`

	// APIKey is sent as Authorization Bearer. Optional for local Ollama.
	// Override with DECISIONS_API_KEY.
	APIKey string `json:"apiKey,omitempty"`

	// Model is the decision model id.
	// Examples: nimble, tev1, tev1:0.8b (Ollama); jev-latest (hosted).
	Model string `json:"model"`

	// Timeout for the decision request. Default 5s.
	Timeout time.Duration `json:"timeout"`

	// MaxEntities is the maximum number of top search hits to annotate.
	// Default 10. Hard cap 21 so 3 questions per hit stay within Ollama's
	// 64-question System One limit.
	MaxEntities int `json:"maxEntities"`

	// MinMatch skips hits below this Watchman score. Default 0 (annotate all
	// returned hits, up to MaxEntities).
	MinMatch float64 `json:"minMatch"`
}

func DefaultConfig() Config {
	return Config{
		Enabled:     false,
		BaseURL:     defaultBaseURL,
		Model:       defaultModel,
		Timeout:     defaultTimeout,
		MaxEntities: defaultMaxEntities,
	}
}

func (c *Config) LoadFromEnv() {
	if enabled := strings.TrimSpace(os.Getenv("DECISIONS_ENABLED")); enabled != "" {
		c.Enabled = strx.Yes(enabled)
	}
	if apiKey := os.Getenv("DECISIONS_API_KEY"); apiKey != "" {
		c.APIKey = apiKey
	}
	if baseURL := os.Getenv("DECISIONS_BASE_URL"); baseURL != "" {
		c.BaseURL = baseURL
	}
	if model := os.Getenv("DECISIONS_MODEL"); model != "" {
		c.Model = model
	}
	if timeout := os.Getenv("DECISIONS_TIMEOUT"); timeout != "" {
		if d, err := time.ParseDuration(timeout); err == nil {
			c.Timeout = d
		}
	}
	if maxEntities := os.Getenv("DECISIONS_MAX_ENTITIES"); maxEntities != "" {
		if n, err := strconv.Atoi(maxEntities); err == nil && n > 0 {
			c.MaxEntities = n
		}
	}
	if minMatch := os.Getenv("DECISIONS_MIN_MATCH"); minMatch != "" {
		if n, err := strconv.ParseFloat(minMatch, 64); err == nil && n >= 0 && n <= 1 {
			c.MinMatch = n
		}
	}
}

func (c Config) Validate() error {
	if !c.Enabled {
		return nil
	}
	if strings.TrimSpace(c.BaseURL) == "" {
		return ErrBaseURLRequired
	}
	if strings.TrimSpace(c.Model) == "" {
		return ErrModelRequired
	}
	if c.Timeout < 0 {
		return ErrInvalidTimeout
	}
	if c.MaxEntities < 0 {
		return ErrInvalidMaxEntities
	}
	if c.MinMatch < 0 || c.MinMatch > 1 {
		return ErrInvalidMinMatch
	}
	return nil
}

func (c Config) timeoutOrDefault() time.Duration {
	if c.Timeout > 0 {
		return c.Timeout
	}
	return defaultTimeout
}

func (c Config) maxEntitiesOrDefault() int {
	n := c.MaxEntities
	if n <= 0 {
		n = defaultMaxEntities
	}
	if n > hardMaxEntities {
		return hardMaxEntities
	}
	return n
}
