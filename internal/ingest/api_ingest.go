package ingest

import (
	"context"
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/api"
	pubsearch "github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/senzing"

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
		Name("export-file").
		Methods("GET").
		Path("/v2/export/{fileType}").
		HandlerFunc(c.exportFile)

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

func (c *controller) exportFile(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.StartSpan(context.Background(), "api-export-senzing")
	defer span.End()

	fileType := api.CleanUserInput(mux.Vars(r)["fileType"])
	if fileType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams := api.NewQueryParams(r.URL)
	outputFormat, subformat := api.ChooseEntityFormat(r.Header, queryParams.Get("format"))

	span.SetAttributes(
		attribute.String("file_type", fileType),
		attribute.String("output_format", string(outputFormat)),
	)

	logger := c.logger.With(log.Fields{
		"file_type":     log.String(fileType),
		"output_format": log.String(string(outputFormat)),
	})

	// Read the entities
	entities, err := c.service.GetEntitiesBySource(ctx, fileType)
	if err != nil {
		err = logger.Error().LogErrorf("problem getting entities for export: %v", err).Err()
		span.RecordError(err)

		api.ErrorResponse(w, err)
		return
	}
	logger.Info().Logf("exporting %d entities to senzing format", len(entities))

	switch outputFormat {
	case api.EntityWatchman:
		err = api.JsonResponse(w, entities)

	case api.EntitySenzing:
		// TODO(adam): api.JsonResponse sets these headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		opts := senzing.ExportOptions{
			DataSource: fileType,
			Format:     subformat,
		}
		err = senzing.WriteEntities(w, entities, opts)
	}
	if err != nil {
		err = logger.Error().LogErrorf("problem rendering export response into %v", outputFormat).Err()
		span.RecordError(err)

		api.ErrorResponse(w, err)
		return
	}
}
