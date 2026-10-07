// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	csvPath := flag.String("csv", "pkg/search/testdata/cascade-name-pairs-2.csv", "cascade name-pairs CSV")
	models := flag.String("models", "qwen3-embedding:0.6b,qwen3-embedding:4b,qwen3-embedding:latest,bge-m3,nomic-embed-text", "comma-separated Ollama embedding models")
	baseURL := flag.String("ollama", "http://localhost:11434", "Ollama base URL")
	prefix := flag.String("prefix", "", "optional task prefix prepended to every name (EmbeddingGemma 2 STS: \"task: sentence similarity | query: \")")
	batch := flag.Int("batch", 16, "embed batch size")
	out := flag.String("out", "", "write markdown report to this path (stdout if empty)")
	flag.Parse()

	pairs, err := loadPairs(*csvPath)
	if err != nil {
		return err
	}
	names := uniqueNames(pairs)
	fmt.Fprintf(os.Stderr, "loaded %d pairs, %d unique names from %s\n", len(pairs), len(names), *csvPath)

	reports := []modelReport{scoreModel("jaro-winkler", 0, pairs, nil, "")}

	ctx := context.Background()
	for _, model := range splitCSV(*models) {
		client := newEmbedClient(*baseURL, model, *prefix)
		start := time.Now()
		if err := client.ensure(ctx, names, *batch); err != nil {
			fmt.Fprintf(os.Stderr, "skip %s: %v\n", model, err)
			continue
		}
		label := model + " hybrid"
		if *prefix != "" {
			label += " +prefix"
		}
		fmt.Fprintf(os.Stderr, "%s dim=%d in %s\n", label, client.dim, time.Since(start).Round(time.Millisecond))
		reports = append(reports, scoreModel(label, client.dim, pairs, client.cosine, "hybrid"))
	}

	body := "# Cascade embedding model comparison\n\n" +
		"Pairwise `pkg/search.Similarity` plus Ollama cosine on names. " +
		"This is the OpenSanctions embed-hybrid mix, not `/v2/search`.\n\n" +
		fmt.Sprintf("Fixture: `%s` (%d pairs).\n", *csvPath, len(pairs)) +
		fmt.Sprintf("Ollama: `%s`. Prefix: `%s`.\n\n", *baseURL, *prefix) +
		renderReports("Results", reports)

	if *out == "" {
		fmt.Fprint(os.Stdout, body)
		return nil
	}
	return os.WriteFile(*out, []byte(body), 0o600)
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
