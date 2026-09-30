package search

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/api"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/senzing"

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

	indexedLists := index.NewLists(nil) // only in-mem

	searchConfig := DefaultConfig()
	service, err := NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(tb, err)

	dl := ofactest.GetDownloader(tb)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(tb, err)

	indexedLists.Update(stats)

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

type stubAddressParser struct {
	addr search.Address
	err  error
}

func (s stubAddressParser) ParseAddress(ctx context.Context, input string) (search.Address, error) {
	return s.addr, s.err
}

func TestParseSearchDebug(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in          string
		debug       bool
		minMatch    float64
		errContains string
	}{
		{in: ""},
		{in: "yes", debug: true},
		{in: "true", debug: true},
		{in: "TRUE", debug: true},
		{in: "1", debug: true},
		{in: "t", debug: true},
		{in: "false"},
		{in: "0"},
		{in: "no"},
		{in: "off"},
		{in: "0.80", debug: true, minMatch: 0.80},
		{in: "0.8", debug: true, minMatch: 0.8},
		{in: ".8", debug: true, minMatch: 0.8},
		{in: "1.0", debug: true, minMatch: 1.0},
		{in: "0.0", debug: true, minMatch: 0.0},
		{in: " 0.75 ", debug: true, minMatch: 0.75},
		{in: "maybe", errContains: "invalid debug value"},
		{in: "80", errContains: "invalid debug value"},
		{in: "1.5", errContains: "debug threshold must be between"},
		{in: "-0.1", errContains: "debug threshold must be between"},
		{in: "NaN.", errContains: "invalid debug threshold"},
	}
	for _, tc := range cases {
		name := tc.in
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			debug, minMatch, err := parseSearchDebug(tc.in)
			if tc.errContains != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errContains)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.debug, debug)
			require.InDelta(t, tc.minMatch, minMatch, 0.0000001)
		})
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

	t.Run("name without type", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=adam", nil)
		q := &api.QueryParams{Values: req.URL.Query()}

		query, err := readSearchRequest(ctx, nil, q)
		require.NoError(t, err)

		require.Equal(t, "adam", query.Name)
		require.Empty(t, query.Type)
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

	t.Run("address with parser", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?type=person&name=Jane&address=raw-input", nil)
		q := &api.QueryParams{Values: req.URL.Query()}

		parser := stubAddressParser{addr: search.Address{Line1: "350 rue des lilas", City: "quebec city", PostalCode: "g1l 1b6"}}
		query, err := readSearchRequest(ctx, parser, q)
		require.NoError(t, err)
		require.Equal(t, parser.addr, query.Addresses[0])
	})

	t.Run("government id (US Passport)", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?type=person&gov_passport=US:123456789", nil)
		q := &api.QueryParams{Values: req.URL.Query()}

		query, err := readSearchRequest(ctx, nil, q)
		require.NoError(t, err)
		require.NotNil(t, query.Person)

		govIDs := query.Person.GovernmentIDs
		require.Len(t, govIDs, 1)

		expected := search.GovernmentID{
			Type:       search.GovernmentIDPassport,
			Country:    "US",
			Identifier: "123456789",
		}
		require.Equal(t, expected, govIDs[0])
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

	t.Run("debug=1 boolean still attaches all", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&debug=1", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var response search.SearchResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		require.Len(t, response.Entities, 2)
		require.NotEmpty(t, response.Entities[0].Debug)
		require.NotEmpty(t, response.Entities[1].Debug)
	})

	t.Run("debug threshold attaches only qualifying hits", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=5&debug=0.90", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var response search.SearchResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		require.NotEmpty(t, response.Entities)

		for _, ent := range response.Entities {
			if ent.Match >= 0.90 {
				require.NotEmpty(t, ent.Debug, "match=%.4f should include debug", ent.Match)
				require.NotEmpty(t, ent.Details.Pieces)
			} else {
				require.Empty(t, ent.Debug, "match=%.4f should omit debug", ent.Match)
				require.Empty(t, ent.Details.Pieces)
			}
		}
	})

	t.Run("minMatch 0.75 with debug 0.80", func(t *testing.T) {
		searchJSON := func(url string) search.SearchResponse {
			t.Helper()
			req := httptest.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()
			env.router.ServeHTTP(w, req)
			require.Equal(t, http.StatusOK, w.Code)

			var resp search.SearchResponse
			require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
			return resp
		}

		assertDebugByThreshold := func(resp search.SearchResponse) (belowDebug, withPieces int) {
			t.Helper()
			require.NotEmpty(t, resp.Entities)
			for _, ent := range resp.Entities {
				require.GreaterOrEqual(t, ent.Match, 0.75)
				if ent.Match >= 0.80 {
					require.NotEmpty(t, ent.Debug, "match=%.4f should include debug", ent.Match)
					require.NotEmpty(t, ent.Details.Pieces)
					withPieces++
				} else {
					require.Empty(t, ent.Debug, "match=%.4f should omit debug", ent.Match)
					require.Empty(t, ent.Details.Pieces)
					belowDebug++
				}
			}
			return belowDebug, withPieces
		}

		withoutDebug := searchJSON("/v2/search?name=Dmitry+Khoroshev&type=person&limit=10&minMatch=0.75")
		withDebug := searchJSON("/v2/search?name=Dmitry+Khoroshev&type=person&limit=10&minMatch=0.75&debug=0.80")
		require.Len(t, withDebug.Entities, len(withoutDebug.Entities))
		for i, ent := range withDebug.Entities {
			require.Equal(t, withoutDebug.Entities[i].SourceID, ent.SourceID)
			require.InDelta(t, withoutDebug.Entities[i].Match, ent.Match, 0.0001)
			require.Empty(t, withoutDebug.Entities[i].Debug)
		}
		belowDebug, withPieces := assertDebugByThreshold(withDebug)
		require.Greater(t, belowDebug, 0, "name-only Khoroshev should return a hit in [0.75, 0.80)")
		require.Equal(t, 0, withPieces)
	})

	t.Run("debug=1.0 only exact matches", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&debug=1.0", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var response search.SearchResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		require.Len(t, response.Entities, 2)

		for _, ent := range response.Entities {
			if ent.Match >= 1.0 {
				require.NotEmpty(t, ent.Debug)
			} else {
				require.Empty(t, ent.Debug)
			}
		}
	})

	t.Run("debug invalid value", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&debug=maybe", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Contains(t, w.Body.String(), "invalid debug value")
	})

	t.Run("debug threshold out of range", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&debug=1.5", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Contains(t, w.Body.String(), "debug threshold must be between 0.0 and 1.0")
	})

	t.Run("algorithm soundex", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=soundex", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var response search.SearchResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		require.Len(t, response.Entities, 2)
	})

	t.Run("algorithm jaro-winkler", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=jaro-winkler", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm soft-bidist", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=soft-bidist", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm soft-bisim", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=soft-bisim", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm editex", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=editex", nil)
		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm nsim", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=nsim", nil)
		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm double-metaphone", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=double-metaphone", nil)
		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm beider-morse", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=beider-morse", nil)
		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm soft-bigram alias", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&algorithm=soft-bigram", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("algorithm invalid", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&algorithm=levenshtein", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Contains(t, w.Body.String(), "unknown algorithm")
	})
}

