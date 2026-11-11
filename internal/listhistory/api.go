// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/api"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
)

type Controller interface {
	AppendRoutes(router *mux.Router) *mux.Router
}

func NewController(logger log.Logger, service *Service) Controller {
	return &controller{
		logger:  logger,
		service: service,
	}
}

var _ Controller = (*controller)(nil)

type controller struct {
	logger  log.Logger
	service *Service
}

func (c *controller) AppendRoutes(router *mux.Router) *mux.Router {
	router.
		Name("ListHistory.lists").
		Methods("GET").
		Path("/v2/lists").
		HandlerFunc(c.listLatest)

	router.
		Name("ListHistory.refs").
		Methods("GET").
		Path("/v2/lists/{listName}/refs").
		HandlerFunc(c.listRefs)

	router.
		Name("ListHistory.entities").
		Methods("GET").
		Path("/v2/lists/{listName}/entities").
		HandlerFunc(c.listEntities)

	router.
		Name("ListHistory.entity").
		Methods("GET").
		Path("/v2/lists/{listName}/entities/{sourceID}").
		HandlerFunc(c.getEntity)

	router.
		Name("ListHistory.diff").
		Methods("GET").
		Path("/v2/lists/{listName}/diff").
		HandlerFunc(c.diff)

	return router
}

func (c *controller) listLatest(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.StartSpan(r.Context(), "api-list-history-lists")
	defer span.End()

	if extra := unusedParams(r); len(extra) > 0 {
		api.ErrorResponse(w, extraQueryParamsError(extra))
		return
	}

	refs, err := c.service.ListLatest(ctx)
	if err != nil {
		c.writeHistoryError(w, err)
		return
	}
	api.JsonResponse(w, ListsResponse{Lists: refs})
}

func (c *controller) listRefs(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.StartSpan(r.Context(), "api-list-history-refs")
	defer span.End()

	listName, ok := parseListName(w, r)
	if !ok {
		return
	}
	span.SetAttributes(attribute.String("list_name", listName))

	if extra := unusedParams(r); len(extra) > 0 {
		api.ErrorResponse(w, extraQueryParamsError(extra))
		return
	}

	refs, err := c.service.ListRefs(ctx, listName)
	if err != nil {
		c.writeHistoryError(w, err)
		return
	}
	if len(refs) == 0 {
		c.writeHistoryError(w, ErrNotFound)
		return
	}
	api.JsonResponse(w, RefsResponse{ListName: listName, Refs: refs})
}

func (c *controller) listEntities(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.StartSpan(r.Context(), "api-list-history-entities")
	defer span.End()

	listName, ok := parseListName(w, r)
	if !ok {
		return
	}
	span.SetAttributes(attribute.String("list_name", listName))

	q := api.NewQueryParams(r.URL)
	refToken := q.Get("ref")
	query := q.Get("q")
	entityType := q.Get("type")
	limit, err := parseLimit(q.Get("limit"), defaultListLimit, maxListLimit)
	if err != nil {
		api.ErrorResponse(w, err)
		return
	}
	after := q.Get("after")
	if extra := q.UnusedQueryParams(); len(extra) > 0 {
		api.ErrorResponse(w, extraQueryParamsError(extra))
		return
	}

	if !validRefToken(refToken) {
		api.ErrorResponse(w, errInvalidRef)
		return
	}
	if !validEntityType(entityType) {
		api.ErrorResponse(w, errInvalidType)
		return
	}

	out, err := c.service.ListEntities(ctx, listName, refToken, EntityFilter{
		Query:      query,
		EntityType: entityType,
		After:      after,
		Limit:      limit,
	})
	if err != nil {
		c.writeHistoryError(w, err)
		return
	}
	api.JsonResponse(w, out)
}

