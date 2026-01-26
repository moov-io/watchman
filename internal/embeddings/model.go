//go:build embeddings

package embeddings

import (
	"context"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/knights-analytics/hugot"
	"github.com/knights-analytics/hugot/options"
	"github.com/knights-analytics/hugot/pipelines"
)

// model wraps the hugot ONNX model for embedding generation.
type model struct {
	session   *hugot.Session
	pipeline  *pipelines.FeatureExtractionPipeline
	dimension int
	backend   string // "ort" or "go"
}

// loadModel loads the ONNX model from the specified path.
func loadModel(ctx context.Context, config Config) (*model, error) {
	// Verify model directory exists
	if _, err := os.Stat(config.ModelPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("model directory not found: %s", config.ModelPath)
	}

	// Verify model.onnx exists
	modelFile := filepath.Join(config.ModelPath, "model.onnx")
	if _, err := os.Stat(modelFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("model.onnx not found in %s", config.ModelPath)
	}

	// Create session - try ONNX Runtime first, fall back to native Go
	var session *hugot.Session
	var err error
	var backend string

	// Try ONNX Runtime (faster, but requires ORT build tag and library)
	var opts []options.WithOption
	onnxLibraryPath := os.Getenv("ONNX_LIBRARY_PATH")
	if onnxLibraryPath != "" {
		opts = append(opts, options.WithOnnxLibraryPath(onnxLibraryPath))
	}
	session, err = hugot.NewORTSession(opts...)
	if err != nil {
		// Fall back to native Go backend
		session, err = hugot.NewGoSession()
		if err != nil {
			return nil, fmt.Errorf("failed to create session: %w", err)
		}
		backend = "go"
	} else {
		backend = "ort"
	}

	// Create feature extraction pipeline with normalization
	pipelineConfig := hugot.FeatureExtractionConfig{
		ModelPath: config.ModelPath,
		Name:      "embeddings",
		Options: []hugot.FeatureExtractionOption{
			pipelines.WithNormalization(), // L2 normalize embeddings
		},
	}

	pipeline, err := hugot.NewPipeline(session, pipelineConfig)
	if err != nil {
		session.Destroy()
		return nil, fmt.Errorf("failed to create pipeline: %w", err)
	}

	return &model{
		session:   session,
		pipeline:  pipeline,
		dimension: 384, // MiniLM dimension
		backend:   backend,
	}, nil
}

// encode converts texts to normalized embedding vectors.
func (m *model) encode(texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	// Run inference
	result, err := m.pipeline.RunPipeline(texts)
	if err != nil {
		return nil, fmt.Errorf("inference failed: %w", err)
	}

	// Embeddings are already normalized due to WithNormalization() option
	return result.Embeddings, nil
}

// close releases model resources.
func (m *model) close() {
	if m.session != nil {
		m.session.Destroy()
	}
}

// normalizeL2 applies L2 normalization to a vector.
// After normalization, dot product equals cosine similarity.
func normalizeL2(vec []float32) []float32 {
	var norm float64
	for _, v := range vec {
		norm += float64(v) * float64(v)
	}
	norm = math.Sqrt(norm)

	if norm == 0 {
		return vec
	}

	result := make([]float32, len(vec))
	for i, v := range vec {
		result[i] = float32(float64(v) / norm)
	}
	return result
}

// dotProduct computes the dot product of two vectors.
// For L2-normalized vectors, this equals cosine similarity.
func dotProduct(a, b []float32) float64 {
	var sum float64
	for i := range a {
		sum += float64(a[i]) * float64(b[i])
	}
	return sum
}
