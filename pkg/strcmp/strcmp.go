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

// JaroWinkler is a more accurate version of the Jaro algorithm. It works by boosting the
// score of exact matches at the beginning of the strings. By doing this, Winkler says that
// typos are less common to happen at the beginning.
//
// For this to happen, it introduces two more parameters: the boostThreshold and the prefixSize.
// These are commonly set to 0.7 and 4, respectively.
//
// From: https://godoc.org/github.com/xrash/smetrics
func JaroWinkler(a, b string) float64 {
	if a == "" || b == "" {
		return 0.00
	}
	return smetrics.JaroWinkler(a, b, 0.7, 4)
}

// Levenshtein is the "edit distance" between two strings. This is the count of operations
// (insert, delete, replace) needed for two strings to be equal.
func Levenshtein(a, b string) float64 {
	if a == "" || b == "" {
		return 0.00
	}

	length := utf8.RuneCountInString(a)
	if n := utf8.RuneCountInString(b); n > length {
		length = n // set length to larger value of the two strings
	}
	ed := float64(smetrics.WagnerFischer(a, b, 1, 1, 2))

	score := ed / float64(length)
	if score > 1.00 {
		// If more edits are required than the string's length a and b aren't equal.
		return 0.00
	}
	return 1 - score
}

// Soundex is a phonetic algorithm that considers how the words sound in english.
// Soundex maps a name to a 4-byte string consisting of the first letter of the original string and three numbers. Strings that sound similar should map to the same thing.
//
// Retruned is Hamming computed over both phonetic outputs.
func Soundex(a, b string) float64 {
	if a == "" || b == "" {
		return 0.00
	}
	return Hamming(smetrics.Soundex(a), smetrics.Soundex(b))
}

// Hamming distance is the minimum number of substitutions required to change one string into the other.
func Hamming(a, b string) float64 {
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
