package ingest_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/compress"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/internal/search"
	pubsearch "github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/senzing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestIngest_API(t *testing.T) {
	setupIngestAPITest(t, func(scope ingestApiSetup) {
		file, err := os.Open(filepath.Join("testdata", "fincen-person.csv"))
		require.NoError(t, err)
		t.Cleanup(func() { file.Close() })

		ctx := context.Background()
		ingestResponse, err := scope.client.IngestFile(ctx, "fincen-person", file)
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
		searchResponse, err := scope.searchService.Search(ctx, query.Normalize(), search.SearchOpts{
			Limit: 1,
			Debug: true,
		})
		require.NoError(t, err)
		require.Len(t, searchResponse, 1)

		// Sanity check the response
		require.Equal(t, "John Jr K Doe1", searchResponse[0].Name)
		require.Equal(t, pubsearch.SourceList("fincen-person"), searchResponse[0].Source)
		require.InDelta(t, searchResponse[0].Match, 0.839, 0.001)
	})
}

func TestIngest_API_Gzip(t *testing.T) {
	setupIngestAPITest(t, func(scope ingestApiSetup) {
		// Replace linux line endings with windows (\r\n)
		bs, err := os.ReadFile(filepath.Join("testdata", "Person12.09.2025.csv"))
		require.NoError(t, err)

		lines := bytes.Split(bs, []byte("\n"))
		require.Len(t, lines, 4)

		var file bytes.Buffer
		for i := range lines {
			file.Write(lines[i])
			file.WriteString("\r\n")
		}

		// Make ingest API request
		ctx := context.Background()
		body := compress.GzipTestFile(t, &file)

		ingestResponse, err := scope.client.IngestFile(ctx, "fincen-person", body)
		require.NoError(t, err)

		require.Equal(t, "fincen-person", ingestResponse.FileType)
		require.Len(t, ingestResponse.Entities, 2)

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
		searchResponse, err := scope.searchService.Search(ctx, query.Normalize(), search.SearchOpts{
			Limit: 1,
			Debug: true,
		})
		require.NoError(t, err)
		require.Len(t, searchResponse, 1)

		// Sanity check the response
		require.Equal(t, "John Doe", searchResponse[0].Name)
		require.Equal(t, pubsearch.SourceList("fincen-person"), searchResponse[0].Source)
		require.InDelta(t, searchResponse[0].Match, 0.8198, 0.001)
	})
}

func TestAPI_Senzing(t *testing.T) {
	setupIngestAPITest(t, func(scope ingestApiSetup) {
		file, err := os.Open(filepath.Join("testdata", "fincen-person.csv"))
		require.NoError(t, err)
		t.Cleanup(func() { file.Close() })

		ctx := context.Background()
		ingestResponse, err := scope.client.IngestFile(ctx, "fincen-person", file)
		require.NoError(t, err)

		require.Equal(t, "fincen-person", ingestResponse.FileType)
		require.Len(t, ingestResponse.Entities, 3)

		// Export the data
		req, err := http.NewRequest("GET", scope.server.URL+"/v2/export/fincen-person", nil)
		require.NoError(t, err)

		req.Header.Set("Accept", "senzing/jsonl")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		bs, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		// Make sure the response is JSON Lines
		require.True(t, bytes.HasPrefix(bs, []byte("{")))

		entities, err := senzing.ReadEntities(bytes.NewReader(bs), pubsearch.SourceList("senzing"))
		require.NoError(t, err)
		require.Len(t, entities, 3)
	})
}

type ingestApiSetup struct {
	client pubsearch.Client
	server *httptest.Server

	searchService search.Service
}

func setupIngestAPITest(t *testing.T, fn func(ingestApiSetup)) {
	t.Helper()

	db.ForEachDatabase(t, func(db db.DB) {
		logger := log.NewTestLogger()

		ingestRepository := ingest.NewRepository(db)
		indexedLists := index.NewLists(ingestRepository)

		searchConfig := search.DefaultConfig()
		searchService, err := search.NewService(logger, searchConfig, indexedLists)
		require.NoError(t, err)

		t.Setenv("APP_CONFIG_SECRETS", filepath.Join("testdata", "fincen-config.yml"))
		ingestConf, err := config.LoadConfig(logger)
		require.NoError(t, err)

		ingestService := ingest.NewService(logger, ingestConf.Ingest, ingestRepository)
		controller := ingest.NewController(logger, ingestService)

		router := mux.NewRouter()
		controller.AppendRoutes(router)

		server := httptest.NewServer(router)
		t.Cleanup(func() { server.Close() })

		// Setup our client
		client := pubsearch.NewClient(nil, server.URL)

		fn(ingestApiSetup{
			client:        client,
			server:        server,
			searchService: searchService,
		})
	})
}
