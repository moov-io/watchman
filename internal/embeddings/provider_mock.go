package embeddings

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"sync/atomic"
	"time"

	"github.com/ccoveille/go-safecast/v2"
)

// MockProvider generates deterministic embeddings for testing.
// It produces consistent embeddings based on text hash, allowing predictable tests.
type MockProvider struct {
	dimension int
	delay     time.Duration
	failAfter int32 // Fail after N calls (0 = never fail)
	callCount int32
}

// MockProviderOption configures a MockProvider.
type MockProviderOption func(*MockProvider)

// WithMockDelay adds simulated latency to the mock provider.
func WithMockDelay(delay time.Duration) MockProviderOption {
	return func(m *MockProvider) {
		m.delay = delay
	}
}

// WithMockFailAfter causes the mock to fail after N successful calls.
func WithMockFailAfter(n int) MockProviderOption {
	return func(m *MockProvider) {
		i, err := safecast.Convert[int32](n)
		if err == nil {
			m.failAfter = i
		}
	}
}

// NewMockProvider creates a mock provider for testing.
func NewMockProvider(dimension int, opts ...MockProviderOption) *MockProvider {
	m := &MockProvider{
		dimension: dimension,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// Embed generates deterministic embeddings based on text hash.
func (m *MockProvider) Embed(ctx context.Context, texts []string) ([][]float64, error) {
	count := atomic.AddInt32(&m.callCount, 1)
	if m.failAfter > 0 && count > m.failAfter {
		return nil, ErrProviderFailure
	}

	if m.delay > 0 {
		select {
		case <-time.After(m.delay):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	result := make([][]float64, len(texts))
	for i, text := range texts {
		result[i] = m.deterministicEmbedding(text)
	}
	return result, nil
}

// deterministicEmbedding generates a consistent, normalized embedding from text hash.
func (m *MockProvider) deterministicEmbedding(text string) []float64 {
	hash := sha256.Sum256([]byte(text))
	vec := make([]float64, m.dimension)

	// Use hash bytes to seed each dimension
	for i := 0; i < m.dimension; i++ {
		// Cycle through hash bytes for dimensions > 32
		idx := i % 32
		// Use multiple hash bytes to generate more variance
		val := binary.BigEndian.Uint32([]byte{
			hash[idx],
			hash[(idx+1)%32],
			hash[(idx+2)%32],
			hash[(idx+3)%32],
		})
		// Map to [-1, 1] range
		vec[i] = float64(val)/float64(math.MaxUint32)*2 - 1
	}

	return normalizeL2(vec)
}

// Dimension returns the configured embedding dimension.
func (m *MockProvider) Dimension() int {
	return m.dimension
}

// Name returns the provider name.
func (m *MockProvider) Name() string {
	return "mock"
}

// Close is a no-op for the mock provider.
func (m *MockProvider) Close() error {
	return nil
}

// CallCount returns the number of times Embed was called.
func (m *MockProvider) CallCount() int {
	return int(atomic.LoadInt32(&m.callCount))
}

// Reset resets the call counter.
func (m *MockProvider) Reset() {
	atomic.StoreInt32(&m.callCount, 0)
}
