package mcp

import (
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/moov-io/base/log"
	watchman "github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/search"
)

type Server struct {
	logger  log.Logger
	service search.Service
}

func NewServer(logger log.Logger, service search.Service) *Server {
	return &Server{
		logger:  logger,
		service: service,
	}
}

func (s *Server) Handler() http.Handler {
	impl := &mcp.Implementation{
		Name:    "watchman-mcp",
		Version: watchman.Version,
	}

	server := mcp.NewServer(impl, nil)
	s.logger.Info().Log("starting MCP server over HTTP")

	// Add the search_entities tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_entities",
		Description: "Search for entities in sanctions lists, same as /v2/search endpoint",
	}, s.HandleSearchEntities)

	// Create the streamable HTTP handler with stateless mode for simplicity
	opts := &mcp.StreamableHTTPOptions{
		Stateless:    true,
		JSONResponse: true, // Use JSON responses instead of SSE for easier testing
	}
	return mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, opts)
}
