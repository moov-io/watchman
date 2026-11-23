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
	punctuationReplacer = strings.NewReplacer(".", " ", ",", " ", "-", " ")
)

// LowerAndRemovePunctuation will lowercase each substring and remove punctuation
//
// This function is called on every record from the flat files and all
// search requests (i.e. HTTP and searcher.TopNNNs methods).
// See: https://godoc.org/golang.org/x/text/unicode/norm#Form
// See: https://withblue.ink/2019/03/11/why-you-need-to-normalize-unicode-strings.html
func LowerAndRemovePunctuation(s string) string {
	trimmed := strings.TrimSpace(strings.ToLower(punctuationReplacer.Replace(s)))
	chars := []rune(trimmed)

	// Remove all punctuation and normalize whitespace
	var normalized strings.Builder
	lastWasSpace := true // Start with true to trim leading spaces

	for i := 0; i < len(chars); i++ {
		r := chars[i]
		if isApostrophe(r) {
			if isPossessiveS(chars, i) {
				continue
			}
			if !lastWasSpace {
				normalized.WriteRune(' ')
				lastWasSpace = true
			}
			continue
		}
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			if !lastWasSpace {
				normalized.WriteRune(' ')
				lastWasSpace = true
			}
			continue
		}

		if unicode.IsSpace(r) {
			if !lastWasSpace {
				normalized.WriteRune(' ')
				lastWasSpace = true
			}
			continue
		}

		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			normalized.WriteRune(r)
			lastWasSpace = false
		}
	}

	// UTF-8 normalization
	chain := getTransformChain()
	defer saveBuffer(chain)

	result, _, _ := transform.String(chain, strings.TrimSpace(normalized.String()))
	return result
}

func isApostrophe(r rune) bool {
	switch r {
	case '\'', '\u2018', '\u2019', '\u02BC', '`':
		return true
	}
	return false
}

func isPossessiveS(runes []rune, apostropheIdx int) bool {
	// Elide only token-final 's so "GINO'S" matches "Ginos" while "O'Doyle" stays "o doyle".
	i := apostropheIdx + 1
	if i >= len(runes) || runes[i] != 's' {
		return false
	}
	if i+1 >= len(runes) {
		return true
	}
	next := runes[i+1]
	return !unicode.IsLetter(next) && !unicode.IsNumber(next)
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
