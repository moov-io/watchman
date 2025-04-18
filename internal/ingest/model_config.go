package ingest

type Config struct {
	Files map[string]File
}

type File struct {
	Format  Format
	Mapping Mapping
}

type Format string

var (
	FormatCSV Format = "csv"
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
