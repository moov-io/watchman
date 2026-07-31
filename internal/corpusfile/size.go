package corpusfile

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"

	"github.com/moov-io/watchman/pkg/search"
)

// SizeReport is a stable, GC-independent estimate of corpus footprint.
// Sizes are gob-encoded byte lengths (a consistent proxy, not RSS).
type SizeReport struct {
	Entities int

	// FullEntityBytes is gob size of the production []Entity (with SourceData).
	FullEntityBytes int64

	// HotBytes is gob size of entities with SourceData cleared.
	HotBytes int64

	// ColdBytes is gob size of cold payloads only (pre-file framing).
	ColdBytes int64

	// SnapshotFileBytes is the on-disk snapshot size when built (0 if not built).
	SnapshotFileBytes int64
}

func (r SizeReport) String() string {
	n := float64(r.Entities)
	if n < 1 {
		n = 1
	}
	return fmt.Sprintf(
		"entities=%d full=%.2fMB (%.0fB/e) hot=%.2fMB (%.0fB/e) cold=%.2fMB (%.0fB/e) snap=%.2fMB cold_frac=%.1f%%",
		r.Entities,
		float64(r.FullEntityBytes)/(1<<20), float64(r.FullEntityBytes)/n,
		float64(r.HotBytes)/(1<<20), float64(r.HotBytes)/n,
		float64(r.ColdBytes)/(1<<20), float64(r.ColdBytes)/n,
		float64(r.SnapshotFileBytes)/(1<<20),
		100*float64(r.ColdBytes)/float64(max64(r.FullEntityBytes, 1)),
	)
}

// MeasureSizes gob-encodes full/hot/cold views for a stable size comparison.
func MeasureSizes(entities []search.Entity[search.Value]) (SizeReport, error) {
	var r SizeReport
	r.Entities = len(entities)

	full, err := gobSize(entities)
	if err != nil {
		return r, fmt.Errorf("full: %w", err)
	}
	r.FullEntityBytes = full

	hot := make([]search.Entity[search.Value], len(entities))
	var coldTotal int64
	for i := range entities {
		split := SplitEntity(entities[i])
		hot[i] = split.Hot
		raw, err := EncodeCold(split.Cold)
		if err != nil {
			return r, fmt.Errorf("cold %d: %w", i, err)
		}
		coldTotal += int64(len(raw))
	}
	r.ColdBytes = coldTotal

	hotBytes, err := gobSize(hot)
	if err != nil {
		return r, fmt.Errorf("hot: %w", err)
	}
	r.HotBytes = hotBytes
	return r, nil
}

func gobSize(v any) (int64, error) {
	if v == nil {
		return 0, nil
	}
	// gob panics on typed nil pointers (e.g. (*Person)(nil) in interface{}).
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface:
		if rv.IsNil() {
			return 0, nil
		}
	}
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		return 0, err
	}
	return int64(buf.Len()), nil
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
