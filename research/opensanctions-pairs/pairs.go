// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// Pair is one labeled OpenSanctions entity-matching example.
type Pair struct {
	Left      FTMEntity `json:"left"`
	Right     FTMEntity `json:"right"`
	Judgement string    `json:"judgement"`
}

// FTMEntity is a FollowTheMoney record as released in OpenSanctions Pairs.
type FTMEntity struct {
	ID         string              `json:"id"`
	Caption    string              `json:"caption"`
	Schema     string              `json:"schema"`
	Referents  []string            `json:"referents"`
	Datasets   []string            `json:"datasets"`
	Properties map[string][]string `json:"properties"`
	Target     bool                `json:"target"`
}

type sampleFile struct {
	Metadata json.RawMessage `json:"metadata"`
	Pairs    []Pair          `json:"pairs"`
}

func (p Pair) schemaKey() string {
	return p.Left.Schema + "/" + p.Right.Schema
}

func (p Pair) isPositive() bool {
	return strings.EqualFold(strings.TrimSpace(p.Judgement), "positive")
}

func (e FTMEntity) prop(keys ...string) []string {
	if e.Properties == nil {
		return nil
	}
	var out []string
	seen := make(map[string]struct{})
	for _, key := range keys {
		for _, v := range e.Properties[key] {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

func (e FTMEntity) firstProp(keys ...string) string {
	vals := e.prop(keys...)
	if len(vals) == 0 {
		return ""
	}
	return vals[0]
}

func loadPairs(path string) ([]Pair, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open pairs file: %w", err)
	}
	defer f.Close()

	var r io.Reader = f
	if strings.HasSuffix(strings.ToLower(path), ".gz") {
		gz, err := gzip.NewReader(f)
		if err != nil {
			return nil, fmt.Errorf("gzip reader: %w", err)
		}
		defer gz.Close()
		r = gz
	}

	return decodePairs(r)
}

func decodePairs(r io.Reader) ([]Pair, error) {
	br := bufio.NewReader(r)
	for {
		b, err := br.Peek(1)
		if err == io.EOF {
			return nil, fmt.Errorf("empty pairs file")
		}
		if err != nil {
			return nil, fmt.Errorf("peek pairs file: %w", err)
		}
		if !unicode.IsSpace(rune(b[0])) {
			break
		}
		if _, err := br.ReadByte(); err != nil {
			return nil, fmt.Errorf("skip whitespace: %w", err)
		}
	}

	first, err := br.Peek(1)
	if err != nil {
		return nil, fmt.Errorf("peek first byte: %w", err)
	}

	dec := json.NewDecoder(br)
	if first[0] == '[' {
		var pairs []Pair
		if err := dec.Decode(&pairs); err != nil {
			return nil, fmt.Errorf("decode pair array: %w", err)
		}
		return pairs, nil
	}

	var firstObj json.RawMessage
	if err := dec.Decode(&firstObj); err != nil {
		return nil, fmt.Errorf("decode first object: %w", err)
	}

	var wrapped sampleFile
	if err := json.Unmarshal(firstObj, &wrapped); err == nil && len(wrapped.Pairs) > 0 {
		return wrapped.Pairs, nil
	}

	var one Pair
	if err := json.Unmarshal(firstObj, &one); err != nil {
		return nil, fmt.Errorf("decode pair: %w", err)
	}
	pairs := []Pair{one}
	for {
		var p Pair
		err := dec.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decode jsonl pair: %w", err)
		}
		pairs = append(pairs, p)
		if len(pairs)%50000 == 0 {
			fmt.Fprintf(os.Stderr, "loading pairs: %d\n", len(pairs))
		}
	}
	return pairs, nil
}

func defaultDataDir() string {
	if st, err := os.Stat(filepath.Join("research", "opensanctions-pairs")); err == nil && st.IsDir() {
		return filepath.Join("research", "opensanctions-pairs", "data")
	}
	return "data"
}
