package mcp

import (
	"net/http"
	"os"

	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/moov-io/base/log"
	watchman "github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/search"
	mcps "github.com/razashariff/mcps-go"
)

type Server struct {
	logger  log.Logger
	service search.Service
	keyPair *mcps.KeyPair
	signing bool
}

func NewServer(logger log.Logger, service search.Service, signingConf config.MCPSigning) *Server {
	s := &Server{
		logger:  logger,
		service: service,
		signing: signingConf.Enabled,
	}

	if signingConf.Enabled {
		kp, err := loadOrGenerateKeys(logger, signingConf)
		if err != nil {
			logger.Error().LogErrorf("MCPS: failed to initialise signing keys: %v", err)
			s.signing = false
		} else {
			s.keyPair = kp
			logger.Info().Log("MCPS: message signing enabled")
		}
	}

	return s
}

func loadOrGenerateKeys(logger log.Logger, conf config.MCPSigning) (*mcps.KeyPair, error) {
	// Try environment variables first
	if privEnv := os.Getenv("MCPS_PRIVATE_KEY"); privEnv != "" {
		logger.Info().Log("MCPS: loading signing keys from environment variables")
		return mcps.LoadKeyPairFromEnv("MCPS_PRIVATE_KEY", "MCPS_PUBLIC_KEY")
	}

	keyPath := conf.KeyPath
	pubPath := conf.PubPath

	// Defaults
	if keyPath == "" {
		keyPath = "watchman-mcps.key"
	}
	if pubPath == "" {
		pubPath = "watchman-mcps.pub"
	}

	// Try loading existing keys
	if _, err := os.Stat(keyPath); err == nil {
		logger.Info().Logf("MCPS: loading signing keys from %s", keyPath)
		return mcps.LoadKeyPair(keyPath, pubPath)
	}

	// Generate new keys
	logger.Info().Logf("MCPS: generating new signing keys at %s", keyPath)
	return mcps.GenerateAndSaveKeyPair(keyPath, pubPath)
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
	return mcpsdk.NewStreamableHTTPHandler(func(req *http.Request) *mcpsdk.Server {
		return server
	}, opts)
}
