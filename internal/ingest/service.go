package ingest

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strings"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/senzing"
	"github.com/moov-io/watchman/pkg/search"
)

type Service interface {
	ReadEntitiesFromFile(ctx context.Context, name string, contents io.Reader) (FileEntities, error)
	ReplaceEntities(ctx context.Context, fileType string, entities []search.Entity[search.Value]) error
	GetEntitiesBySource(ctx context.Context, source string) ([]search.Entity[search.Value], error)
}

func NewService(logger log.Logger, conf Config, repo Repository) Service {
	logger.Info().Logf("found %d ingest fileTypes, using %T", len(conf.Files), repo)

	return &service{
		logger: logger,
		conf:   conf,
		repo:   repo,
	}
}

type service struct {
	logger log.Logger
	conf   Config
	repo   Repository
}

type FileEntities struct {
	FileType string
	Entities []search.Entity[search.Value]
}

func (s *service) ReadEntitiesFromFile(ctx context.Context, name string, contents io.Reader) (FileEntities, error) {
	var out FileEntities
	var err error

	for fileType, schema := range s.conf.Files {
		if strings.EqualFold(name, fileType) {
			// Process the file according to the schema type
			switch Format(strings.ToLower(string(schema.Format))) {
			case FormatCSV:
				out, err = s.readEntitiesFromCSVFile(ctx, fileType, schema, contents)

			case FormatSenzing, FormatSenzingJSON, FormatSenzingJSONL:
				out, err = s.readEntitiesFromSenzingFile(ctx, fileType, contents)

			default:
				return out, fmt.Errorf("unknown format %v", schema.Format)
			}

			if err == nil {
				s.logger.With(log.Fields{
					"name":      log.String(name),
					"file_type": log.String(fileType),
				}).Logf("read %d entities", len(out.Entities))
			}
		}
	}
	if err != nil {
		return out, fmt.Errorf("reading entities from %s failed: %w", name, err)
	}

	// Merge the entities
	out.Entities = search.Merge(out.Entities)

	s.logger.With(log.Fields{
		"name": log.String(name),
	}).Logf("after merge have %d entities", len(out.Entities))

	if len(out.Entities) == 0 {
		return out, fmt.Errorf("schema %s not found", name)
	}
	return out, nil
}

func maybeDecompressBody(r io.Reader) io.Reader {
	header := make([]byte, 10)
	n, err := io.ReadFull(r, header)
	if err != nil || n < 2 {
		if n > 0 {
			return io.MultiReader(bytes.NewReader(header[:n]), r)
		}
		return r
	}

	gz, err := gzip.NewReader(io.MultiReader(bytes.NewReader(header), r))
	if err == nil {
		return gz
	}

	return io.MultiReader(bytes.NewReader(header), r)
}

func (s *service) readEntitiesFromCSVFile(ctx context.Context, name string, schema File, contents io.Reader) (FileEntities, error) {
	out := FileEntities{
		FileType: name,
	}

	r := csv.NewReader(maybeDecompressBody(contents))

	headers, err := r.Read()
	if err != nil {
		return out, fmt.Errorf("problem reading headers: %w", err)
	}
	for idx := range headers {
		// Sometimes we get UTF-8 replacement characters (U+FFFD) in files which can be ignored
		headers[idx] = strings.ReplaceAll(headers[idx], "\uFFFD", "")
	}

	for {
		// Read each row until we run out
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return out, fmt.Errorf("problem reading row: %w", err)
		}
		for idx := range row {
			// ignore UTF-8 replacement characters
			row[idx] = strings.ReplaceAll(row[idx], "\uFFFD", "")
		}

		var entity search.Entity[search.Value]

		// Read First set of common fields
		entity.Name = readColumnDef(headers, schema.Mapping.Name, row)
		entity.Type = search.EntityType(readType(headers, schema.Mapping.Type, row))
		entity.Source = search.SourceList(name)
		entity.SourceID = readColumnDef(headers, schema.Mapping.SourceID, row)

		// Read Business, Person, etc fields
		if schema.Mapping.Person != nil {
			entity.Person = &search.Person{
				Name:          entity.Name,
				AltNames:      readColumnArrayDef(headers, schema.Mapping.Person.AltNames, row),
				GovernmentIDs: readGovernmentIDs(headers, schema.Mapping.Person.GovernmentIDs, row),
			}

			birthDate, err := readTime(readColumnDef(headers, schema.Mapping.Person.BirthDate, row))
			if err != nil {
				return out, fmt.Errorf("reading person birth date: %w", err)
			}
			if !birthDate.IsZero() {
				entity.Person.BirthDate = &birthDate
			}
		}
		if schema.Mapping.Business != nil {
			entity.Business = &search.Business{
				Name:          entity.Name,
				AltNames:      readColumnArrayDef(headers, schema.Mapping.Business.AltNames, row),
				GovernmentIDs: readGovernmentIDs(headers, schema.Mapping.Business.GovernmentIDs, row),
			}

			created, err := readTime(readColumnDef(headers, schema.Mapping.Business.Created, row))
			if err != nil {
				return out, fmt.Errorf("reading business creation time: %w", err)
			}
			if !created.IsZero() {
				entity.Business.Created = &created
			}
		}

		// Read More common fields
		entity.Contact.PhoneNumbers = readColumnArrayDef(headers, schema.Mapping.Contact.PhoneNumbers, row)

		entity.Addresses = readAddresses(headers, schema.Mapping.Addresses, row)

		out.Entities = append(out.Entities, entity.Normalize())
	}

	return out, nil
}

