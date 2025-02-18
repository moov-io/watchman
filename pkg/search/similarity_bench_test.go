package search_test

import (
	"bytes"
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"
)

func BenchmarkDebugSimilarity(b *testing.B) {
	// Load OFAC records once, before any subtest starts
	ofactest.FindEntity(b, "29702")

	bench := func(b *testing.B, name string, entities []search.Entity[search.Value], debug bool) {
		b.Helper()

		if debug {
			name += "-debug"
		}
		b.ResetTimer()

		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				query := i % (len(entities) - 1)
				index := (i + 1) % (len(entities) - 1)

				if debug {
					var buf bytes.Buffer
					search.DebugSimilarity(&buf, entities[query], entities[index])
				} else {
					search.Similarity(entities[query], entities[index])
				}
			}
		})
	}

	b.ResetTimer()

	// individuals
	individuals := []search.Entity[search.Value]{
		ofactest.FindEntity(b, "15102"),
		ofactest.FindEntity(b, "29702"),
		ofactest.FindEntity(b, "48603"),
	}
	bench(b, "individuals", individuals, false)
	bench(b, "individuals", individuals, true)

	// businesses
	businesses := []search.Entity[search.Value]{
		ofactest.FindEntity(b, "12685"),
		ofactest.FindEntity(b, "28603"),
		ofactest.FindEntity(b, "33151"),
		ofactest.FindEntity(b, "44525"),
		ofactest.FindEntity(b, "50544"),
	}
	bench(b, "businesses", businesses, false)
	bench(b, "businesses", businesses, true)

	// vessel
	vessels := []search.Entity[search.Value]{
		ofactest.FindEntity(b, "47371"),
		ofactest.FindEntity(b, "50972"),
		ofactest.FindEntity(b, "52327"),
	}
	bench(b, "vessels", vessels, false)
	bench(b, "vessels", vessels, true)

	// aircraft
	aircraft := []search.Entity[search.Value]{
		ofactest.FindEntity(b, "11195"),
		ofactest.FindEntity(b, "19709"),
		ofactest.FindEntity(b, "48727"),
	}
	bench(b, "aircraft", aircraft, false)
	bench(b, "aircraft", aircraft, true)
}
