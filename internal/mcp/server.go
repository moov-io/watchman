package mcp

import (
	"net/http"

	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/moov-io/base/log"
	watchman "github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/search"
)

type Server struct {
	logger  log.Logger
	service search.Service
}

func NewServer(logger log.Logger, service search.Service, config config.MCPConfig) (*Server, error) {
	return &Server{
		logger:  logger,
		service: service,
	}, nil
}

func (s *Server) Handler() http.Handler {
	impl := &mcpsdk.Implementation{
		Name:    "watchman-mcp",
		Version: watchman.Version,
	}

	server := mcpsdk.NewServer(impl, nil)
	s.logger.Info().Log("starting MCP server over HTTP")

	// Add the search_entities tool
	mcpsdk.AddTool(server, &mcpsdk.Tool{
		Name:        "search_entities",
		Description: "Search for entities in sanctions lists, same as /v2/search endpoint",
	}, s.HandleSearchEntities)

	// Create the streamable HTTP handler with stateless mode for simplicity
	opts := &mcpsdk.StreamableHTTPOptions{
		Stateless:    true,
		JSONResponse: true, // Use JSON responses instead of SSE for easier testing
	}
	handler := mcpsdk.NewStreamableHTTPHandler(func(req *http.Request) *mcpsdk.Server {
		return server
	}, opts)

	return handler
}