func readType(headers []string, def Type, row []string) string {
	return def.Default
}

func readColumnDef(headers []string, def ColumnDef, row []string) string {
	var fields []string

	if def.Column != "" {
		idx := slices.Index(headers, def.Column)
		if idx >= 0 && idx < len(row) {
			value := strings.TrimSpace(row[idx])
			if value != "" {
				fields = append(fields, value)
			}
		}
	}
	for _, col := range def.Merge {
		idx := slices.Index(headers, col)
		if idx >= 0 && idx < len(row) {
			value := strings.TrimSpace(row[idx])
			if value != "" {
				fields = append(fields, value)
			}
		}
	}

	return strings.Join(fields, " ")
}

func readColumnArrayDef(headers []string, def ColumnArrayDef, row []string) []string {
	var fields []string

	if def.Columns != "" {
		idx := slices.Index(headers, def.Columns)
		if idx >= 0 && idx < len(row) {
			value := strings.TrimSpace(row[idx])
			if value != "" {
				fields = append(fields, value)
			}
		}
	}

	var name string
	for _, col := range def.Merge {
		idx := slices.Index(headers, col)
		if idx >= 0 && idx < len(row) {
			value := strings.TrimSpace(row[idx])
			if value != "" {
				name += fmt.Sprintf(" %s", value)
			}
		}
	}

	name = strings.TrimSpace(name)

	if name != "" {
		fields = append(fields, name)
	}

	return fields
}

var (
	acceptedTimeFormats = []string{
		"2006-01-02",
		"1/2/2006",
		"01/02/2006",
	}
)

func readTime(value string) (tt time.Time, err error) {
	if value == "" {
		return
	}
	for _, fmt := range acceptedTimeFormats {
		tt, err = time.Parse(fmt, value)
		if err == nil {
			return tt, nil
		}
	}
	return
}

func readGovernmentIDs(headers []string, def GovernmentIDs, row []string) (out []search.GovernmentID) {
	var id search.GovernmentID
	// id.Type = search.GovernmentIDType()
	// id.Country = ""
	id.Identifier = readColumnDef(headers, def.Identifier, row)

	if id.Identifier != "" {
		return append(out, id)
	}
	return
}

func readAddresses(headers []string, def Addresses, row []string) (out []search.Address) {
	line1 := readColumnArrayDef(headers, def.Line1, row)
	line2 := readColumnArrayDef(headers, def.Line2, row)
	city := readColumnArrayDef(headers, def.City, row)
	state := readColumnArrayDef(headers, def.State, row)
	postalCode := readColumnArrayDef(headers, def.PostalCode, row)
	country := readColumnArrayDef(headers, def.Country, row)

	for idx := range line1 {
		var addr search.Address

		if idx < len(line1) {
			addr.Line1 = line1[idx]
		}
		if idx < len(line2) {
			addr.Line2 = line2[idx]
		}
		if idx < len(city) {
			addr.City = city[idx]
		}
		if idx < len(state) {
			addr.State = state[idx]
		}
		if idx < len(postalCode) {
			addr.PostalCode = postalCode[idx]
		}
		if idx < len(country) {
			addr.Country = country[idx]
		}

		if addr.Line1 != "" {
			out = append(out, addr)
		}
	}

	return
}

func (s *service) ReplaceEntities(ctx context.Context, fileType string, entities []search.Entity[search.Value]) error {
	if s.repo == nil {
		return nil
	}

	err := s.repo.Upsert(ctx, fileType, entities)
	if err != nil {
		return fmt.Errorf("problem replacing %s entities: %w", fileType, err)
	}

	return nil
}

func (s *service) readEntitiesFromSenzingFile(ctx context.Context, name string, contents io.Reader) (FileEntities, error) {
	out := FileEntities{
		FileType: name,
	}

	sourceList := search.SourceList(name)

	entities, err := senzing.ReadEntities(maybeDecompressBody(contents), sourceList)
	if err != nil {
		return out, fmt.Errorf("reading senzing entities: %w", err)
	}

	out.Entities = entities

	return out, nil
}

const defaultPaginationLimit = 1000

func (s *service) GetEntitiesBySource(ctx context.Context, source string) ([]search.Entity[search.Value], error) {
	if s.repo == nil {
		return nil, fmt.Errorf("no repository configured")
	}

	var allEntities []search.Entity[search.Value]
	lastSourceID := ""
	limit := s.conf.PaginationLimit
	if limit <= 0 {
		limit = defaultPaginationLimit
	}

	for {
		entities, err := s.repo.ListBySource(ctx, lastSourceID, search.SourceList(source), limit)
		if err != nil {
			return nil, fmt.Errorf("listing entities by source %s: %w", source, err)
		}

		if len(entities) == 0 {
			break
		}

		allEntities = append(allEntities, entities...)
		lastSourceID = entities[len(entities)-1].SourceID

		if len(entities) < limit {
			break
		}
	}

	return allEntities, nil
}
