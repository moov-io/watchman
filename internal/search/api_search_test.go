package search

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/api"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestAPI_ListInfo(t *testing.T) {
	env := testAPI(t)

	req := httptest.NewRequest("GET", "/v2/listinfo", nil)

	w := httptest.NewRecorder()
	env.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body := w.Body.String()
	require.Contains(t, body, `"lists":{"us_ofac":17}`)
	require.Contains(t, body, `"listHashes":{"us_ofac":"b9d56301`)
}

type testSetup struct {
	logger     log.Logger
	service    Service
	router     *mux.Router
	controller Controller
}

func testAPI(tb testing.TB) testSetup {
	tb.Helper()

	logger := log.NewTestLogger()

	searchConfig := DefaultConfig()
	service, err := NewService(logger, searchConfig)
	require.NoError(tb, err)

	dl := ofactest.GetDownloader(tb)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(tb, err)

	service.UpdateEntities(stats)

	controller := NewController(logger, service, nil)

	router := mux.NewRouter()
	controller.AppendRoutes(router)

	return testSetup{
		logger:     logger,
		service:    service,
		router:     router,
		controller: controller,
	}
}

func TestAPI_readSearchRequest(t *testing.T) {
	ctx := context.Background()

	t.Run("basic", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=adam&type=person&birthDate=2025-01-02", nil)
		q := &api.QueryParams{Values: req.URL.Query()}

		query, err := readSearchRequest(ctx, nil, q)
		require.NoError(t, err)

		require.Equal(t, "adam", query.Name)
		require.Equal(t, search.EntityPerson, query.Type)

		require.NotNil(t, query.Person)
		require.Equal(t, "2025-01-02T00:00:00Z", query.Person.BirthDate.Format(time.RFC3339))

	})

	t.Run("contact info", func(t *testing.T) {
		address := "/v2/search?type=business&emailAddress=a@corp.com&phone=1234567890"
		address += "&faxNumber=3334445566&email=b@corp.com&phone=9876543210"
		address += "&website=corp.com&website=corp2.com"

		req := httptest.NewRequest("GET", address, nil)
		q := &api.QueryParams{Values: req.URL.Query()}

		query, err := readSearchRequest(ctx, nil, q)
		require.NoError(t, err)
		require.Empty(t, query.Name)

		expected := search.ContactInfo{
			EmailAddresses: []string{"b@corp.com", "a@corp.com"},
			PhoneNumbers:   []string{"1234567890", "9876543210"},
			FaxNumbers:     []string{"3334445566"},
			Websites:       []string{"corp.com", "corp2.com"},
		}
		require.ElementsMatch(t, expected.EmailAddresses, query.Contact.EmailAddresses)
		require.ElementsMatch(t, expected.PhoneNumbers, query.Contact.PhoneNumbers)
		require.ElementsMatch(t, expected.FaxNumbers, query.Contact.FaxNumbers)
		require.ElementsMatch(t, expected.Websites, query.Contact.Websites)
	})

	t.Run("crypto addresses", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?type=person&cryptoAddress=xbt:12345&cryptoAddress=eth:54321", nil)
		q := &api.QueryParams{Values: req.URL.Query()}

		query, err := readSearchRequest(ctx, nil, q)
		require.NoError(t, err)
		require.Empty(t, query.Name)

		require.Len(t, query.CryptoAddresses, 2)

		expected := []search.CryptoAddress{
			{Currency: "XBT", Address: "12345"},
			{Currency: "ETH", Address: "54321"},
		}
		require.ElementsMatch(t, expected, query.CryptoAddresses)
	})

	t.Run("address", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?type=person&name=Jane&address=123+Acme+St+Acmetown+KY+54321+US", nil)
		q := &api.QueryParams{Values: req.URL.Query()}

		query, err := readSearchRequest(ctx, nil, q)
		require.NoError(t, err)

		require.Equal(t, "Jane", query.Name)
		require.Equal(t, search.EntityPerson, query.Type)

		expected := search.Address{
			Line1:      "123 ACME ST",
			City:       "ACMETOWN",
			PostalCode: "54321",
			State:      "KY",
			Country:    "United States",
		}
		require.Len(t, query.Addresses, 1)
		require.Equal(t, expected, query.Addresses[0])
	})
}

func TestAPI_Search(t *testing.T) {
	env := testAPI(t)

	t.Run("normal", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		t.Log(w.Body.String())

		require.Equal(t, http.StatusOK, w.Code)

		var response search.SearchResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		require.Len(t, response.Entities, 2)

		require.NotEmpty(t, response.Entities[0].Name)
		require.NotEmpty(t, response.Entities[1].Name)

		require.Empty(t, response.Entities[0].Debug)
		require.Empty(t, response.Entities[1].Debug)
	})

	t.Run("debug", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&debug=yes", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var response search.SearchResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		require.Len(t, response.Entities, 2)

		require.NotEmpty(t, response.Entities[0].Name)
		require.NotEmpty(t, response.Entities[1].Name)

		require.NotEmpty(t, response.Entities[0].Debug)
		require.NotEmpty(t, response.Entities[1].Debug)

		raw, err := base64.StdEncoding.DecodeString(response.Entities[0].Debug)
		require.NoError(t, err)
		require.NotEmpty(t, raw)

		if testing.Verbose() {
			fmt.Println(string(raw))
		}
	})
}

func BenchmarkAPI_Search(b *testing.B) {
	env := testAPI(b)
	b.ResetTimer()

	b.Run("normal", func(b *testing.B) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=5", nil)

		for b.Loop() {
			w := httptest.NewRecorder()
			env.router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				b.Fatalf("unexpected %v status code", w.Code)
			}
		}
	})

	b.Run("debug", func(b *testing.B) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=5&debug=true", nil)

		for b.Loop() {
			w := httptest.NewRecorder()
			env.router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				b.Fatalf("unexpected %v status code", w.Code)
			}
		}
	})

	b.Run("name address", func(b *testing.B) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=5", nil)

		for b.Loop() {
			w := httptest.NewRecorder()
			env.router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				b.Fatalf("unexpected %v status code", w.Code)
			}
		}
	})

	b.Run("name email", func(b *testing.B) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=5", nil)

		for b.Loop() {
			w := httptest.NewRecorder()
			env.router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				b.Fatalf("unexpected %v status code", w.Code)
			}
		}
	})

	b.Run("name address email", func(b *testing.B) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=5", nil)

		for b.Loop() {
			w := httptest.NewRecorder()
			env.router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				b.Fatalf("unexpected %v status code", w.Code)
			}
		}
	})
}
