package ingest

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
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
		Path("/v2/ingest/{type}").
		HandlerFunc(c.ingestFile)

	return router
}

func (c *controller) ingestFile(w http.ResponseWriter, r *http.Request) {
	fileType := api.CleanUserInput(mux.Vars(r)["type"])
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

	q := r.URL.Query()
	debug := strx.Yes(r.URL.Query().Get("debug"))

	searchOpts := search.SearchOpts{
		Limit:          1,
		RequestID:      q.Get("requestID"),
		Debug:          debug,
		DebugSourceIDs: strings.Split(q.Get("debugSourceIDs"), ","),
	}

	// Concurrently search for each
	var wg sync.WaitGroup
	wg.Add(len(parsedFile.Entities))

	errorCh := make(chan error, 1)
	responses := make([]pubsearch.IngestedEntities, len(parsedFile.Entities))

	for idx := range parsedFile.Entities {
		go func(idx int) {
			defer wg.Done()

			// Concurrently run each search
			query := parsedFile.Entities[idx]

			entities, err := c.searchService.Search(ctx, query, searchOpts)
			if err != nil {
				errorCh <- fmt.Errorf("searching %v/%v failed: %w", query.Source, query.SourceID, err)
			}

			responses[idx] = pubsearch.IngestedEntities{
				Query:    query,
				Entities: entities,
			}
		}(idx)
	}

	// Wait for all searches to complete
	wg.Wait()

	// Send at least an empty error
	go func() {
		errorCh <- nil
	}()

	// Check for search errors
	err = <-errorCh
	if err != nil {
		err = logger.Error().LogErrorf("problem running ingest search from %s file: %w", parsedFile.FileType, err).Err()
		api.ErrorResponse(w, err)
		return
	}

	err = api.JsonResponse(w, pubsearch.IngestSearchResponse{
		FileType: parsedFile.FileType,
		Records:  responses,
	})
	if err != nil {
		span.RecordError(err)
	}
}
