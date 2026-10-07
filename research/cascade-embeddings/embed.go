// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

type embedClient struct {
	url    string
	model  string
	prefix string
	http   *http.Client
	vecs   map[string][]float32
	dim    int
}

type ollamaEmbedRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type ollamaEmbedResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
}

func newEmbedClient(baseURL, model, prefix string) *embedClient {
	return &embedClient{
		url:    strings.TrimRight(baseURL, "/"),
		model:  model,
		prefix: prefix,
		http:   &http.Client{Timeout: 3 * time.Minute},
		vecs:   make(map[string][]float32),
	}
}

func (c *embedClient) ensure(ctx context.Context, names []string, batch int) error {
	if batch < 1 {
		batch = 16
	}
	missing := make([]string, 0)
	for _, n := range names {
		if _, ok := c.vecs[n]; !ok {
			missing = append(missing, n)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	fmt.Fprintf(os.Stderr, "embedding %d names with %s (batch=%d)\n", len(missing), c.model, batch)
	for i := 0; i < len(missing); i += batch {
		if err := ctx.Err(); err != nil {
			return err
		}
		end := i + batch
		if end > len(missing) {
			end = len(missing)
		}
		chunk := missing[i:end]
		vecs, err := c.embedBatch(ctx, chunk)
		if err != nil {
			return fmt.Errorf("%s batch at %d: %w", c.model, i, err)
		}
		for j, name := range chunk {
			c.vecs[name] = vecs[j]
		}
		fmt.Fprintf(os.Stderr, "  %s %d/%d\n", c.model, end, len(missing))
	}
	return nil
}

func (c *embedClient) embedBatch(ctx context.Context, names []string) ([][]float32, error) {
	inputs := names
	if c.prefix != "" {
		inputs = make([]string, len(names))
		for i, n := range names {
			inputs[i] = c.prefix + n
		}
	}
	body, err := json.Marshal(ollamaEmbedRequest{Model: c.model, Input: inputs})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url+"/api/embed", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 64<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama embed %s: %s", resp.Status, truncate(string(raw), 200))
	}
	var out ollamaEmbedResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("decode ollama embed: %w", err)
	}
	if len(out.Embeddings) != len(names) {
		return nil, fmt.Errorf("ollama embed count %d != %d", len(out.Embeddings), len(names))
	}
	vecs := make([][]float32, len(out.Embeddings))
	for i, src := range out.Embeddings {
		v := l2normalize(src)
		if c.dim == 0 {
			c.dim = len(v)
		} else if len(v) != c.dim {
			return nil, fmt.Errorf("dim %d != %d", len(v), c.dim)
		}
		vecs[i] = v
	}
	return vecs, nil
}

func (c *embedClient) cosine(a, b string) float64 {
	if a == "" || b == "" {
		return 0
	}
	if a == b {
		return 1
	}
	va, oka := c.vecs[a]
	vb, okb := c.vecs[b]
	if !oka || !okb || len(va) == 0 || len(vb) == 0 || len(va) != len(vb) {
		return 0
	}
	var sum float64
	for i := range va {
		sum += float64(va[i]) * float64(vb[i])
	}
	if sum < 0 {
		return 0
	}
	if sum > 1 {
		return 1
	}
	return sum
}

func l2normalize(src []float64) []float32 {
	var ss float64
	for _, x := range src {
		ss += x * x
	}
	out := make([]float32, len(src))
	if ss == 0 {
		return out
	}
	inv := 1 / math.Sqrt(ss)
	for i, x := range src {
		out[i] = float32(x * inv)
	}
	return out
}

func mixEmbed(sim, cos float64, mode string, crossScript bool) float64 {
	switch mode {
	case "only":
		return cos
	case "max":
		if cos > sim {
			return cos
		}
		return sim
	case "hybrid":
		if crossScript && cos > sim {
			return cos
		}
		return sim
	default:
		return sim
	}
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
}
