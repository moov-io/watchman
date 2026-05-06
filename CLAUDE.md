# CLAUDE.md

Project-specific guidance for Claude Code (and humans skimming for context). This file is loaded into every session — keep it short, concrete, and accurate. Do not duplicate `README.md` or `CONTRIBUTING.md`; reference them instead.

## Project at a glance

`github.com/moov-io/watchman` — high-performance sanctions / watchlist screening service. Three surfaces, one engine:

- **HTTP API** (`cmd/server`, default `:8084`, admin `:9094`) — primary product
- **Go library** (`pkg/search.Client`) — embeddable in other Go services
- **MCP server** (`/mcp` path prefix, optional) — secure agent access via MCPS signing + AgentPass

Sanctions lists (US OFAC, US CSL, US Non-SDN, FinCEN 311, EU CSL, UK OFSI, UN, OpenSanctions/Senzing) are downloaded, normalized, and held **in memory**. The optional database is for *ingested custom records* and the geocoding L2 cache only — it does not back search.

## Repo layout

| Path | Purpose |
| --- | --- |
| `cmd/server/` | Main HTTP API binary. `main.go` wires config → search → download → MCP; `download.go` schedules refreshes. |
| `cmd/postal-server/` | Standalone libpostal address-parser daemon (used by `internal/postalpool`). |
| `cmd/ui/` | Fyne-based desktop / WASM UI. Most generated assets are gitignored. |
| `pkg/search/` | **Public API.** `Client`, request/response models. External consumers depend on this — treat changes as semver-significant. |
| `pkg/sources/` | Per-list parsers: `ofac/`, `csl_us/`, `csl_eu/`, `csl_uk/`, `csl_un/`, `fincen_311/`, `us_non_sdn/`, `opensanctions/`, `senzing/`, `display/`. |
| `pkg/{address,usaddress,download}/` | Public helpers used by callers. |
| `internal/search/` | Service layer (`api_search.go`, `service.go`). Wraps `pkg/search` with concurrency, embeddings, cross-script logic. |
| `internal/mcp/` | MCP server: `server.go` (boot + key mgmt), `mcp_search_entities.go` (tool), `agentpass.go` (auth middleware). |
| `internal/{prepare,index,indices,norm,stringscore,tfidf}/` | Data pipeline → in-memory indexing → matching/scoring. |
| `internal/{download,ingest,db,api,webui,config}/` | Download orchestration, custom ingestion, SQL layer, HTTP handlers, UI handlers, config loader. |
| `internal/{embeddings,geocoding,postalpool}/` | Optional features (all opt-in via config). |
| `internal/concurrencychamp/` | `ConcurrencyManager` used by search to bound goroutines. |
| `internal/{entitytest,ofactest,cslustest,model_validation}/` | Test helpers and fixtures. |
| `migrations/` | Dual MySQL + Postgres SQL files; embedded into the binary via `package.go`. |
| `configs/config.default.yml` | Embedded default config. **Do not put secrets here.** |
| `docs/` | Markdown user docs (intro, search, mcp, config, pipeline, performance, …). |
| `build/` | `Dockerfile`, `Dockerfile.openshift`, `Dockerfile.static`. |
| `test/` | Integration test helpers / fixtures. |

## Build, run, test

All commands live in `makefile` — prefer them over re-deriving `go build` invocations.

```
make run              # go run ./cmd/server
make build            # builds server + postal-server + webui (./bin/)
make build-server     # server only — fastest iteration loop
make setup-webui      # one-time: go install fyne.io/tools/cmd/fyne@latest (required for build-webui)
make check            # canonical lint + test gate (matches CI). Coverage floor 50%.
                      # Linux/macOS only — fetches lint-project.sh from moov-io/infra.
                      # On Windows falls back to: go test ./... -short
make setup            # docker compose up -d  (Postgres + MySQL for local dev / tests)
make teardown         # docker compose down
make install          # build & install libpostal (only needed when touching address parsing)
make docker           # docker-hub + docker-openshift images
```

Prereqs not auto-installed by `make`: Go (≥1.25), `wget`, `docker` with the **Compose v2 plugin** (`docker compose`, not legacy `docker-compose`), and `fyne` (only for `build-webui`, install via `make setup-webui`).

Quick iteration: `go test ./internal/search/...` is fine for tight loops, but **`make check` is the merge gate** — run it before declaring work done.

If you only want to start the server with no lists (fast boot), set `INCLUDED_LISTS=` empty (the default). To exercise real lookups locally, e.g. `INCLUDED_LISTS=us_ofac` or use the `moov/watchman:v2-static` Docker image (frozen 2019 data).

## Configuration model

