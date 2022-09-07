// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/stretchr/testify/require"
)

func BenchmarkSearch__All(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buildFullSearchResponse(searcher, filterRequest{}, 10, 0.0, randomName())
	}
}

func BenchmarkSearch__Addresses(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopAddresses(10, 0.0, randomName())
	}
}

func BenchmarkSearch__BISEntities(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopBISEntities(10, 0.0, randomName())
	}
}

func BenchmarkSearch__DPs(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopDPs(10, 0.0, randomName())
	}
}

func BenchmarkSearch__SDNsBasic(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	keeper := keepSDN(filterRequest{})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSDNs(10, 0.0, randomName(), keeper)
	}
}

func BenchmarkSearch__SDNsMinMatch50(b *testing.B) {
	minMatch := 0.50
	searcher := createBenchmarkSearcher(b)
	keeper := keepSDN(filterRequest{})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSDNs(10, minMatch, randomName(), keeper)
	}
}

func BenchmarkSearch__SDNsMinMatch95(b *testing.B) {
	minMatch := 0.95
	searcher := createBenchmarkSearcher(b)
	keeper := keepSDN(filterRequest{})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSDNs(10, minMatch, randomName(), keeper)
	}
}

func BenchmarkSearch__SDNsEntity(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	keeper := keepSDN(filterRequest{
		sdnType: "entity",
	})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSDNs(10, 0.0, randomName(), keeper)
	}
}

func BenchmarkSearch__SDNsComplex(b *testing.B) {
	minMatch := 0.95
	searcher := createBenchmarkSearcher(b)
	keeper := keepSDN(filterRequest{
		sdnType: "entity",
	})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSDNs(10, minMatch, randomName(), keeper)
	}
}

func BenchmarkSearch__SSIs(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searcher.TopSSIs(10, 0.0, randomName())
	}
}

func BenchmarkSearch__CSL(b *testing.B) {
	searcher := createBenchmarkSearcher(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resp := buildFullSearchResponseWith(searcher, cslGatherings, filterRequest{}, 10, 0.0, randomName())

		b.StopTimer()
		require.Greater(b, len(resp.BISEntities), 1)
		require.Greater(b, len(resp.MilitaryEndUsers), 1)
		require.Greater(b, len(resp.SectoralSanctions), 1)
		require.Greater(b, len(resp.Unverified), 1)
		require.Greater(b, len(resp.NonproliferationSanctions), 1)
		require.Greater(b, len(resp.ForeignSanctionsEvaders), 1)
		require.Greater(b, len(resp.PalestinianLegislativeCouncil), 1)
		require.Greater(b, len(resp.CaptaList), 1)
		require.Greater(b, len(resp.ITARDebarred), 1)
		require.Greater(b, len(resp.NonSDNChineseMilitaryIndustrialComplex), 1)
		require.Greater(b, len(resp.NonSDNMenuBasedSanctionsList), 1)
		b.StartTimer()
	}
}

func randomName() string {
	return namesgenerator.GetRandomName(0)
}
