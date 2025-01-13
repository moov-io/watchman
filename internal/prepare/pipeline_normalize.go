// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"strings"
	"sync"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	punctuationReplacer = strings.NewReplacer(".", "", ",", "", "-", " ", "  ", " ")
)

// LowerAndRemovePunctuation will lowercase each substring and remove punctuation
//
// This function is called on every record from the flat files and all
// search requests (i.e. HTTP and searcher.TopNNNs methods).
// See: https://godoc.org/golang.org/x/text/unicode/norm#Form
// See: https://withblue.ink/2019/03/11/why-you-need-to-normalize-unicode-strings.html
func LowerAndRemovePunctuation(s string) string {
	trimmed := strings.TrimSpace(strings.ToLower(punctuationReplacer.Replace(s)))

	// UTF-8 normalization
	chain := getTransformChain()
	defer saveBuffer(chain)

	result, _, _ := transform.String(chain, trimmed)
	return result
}

var (
	transformChainPool = sync.Pool{
		New: func() any {
			return newTransformChain()
		},
	}
)

func newTransformChain() transform.Transformer {
	nonspacingMarksRemover := runes.Remove(runes.In(unicode.Mn)) // Mn: nonspacing marks
	return transform.Chain(norm.NFD, nonspacingMarksRemover, norm.NFC)
}

func getTransformChain() transform.Transformer {
	t, ok := transformChainPool.Get().(transform.Transformer)
	if !ok {
		return newTransformChain()
	}
	return t
}

func saveBuffer(t transform.Transformer) {
	t.Reset()
	transformChainPool.Put(t)
}