- File: YAML at `configs/config.default.yml` is **embedded** into the binary (`package.go`); root key is `Watchman:`. Override on disk by passing a config file or with env vars.
- Env-var conventions used today: `INCLUDED_LISTS`, `OPENSANCTIONS_API_KEY`, `EMBEDDINGS_API_KEY`, `GEOCODING_API_KEY`, `POSTAL_SERVER_BIN_PATH`, `MCPS_PRIVATE_KEY`, `MCPS_PUBLIC_KEY`, `MCPS_KEY_DIR`.
- Opt-in features (all default off): `Database`, `Embeddings`, `Geocoding`, `PostalPool`, `MCP`. Don't change defaults — additions belong behind their own flag.

## Architecture cheatsheet

- **Search hot path**: `cmd/server/main.go` → `internal/search/service.go` (`concurrencychamp.ConcurrencyManager` bounds goroutines per `Search.Goroutines`) → `pkg/search` matching (Jaro-Winkler via `xrash/smetrics`) → response. Don't spawn unbounded goroutines; reuse the manager.
- **Download/refresh**: `cmd/server/download.go` schedules; `internal/download/` orchestrates; per-list parsing in `pkg/sources/<list>/`. Refreshes at `Download.RefreshInterval` (12h default).
- **Pipeline**: parsers produce raw entities → `internal/prepare/` (name reorder, suffix cleanup, stopwords, UTF-8 norm) → `internal/index/` (in-memory) → optional embeddings index in `internal/embeddings/`.
- **MCP**: mounted in `cmd/server/main.go:120` as `router.PathPrefix("/mcp")`. `internal/mcp/server.go` boots the SDK server; `mcp_search_entities.go` exposes the `search_entities` tool; `agentpass.go` is the verification middleware. Most-recent contract change: `agentContext` field is stamped into every `search_entities` response (commit `6ed132b`) — preserve it.

## Conventions

- **Logging**: use `github.com/moov-io/base/log` (already pervasive). Don't introduce stdlib `log` or `log/slog`.
- **Telemetry**: OpenTelemetry via `go.opentelemetry.io/otel`; spans/metrics flow through `moov-io/base`. Don't wire up a second tracer.
- **Tests**: `stretchr/testify` only. **No mocking framework** is in use — write hand-rolled fakes, or run real Postgres/MySQL via `make setup`. Don't add `sqlmock`, `gomock`, etc. without discussion.
- **Errors**: wrap with `fmt.Errorf("...: %w", err)`. Style follows Go Code Review Comments and Go Proverbs (per `CONTRIBUTING.md`).
- **Public API stability**: anything exported under `pkg/` is consumed externally (`pkg.go.dev/github.com/moov-io/watchman/pkg/search#Client`). Breaking changes require version bump + changelog.
- **Migrations**: every schema change ships **both** dialects under `migrations/NNN_<name>.up.{mysql,postgres}.sql` (and the matching `.down.` files). They are embedded via `package.go`.
- **Concurrency in search**: respect `Search.Goroutines` (Default/Min/Max) — `internal/concurrencychamp` adapts dynamically; bypassing it skews benchmarks and pollutes prod tuning.

## MCP specifics (recent area, changes often)

- Endpoint: any path under `/mcp` (PathPrefix mount), `POST` for tool calls.
- **MCPS keys**: precedence is `MCPS_PRIVATE_KEY`/`MCPS_PUBLIC_KEY` env vars → `MCPS_KEY_DIR` → `$XDG_CONFIG_HOME/watchman` → `$HOME/.watchman` → `/etc/watchman` (see `internal/mcp/server.go`). Keys are auto-generated on first run if absent. **Never commit generated keys.**
- **AgentPass**: `internal/mcp/agentpass.go` enforces `MinTrustLevel` and `RequiredScopes` from `MCP.AgentPass` config; rejects requests with 401.
- **Response contract**: keep the `agentContext` field on `search_entities` responses populated when AgentPass is enabled — downstream demos rely on it.

## Don'ts

- Don't bypass `make check` (`--no-verify`, `-skip` flags, etc.). Same gate as CI.
- Don't put secrets / API keys in `configs/config.default.yml`. It's embedded into every binary.
- Don't commit `embeddings.test` or other `*.test` binaries (already in `.gitignore`, but easy to force-add).
- Don't add features that change defaults silently (network calls, new ports, new disk writes). Make them opt-in.
- Don't duplicate sanctions data in tests — use fixtures under `internal/{ofactest,cslustest,entitytest}/`.
- Don't replace fuzzy matching or normalization without benchmarks. There are bench results committed under `docs/bench/`.

## Reference docs

- `README.md` — what the project is, Docker quickstart, list of integrated lists.
- `CONTRIBUTING.md` — fork/PR flow, style guides.
- `docs/intro.md`, `docs/search.md`, `docs/pipeline.md`, `docs/config.md`, `docs/mcp.md`, `docs/performance.md`, `docs/cross-script-matching.md`.
- Live API spec: `docs/api.yaml`.
