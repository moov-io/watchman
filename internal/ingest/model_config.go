package ingest

import (
	"os"
	"strconv"
	"strings"
)

const defaultMaxBodyBytes int64 = 32 << 20 // 32 MiB

type Config struct {
	Files map[string]File

	// PaginationLimit controls the batch size when listing entities from the database.
	// Defaults to 1000 if not set.
	PaginationLimit int

	// MaxBodyBytes is the maximum POST /v2/ingest/{fileType} request body size.
	// Zero or negative uses defaultMaxBodyBytes (32 MiB). Override with INGEST_MAX_BODY_BYTES.
	MaxBodyBytes int64
}

func (c Config) maxBodyBytes() int64 {
	if v := strings.TrimSpace(os.Getenv("INGEST_MAX_BODY_BYTES")); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil && n > 0 {
			return n
		}
	}
	if c.MaxBodyBytes > 0 {
		return c.MaxBodyBytes
	}
	return defaultMaxBodyBytes
}

type File struct {
	Format  Format
	Mapping Mapping
}

type Format string

var (
	FormatCSV Format = "csv"

	// Senzing entity format support
	// Reference: https://www.senzing.com/docs/entity_specification/index.html
	FormatSenzing      Format = "senzing"       // auto-detect JSON Lines or JSON Array
	FormatSenzingJSON  Format = "senzing-json"  // JSON Array format
	FormatSenzingJSONL Format = "senzing-jsonl" // JSON Lines format
)

type Mapping struct {
	Name     ColumnDef
	SourceID ColumnDef
	Type     Type

	Person   *Person
	Business *Business

	Contact   Contact
	Addresses Addresses
}

type ColumnDef struct {
	Column string
	Merge  []string
}

type ColumnArrayDef struct {
	Columns string
	Merge   []string
}

type Type struct {
	Default string
}

type Person struct {
	Name          ColumnDef
	AltNames      ColumnArrayDef
	BirthDate     ColumnDef
	GovernmentIDs GovernmentIDs
}

type Business struct {
	Name          ColumnDef
	AltNames      ColumnArrayDef
	Created       ColumnDef
	GovernmentIDs GovernmentIDs
}

type GovernmentIDs struct {
	Type       ColumnDef
	Identifier ColumnDef
}

type Contact struct {
	PhoneNumbers ColumnArrayDef
}

type Addresses struct {
	Line1      ColumnArrayDef
	Line2      ColumnArrayDef
	City       ColumnArrayDef
	State      ColumnArrayDef
	PostalCode ColumnArrayDef
	Country    ColumnArrayDef
}
