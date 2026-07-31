package corpusfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/moov-io/watchman/pkg/search"
)

// Snapshot is an opened corpus: hot entities in RAM, cold blobs on disk.
//
// This is the tiered prototype used by BenchmarkCorpus_Tiered_*. Production
// wiring would replace index.corpus.entities with Snapshot.Hot (or a denser
// binary hot table) and call Hydrate only for top-N API results.
type Snapshot struct {
	path string
	hdr  Header
	hot  []search.Entity[search.Value]
	dir  []ColdDirEntry

	// f is held so cold reads stay valid; Close releases it.
	f *os.File
}

// BuildSnapshot splits entities, writes cold data to path, and returns an open Snapshot.
// Hot records stay in process memory with SourceData cleared.
func BuildSnapshot(path string, entities []search.Entity[search.Value]) (*Snapshot, error) {
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, fmt.Errorf("mkdir snapshot dir: %w", err)
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create snapshot: %w", err)
	}
	ok := false
	defer func() {
		if !ok {
			f.Close()
			os.Remove(path)
		}
	}()

	n := len(entities)
	hot := make([]search.Entity[search.Value], n)
	dir := make([]ColdDirEntry, n)

	// Layout:
	//   [Header 128]
	//   [Cold directory n*16]
	//   [Cold blobs …]
	coldDirOff := uint64(HeaderSize)
	coldDirLen := uint64(n * coldDirEntrySize)
	coldBlobOff := coldDirOff + coldDirLen

	if _, err := f.Seek(int64(coldBlobOff), io.SeekStart); err != nil {
		return nil, err
	}

	var blobLen uint64
	for i := range entities {
		split := SplitEntity(entities[i])
		hot[i] = split.Hot

		raw, err := EncodeCold(split.Cold)
		if err != nil {
			return nil, fmt.Errorf("entity %d: %w", i, err)
		}
		dir[i] = ColdDirEntry{
			Offset: blobLen,
			Length: uint32(len(raw)),
			CRC32:  blobCRC(raw),
		}
		if _, err := f.Write(raw); err != nil {
			return nil, fmt.Errorf("write cold %d: %w", i, err)
		}
		blobLen += uint64(len(raw))
	}

	// Write cold directory
	dirBuf := make([]byte, coldDirLen)
	for i := range dir {
		putColdDirEntry(dirBuf[i*coldDirEntrySize:], dir[i])
	}
	if _, err := f.WriteAt(dirBuf, int64(coldDirOff)); err != nil {
		return nil, fmt.Errorf("write cold dir: %w", err)
	}

	hdr := Header{
		Version:         Version,
		EntityCount:     uint32(n),
		ColdDirOff:      coldDirOff,
		ColdDirLen:      coldDirLen,
		ColdBlobOff:     coldBlobOff,
		ColdBlobLen:     blobLen,
		CreatedUnixNano: nowUnixNano(),
	}
	copy(hdr.Magic[:], Magic)

	var hdrBuf [HeaderSize]byte
	PutHeader(hdrBuf[:], hdr)
	if _, err := f.WriteAt(hdrBuf[:], 0); err != nil {
		return nil, fmt.Errorf("write header: %w", err)
	}
	if err := f.Sync(); err != nil {
		return nil, err
	}

	ok = true
	return &Snapshot{
		path: path,
		hdr:  hdr,
		hot:  hot,
		dir:  dir,
		f:    f,
	}, nil
}

// OpenSnapshot loads hot entities from a side-car is not supported yet; use BuildSnapshot.
// Open reopens cold storage for an existing file when hot is provided separately.
func OpenSnapshot(path string, hot []search.Entity[search.Value]) (*Snapshot, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var hdrBuf [HeaderSize]byte
	if _, err := io.ReadFull(f, hdrBuf[:]); err != nil {
		f.Close()
		return nil, err
	}
	hdr, err := ReadHeader(hdrBuf[:])
	if err != nil {
		f.Close()
		return nil, err
	}
	if int(hdr.EntityCount) != len(hot) {
		f.Close()
		return nil, errEntityCount(int(hdr.EntityCount), len(hot))
	}

	dirBuf := make([]byte, hdr.ColdDirLen)
	if _, err := f.ReadAt(dirBuf, int64(hdr.ColdDirOff)); err != nil {
		f.Close()
		return nil, fmt.Errorf("read cold dir: %w", err)
	}
	dir := make([]ColdDirEntry, hdr.EntityCount)
	for i := range dir {
		dir[i] = readColdDirEntry(dirBuf[i*coldDirEntrySize:])
	}

	return &Snapshot{path: path, hdr: hdr, hot: hot, dir: dir, f: f}, nil
}

// Close releases the underlying file.
func (s *Snapshot) Close() error {
	if s == nil || s.f == nil {
		return nil
	}
	err := s.f.Close()
	s.f = nil
	return err
}

// Len returns entity count.
func (s *Snapshot) Len() int { return len(s.hot) }

// Hot returns the in-memory scoring entity at i (SourceData is nil).
func (s *Snapshot) Hot(i int) search.Entity[search.Value] { return s.hot[i] }

// HotSlice returns the backing hot slice (read-only by convention).
func (s *Snapshot) HotSlice() []search.Entity[search.Value] { return s.hot }

// FileSize returns the snapshot file size in bytes.
func (s *Snapshot) FileSize() (int64, error) {
	st, err := os.Stat(s.path)
	if err != nil {
		return 0, err
	}
	return st.Size(), nil
}

// Header returns the parsed file header.
func (s *Snapshot) Header() Header { return s.hdr }

// ReadCold loads and decodes the cold payload for entity i.
func (s *Snapshot) ReadCold(i int) (ColdPayload, error) {
	if i < 0 || i >= len(s.dir) {
		return ColdPayload{}, fmt.Errorf("corpusfile: index %d out of range", i)
	}
	ent := s.dir[i]
	buf := make([]byte, ent.Length)
	off := int64(s.hdr.ColdBlobOff + ent.Offset)
	if _, err := s.f.ReadAt(buf, off); err != nil {
		return ColdPayload{}, fmt.Errorf("read cold %d: %w", i, err)
	}
	if crc := blobCRC(buf); crc != ent.CRC32 {
		return ColdPayload{}, fmt.Errorf("corpusfile: cold %d crc mismatch", i)
	}
	return DecodeCold(buf)
}

// Hydrate returns a full entity (hot + cold SourceData) for API responses.
func (s *Snapshot) Hydrate(i int) (search.Entity[search.Value], error) {
	cold, err := s.ReadCold(i)
	if err != nil {
		return search.Entity[search.Value]{}, err
	}
	return Reattach(s.hot[i], cold), nil
}

// Path returns the snapshot file path.
func (s *Snapshot) Path() string { return s.path }
