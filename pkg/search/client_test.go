package search_test

import (
	"context"
	"math"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/internal/search"
	public "github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type testSetup struct {
	logger     log.Logger
	service    search.Service
	router     *mux.Router
	controller search.Controller
	client     public.Client
}

func testAPI(tb testing.TB) testSetup {
	tb.Helper()

	logger := log.NewTestLogger()

	searchConfig := search.DefaultConfig()
	service, err := search.NewService(logger, searchConfig)
	require.NoError(tb, err)

	dl := ofactest.GetDownloader(tb)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(tb, err)

	service.UpdateEntities(stats)

	controller := search.NewController(logger, service, nil)

	router := mux.NewRouter()
	controller.AppendRoutes(router)

	server := httptest.NewServer(router)
	tb.Cleanup(func() {
		server.Close()
	})

	client := public.NewClient(nil, server.URL)

	return testSetup{
		logger:     logger,
		service:    service,
		router:     router,
		controller: controller,
		client:     client,
	}
}

func TestClient_SearchByEntity(t *testing.T) {
	scope := testAPI(t)

	t.Run("normal", func(t *testing.T) {
		ctx := context.Background()
		query := public.Entity[public.Value]{
			Name: "Flight",
			Type: public.EntityAircraft,
		}
		var opts public.SearchOpts

		response, err := scope.client.SearchByEntity(ctx, query, opts)
		require.NoError(t, err)
		require.NotEmpty(t, response.Entities)

		require.Empty(t, response.Entities[0].Details)
		require.InDelta(t, response.Entities[0].Details.FinalScore, 0.00, 0.001)
	})

	t.Run("debug", func(t *testing.T) {
		ctx := context.Background()
		query := public.Entity[public.Value]{
			Name: "Flight",
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

	t.Run("missing type", func(t *testing.T) {
		ctx := context.Background()
		query := public.Entity[public.Value]{
			Name: "John",
		}
		var opts public.SearchOpts

		response, err := scope.client.SearchByEntity(ctx, query, opts)
		require.ErrorContains(t, err, "missing type")
		require.Empty(t, response.Entities)
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
