package mcp

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/moov-io/base/log"
	watchman "github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/search"
	mcps "github.com/razashariff/mcps-go"
)

type Server struct {
	logger        log.Logger
	service       search.Service
	keyPair       *mcps.KeyPair
	signing       bool
	agentPassGate *agentPassGate
}

func NewServer(logger log.Logger, service search.Service, config config.MCPConfig) (*Server, error) {
	s := &Server{
		logger:  logger,
		service: service,
		signing: config.Signing.Enabled,
	}

	if config.Signing.Enabled {
		kp, err := loadOrGenerateKeys(logger, config.Signing)
		if err != nil {
			return nil, logger.Error().LogErrorf("MCPS: failed to initialise signing keys: %v", err).Err()
		} else {
			s.keyPair = kp
			logger.Info().Log("MCPS: message signing enabled")
		}
	}

	gate, err := initAgentPass(logger, config.AgentPass)
	if err != nil {
		return nil, err
	}
	s.agentPassGate = gate

	return s, nil
}

func loadOrGenerateKeys(logger log.Logger, conf config.MCPSigning) (*mcps.KeyPair, error) {
	var envKeyPaths int
	if os.Getenv("MCPS_PUBLIC_KEY") != "" {
		envKeyPaths++
	}
	if os.Getenv("MCPS_PRIVATE_KEY") != "" {
		envKeyPaths++
	}
	if envKeyPaths == 1 {
		return nil, errors.New("MCPS: both env vars MCPS_PRIVATE_KEY and MCPS_PUBLIC_KEY are required")
	}
	if envKeyPaths == 2 {
		logger.Info().Log("MCPS: loading signing keys from environment variables")
		return mcps.LoadKeyPairFromEnv("MCPS_PRIVATE_KEY", "MCPS_PUBLIC_KEY")
	}

	keyPath := conf.KeyPath
	pubPath := conf.PubPath

	// Defaults -- avoid relative paths in production by resolving against a predictable base
	if keyPath == "" {
		keyPath = defaultKeyPath("watchman-mcps.key")
	}
	if pubPath == "" {
		pubPath = defaultKeyPath("watchman-mcps.pub")
	}

	// If a relative path was supplied (e.g. via config), resolve it to an absolute path
	// so key storage is not tied to the process working directory.
	absKey, err := filepath.Abs(keyPath)
	if err != nil {
		return nil, fmt.Errorf("MCPS: cannot resolve absolute key path %q: %w", keyPath, err)
	}
	absPub, err := filepath.Abs(pubPath)
	if err != nil {
		return nil, fmt.Errorf("MCPS: cannot resolve absolute public key path %q: %w", pubPath, err)
	}
	if absKey != keyPath {
		logger.Info().Logf("MCPS: resolved relative key path to %s", absKey)
	}
	keyPath = absKey
	pubPath = absPub

	// Try loading existing keys
	if _, err := os.Stat(keyPath); err == nil {
		logger.Info().Logf("MCPS: loading signing keys from %s", keyPath)
		return mcps.LoadKeyPair(keyPath, pubPath)
	}

	// Generate new keys
	logger.Info().Logf("MCPS: generating new signing keys at %s", keyPath)
	return mcps.GenerateAndSaveKeyPair(keyPath, pubPath)
}

// defaultKeyPath returns a sensible default absolute location for a signing key file.
// Precedence: $MCPS_KEY_DIR, then $XDG_CONFIG_HOME/watchman, then $HOME/.watchman, else /etc/watchman.
func defaultKeyPath(filename string) string {
	if dir := os.Getenv("MCPS_KEY_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		return filepath.Join(dir, "watchman", filename)
	}
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		return filepath.Join(home, ".watchman", filename)
	}
	return filepath.Join("/etc/watchman", filename)
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

	// Wrap with AgentPass verification middleware if enabled.
	// When active, every MCP request must carry a valid agent
	// certificate in the X-AgentPass-Certificate header or it
	// receives a 401 before the entity screen runs.
	if s.agentPassGate != nil {
		return s.agentPassGate.middleware(handler)
	}
	return handler
}
