package search

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestAPI_ListInfo(t *testing.T) {
	logger := log.NewTestLogger()
	service := NewService(logger)

	stats := download.Stats{
		Lists:      make(map[string]int),
		ListHashes: make(map[string]string),
		StartedAt:  time.Now().In(time.UTC),
	}
	stats.Lists[string(search.SourceUSOFAC)] = 123
	stats.ListHashes[string(search.SourceUSOFAC)] = "abc1234"
	service.UpdateEntities(stats)

	controller := NewController(logger, service)

	router := mux.NewRouter()
	controller.AppendRoutes(router)

	req := httptest.NewRequest("GET", "/v2/listinfo", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body := w.Body.String()
	require.Contains(t, body, `"lists":{"us_ofac":123}`)
	require.Contains(t, body, `"listHashes":{"us_ofac":"abc1234"}`)
}

func TestAPI_readSearchRequest(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v2/search?name=adam&type=person&birthDate=2025-01-02", nil)

		query, err := readSearchRequest(req)
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

		query, err := readSearchRequest(req)
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

		query, err := readSearchRequest(req)
		require.NoError(t, err)
		require.Empty(t, query.Name)

		require.Len(t, query.CryptoAddresses, 2)

		expected := []search.CryptoAddress{
			{Currency: "XBT", Address: "12345"},
			{Currency: "ETH", Address: "54321"},
		}
		require.ElementsMatch(t, expected, query.CryptoAddresses)
	})
}
