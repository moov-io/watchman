// Package corpusfile sketches an on-disk corpus layout that keeps only the
// search hot path in process memory and parks cold fields (primarily SourceData)
// in an immutable snapshot file.
//
// # Motivation
//
// Watchman today holds []search.Entity in RAM, including SourceData (original
// list payloads) that similarity scoring never reads. Refresh also double-buffers
// the full object graph. This package explores a tiered layout:
//
//	HOT (RAM)     — fields Similarity / candidate indexes need
//	COLD (disk)   — SourceData and other response-only payload
//	INDEX (RAM)   — partitions, name tokens, crypto keys (unchanged model)
//
// # Snapshot file layout (v0 sketch)
//
//	Offset 0
//	┌─────────────────────────────────────────────────────────────┐
//	│ Header (128 bytes)                                          │
//	│   magic[8] = "WMCORPUS"                                     │
//	│   version u32, flags u32, entity_count u32                  │
//	│   section offsets/lengths for string table, hot table,      │
//	│   cold directory, cold blob region                          │
//	├─────────────────────────────────────────────────────────────┤
//	│ String table (optional densification)                       │
//	│   u32 count, then repeated: u32 len + utf8 bytes            │
//	│   Shared tokens: countries, programs, currencies, …         │
//	├─────────────────────────────────────────────────────────────┤
//	│ Hot table                                                   │
//	│   entity_count records, each:                               │
//	│     fixed header (type, source ids as string-table refs)    │
//	│     varint section lengths + packed scoring fields          │
//	│     cold_offset u64, cold_length u32  → into cold region    │
//	├─────────────────────────────────────────────────────────────┤
//	│ Cold directory                                              │
//	│   entity_count × { offset u64, length u32, crc32 u32 }      │
//	├─────────────────────────────────────────────────────────────┤
//	│ Cold blob region                                            │
//	│   per entity: length-prefixed payload (see ColdPayload)     │
//	│   Payload encoding: codec flag + bytes (gob/zstd+gob/…)     │
//	└─────────────────────────────────────────────────────────────┘
//
// Refresh writes corpus-vN.snap, fsyncs, atomically renames, then swaps the
// open reader (same lifecycle as today's index.Lists.Update).
//
// # Field placement
//
// Hot (needed by pkg/search Similarity* and inverted indexes):
//   - Name, Type, Source, SourceID
//   - Person / Business / Organization / Aircraft / Vessel (IDs, titles, dates)
//   - CryptoAddresses, Affiliations, SanctionsInfo, HistoricalInfo
//   - PreparedFields (names, tokens, weights, prepared addresses, contact)
//   - Addresses used only via PreparedFields.Addresses on the score path; the
//     raw Addresses slice is still kept hot for API parity until a thinner
//     response DTO exists
//
// Cold (not read during scoreSimilarityFast):
//   - SourceData — original SDN/CSL/etc. structs
//   - Raw Addresses (scoring uses PreparedFields.Addresses)
//   - Raw phones/faxes/websites (phones/faxes score via PreparedFields.Contact)
//   - SanctionsInfo.Description, Affiliation.Details, HistoricalInfo.Date
//   - Type-level Name/AltNames mirrored by Entity.Name + PreparedFields
//
// Default cold codec is gob+zstd (see DefaultColdCodec).
//
// # Benchmarks
//
// Paired benches live in this package:
//
//	BenchmarkCorpus_Baseline_*  — today's full in-memory []Entity
//	BenchmarkCorpus_Tiered_*    — hot slice in RAM + cold blobs on disk
//
// Run both and compare heap, file size, score latency, and hydrate cost:
//
//	go test ./internal/corpusfile/ -bench='BenchmarkCorpus_' -benchmem -count=5
//
// See layout.go and split.go for the concrete types used by the prototype.
package corpusfile
