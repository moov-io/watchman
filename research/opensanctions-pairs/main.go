// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "opensanctions-pairs: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("opensanctions-pairs", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	input := fs.String("input", "", "Path to OpenSanctions pairs JSON, JSONL, or .json.gz")
	mode := fs.String("mode", "pairwise", "pairwise (in-process Similarity) or search (Watchman HTTP)")
	watchman := fs.String("watchman", "http://localhost:8084", "Watchman base URL for -mode=search")
	algorithm := fs.String("algorithm", "jaro-winkler", "Name algorithm (jaro-winkler, soundex, soft-bidist, ...)")
	compare := fs.String("compare", "", "Comma-separated configs or 'default' for the standard matrix")
	useTFIDF := fs.Bool("tfidf", false, "Weight name tokens with a TF-IDF index built from this corpus")
	embedMode := fs.String("embed", "", "Mix Ollama embeddings into the score: max, hybrid, or only")
	embedModel := fs.String("embed-model", "qwen3-embedding:0.6b", "Ollama embedding model")
	embedURL := fs.String("embed-url", "http://127.0.0.1:11434", "Ollama base URL")
	threshold := fs.Float64("threshold", 0.80, "Score cutoff for a positive prediction")
	sweep := fs.Bool("sweep", true, "Print F1 across thresholds 0.50–0.99")
	limit := fs.Int("limit", 0, "Max pairs after filtering (0 = all)")
	offset := fs.Int("offset", 0, "Skip this many pairs after filtering")
	workers := fs.Int("workers", runtime.NumCPU(), "Parallel workers")
	subjectsOnly := fs.Bool("subjects-only", true, "Drop Occupancy/Position/auto-merge pairs")
	schemas := fs.String("schemas", "", "Comma-separated schemas to keep (person,company,...)")
	nameOnly := fs.Bool("name-only", false, "Score names only (no dates, IDs, addresses)")
	maxAlt := fs.Int("max-alt-names", defaultMaxAltNames, "Cap aliases copied onto each entity")
	debugN := fs.Int("debug-errors", 5, "Print this many false positives and false negatives")
	searchN := fs.Int("search-limit", 10, "Watchman result page size for -mode=search")
	minMatch := fs.Float64("min-match", 0, "Watchman minMatch for -mode=search")
	download := fs.String("download", "", "Download dataset first: sample (1k) or full (~756k)")
	dataDir := fs.String("data-dir", defaultDataDir(), "Directory for downloads")
	outPath := fs.String("out", "", "Write per-pair JSONL results")
	reportPath := fs.String("report", "", "Write JSON summary")

	if err := fs.Parse(args); err != nil {
		return err
	}

	path := *input
	if *download != "" {
		got, err := downloadDataset(*download, *dataDir)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "downloaded %s\n", got)
		if path == "" {
			path = got
		}
	}
	if path == "" {
		return fmt.Errorf("need -input or -download (see -h)")
	}

	algo, err := search.ParseStringMatchAlgorithm(*algorithm)
	if err != nil {
		return err
	}

	pairs, err := loadPairs(path)
	if err != nil {
		return err
	}
	schemaFilter := parseSchemaFilter(*schemas)
	filtered := filterPairs(pairs, *subjectsOnly, schemaFilter, *offset, *limit)
	fmt.Fprintf(os.Stderr, "loaded %d pairs from %s; using %d after filters\n", len(pairs), path, len(filtered))
	if len(filtered) == 0 {
		return fmt.Errorf("no pairs left after filters")
	}

	compareSpecs := parseCompareList(*compare)
	needTFIDF := *useTFIDF
	needEmbed := *embedMode != ""
	for _, spec := range compareSpecs {
		if strings.Contains(spec, "tfidf") {
			needTFIDF = true
		}
		if strings.HasPrefix(spec, "embed-") {
			needEmbed = true
		}
	}

	var tfidfIndex *tfidf.Index
	if needTFIDF {
		tfidfIndex = buildTFIDF(filtered, *maxAlt)
	}

	var emb *embedCache
	if needEmbed {
		cachePath := filepath.Join(*dataDir, "embed-"+strings.ReplaceAll(*embedModel, ":", "_")+".gob")
		emb = newEmbedCache(*embedURL, *embedModel, cachePath, 0)
		if err := emb.load(); err != nil {
			return err
		}
		names := collectPrimaryNames(filtered, *maxAlt)
		ctx, cancel := context.WithTimeout(context.Background(), 6*time.Hour)
		defer cancel()
		if err := emb.ensure(ctx, names, 32); err != nil {
			return fmt.Errorf("embeddings: %w", err)
		}
	}

	cfg := evalConfig{
		algorithm:   algo,
		threshold:   *threshold,
		nameOnly:    *nameOnly,
		maxAltNames: *maxAlt,
		workers:     *workers,
		debugErrors: *debugN,
	}
	if *useTFIDF {
		cfg.tfidf = tfidfIndex
	}
	if *embedMode != "" {
		cfg.embed = emb
		cfg.embedMode = *embedMode
	}

	if len(compareSpecs) > 0 {
		rows, err := runCompare(filtered, compareSpecs, cfg, tfidfIndex, emb)
		if err != nil {
			return err
		}
		printCompareTable(os.Stdout, rows, *threshold)
		if *reportPath != "" {
			if err := writeJSON(*reportPath, rows); err != nil {
				return err
			}
		}
		return nil
	}

	var results []pairResult
	switch strings.ToLower(*mode) {
	case "compare":
		return fmt.Errorf("use -compare default (or a comma-separated list) with -mode pairwise")
	case "pairwise":
		results = scorePairwise(filtered, cfg)
	case "search":
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Hour)
		defer cancel()
		results, err = scoreSearch(ctx, filtered, searchConfig{
			evalConfig: cfg,
			address:    *watchman,
			searchN:    *searchN,
			minMatch:   *minMatch,
		})
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown -mode %q (pairwise or search)", *mode)
	}

	applyThreshold(results, cfg.threshold)
	if strings.ToLower(*mode) == "pairwise" {
		attachDebugPieces(filtered, results, cfg)
	}
	printReport(os.Stdout, results, cfg, strings.ToLower(*mode), *sweep)

	if *outPath != "" {
		f, err := os.Create(*outPath)
		if err != nil {
			return err
		}
		if err := writeJSONL(f, results); err != nil {
			f.Close()
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	if *reportPath != "" {
		if err := writeJSON(*reportPath, buildReport(results, cfg, strings.ToLower(*mode))); err != nil {
			return err
		}
	}
	return nil
}

func parseCompareList(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	if strings.EqualFold(raw, "default") {
		return defaultCompareSpecs()
	}
	var out []string
	for _, p := range strings.Split(raw, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func parseSchemaFilter(raw string) map[string]struct{} {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	out := make(map[string]struct{})
	for _, p := range strings.Split(raw, ",") {
		p = strings.ToLower(strings.TrimSpace(p))
		if p == "" {
			continue
		}
		out[p] = struct{}{}
		if p == "company" || p == "organization" || p == "legalentity" {
			out["company"] = struct{}{}
			out["organization"] = struct{}{}
			out["legalentity"] = struct{}{}
			out["publicbody"] = struct{}{}
		}
	}
	return out
}
