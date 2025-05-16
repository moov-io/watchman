package ingest

import (
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/api"
	"github.com/moov-io/watchman/internal/search"
	pubsearch "github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
)

type Controller interface {
	AppendRoutes(router *mux.Router) *mux.Router
}

func NewController(logger log.Logger, service Service, searchService search.Service) Controller {
	return &controller{
		logger:        logger,
		service:       service,
		searchService: searchService,
	}
}

type controller struct {
	logger        log.Logger
	service       Service
	searchService search.Service
}

func (c *controller) AppendRoutes(router *mux.Router) *mux.Router {
	router.
		Name("ingest-file").
		Methods("POST").
		Path("/v2/ingest/{fileType}").
		HandlerFunc(c.ingestFile)

	return router
}

func (c *controller) ingestFile(w http.ResponseWriter, r *http.Request) {
	fileType := api.CleanUserInput(mux.Vars(r)["fileType"])
	if fileType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, span := telemetry.StartSpan(r.Context(), "api-ingest-file")
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
		err = logger.Error().LogErrorf("problem reading entities from %s file: %w", fileType, err).Err()
		api.ErrorResponse(w, err)
		return
	}

	logger.Info().Logf("found %d entities from %s file", len(parsedFile.Entities), parsedFile.FileType)

	// Set the parsed entities in-memory
	ingestKey := pubsearch.SourceList(parsedFile.FileType)
	c.searchService.SetIngestedEntities(ingestKey, parsedFile.Entities)

	err = api.JsonResponse(w, pubsearch.IngestFileResponse{
		FileType: parsedFile.FileType,
		Entities: parsedFile.Entities,
	})
	if err != nil {
		span.RecordError(err)
	}
}
