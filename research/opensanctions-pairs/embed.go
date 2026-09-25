// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type embedCache struct {
	mu    sync.RWMutex
	dim   int
	model string
	url   string
	path  string
	vecs  map[string][]float32
	http  *http.Client
}

type gobStore struct {
	Dim   int
	Model string
	Vecs  map[string][]float32
}

func newEmbedCache(baseURL, model, cachePath string, dim int) *embedCache {
	return &embedCache{
		dim:   dim,
		model: model,
		url:   baseURL,
		path:  cachePath,
		vecs:  make(map[string][]float32),
		http:  &http.Client{Timeout: 3 * time.Minute},
	}
}

func (c *embedCache) load() error {
	f, err := os.Open(c.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()
	var store gobStore
	if err := gob.NewDecoder(f).Decode(&store); err != nil {
		fmt.Fprintf(os.Stderr, "embed cache unreadable, starting empty: %v\n", err)
		return nil
	}
	if store.Model != c.model || (c.dim > 0 && store.Dim != c.dim) {
		fmt.Fprintf(os.Stderr, "embed cache model/dim mismatch (have %s/%d want %s/%d), starting empty\n",
			store.Model, store.Dim, c.model, c.dim)
		return nil
	}
	if c.dim == 0 {
		c.dim = store.Dim
	}
	c.vecs = store.Vecs
	if c.vecs == nil {
		c.vecs = make(map[string][]float32)
	}
	fmt.Fprintf(os.Stderr, "embed cache: loaded %d vectors (%s)\n", len(c.vecs), c.path)
	return nil
}

func (c *embedCache) save() error {
	c.mu.RLock()
	store := gobStore{Dim: c.dim, Model: c.model, Vecs: c.vecs}
	c.mu.RUnlock()
	if err := os.MkdirAll(filepath.Dir(c.path), 0o755); err != nil {
		return err
	}
	tmp := c.path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	if err := gob.NewEncoder(f).Encode(store); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	if err := f.Close(); err != nil {
		os.Remove(tmp)
		return err
	}
	return os.Rename(tmp, c.path)
}

func (c *embedCache) cosine(a, b string) float64 {
	if a == "" || b == "" {
		return 0
	}
	if a == b {
		return 1
	}
	c.mu.RLock()
	va, oka := c.vecs[a]
	vb, okb := c.vecs[b]
	c.mu.RUnlock()
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

func (c *embedCache) ensure(ctx context.Context, names []string, batch int) error {
	if batch < 1 {
		batch = 16
	}
	missing := make([]string, 0)
	seen := make(map[string]struct{})
	c.mu.RLock()
	for _, n := range names {
		n = normEmbedKey(n)
		if n == "" {
			continue
		}
		if _, dup := seen[n]; dup {
			continue
		}
		seen[n] = struct{}{}
		if _, ok := c.vecs[n]; !ok {
			missing = append(missing, n)
		}
	}
	c.mu.RUnlock()
	if len(missing) == 0 {
		fmt.Fprintf(os.Stderr, "embed cache: all %d unique names present\n", len(seen))
		return nil
	}
	fmt.Fprintf(os.Stderr, "embedding %d new names with %s (batch=%d)\n", len(missing), c.model, batch)

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
			return fmt.Errorf("embed batch at %d: %w", i, err)
		}
		c.mu.Lock()
		for j, name := range chunk {
			c.vecs[name] = vecs[j]
		}
		c.mu.Unlock()
		if (i/batch)%10 == 0 || end == len(missing) {
			fmt.Fprintf(os.Stderr, "embedded %d/%d names\n", end, len(missing))
		}
	}
	return c.save()
}

type ollamaEmbedRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type ollamaEmbedResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
}

func (c *embedCache) embedBatch(ctx context.Context, names []string) ([][]float32, error) {
	body, err := json.Marshal(ollamaEmbedRequest{Model: c.model, Input: names})
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
	for i, v := range out.Embeddings {
		if c.dim == 0 && len(v) > 0 {
			c.dim = len(v)
		}
		if c.dim > 0 && len(v) != c.dim {
			return nil, fmt.Errorf("embed dim %d != %d", len(v), c.dim)
		}
		vecs[i] = l2float32(v)
	}
	return vecs, nil
}

func l2float32(v []float64) []float32 {
	var n float64
	for _, x := range v {
		n += x * x
	}
	n = math.Sqrt(n)
	out := make([]float32, len(v))
	if n == 0 {
		return out
	}
	for i, x := range v {
		out[i] = float32(x / n)
	}
	return out
}

func collectPrimaryNames(pairs []Pair, maxAlt int) []string {
	seen := make(map[string]struct{}, len(pairs))
	conv := convertOpts{maxAltNames: maxAlt}
	var names []string
	add := func(s string) {
		s = normEmbedKey(s)
		if s == "" {
			return
		}
		if _, ok := seen[s]; ok {
			return
		}
		seen[s] = struct{}{}
		names = append(names, s)
	}
	for _, p := range pairs {
		left, right := convertPair(p, conv)
		add(left.Name)
		add(right.Name)
	}
	return names
}

func normEmbedKey(s string) string {
	return strings.TrimSpace(s)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
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
