// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/docker/docker/pkg/namesgenerator"
)

func BenchmarkSearch__Addresses(b *testing.B) {
	b.StopTimer()
	searcher := createBenchmarkSearcher(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopAddresses(10, 0.0, randomName())
	}
}

func BenchmarkSearch__BISEntities(b *testing.B) {
	b.StopTimer()
	searcher := createBenchmarkSearcher(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopBISEntities(10, 0.0, randomName())
	}
}

func BenchmarkSearch__DPs(b *testing.B) {
	b.StopTimer()
	searcher := createBenchmarkSearcher(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopDPs(10, 0.0, randomName())
	}
}

func BenchmarkSearch__SDNs(b *testing.B) {
	b.StopTimer()
	searcher := createBenchmarkSearcher(b)
	keeper := keepSDN(filterRequest{})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSDNs(10, 0.0, randomName(), keeper)
	}
}

func BenchmarkSearch__SSIs(b *testing.B) {
	b.StopTimer()
	searcher := createBenchmarkSearcher(b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSSIs(10, 0.0, randomName())
	}
}

func randomName() string {
	return namesgenerator.GetRandomName(0)
}
