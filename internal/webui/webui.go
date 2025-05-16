package webui

import (
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/cmd/ui/wasm"

	"github.com/gorilla/mux"
)

type Controller interface {
	AppendRoutes(router *mux.Router) *mux.Router
}

func NewController(logger log.Logger, config Config) Controller {
	return &controller{
		logger: logger,
		config: config,
	}
}

type controller struct {
	logger log.Logger
	config Config
}

func (c *controller) AppendRoutes(router *mux.Router) *mux.Router {
	staticFS := http.FileServer(http.FS(wasm.WebRoot))
	router.PathPrefix(c.config.BasePath).Handler(http.StripPrefix(c.config.BasePath, staticFS))

	return router
}
