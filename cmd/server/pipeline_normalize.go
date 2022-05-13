// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	punctuationReplacer = strings.NewReplacer(".", "", ",", "", "-", " ", "  ", " ")
)

type normalizeStep struct {
}

func (s *normalizeStep) apply(in *Name) error {
	in.Processed = precompute(in.Processed)
	return nil
}

// precompute will lowercase each substring and remove punctuation
//
// This function is called on every record from the flat files and all
// search requests (i.e. HTTP and searcher.TopNNNs methods).
// See: https://godoc.org/golang.org/x/text/unicode/norm#Form
// See: https://withblue.ink/2019/03/11/why-you-need-to-normalize-unicode-strings.html
func precompute(s string) string {
	trimmed := strings.TrimSpace(strings.ToLower(punctuationReplacer.Replace(s)))

	// UTF-8 normalization
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC) // Mn: nonspacing marks
	result, _, _ := transform.String(t, trimmed)
	return result
}
