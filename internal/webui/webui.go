package webui

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"

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
	handler := newStaticHandler(wasm.WebRoot, c.logger)
	handler.warmWASM()
	router.PathPrefix(c.config.BasePath).Handler(http.StripPrefix(c.config.BasePath, handler))

	return router
}

// staticHandler serves the Fyne web UI. The wasm binary is tens of megabytes
// uncompressed; browsers send Accept-Encoding: gzip, so the compressed body is
// built once and reused.
type staticHandler struct {
	fsys   fs.FS
	files  http.Handler
	logger log.Logger

	caches sync.Map // filename -> *wasmCache
}

type wasmCache struct {
	once sync.Once
	body []byte
	err  error
}

func newStaticHandler(fsys fs.FS, logger log.Logger) *staticHandler {
	return &staticHandler{
		fsys:   fsys,
		files:  http.FileServer(http.FS(fsys)),
		logger: logger,
	}
}

func (h *staticHandler) warmWASM() {
	entries, err := fs.ReadDir(h.fsys, ".")
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() || !isWASM(entry.Name()) {
			continue
		}
		name := entry.Name()
		go h.compressed(name)
	}
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=3600")

	name := path.Base(r.URL.Path)
	if isWASM(name) {
		w.Header().Set("Vary", "Accept-Encoding")
		if acceptsGzip(r) {
			body, err := h.compressed(name)
			if err == nil {
				h.serveGzipWASM(w, r, body)
				return
			}
		}
	}

	h.files.ServeHTTP(w, r)
}

func (h *staticHandler) serveGzipWASM(w http.ResponseWriter, r *http.Request, body []byte) {
	w.Header().Set("Content-Type", "application/wasm")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}
	if _, err := w.Write(body); err != nil && h.logger != nil {
		h.logger.Debug().Logf("writing wasm response: %v", err)
	}
}

func (h *staticHandler) compressed(name string) ([]byte, error) {
	if !isWASM(name) {
		return nil, fs.ErrNotExist
	}
	loaded, _ := h.caches.LoadOrStore(name, &wasmCache{})
	cache, ok := loaded.(*wasmCache)
	if !ok {
		return nil, fmt.Errorf("cached wasm %s has unexpected type %T", name, loaded)
	}
	cache.once.Do(func() {
		cache.body, cache.err = gzipFile(h.fsys, name)
		if cache.err != nil || h.logger == nil {
			return
		}
		info, err := fs.Stat(h.fsys, name)
		if err != nil {
			return
		}
		// Omit the request path so this log is not a log-injection source (CWE-117).
		h.logger.Info().Logf("compressed wasm from %d to %d bytes", info.Size(), len(cache.body))
	})
	return cache.body, cache.err
}

func gzipFile(fsys fs.FS, name string) ([]byte, error) {
	raw, err := fs.ReadFile(fsys, name)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	zw, err := gzip.NewWriterLevel(&buf, gzip.BestSpeed)
	if err != nil {
		return nil, err
	}
	if _, err := zw.Write(raw); err != nil {
		_ = zw.Close()
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func isWASM(name string) bool {
	return strings.EqualFold(path.Ext(name), ".wasm")
}

func acceptsGzip(r *http.Request) bool {
	for _, part := range strings.Split(r.Header.Get("Accept-Encoding"), ",") {
		part = strings.TrimSpace(part)
		encoding, _, _ := strings.Cut(part, ";")
		if strings.EqualFold(strings.TrimSpace(encoding), "gzip") {
			return true
		}
	}
	return false
}
