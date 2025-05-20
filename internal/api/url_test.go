package api_test

import (
	"net/http/httptest"
	"testing"

	"github.com/moov-io/watchman/internal/api"

	"github.com/stretchr/testify/require"
)

func TestQueryParams(t *testing.T) {
	req := httptest.NewRequest("GET", "/v2/search?name=adam&type=person&addresses=123+first+St", nil)

	qp := api.QueryParams{Values: req.URL.Query()}

	// simulate reading some of the query params
	qp.Get("name")
	qp.Get("type")

	// query param provided that we didn't read
	expected := []string{"addresses"}
	require.ElementsMatch(t, expected, qp.UnusedQueryParams())
}
