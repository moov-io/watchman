package ingest

import (
	"net/http"

	"github.com/moov-io/base/log"

	"github.com/gorilla/mux"
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
		Path("/v2/ingest/{name}").
		HandlerFunc(c.ingestFile)

	return router
}

func (c *controller) ingestFile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