func TestAPI_Senzing(t *testing.T) {
	env := testAPI(t)

	t.Run("json/query", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&format=senzing", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		// Make sure the response is a JSON object
		require.True(t, strings.HasPrefix(w.Body.String(), "["))

		entities, err := senzing.ReadEntities(w.Body, search.SourceList("senzing"))
		require.NoError(t, err)
		require.Len(t, entities, 2)
	})

	t.Run("json/header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2", nil)
		req.Header.Set("Accept", "senzing")

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		// Make sure the response is a JSON object
		require.True(t, strings.HasPrefix(w.Body.String(), "["))

		entities, err := senzing.ReadEntities(w.Body, search.SourceList("senzing"))
		require.NoError(t, err)
		require.Len(t, entities, 2)
	})

	t.Run("jsonl/query", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2&format=senzing/jsonl", nil)

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		// Make sure the response is JSON Lines
		require.True(t, strings.HasPrefix(w.Body.String(), "{"))

		entities, err := senzing.ReadEntities(w.Body, search.SourceList("senzing"))
		require.NoError(t, err)
		require.Len(t, entities, 2)
	})

	t.Run("jsonl/header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=2", nil)
		req.Header.Set("Accept", "senzing/jsonl")

		w := httptest.NewRecorder()
		env.router.ServeHTTP(w, req)

		// Make sure the response is JSON Lines
		require.True(t, strings.HasPrefix(w.Body.String(), "{"))

		entities, err := senzing.ReadEntities(w.Body, search.SourceList("senzing"))
		require.NoError(t, err)
		require.Len(t, entities, 2)
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

	b.Run("debug-threshold", func(b *testing.B) {
		req := httptest.NewRequest("GET", "/v2/search?name=Mohammad&type=person&limit=5&debug=0.80", nil)

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
