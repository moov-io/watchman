package ingest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"sort"

	"github.com/moov-io/watchman/pkg/search"
)

// SourceChecksum is a cheap snapshot of one ingested fileType.
// Search compares these rows to the in-memory corpus instead of scanning entity JSON.
type SourceChecksum struct {
	Source      string
	EntityCount int
	Checksum    string
}

// Fingerprint hashes a set of per-source checksums into one cache key.
func Fingerprint(sums []SourceChecksum) string {
	if len(sums) == 0 {
		return ""
	}
	sorted := append([]SourceChecksum(nil), sums...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Source < sorted[j].Source
	})
	h := sha256.New()
	for _, s := range sorted {
		writeChecksum(h, s.Source, []byte(fmt.Sprintf("%d\x00%s", s.EntityCount, s.Checksum)))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func checksumEntities(entities []search.Entity[search.Value]) (string, error) {
	indexed := append([]search.Entity[search.Value](nil), entities...)
	sort.Slice(indexed, func(i, j int) bool {
		if indexed[i].Source != indexed[j].Source {
			return indexed[i].Source < indexed[j].Source
		}
		return indexed[i].SourceID < indexed[j].SourceID
	})

	h := sha256.New()
	for i := range indexed {
		raw, err := json.Marshal(indexed[i])
		if err != nil {
			return "", fmt.Errorf("json marshal: %w", err)
		}
		writeChecksum(h, indexed[i].SourceID, raw)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func writeChecksum(h hash.Hash, sourceID string, raw []byte) {
	h.Write([]byte(sourceID))
	h.Write([]byte{0})
	h.Write(raw)
	h.Write([]byte{0})
}
