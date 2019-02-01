// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package strcmp

import (
	"flag"
	"math"
	"strings"
	"testing"

	"github.com/docker/docker/pkg/namesgenerator"
)

var (
	flagIterations = flag.Int("iterations", 1000, "How many iterations of each algorithm to test")
)

func TestLevenshtein(t *testing.T) {
	cases := []struct {
		input []string
		score float64
	}{
		{[]string{"abc", "abb"}, 0.66},
		{[]string{"abc", "abc"}, 1.0},
		{[]string{"abcd", "aabc"}, 0.50},
	}
	for i := range cases {
		out := Levenshtein(cases[i].input[0], cases[i].input[1])
		if !eql(out, cases[i].score) {
			t.Errorf("(%s, %s)=%.3f but expected %.3f", cases[i].input[0], cases[i].input[1], out, cases[i].score)
		}
	}
}

func TestSoundex(t *testing.T) {
	for i := 0; i < 500; i += 1 {
		parts := strings.Split(namesgenerator.GetRandomName(0), "_")
		if len(parts) < 2 {
			continue // invalid random name
		}

		a, b := parts[0], parts[1]
		score := Soundex(a, b)
		if score > 1.0 || score < 0.0 {
			t.Fatalf("a=%q b=%q got score %.2f", a, b, score)
		}

		score = Soundex(a, a)
		if !eql(score, 1.0) {
			t.Fatalf("a=%q b=%q got score: %.2f", a, a, score)
		}
	}

	type test struct {
		a, b  string
		score float64
	}

	// Static tests
	cases := []test{
		{"Adam", "Bob", 0.25},
		{"Euler", "Ellery", 1.0},
		{"Lloyd", "Ladd", 1.0},
	}
	for i := range cases {
		score := Soundex(cases[i].a, cases[i].b)
		if !eql(score, cases[i].score) {
			t.Fatalf("a=%q b=%q got score: %.2f and expected: %.2f", cases[i].a, cases[i].b, score, cases[i].score)
		}
	}
}

func eql(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}
