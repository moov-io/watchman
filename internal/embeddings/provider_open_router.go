package embeddings

import (
	"context"
	"fmt"
	"os"
	"time"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/operations"
	"github.com/ccoveille/go-safecast/v2"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
)

// OpenRouterProvider implements Provider using OpenAI-compatible APIs.
// Compatible with: OpenAI, Ollama, OpenRouter, Azure OpenAI, LMStudio, etc.
type OpenRouterProvider struct {
	config    ProviderConfig
	client    *openrouter.OpenRouter
	limiter   *rate.Limiter
	dimension int
	normalize bool
}

// NewOpenRouterProvider creates a new OpenAI-compatible embedding provider.
func NewOpenRouterProvider(config ProviderConfig) (*OpenRouterProvider, error) {
	if config.BaseURL == "" {
		return nil, ErrBaseURLRequired
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

	var args []openrouter.SDKOption
	if config.APIKey != "" {
		args = append(args, openrouter.WithSecurity(config.APIKey))
	}
	if config.BaseURL != "" {
		args = append(args, openrouter.WithServerURL(config.BaseURL))
	}
	if timeout > time.Second {
		args = append(args, openrouter.WithTimeout(timeout))
	}

	return &OpenRouterProvider{
		config:    config,
		dimension: config.Dimension,
		normalize: config.NormalizeVectors,
		client:    openrouter.New(args...),
		limiter:   rate.NewLimiter(rate.Limit(rps), burst),
	}, nil
}

// Embed generates embeddings for a batch of texts using the OpenAI-compatible API.
func (p *OpenRouterProvider) Embed(ctx context.Context, texts []string) ([][]float64, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	ctx, span := telemetry.StartSpan(ctx, "openai-embed", trace.WithAttributes(
		attribute.String("provider", p.Name()),
		attribute.Int("batch_size", len(texts)),
	))
	defer span.End()

	out := make([][]float64, len(texts))

	g, ctx := errgroup.WithContext(ctx)

	for i := range texts {
		i := i // capture for closure
		text := texts[i]

		g.Go(func() error {
			// rate limit check
			if err := p.limiter.Wait(ctx); err != nil {
				return fmt.Errorf("rate limit: %w", err)
			}

			req := operations.CreateEmbeddingsRequest{
				Input: operations.InputUnion{
					Str: openrouter.String(text),
				},
			}
			if p.config.Model != "" {
				req.Model = p.config.Model
			}

			resp, err := p.client.Embeddings.Generate(ctx, req)
			if err != nil {
				return fmt.Errorf("generating embeddings failed: %w", err)
			}

			// backoff := p.calculateBackoff(attempt) // TODO(adam): ??

			if body := resp.CreateEmbeddingsResponseBody; body != nil {
				if len(body.Data) > 0 {
					out[i] = body.Data[0].Embedding.ArrayOfNumber
				}

				// if body.Usage != nil { // TODO(adam):
				// 	fmt.Printf("  Tokens: Prompt=%.2f  Total=%.2f", body.Usage.PromptTokens, body.Usage.TotalTokens)
				// 	if body.Usage.Cost != nil {
				// 		fmt.Printf("  Cost=%.2f", *body.Usage.Cost)
				// 	}
				// 	fmt.Printf("\n")
				// }
			}

			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	// Normalize if configured
	if p.normalize {
		out = normalizeL2Batch(out)
	}

	// span.SetAttributes(
	// 	attribute.Int("tokens_used", embResp.Usage.TotalTokens), // TODO(adam):
	// )

	return out, nil
}

// calculateBackoff computes the backoff duration for a given retry attempt.
func (p *OpenRouterProvider) calculateBackoff(attempt int) time.Duration {
	initial := p.config.Retry.InitialBackoff
	if initial == 0 {
		initial = time.Second
	}
	maxBackoff := p.config.Retry.MaxBackoff
	if maxBackoff == 0 {
		maxBackoff = 30 * time.Second
	}

	shift, err := safecast.Convert[uint](attempt - 1)
	if err == nil {
		backoff := initial * time.Duration(1<<shift)
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
		return backoff
	}

	return initial
}

// Dimension returns the embedding dimension for this provider.
func (p *OpenRouterProvider) Dimension() int {
	return p.dimension
}

// Name returns the provider name for logging/telemetry.
func (p *OpenRouterProvider) Name() string {
	return fmt.Sprintf("openai-compatible(%s)", p.config.Model)
}

// Close releases any resources held by the provider.
func (p *OpenRouterProvider) Close() error {
	return nil // HTTP client doesn't need explicit cleanup
}
