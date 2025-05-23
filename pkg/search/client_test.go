package search_test

import (
	"context"
	"math"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/internal/search"
	public "github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type testSetup struct {
	logger log.Logger

	searchService search.Service
	ingestService ingest.Service

	router *mux.Router

	client public.Client
}

func testAPI(tb testing.TB) testSetup {
	tb.Helper()

	logger := log.NewTestLogger()

	searchConfig := search.DefaultConfig()
	searchService, err := search.NewService(logger, searchConfig)
	require.NoError(tb, err)

	dl := ofactest.GetDownloader(tb)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(tb, err)

	searchService.UpdateEntities(stats)

	conf, err := config.LoadConfig(logger)
	require.NoError(tb, err)

	ingestService := ingest.NewService(logger, conf.Ingest)

	searchController := search.NewController(logger, searchService, nil)
	ingestController := ingest.NewController(logger, ingestService, searchService)

	router := mux.NewRouter()
	searchController.AppendRoutes(router)
	ingestController.AppendRoutes(router)

	server := httptest.NewServer(router)
	tb.Cleanup(func() {
		server.Close()
	})

	client := public.NewClient(nil, server.URL)

	return testSetup{
		logger:        logger,
		searchService: searchService,
		ingestService: ingestService,
		router:        router,
		client:        client,
	}
}

func TestClient_SearchByEntity(t *testing.T) {
	scope := testAPI(t)

	t.Run("normal", func(t *testing.T) {
		ctx := context.Background()
		query := public.Entity[public.Value]{
			Name: "P-532",
			Type: public.EntityAircraft,
		}
		opts := public.SearchOpts{
			Limit: 2,
		}

		response, err := scope.client.SearchByEntity(ctx, query, opts)
		require.NoError(t, err)
		require.NotEmpty(t, response.Entities)

		require.Empty(t, response.Entities[0].Details)
		require.InDelta(t, response.Entities[0].Details.FinalScore, 0.00, 0.001)
	})

	t.Run("debug", func(t *testing.T) {
		ctx := context.Background()
		query := public.Entity[public.Value]{
			Name: "P-532",
			Type: public.EntityAircraft,
		}
		opts := public.SearchOpts{
			Limit: 10,
			Debug: true,
		}

		response, err := scope.client.SearchByEntity(ctx, query, opts)
		require.NoError(t, err)
		require.NotEmpty(t, response.Entities)

		require.NotEmpty(t, response.Entities[0].Details)
		require.Greater(t, response.Entities[0].Details.FinalScore, 0.01)
	})

	t.Run("error", func(t *testing.T) {
		ctx := context.Background()
		query := public.Entity[public.Value]{
			Name: "Sea",
			Type: public.EntityVessel,
			Vessel: &public.Vessel{
				Tonnage: int(math.MaxInt),
			},
		}
		var opts public.SearchOpts

		response, err := scope.client.SearchByEntity(ctx, query, opts)
		require.ErrorContains(t, err, "reading vessel tonnage: strconv.ParseInt:")
		require.Empty(t, response.Entities)
	})
}

func TestClient_IngestFile(t *testing.T) {
	// Load our ingest config
	t.Setenv("APP_CONFIG_SECRETS", filepath.Join("..", "..", "internal", "ingest", "testdata", "fincen-config.yml"))

	scope := testAPI(t)

	t.Run("fincen-person", func(t *testing.T) {
		ctx := context.Background()

		fd, err := os.Open(filepath.Join("..", "..", "internal", "ingest", "testdata", "fincen-person.csv"))
		require.NoError(t, err)

		response, err := scope.client.IngestFile(ctx, "fincen-person", fd)
		require.NoError(t, err)

		require.Equal(t, "fincen-person", response.FileType)
		require.Len(t, response.Entities, 3)
	})
}
