package corpusfile

import "encoding/binary"

// File format constants for the corpus snapshot sketch.
// The on-disk writer/reader in this package implements a pragmatic subset:
// header + cold directory + cold blobs, with hot entities still held as Go
// values in RAM (SourceData stripped). A denser binary hot table / string
// table can replace that without changing ColdPayload.

const (
	// Magic identifies a Watchman corpus snapshot.
	Magic = "WMCORPUS"

	// Version is the snapshot format version written by this package.
	Version uint32 = 1

	// HeaderSize is the fixed header length in bytes.
	HeaderSize = 128

	// Codec identity for cold payloads.
	CodecGob uint8 = 1
	// CodecGobZstd is reserved for compressed cold blobs (not required for baseline).
	CodecGobZstd uint8 = 2
)

// Header is the fixed 128-byte file preamble.
//
// Binary layout (little-endian):
//
//	0   magic[8]
//	8   version u32
//	12  flags u32
//	16  entity_count u32
//	20  reserved u32
//	24  string_table_off u64
//	32  string_table_len u64
//	40  hot_table_off u64
//	48  hot_table_len u64
//	56  cold_dir_off u64
//	64  cold_dir_len u64
//	72  cold_blob_off u64
//	80  cold_blob_len u64
//	88  created_unix_nano i64
//	96  reserved[32]
type Header struct {
	Magic         [8]byte
	Version       uint32
	Flags         uint32
	EntityCount   uint32
	_             uint32
	StringTableOff uint64
	StringTableLen uint64
	HotTableOff   uint64
	HotTableLen   uint64
	ColdDirOff    uint64
	ColdDirLen    uint64
	ColdBlobOff   uint64
	ColdBlobLen   uint64
	CreatedUnixNano int64
}

// Flag bits for Header.Flags.
const (
	FlagHotInFile uint32 = 1 << iota // hot table present on disk (future)
	FlagHasStringTable
)

// ColdDirEntry locates one entity's cold blob inside the blob region.
// On disk: offset u64, length u32, crc32 u32 (16 bytes).
type ColdDirEntry struct {
	Offset uint64
	Length uint32
	CRC32  uint32
}

const coldDirEntrySize = 16

// PutHeader serializes h into dst, which must be at least HeaderSize bytes.
func PutHeader(dst []byte, h Header) {
	if len(dst) < HeaderSize {
		panic("corpusfile: header buffer too small")
	}
	copy(dst[0:8], h.Magic[:])
	binary.LittleEndian.PutUint32(dst[8:12], h.Version)
	binary.LittleEndian.PutUint32(dst[12:16], h.Flags)
	binary.LittleEndian.PutUint32(dst[16:20], h.EntityCount)
	binary.LittleEndian.PutUint64(dst[24:32], h.StringTableOff)
	binary.LittleEndian.PutUint64(dst[32:40], h.StringTableLen)
	binary.LittleEndian.PutUint64(dst[40:48], h.HotTableOff)
	binary.LittleEndian.PutUint64(dst[48:56], h.HotTableLen)
	binary.LittleEndian.PutUint64(dst[56:64], h.ColdDirOff)
	binary.LittleEndian.PutUint64(dst[64:72], h.ColdDirLen)
	binary.LittleEndian.PutUint64(dst[72:80], h.ColdBlobOff)
	binary.LittleEndian.PutUint64(dst[80:88], h.ColdBlobLen)
	binary.LittleEndian.PutUint64(dst[88:96], uint64(h.CreatedUnixNano))
}

// ReadHeader parses a Header from src.
func ReadHeader(src []byte) (Header, error) {
	var h Header
	if len(src) < HeaderSize {
		return h, errShort("header", HeaderSize, len(src))
	}
	copy(h.Magic[:], src[0:8])
	if string(h.Magic[:]) != Magic {
		return h, errBadMagic(string(h.Magic[:]))
	}
	h.Version = binary.LittleEndian.Uint32(src[8:12])
	h.Flags = binary.LittleEndian.Uint32(src[12:16])
	h.EntityCount = binary.LittleEndian.Uint32(src[16:20])
	h.StringTableOff = binary.LittleEndian.Uint64(src[24:32])
	h.StringTableLen = binary.LittleEndian.Uint64(src[32:40])
	h.HotTableOff = binary.LittleEndian.Uint64(src[40:48])
	h.HotTableLen = binary.LittleEndian.Uint64(src[48:56])
	h.ColdDirOff = binary.LittleEndian.Uint64(src[56:64])
	h.ColdDirLen = binary.LittleEndian.Uint64(src[64:72])
	h.ColdBlobOff = binary.LittleEndian.Uint64(src[72:80])
	h.ColdBlobLen = binary.LittleEndian.Uint64(src[80:88])
	h.CreatedUnixNano = int64(binary.LittleEndian.Uint64(src[88:96]))
	return h, nil
}

func putColdDirEntry(dst []byte, e ColdDirEntry) {
	binary.LittleEndian.PutUint64(dst[0:8], e.Offset)
	binary.LittleEndian.PutUint32(dst[8:12], e.Length)
	binary.LittleEndian.PutUint32(dst[12:16], e.CRC32)
}

func readColdDirEntry(src []byte) ColdDirEntry {
	return ColdDirEntry{
		Offset: binary.LittleEndian.Uint64(src[0:8]),
		Length: binary.LittleEndian.Uint32(src[8:12]),
		CRC32:  binary.LittleEndian.Uint32(src[12:16]),
	}
}
