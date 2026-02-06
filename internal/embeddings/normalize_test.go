package embeddings

import (
	"math"
	"testing"
)

func TestNormalizeL2_Comprehensive(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		wantNorm float64 // expected L2 norm of result (should be 1.0 for non-zero)
	}{
		{
			name:     "simple vector",
			input:    []float64{3, 4},
			wantNorm: 1.0,
		},
		{
			name:     "already normalized",
			input:    []float64{1, 0, 0},
			wantNorm: 1.0,
		},
		{
			name:     "negative values",
			input:    []float64{-3, 4},
			wantNorm: 1.0,
		},
		{
			name:     "larger vector",
			input:    []float64{1, 2, 3, 4, 5},
			wantNorm: 1.0,
		},
		{
			name:     "zero vector",
			input:    []float64{0, 0, 0},
			wantNorm: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeL2(tt.input)

			// Check length preserved
			if len(result) != len(tt.input) {
				t.Errorf("length changed: got %d, want %d", len(result), len(tt.input))
			}

			// Calculate L2 norm of result
			var norm float64
			for _, v := range result {
				norm += float64(v) * float64(v)
			}
			norm = math.Sqrt(norm)

			if math.Abs(norm-tt.wantNorm) > 1e-6 {
				t.Errorf("norm = %v, want %v", norm, tt.wantNorm)
			}
		})
	}
}

func TestNormalizeL2_PreservesDirection(t *testing.T) {
	input := []float64{3, 4}
	result := normalizeL2(input)

	// For [3, 4], normalized should be [0.6, 0.8]
	if math.Abs(float64(result[0])-0.6) > 1e-6 {
		t.Errorf("result[0] = %v, want 0.6", result[0])
	}
	if math.Abs(float64(result[1])-0.8) > 1e-6 {
		t.Errorf("result[1] = %v, want 0.8", result[1])
	}
}

func TestNormalizeL2Batch(t *testing.T) {
	input := [][]float64{
		{3, 4},
		{1, 0},
		{0, 5},
	}

	result := normalizeL2Batch(input)

	if len(result) != len(input) {
		t.Errorf("batch length changed: got %d, want %d", len(result), len(input))
	}

	// Each vector should be normalized
	for i, vec := range result {
		var norm float64
		for _, v := range vec {
			norm += float64(v) * float64(v)
		}
		norm = math.Sqrt(norm)

		if math.Abs(norm-1.0) > 1e-6 {
			t.Errorf("vector %d norm = %v, want 1.0", i, norm)
		}
	}
}

func TestDotProduct_Comprehensive(t *testing.T) {
	tests := []struct {
		name string
		a    []float64
		b    []float64
		want float64
	}{
		{
			name: "orthogonal vectors",
			a:    []float64{1, 0},
			b:    []float64{0, 1},
			want: 0.0,
		},
		{
			name: "same vector",
			a:    []float64{1, 0},
			b:    []float64{1, 0},
			want: 1.0,
		},
		{
			name: "opposite vectors",
			a:    []float64{1, 0},
			b:    []float64{-1, 0},
			want: -1.0,
		},
		{
			name: "normalized similar",
			a:    normalizeL2([]float64{1, 1}),
			b:    normalizeL2([]float64{1, 2}),
			want: 0.9486832980505138, // cos(angle between [1,1] and [1,2])
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dotProduct(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("dotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDotProduct_CosineSimilarity(t *testing.T) {
	// For normalized vectors, dot product equals cosine similarity
	a := normalizeL2([]float64{1, 2, 3})
	b := normalizeL2([]float64{4, 5, 6})

	similarity := dotProduct(a, b)

	// Similarity should be in range [-1, 1]
	if similarity < -1 || similarity > 1 {
		t.Errorf("similarity out of range: %v", similarity)
	}

	// These vectors are similar (same direction tendency), so similarity > 0
	if similarity < 0.9 {
		t.Errorf("expected high similarity for similar vectors, got %v", similarity)
	}
}

func BenchmarkNormalizeL2(b *testing.B) {
	vec := make([]float64, 768) // Common embedding dimension
	for i := range vec {
		vec[i] = float64(i) * 0.001
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		normalizeL2(vec)
	}
}

func BenchmarkDotProduct(b *testing.B) {
	vec1 := make([]float64, 768)
	vec2 := make([]float64, 768)
	for i := range vec1 {
		vec1[i] = float64(i) * 0.001
		vec2[i] = float64(i) * 0.002
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dotProduct(vec1, vec2)
	}
}
