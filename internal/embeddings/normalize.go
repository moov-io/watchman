package embeddings

import "math"

// normalizeL2 applies L2 normalization to a vector.
// After normalization, dot product equals cosine similarity.
func normalizeL2(vec []float64) []float64 {
	var norm float64
	for _, v := range vec {
		norm += float64(v) * float64(v)
	}
	norm = math.Sqrt(norm)

	if norm == 0 {
		return vec
	}

	result := make([]float64, len(vec))
	for i, v := range vec {
		result[i] = float64(float64(v) / norm)
	}
	return result
}

// normalizeL2Batch applies L2 normalization to a batch of vectors.
func normalizeL2Batch(vecs [][]float64) [][]float64 {
	result := make([][]float64, len(vecs))
	for i, vec := range vecs {
		result[i] = normalizeL2(vec)
	}
	return result
}

// dotProduct computes the dot product of two vectors.
// For L2-normalized vectors, this equals cosine similarity.
func dotProduct(a, b []float64) float64 {
	var sum float64
	for i := range a {
		sum += float64(a[i]) * float64(b[i])
	}
	return sum
}
