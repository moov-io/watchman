package embeddings

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"sync/atomic"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// )

// func TestNewOpenAIProvider_Validation(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		config  ProviderConfig
// 		wantErr error
// 	}{
// 		{
// 			name:    "missing base URL",
// 			config:  ProviderConfig{Model: "test", Dimension: 768},
// 			wantErr: ErrBaseURLRequired,
// 		},
// 		{
// 			name:    "missing model",
// 			config:  ProviderConfig{BaseURL: "http://localhost", Dimension: 768},
// 			wantErr: ErrModelRequired,
// 		},
// 		{
// 			name:    "invalid dimension",
// 			config:  ProviderConfig{BaseURL: "http://localhost", Model: "test", Dimension: 0},
// 			wantErr: ErrInvalidDimension,
// 		},
// 		{
// 			name: "valid config",
// 			config: ProviderConfig{
// 				BaseURL:   "http://localhost:11434/v1",
// 				Model:     "nomic-embed-text",
// 				Dimension: 768,
// 			},
// 			wantErr: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := NewOpenAIProvider(tt.config)
// 			if tt.wantErr != nil {
// 				require.ErrorIs(t, err, tt.wantErr)
// 			} else {
// 				require.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestOpenAIProvider_Embed(t *testing.T) {
// 	// Mock server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		require.Equal(t, "/embeddings", r.URL.Path)
// 		require.Equal(t, "application/json", r.Header.Get("Content-Type"))
// 		require.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))

// 		var req embeddingRequest
// 		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
// 		require.Equal(t, "test-model", req.Model)

// 		resp := embeddingResponse{
// 			Object: "list",
// 			Data: make([]struct {
// 				Object    string    `json:"object"`
// 				Index     int       `json:"index"`
// 				Embedding []float64 `json:"embedding"`
// 			}, len(req.Input)),
// 			Usage: struct {
// 				PromptTokens int `json:"prompt_tokens"`
// 				TotalTokens  int `json:"total_tokens"`
// 			}{
// 				PromptTokens: 10,
// 				TotalTokens:  10,
// 			},
// 		}
// 		for i := range req.Input {
// 			resp.Data[i].Index = i
// 			resp.Data[i].Object = "embedding"
// 			resp.Data[i].Embedding = make([]float64, 768)
// 			for j := range resp.Data[i].Embedding {
// 				resp.Data[i].Embedding[j] = float64(i+j) * 0.001
// 			}
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(resp)
// 	}))
// 	defer server.Close()

// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		Name:      "openai",
// 		BaseURL:   server.URL,
// 		APIKey:    "test-key",
// 		Model:     "test-model",
// 		Dimension: 768,
// 		Timeout:   5 * time.Second,
// 	})
// 	require.NoError(t, err)

// 	embeddings, err := provider.Embed(context.Background(), []string{"hello", "world"})
// 	require.NoError(t, err)
// 	require.Len(t, embeddings, 2)
// 	require.Len(t, embeddings[0], 768)
// 	require.Len(t, embeddings[1], 768)
// }

