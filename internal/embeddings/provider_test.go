package embeddings_test

import (
	"os"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/embeddings"

	"github.com/stretchr/testify/require"
)

func TestProviders(t *testing.T) {
	forEachProvider(t, func(p embeddings.Provider) {
		require.NotEmpty(t, p.Name())
		require.Greater(t, p.Dimension(), 0)

		t.Cleanup(func() { p.Close() })

		inputs := []string{
			"hello world",
			"testing12345",
			"the quick brown fox jumped over the lazy dog",
		}

		results, err := p.Embed(t.Context(), inputs)
		require.NoError(t, err)
		require.Len(t, results, len(inputs))

		for _, values := range results {
			require.Len(t, values, p.Dimension())
		}
	})
}

func forEachProvider(t *testing.T, fn func(p embeddings.Provider)) {
	t.Helper()

	run := func(t *testing.T, apiKey string, config embeddings.ProviderConfig) {
		t.Helper()

		if apiKey == "" {
			t.Skip("missing api key")
		}

		p, err := embeddings.NewOpenRouterProvider(config)
		require.NoError(t, err)

		fn(p)
	}

	t.Run("ollama", func(t *testing.T) {
		apiKey := os.Getenv("OLLAMA_API_KEY")

		run(t, apiKey, embeddings.ProviderConfig{
			Name:             "ollama",
			BaseURL:          "http://localhost:11434/v1",
			APIKey:           apiKey,
			Model:            "qwen3-embedding",
			Dimension:        4096,
			NormalizeVectors: true,
			Timeout:          10 * time.Second,
		})
	})

	t.Run("openai", func(t *testing.T) {
		if testing.Short() {
			t.Skip("-short flag provided")
		}

		apiKey := os.Getenv("OPENAI_API_KEY")
		run(t, apiKey, embeddings.ProviderConfig{
			Name:             "openai",
			BaseURL:          "https://api.openai.com/v1/",
			APIKey:           apiKey,
			Model:            "text-embedding-3-small",
			Dimension:        1536,
			NormalizeVectors: true,
			Timeout:          10 * time.Second,
		})
	})

	t.Run("open router", func(t *testing.T) {
		if testing.Short() {
			t.Skip("-short flag provided")
		}

		apiKey := os.Getenv("OPENROUTER_API_KEY")
		run(t, apiKey, embeddings.ProviderConfig{
			Name:             "openrouter",
			BaseURL:          "https://openrouter.ai/api/v1",
			APIKey:           apiKey,
			Model:            "qwen/qwen3-embedding-8b",
			Dimension:        4096,
			NormalizeVectors: true,
			Timeout:          10 * time.Second,
		})
	})
}
