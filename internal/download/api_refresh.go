package download

import (
	"encoding/json"
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/api"

	"github.com/gorilla/mux"
)

type RefreshController interface {
	AppendRoutes(router *mux.Router) *mux.Router
}

func NewRefreshController(logger log.Logger, r *Refresher) RefreshController {
	return &refreshController{
		logger:    logger,
		refresher: r,
	}
}

type refreshController struct {
	logger    log.Logger
	refresher *Refresher
}

func (c *refreshController) AppendRoutes(router *mux.Router) *mux.Router {
	router.
		Name("DataRefreshStatus.v2").
		Methods("GET").
		Path("/v2/data/refresh").
		HandlerFunc(c.status)

	router.
		Name("DataRefresh.v2").
		Methods("POST").
		Path("/v2/data/refresh").
		HandlerFunc(c.refresh)

	return router
}

func (c *refreshController) status(w http.ResponseWriter, r *http.Request) {
	_, span := telemetry.StartSpan(r.Context(), "api-data-refresh-status")
	defer span.End()
	api.JsonResponse(w, c.refresher.Status())
}

func (c *refreshController) refresh(w http.ResponseWriter, r *http.Request) {
	_, span := telemetry.StartSpan(r.Context(), "api-data-refresh")
	defer span.End()

	if started := c.refresher.TriggerAsync(TriggerManual); !started {
		writeStatusResponse(w, http.StatusConflict, c.refresher.Status())
		return
	}
	writeStatusResponse(w, http.StatusAccepted, c.refresher.Status())
}

func writeStatusResponse(w http.ResponseWriter, code int, status Status) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(status)
}
