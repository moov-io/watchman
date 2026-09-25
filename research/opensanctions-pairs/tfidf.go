// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

func buildTFIDF(pairs []Pair, maxAlt int) *tfidf.Index {
	cfg := tfidf.DefaultConfig()
	cfg.Enabled = true
	idx := tfidf.NewIndex(cfg)

	docs := make([][]string, 0, len(pairs)*2)
	conv := convertOpts{maxAltNames: maxAlt}
	for i, p := range pairs {
		left, right := convertPair(p, conv)
		docs = appendNameDocs(docs, left)
		docs = appendNameDocs(docs, right)
		if (i+1)%100000 == 0 {
			fmt.Fprintf(os.Stderr, "tfidf documents: %d pairs\n", i+1)
		}
	}
	idx.Build(docs)
	st := idx.Stats()
	fmt.Fprintf(os.Stderr, "tfidf index: %d documents, %d unique terms\n", st.TotalDocuments, st.UniqueTerms)
	return idx
}

func appendNameDocs(docs [][]string, e search.Entity[search.Value]) [][]string {
	if len(e.PreparedFields.NameFields) > 0 {
		docs = append(docs, e.PreparedFields.NameFields)
	}
	for _, alt := range e.PreparedFields.AltNameFields {
		if len(alt) > 0 {
			docs = append(docs, alt)
		}
	}
	return docs
}
