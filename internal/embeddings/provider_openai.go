//go:build embeddings

package embeddings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/time/rate"
)

// OpenAIProvider implements EmbeddingProvider using OpenAI-compatible APIs.
// Compatible with: OpenAI, Ollama, OpenRouter, Azure OpenAI, LMStudio, etc.
type OpenAIProvider struct {
	config    ProviderConfig
	client    *http.Client
	limiter   *rate.Limiter
	dimension int
	normalize bool
}

// embeddingRequest represents the OpenAI embeddings API request.
type embeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

// embeddingResponse represents the OpenAI embeddings API response.
type embeddingResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Index     int       `json:"index"`
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
	Error *apiError `json:"error,omitempty"`
}

// apiError represents an error from the API.
type apiError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

// NewOpenAIProvider creates a new OpenAI-compatible embedding provider.
func NewOpenAIProvider(config ProviderConfig) (*OpenAIProvider, error) {
	if config.BaseURL == "" {
		return nil, ErrBaseURLRequired
	}
	if config.Model == "" {
		return nil, ErrModelRequired
	}
	if config.Dimension <= 0 {
		return nil, ErrInvalidDimension
	}

	// Apply environment variable overrides
	if apiKey := os.Getenv("EMBEDDINGS_API_KEY"); apiKey != "" && config.APIKey == "" {
		config.APIKey = apiKey
	}

	// Set defaults
	timeout := config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	rps := config.RateLimit.RequestsPerSecond
	if rps <= 0 {
		rps = 10
	}
	burst := config.RateLimit.Burst
	if burst <= 0 {
		burst = 20
	}

	return &OpenAIProvider{
		config:    config,
		dimension: config.Dimension,
		normalize: config.NormalizeVectors,
		client: &http.Client{
			Timeout: timeout,
		},
		limiter: rate.NewLimiter(rate.Limit(rps), burst),
	}, nil
}

// Embed generates embeddings for a batch of texts using the OpenAI-compatible API.
func (p *OpenAIProvider) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	ctx, span := telemetry.StartSpan(ctx, "openai-embed", trace.WithAttributes(
		attribute.String("provider", p.Name()),
		attribute.Int("batch_size", len(texts)),
	))
	defer span.End()

	// Rate limit
	if err := p.limiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limit: %w", err)
	}

	// Build request
	reqBody := embeddingRequest{
		Model: p.config.Model,
		Input: texts,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/embeddings", p.config.BaseURL)

	// Execute request with retry
	var resp *http.Response
	var lastErr error
	maxRetries := p.config.Retry.MaxRetries
	if maxRetries <= 0 {
		maxRetries = 3
	}

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			backoff := p.calculateBackoff(attempt)
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoff):
			}
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			return nil, fmt.Errorf("create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		if p.config.APIKey != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.config.APIKey))
		}

		// Apply custom headers
		for k, v := range p.config.Headers {
			req.Header.Set(k, v)
		}

		resp, lastErr = p.client.Do(req)
		if lastErr == nil && resp.StatusCode < 500 {
			break
		}
		if resp != nil {
			resp.Body.Close()
		}
	}

	if lastErr != nil {
		return nil, fmt.Errorf("request failed: %w", lastErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse response
	var embResp embeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&embResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if embResp.Error != nil {
		return nil, fmt.Errorf("API error: %s", embResp.Error.Message)
	}

	// Extract embeddings in correct order
	result := make([][]float32, len(texts))
	for _, item := range embResp.Data {
		if item.Index >= len(result) {
			continue
		}
		result[item.Index] = item.Embedding
	}

	// Validate all embeddings were returned
	for i, emb := range result {
		if emb == nil {
			return nil, fmt.Errorf("%w: missing embedding at index %d", ErrInvalidResponse, i)
		}
	}

	// Normalize if configured
	if p.normalize {
		result = normalizeL2Batch(result)
	}

	span.SetAttributes(
		attribute.Int("tokens_used", embResp.Usage.TotalTokens),
	)

	return result, nil
}

// calculateBackoff computes the backoff duration for a given retry attempt.
func (p *OpenAIProvider) calculateBackoff(attempt int) time.Duration {
	initial := p.config.Retry.InitialBackoff
	if initial == 0 {
		initial = time.Second
	}
	maxBackoff := p.config.Retry.MaxBackoff
	if maxBackoff == 0 {
		maxBackoff = 30 * time.Second
	}

	backoff := initial * time.Duration(1<<uint(attempt-1))
	if backoff > maxBackoff {
		backoff = maxBackoff
	}
	return backoff
}

// Dimension returns the embedding dimension for this provider.
func (p *OpenAIProvider) Dimension() int {
	return p.dimension
}

// Name returns the provider name for logging/telemetry.
func (p *OpenAIProvider) Name() string {
	return fmt.Sprintf("openai-compatible(%s)", p.config.Model)
}

// Close releases any resources held by the provider.
func (p *OpenAIProvider) Close() error {
	return nil // HTTP client doesn't need explicit cleanup
}
