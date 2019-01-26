// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package strcmp

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"math"
	mrand "math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

var (
	flagIterations = flag.Int("iterations", 1000, "How many iterations of each algorithm to test")
)

func init() {
	mrand.Seed(time.Now().Unix())
}

func randString() string {
	size := mrand.Uint32() % 1000 // max string size of 1k
	bs := make([]byte, size)
	n, err := rand.Read(bs)
	if err != nil || n == 0 {
		return ""
	}
	return strings.ToLower(hex.EncodeToString(bs))
}

func TestJaroWinkler(t *testing.T) {
	for i := 0; i < *flagIterations; i += 1 {
		a, b := randString(), randString()
		check(t, a, b, JaroWinkler(a, b))
	}
}

func TestLevenshtein(t *testing.T) {
	for i := 0; i < *flagIterations; i += 1 {
		a, b := randString(), randString()
		check(t, a, b, Levenshtein(a, b))
	}
}

func TestHamming(t *testing.T) {
	for i := 0; i < *flagIterations; i += 1 {
		a, b := randString(), randString()
		check(t, a, b, Hamming(a, b))
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

func one(n float64) bool {
	return eql(n, 1.0)
}

func zero(n float64) bool {
	return eql(n, 0.0)
}

func eql(a, b float64) bool {
	return math.Abs(a-b) < 0.001
}

func check(t *testing.T, a, b string, score float64) {
	t.Helper()

	if one(score) && a != b {
		t.Fatalf("a=%q b=%q matched", a, b)
	}
	if zero(score) && a == b {
		t.Fatalf("a=%q b=%q didn't match", a, b)
	}
	if score > 1.0 || score < 0.0 {
		t.Fatalf("a=%q b=%q got score %.2f", a, b, score)
	}
}