// func TestOpenAIProvider_Embed_WithNormalization(t *testing.T) {
// 	// Mock server returns unnormalized vectors
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		resp := embeddingResponse{
// 			Object: "list",
// 			Data: []struct {
// 				Object    string    `json:"object"`
// 				Index     int       `json:"index"`
// 				Embedding []float64 `json:"embedding"`
// 			}{
// 				{Index: 0, Embedding: []float64{3, 4}}, // L2 norm = 5, normalized = [0.6, 0.8]
// 			},
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(resp)
// 	}))
// 	defer server.Close()

// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:          server.URL,
// 		Model:            "test",
// 		Dimension:        2,
// 		NormalizeVectors: true, // Enable normalization
// 	})
// 	require.NoError(t, err)

// 	embeddings, err := provider.Embed(context.Background(), []string{"test"})
// 	require.NoError(t, err)
// 	require.Len(t, embeddings, 1)

// 	// Check normalization was applied
// 	require.InDelta(t, 0.6, float64(embeddings[0][0]), 0.001)
// 	require.InDelta(t, 0.8, float64(embeddings[0][1]), 0.001)
// }

// func TestOpenAIProvider_Embed_EmptyInput(t *testing.T) {
// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   "http://localhost",
// 		Model:     "test",
// 		Dimension: 768,
// 	})
// 	require.NoError(t, err)

// 	embeddings, err := provider.Embed(context.Background(), []string{})
// 	require.NoError(t, err)
// 	require.Nil(t, embeddings)
// }

// func TestOpenAIProvider_Embed_CustomHeaders(t *testing.T) {
// 	receivedHeaders := make(map[string]string)

// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		receivedHeaders["X-Custom-Header"] = r.Header.Get("X-Custom-Header")
// 		receivedHeaders["HTTP-Referer"] = r.Header.Get("HTTP-Referer")

// 		resp := embeddingResponse{
// 			Object: "list",
// 			Data: []struct {
// 				Object    string    `json:"object"`
// 				Index     int       `json:"index"`
// 				Embedding []float64 `json:"embedding"`
// 			}{
// 				{Index: 0, Embedding: make([]float64, 768)},
// 			},
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(resp)
// 	}))
// 	defer server.Close()

// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   server.URL,
// 		Model:     "test",
// 		Dimension: 768,
// 		Headers: map[string]string{
// 			"X-Custom-Header": "custom-value",
// 			"HTTP-Referer":    "https://myapp.com",
// 		},
// 	})
// 	require.NoError(t, err)

// 	_, err = provider.Embed(context.Background(), []string{"test"})
// 	require.NoError(t, err)

// 	require.Equal(t, "custom-value", receivedHeaders["X-Custom-Header"])
// 	require.Equal(t, "https://myapp.com", receivedHeaders["HTTP-Referer"])
// }

// func TestOpenAIProvider_Embed_APIError(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(embeddingResponse{
// 			Error: &apiError{
// 				Message: "Invalid model",
// 				Type:    "invalid_request_error",
// 				Code:    "model_not_found",
// 			},
// 		})
// 	}))
// 	defer server.Close()

// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   server.URL,
// 		Model:     "invalid-model",
// 		Dimension: 768,
// 	})
// 	require.NoError(t, err)

// 	_, err = provider.Embed(context.Background(), []string{"test"})
// 	require.Error(t, err)
// 	require.Contains(t, err.Error(), "400")
// }

// func TestOpenAIProvider_Retry(t *testing.T) {
// 	var attempts int32
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		count := atomic.AddInt32(&attempts, 1)
// 		if count < 3 {
// 			w.WriteHeader(http.StatusServiceUnavailable)
// 			return
// 		}
// 		resp := embeddingResponse{
// 			Object: "list",
// 			Data: []struct {
// 				Object    string    `json:"object"`
// 				Index     int       `json:"index"`
// 				Embedding []float64 `json:"embedding"`
// 			}{
// 				{Index: 0, Embedding: make([]float64, 768)},
// 			},
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(resp)
// 	}))
// 	defer server.Close()

// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   server.URL,
// 		Model:     "test",
// 		Dimension: 768,
// 		Retry: RetryConfig{
// 			MaxRetries:     3,
// 			InitialBackoff: 10 * time.Millisecond,
// 			MaxBackoff:     100 * time.Millisecond,
// 		},
// 	})
// 	require.NoError(t, err)

// 	_, err = provider.Embed(context.Background(), []string{"test"})
// 	require.NoError(t, err)
// 	require.Equal(t, int32(3), atomic.LoadInt32(&attempts))
// }

// func TestOpenAIProvider_RateLimiting(t *testing.T) {
// 	var callCount int32
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		atomic.AddInt32(&callCount, 1)
// 		resp := embeddingResponse{
// 			Object: "list",
// 			Data: []struct {
// 				Object    string    `json:"object"`
// 				Index     int       `json:"index"`
// 				Embedding []float64 `json:"embedding"`
// 			}{
// 				{Index: 0, Embedding: make([]float64, 768)},
// 			},
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(resp)
// 	}))
// 	defer server.Close()

// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   server.URL,
// 		Model:     "test",
// 		Dimension: 768,
// 		RateLimit: RateLimitConfig{
// 			RequestsPerSecond: 2,
// 			Burst:             1,
// 		},
// 	})
// 	require.NoError(t, err)

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	start := time.Now()
// 	for i := 0; i < 3; i++ {
// 		_, err := provider.Embed(ctx, []string{"test"})
// 		require.NoError(t, err)
// 	}

// 	// With 2 RPS and burst 1, 3 requests should take at least 0.5 seconds
// 	require.GreaterOrEqual(t, time.Since(start), 400*time.Millisecond)
// }

// func TestOpenAIProvider_ContextCancellation(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(5 * time.Second) // Simulate slow response
// 	}))
// 	defer server.Close()

// 	provider, err := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   server.URL,
// 		Model:     "test",
// 		Dimension: 768,
// 		Timeout:   10 * time.Second,
// 	})
// 	require.NoError(t, err)

// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
// 	defer cancel()

// 	_, err = provider.Embed(ctx, []string{"test"})
// 	require.Error(t, err)
// }

// func TestOpenAIProvider_Name(t *testing.T) {
// 	provider, _ := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   "http://localhost",
// 		Model:     "nomic-embed-text",
// 		Dimension: 768,
// 	})

// 	require.Equal(t, "openai-compatible(nomic-embed-text)", provider.Name())
// }

// func TestOpenAIProvider_Dimension(t *testing.T) {
// 	provider, _ := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   "http://localhost",
// 		Model:     "test",
// 		Dimension: 1536,
// 	})

// 	require.Equal(t, 1536, provider.Dimension())
// }

// func TestOpenAIProvider_Close(t *testing.T) {
// 	provider, _ := NewOpenAIProvider(ProviderConfig{
// 		BaseURL:   "http://localhost",
// 		Model:     "test",
// 		Dimension: 768,
// 	})

// 	err := provider.Close()
// 	require.NoError(t, err)
// }
