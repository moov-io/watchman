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
		ofactest.EntityForBenchmark(b, "15102"),
		ofactest.EntityForBenchmark(b, "29702"),
		ofactest.EntityForBenchmark(b, "48603"),
	}
	bench(b, "individuals", individuals, false)
	bench(b, "individuals", individuals, true)

	// businesses
	businesses := []search.Entity[search.Value]{
		ofactest.EntityForBenchmark(b, "12685"),
		ofactest.EntityForBenchmark(b, "28603"),
		ofactest.EntityForBenchmark(b, "33151"),
		ofactest.EntityForBenchmark(b, "44525"),
		ofactest.EntityForBenchmark(b, "50544"),
	}
	bench(b, "businesses", businesses, false)
	bench(b, "businesses", businesses, true)

	// vessel
	vessels := []search.Entity[search.Value]{
		ofactest.EntityForBenchmark(b, "47371"),
		ofactest.EntityForBenchmark(b, "50972"),
		ofactest.EntityForBenchmark(b, "52327"),
	}
	bench(b, "vessels", vessels, false)
	bench(b, "vessels", vessels, true)

	// aircraft
	aircraft := []search.Entity[search.Value]{
		ofactest.EntityForBenchmark(b, "11195"),
		ofactest.EntityForBenchmark(b, "19709"),
		ofactest.EntityForBenchmark(b, "48727"),
	}
	bench(b, "aircraft", aircraft, false)
	bench(b, "aircraft", aircraft, true)
}

func BenchmarkSimilarity_TypeRecast(b *testing.B) {
	person := ofactest.EntityForBenchmark(b, "48603") // Dmitry Khoroshev
	same := person
	asBusiness := person
	asBusiness.Type = search.EntityBusiness
	asBusiness.Person = nil
	asBusiness.Business = &search.Business{
		Name:          person.Name,
		AltNames:      nil,
		GovernmentIDs: nil,
	}
	if person.Person != nil {
		asBusiness.Business.Name = person.Person.Name
		asBusiness.Business.AltNames = person.Person.AltNames
		asBusiness.Business.GovernmentIDs = person.Person.GovernmentIDs
	}
	asBusiness = asBusiness.Normalize()

	vessel := ofactest.EntityForBenchmark(b, "47371")

	b.ReportAllocs()

	b.Run("same-type-person", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.Similarity(same, person)
		}
	})
	b.Run("legalentity-as-business-vs-person", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.Similarity(asBusiness, person)
		}
	})
	b.Run("person-vs-vessel-hard-zero", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.Similarity(same, vessel)
		}
	})
}
