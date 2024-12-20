package search

import (
	"net/http"

	"github.com/moov-io/base/log"

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

func (c *controller) search(w http.ResponseWriter, r *http.Request) {
	c.service.Search(r.Context())
}
