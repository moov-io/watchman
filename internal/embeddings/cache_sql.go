package embeddings

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/moov-io/watchman/internal/db"

	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type sqlRepository struct {
	db    db.DB
	table string
}

func newSqlRepository(ctx context.Context, config Config, database db.DB) (*sqlRepository, error) {
	table := createTableName(config)

	ctx, span := telemetry.StartSpan(ctx, "embeddings-cache-sql", trace.WithAttributes(
		attribute.String("table", table),
	))
	defer span.End()

	err := setupTable(ctx, table, database)
	if err != nil {
		err = fmt.Errorf("creating %s failed: %w", table, err)
		span.RecordError(err)
		return nil, err
	}

	return &sqlRepository{
		db:    database,
		table: table,
	}, nil
}

const (
	maxDataSize = 512
)

func createTableName(config Config) string {
	table := fmt.Sprintf("cache_%s_dim%d", config.Provider.Model, config.Provider.Dimension)

	return strings.ReplaceAll(table, "-", "_")
}

func setupTable(ctx context.Context, table string, database db.DB) error {
	if database == nil {
		return fmt.Errorf("missing database: %v", database)
	}

	var create string

	switch database.Dialect() {
	case "mysql":
		// MySQL 9.5+ supports native JSON for storing vectors
		create = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			data VARCHAR(%d) PRIMARY KEY,
			embedding JSON NOT NULL
		)`, table, maxDataSize)

	case "postgres":
		// PostgreSQL using ARRAY type for native vector storage
		create = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			data VARCHAR(%d) PRIMARY KEY,
			embedding DOUBLE PRECISION[] NOT NULL
		)`, table, maxDataSize)

	default:
		return fmt.Errorf("unknown dialect: %v", database.Dialect())
	}

	_, err := database.ExecContext(ctx, create)
	return err
}

var _ Cache = (&sqlRepository{})

func (r *sqlRepository) Get(ctx context.Context, data string) ([]float64, bool) {
	ctx, span := telemetry.StartSpan(ctx, "embeddings-cache-sql-get", trace.WithAttributes(
		attribute.String("table", r.table),
	))
	defer span.End()

	var query string
	switch r.db.Dialect() {
	case "mysql":
		query = fmt.Sprintf(`SELECT embedding FROM %s WHERE data = ?`, r.table)
	case "postgres":
		query = fmt.Sprintf(`SELECT embedding FROM %s WHERE data = $1`, r.table)
	default:
		span.RecordError(fmt.Errorf("unknown dialect: %v", r.db.Dialect()))
		return nil, false
	}

	var embStr string
	err := r.db.QueryRowContext(ctx, query, data).Scan(&embStr)
	if err != nil {
		span.RecordError(err)
		return nil, false
	}

	var emb []float64
	switch r.db.Dialect() {
	case "mysql":
		// MySQL returns JSON array like "[0.1,0.2,0.3]"
		emb, err = parseVector(embStr)
	case "postgres":
		// PostgreSQL returns array like "{0.1,0.2,0.3}"
		emb, err = parsePostgresArray(embStr)
	default:
		return nil, false
	}

	if err != nil {
		span.RecordError(err)
		return nil, false
	}
	return emb, true
}

// parseVector converts a JSON array string like "[0.1,0.2,0.3]" to a slice of float64 (for MySQL).
func parseVector(embStr string) ([]float64, error) {
	if len(embStr) < 2 || embStr[0] != '[' || embStr[len(embStr)-1] != ']' {
		return nil, fmt.Errorf("invalid vector format: %s", embStr)
	}
	trimmed := embStr[1 : len(embStr)-1]
	if trimmed == "" {
		return []float64{}, nil
	}
	parts := strings.Split(trimmed, ",")
	emb := make([]float64, len(parts))
	for i, part := range parts {
		val, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse float at index %d: %w", i, err)
		}
		emb[i] = val
	}
	return emb, nil
}

// parsePostgresArray converts a PostgreSQL array string like "{0.1,0.2,0.3}" to a slice of float64.
func parsePostgresArray(embStr string) ([]float64, error) {
	if len(embStr) < 2 || embStr[0] != '{' || embStr[len(embStr)-1] != '}' {
		return nil, fmt.Errorf("invalid postgres array format: %s", embStr)
	}
	trimmed := embStr[1 : len(embStr)-1]
	if trimmed == "" {
		return []float64{}, nil
	}
	parts := strings.Split(trimmed, ",")
	emb := make([]float64, len(parts))
	for i, part := range parts {
		val, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse float at index %d: %w", i, err)
		}
		emb[i] = val
	}
	return emb, nil
}

func (r *sqlRepository) Put(ctx context.Context, data string, embedding []float64) {
	ctx, span := telemetry.StartSpan(ctx, "embeddings-cache-sql-put", trace.WithAttributes(
		attribute.String("table", r.table),
	))
	defer span.End()

	key := data
	if len(data) > maxDataSize {
		key = data[:maxDataSize]
	}

	var query string
	var embStr string

	switch r.db.Dialect() {
	case "mysql":
		query = fmt.Sprintf(`INSERT INTO %s (data, embedding) VALUES (?, ?) ON DUPLICATE KEY UPDATE embedding = VALUES(embedding)`, r.table)
		embStr = formatVector(embedding)
		_, err := r.db.ExecContext(ctx, query, key, embStr)
		if err != nil {
			span.RecordError(err)
		}

	case "postgres":
		query = fmt.Sprintf(`INSERT INTO %s (data, embedding) VALUES ($1, $2) ON CONFLICT (data) DO UPDATE SET embedding = EXCLUDED.embedding`, r.table)
		embStr = formatPostgresArray(embedding)
		_, err := r.db.ExecContext(ctx, query, key, embStr)
		if err != nil {
			span.RecordError(err)
		}

	default:
		span.RecordError(fmt.Errorf("unknown dialect: %v", r.db.Dialect()))
	}
}

// formatVector converts a slice of float64 to a JSON array string, e.g., "[0.1,0.2,0.3]" (for MySQL).
func formatVector(embedding []float64) string {
	if len(embedding) == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i, val := range embedding {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%f", val))
	}
	sb.WriteString("]")
	return sb.String()
}

// formatPostgresArray converts a slice of float64 to a PostgreSQL array string, e.g., "{0.1,0.2,0.3}".
func formatPostgresArray(embedding []float64) string {
	if len(embedding) == 0 {
		return "{}"
	}
	var sb strings.Builder
	sb.WriteString("{")
	for i, val := range embedding {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%f", val))
	}
	sb.WriteString("}")
	return sb.String()
}
