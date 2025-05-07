package ingest_test

import (
	"context"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/internal/search"
	pubsearch "github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestIngest_API(t *testing.T) {
	logger := log.NewTestLogger()

	searchConfig := search.DefaultConfig()
	searchService, err := search.NewService(logger, searchConfig)
	require.NoError(t, err)

	dl := ofactest.GetDownloader(t)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)

	searchService.UpdateEntities(stats)

	t.Setenv("APP_CONFIG_SECRETS", filepath.Join("testdata", "fincen-config.yml"))
	ingestConf, err := config.LoadConfig(logger)
	require.NoError(t, err)

	ingestService := ingest.NewService(logger, ingestConf.Ingest)
	controller := ingest.NewController(logger, ingestService, searchService)

	router := mux.NewRouter()
	controller.AppendRoutes(router)

	svc := httptest.NewServer(router)
	t.Cleanup(func() { svc.Close() })

	// Setup our client
	client := pubsearch.NewClient(nil, svc.URL)

	file, err := os.Open(filepath.Join("testdata", "fincen-person.csv"))
	require.NoError(t, err)
	t.Cleanup(func() { file.Close() })

	ctx := context.Background()
	response, err := client.IngestFile(ctx, "fincen-person", file, pubsearch.SearchOpts{
		Limit: 1,
	})
	require.NoError(t, err)

	require.Equal(t, "fincen-person", response.FileType)
	require.Len(t, response.Records, 3)

	// The number of records should match the incoming CSV
	for _, rec := range response.Records {
		// Make sure the query is passed through
		require.NotEmpty(t, rec.Query.Name)

		// Make sure ?limit was enforced
		require.Len(t, rec.Entities, 1)

		// Make sure scoring was actaully performed
		for _, entity := range rec.Entities {
			require.Greater(t, entity.Match, 0.00)
		}
	}
}
