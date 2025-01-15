package search

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

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
