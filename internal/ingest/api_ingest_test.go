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
	ingestResponse, err := client.IngestFile(ctx, "fincen-person", file)
	require.NoError(t, err)

	require.Equal(t, "fincen-person", ingestResponse.FileType)
	require.Len(t, ingestResponse.Entities, 3)

	// Perform a search against the ingested file
	query := pubsearch.Entity[pubsearch.Value]{
		Name:   "John K Doe1",
		Type:   pubsearch.EntityPerson,
		Source: pubsearch.SourceList("fincen-person"),
		Addresses: []pubsearch.Address{
			{
				Line1:      "193 Southfield Lane",
				City:       "Anytown",
				PostalCode: "90210",
				State:      "CA",
				Country:    "US",
			},
		},
	}
	searchResponse, err := searchService.Search(ctx, query.Normalize(), search.SearchOpts{
		Limit: 1,
		Debug: true,
	})
	require.NoError(t, err)
	require.Len(t, searchResponse, 1)

	// Sanity check the response
	require.Equal(t, "John Jr K Doe1", searchResponse[0].Name)
	require.Equal(t, pubsearch.SourceList("fincen-person"), searchResponse[0].Source)
	require.InDelta(t, searchResponse[0].Match, 0.883, 0.001)
}
