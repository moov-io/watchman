package ingest

import (
	"context"
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/api"
	"github.com/moov-io/watchman/internal/senzing"
	pubsearch "github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
)

type Controller interface {
	AppendRoutes(router *mux.Router) *mux.Router
}

func NewController(logger log.Logger, service Service) Controller {
	return &controller{
		logger:  logger,
		service: service,
	}
}

type controller struct {
	logger  log.Logger
	service Service
}

func (c *controller) AppendRoutes(router *mux.Router) *mux.Router {
	router.
		Name("ingest-file").
		Methods("POST").
		Path("/v2/ingest/{fileType}").
		HandlerFunc(c.ingestFile)

	router.
		Name("export-senzing").
		Methods("GET").
		Path("/v2/export/{fileType}/senzing").
		HandlerFunc(c.exportSenzing)

	return router
}

func (c *controller) ingestFile(w http.ResponseWriter, r *http.Request) {
	fileType := api.CleanUserInput(mux.Vars(r)["fileType"])
	if fileType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, span := telemetry.StartSpan(context.Background(), "api-ingest-file")
	defer span.End()

	span.SetAttributes(attribute.String("file_type", fileType))

	logger := c.logger.With(log.Fields{
		"file_type": log.String(fileType),
	})

	if r.Body != nil {
		defer r.Body.Close()
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	parsedFile, err := c.service.ReadEntitiesFromFile(ctx, fileType, r.Body)
	if err != nil {
		err = logger.Error().LogErrorf("problem reading entities from %s file: %v", fileType, err).Err()
		api.ErrorResponse(w, err)
		return
	}

	logger.Info().Logf("found %d entities from %s file", len(parsedFile.Entities), parsedFile.FileType)

	// Save the parsed entities
	err = c.service.ReplaceEntities(ctx, parsedFile.FileType, parsedFile.Entities)
	if err != nil {
		err = logger.Error().LogErrorf("problem updating %s entities: %v", fileType, err).Err()
		api.ErrorResponse(w, err)
		return
	}

	// Marshal the response
	err = api.JsonResponse(w, pubsearch.IngestFileResponse{
		FileType: parsedFile.FileType,
		Entities: parsedFile.Entities,
	})
	if err != nil {
		span.RecordError(err)
	}
}

func (c *controller) exportSenzing(w http.ResponseWriter, r *http.Request) {
	fileType := api.CleanUserInput(mux.Vars(r)["fileType"])
	if fileType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get format from query parameter (default: jsonl)
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "jsonl"
	}

	ctx, span := telemetry.StartSpan(context.Background(), "api-export-senzing")
	defer span.End()

	span.SetAttributes(
		attribute.String("file_type", fileType),
		attribute.String("format", format),
	)

	logger := c.logger.With(log.Fields{
		"file_type": log.String(fileType),
		"format":    log.String(format),
	})

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set content type based on format
	switch format {
	case "jsonl", "json-lines", "ndjson":
		w.Header().Set("Content-Type", "application/x-ndjson")
	default:
		w.Header().Set("Content-Type", "application/json")
	}

	entities, err := c.service.GetEntitiesBySource(ctx, fileType)
	if err != nil {
		err = logger.Error().LogErrorf("problem getting entities for export: %v", err).Err()
		api.ErrorResponse(w, err)
		return
	}

	logger.Info().Logf("exporting %d entities to senzing format", len(entities))

	opts := senzing.ExportOptions{
		DataSource: fileType,
		Format:     format,
	}

	if err := senzing.WriteEntities(w, entities, opts); err != nil {
		span.RecordError(err)
		logger.Error().LogErrorf("problem writing senzing export: %v", err)
	}
}
