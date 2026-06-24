package refresh

import (
	"encoding/json"
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/api"

	"github.com/gorilla/mux"
)

type Controller interface {
	AppendRoutes(router *mux.Router) *mux.Router
}

func NewController(logger log.Logger, manager Manager, allowManualRefresh bool) Controller {
	return &controller{
		logger:             logger,
		manager:            manager,
		allowManualRefresh: allowManualRefresh,
	}
}

type controller struct {
	logger             log.Logger
	manager            Manager
	allowManualRefresh bool
}

func (c *controller) AppendRoutes(router *mux.Router) *mux.Router {
	// The status endpoint is always available.
	router.
		Name("DataRefreshStatus.v2").
		Methods("GET").
		Path("/v2/data/refresh").
		HandlerFunc(c.status)

	// The trigger endpoint is opt-in (Download.AllowManualRefresh). When disabled
	// a POST to the path yields 405 Method Not Allowed (the path exists for GET).
	if c.allowManualRefresh {
		router.
			Name("DataRefresh.v2").
			Methods("POST").
			Path("/v2/data/refresh").
			HandlerFunc(c.refresh)
	}

	return router
}

func (c *controller) status(w http.ResponseWriter, r *http.Request) {
	_, span := telemetry.StartSpan(r.Context(), "api-data-refresh-status")
	defer span.End()

	api.JsonResponse(w, c.manager.Status())
}

func (c *controller) refresh(w http.ResponseWriter, r *http.Request) {
	_, span := telemetry.StartSpan(r.Context(), "api-data-refresh")
	defer span.End()

	// The refresh runs in the background; clients poll GET /v2/data/refresh for status.
	if started := c.manager.TriggerAsync(TriggerManual); !started {
		writeStatusResponse(w, http.StatusConflict, c.manager.Status())
		return
	}
	writeStatusResponse(w, http.StatusAccepted, c.manager.Status())
}

// writeStatusResponse writes a Status as JSON with an explicit status code.
// api.JsonResponse always responds with 200 OK, so refresh outcomes that need a
// different status code (202, 409, 500) are written here.
func writeStatusResponse(w http.ResponseWriter, code int, status Status) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(status)
}
