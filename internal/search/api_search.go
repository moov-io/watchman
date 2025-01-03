package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
)

// GET /v2/search

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
		Name("Search.v2").
		Methods("GET").
		Path("/v2/search").
		HandlerFunc(c.search)

	return router
}

type searchResponse struct {
	Entities []SearchedEntity[search.Value] `json:"entities"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (c *controller) search(w http.ResponseWriter, r *http.Request) {
	req, err := readSearchRequest(r)
	if err != nil {
		err = fmt.Errorf("problem reading v2 search request: %w", err)
		c.logger.Error().LogError(err)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse{
			Error: err.Error(),
		})

		return
	}

	entities, err := c.service.Search(r.Context(), req)
	if err != nil {
		c.logger.Error().LogErrorf("problem with v2 search: %v", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResponse{
		Entities: entities,
	})
}

func readSearchRequest(r *http.Request) (search.Entity[search.Value], error) {
	q := r.URL.Query()

	var req search.Entity[search.Value]

	req.Name = strings.TrimSpace(q.Get("name"))
	req.Type = search.EntityType(strings.TrimSpace(strings.ToLower(q.Get("entityType"))))
	req.Source = search.SourceAPIRequest
	req.SourceID = strings.TrimSpace(q.Get("requestID"))

	switch req.Type {
	case search.EntityPerson: // "person"

	case search.EntityBusiness: // "business"

	case search.EntityAircraft: // "aircraft"

	case search.EntityVessel: // "vessel"

	default:
		return req, fmt.Errorf("unsupported entityType: %v", req.Type)
	}

	// Person       *Person       `json:"person"`
	// Business     *Business     `json:"business"`
	// Organization *Organization `json:"organization"`
	// Aircraft     *Aircraft     `json:"aircraft"`
	// Vessel       *Vessel       `json:"vessel"`

	// CryptoAddresses []CryptoAddress `json:"cryptoAddresses"`
	// TODO(adam): support multiple values? How does Go handle that?

	// Addresses []Address `json:"addresses"`

	// Affiliations   []Affiliation    `json:"affiliations"`
	// SanctionsInfo  *SanctionsInfo   `json:"sanctionsInfo"`
	// HistoricalInfo []HistoricalInfo `json:"historicalInfo"`
	// Titles         []string         `json:"titles"`

	return req, nil
}