func (c *controller) getEntity(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.StartSpan(r.Context(), "api-list-history-entity")
	defer span.End()

	listName, ok := parseListName(w, r)
	if !ok {
		return
	}
	sourceID := api.CleanUserInput(mux.Vars(r)["sourceID"])
	if sourceID == "" || len(sourceID) > 255 {
		api.ErrorResponse(w, errInvalidSourceID)
		return
	}
	span.SetAttributes(attribute.String("list_name", listName))

	q := api.NewQueryParams(r.URL)
	refToken := q.Get("ref")
	if extra := q.UnusedQueryParams(); len(extra) > 0 {
		api.ErrorResponse(w, extraQueryParamsError(extra))
		return
	}
	if !validRefToken(refToken) {
		api.ErrorResponse(w, errInvalidRef)
		return
	}

	rec, err := c.service.GetEntity(ctx, listName, refToken, sourceID)
	if err != nil {
		c.writeHistoryError(w, err)
		return
	}
	api.JsonResponse(w, rec)
}

func (c *controller) diff(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.StartSpan(r.Context(), "api-list-history-diff")
	defer span.End()

	listName, ok := parseListName(w, r)
	if !ok {
		return
	}
	span.SetAttributes(attribute.String("list_name", listName))

	q := api.NewQueryParams(r.URL)
	fromToken := q.Get("from")
	toToken := q.Get("to")
	detail := q.Get("detail")
	limit, err := parseLimit(q.Get("limit"), defaultDiffLimit, maxDiffLimit)
	if err != nil {
		api.ErrorResponse(w, err)
		return
	}
	after := q.Get("after")
	if extra := q.UnusedQueryParams(); len(extra) > 0 {
		api.ErrorResponse(w, extraQueryParamsError(extra))
		return
	}

	if !validRefToken(fromToken) || !validRefToken(toToken) {
		api.ErrorResponse(w, errInvalidRef)
		return
	}
	if detail != "" && !strings.EqualFold(detail, DetailSummary) && !strings.EqualFold(detail, DetailFull) {
		api.ErrorResponse(w, errInvalidDetail)
		return
	}
	if after != "" {
		if _, _, ok := splitCursor(after); !ok {
			api.ErrorResponse(w, errInvalidCursor)
			return
		}
	}

	out, err := c.service.Diff(ctx, listName, fromToken, toToken, after, detail, limit)
	if err != nil {
		c.writeHistoryError(w, err)
		return
	}
	api.JsonResponse(w, out)
}

func parseListName(w http.ResponseWriter, r *http.Request) (string, bool) {
	listName := api.CleanUserInput(mux.Vars(r)["listName"])
	if !validListName(listName) {
		api.ErrorResponse(w, errInvalidListName)
		return "", false
	}
	return listName, true
}

func unusedParams(r *http.Request) []string {
	return api.NewQueryParams(r.URL).UnusedQueryParams()
}

func extraQueryParamsError(extra []string) error {
	return fmt.Errorf("extra/unused query parameters in request: %s", strings.Join(extra, ","))
}

func parseLimit(raw string, def, max int) (int, error) {
	if strings.TrimSpace(raw) == "" {
		return def, nil
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 {
		return 0, errInvalidLimit
	}
	if n > max {
		return max, nil
	}
	return n, nil
}

func (c *controller) writeHistoryError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrNotConfigured):
		api.ErrorResponseStatus(w, http.StatusNotImplemented, ErrNotConfigured)
	case errors.Is(err, ErrNotFound):
		api.ErrorResponseStatus(w, http.StatusNotFound, ErrNotFound)
	default:
		c.logger.Error().Logf("problem handling list history request: %v", err)
		api.ErrorResponseStatus(w, http.StatusInternalServerError, errors.New("internal error"))
	}
}

var (
	errInvalidListName = errors.New("invalid list name")
	errInvalidRef      = errors.New("invalid ref")
	errInvalidType     = errors.New("invalid type")
	errInvalidLimit    = errors.New("invalid limit")
	errInvalidDetail   = errors.New("invalid detail")
	errInvalidCursor   = errors.New("invalid after cursor")
	errInvalidSourceID = errors.New("invalid sourceID")
)
