// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package strcomp defines functions to return a percent match between two
// strings. Many algorithms are available, but all are normalized to [0.00, 1.00]
// for use as percentages.
//
// Any value outside of the normalized range represents a bug and should be fixed.
package strcmp

import (
	"unicode/utf8"

	"github.com/xrash/smetrics"
)

// Levenshtein is the "edit distance" between two strings. This is the count of operations
// (insert, delete, replace) needed for two strings to be equal.
//
// The returend value represents a score between 0-1 for how similar the strings are.
//  Levenshtein("abc", "abc") == 1
//  Levenshtein("abb", "abb") == 0.66
func Levenshtein(a, b string) float64 {
	if a == "" || b == "" {
		return 0.00
	}

	length := utf8.RuneCountInString(a)
	if n := utf8.RuneCountInString(b); n > length {
		length = n // set length to larger value of the two strings
	}

	ed := float64(smetrics.WagnerFischer(a, b, 1, 1, 2))
	if ed == 0 {
		return 1.0 // strings are equal
	}
	score := ed / float64(length)
	if score > 1.00 {
		// If more edits are required than the string's length a and b aren't equal.
		return 0.00
	}
	return score
}

// Soundex is a phonetic algorithm that considers how the words sound in english.
// Soundex maps a name to a 4-byte string consisting of the first letter of the original string and three numbers. Strings that sound similar should map to the same thing.
//
// Retruned is Hamming computed over both phonetic outputs.
func Soundex(a, b string) float64 {
	if a == "" || b == "" {
		return 0.00
	}
	return hamming(smetrics.Soundex(a), smetrics.Soundex(b))
}

// hamming distance is the minimum number of substitutions required to change one string into the other.
func hamming(a, b string) float64 {
	if a == "" || b == "" {
		return 0.00
	}

	length := utf8.RuneCountInString(a)
	if n := utf8.RuneCountInString(b); n != length {
		return 0.0
	}
	n, _ := smetrics.Hamming(a, b)
	return 1 - (float64(n) / float64(length))
}
