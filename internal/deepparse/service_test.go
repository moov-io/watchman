package deepparse

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	deepparsego "github.com/adamdecaf/deepparse-go"
	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestNewService_Disabled(t *testing.T) {
	svc, err := NewService(log.NewTestLogger(), Config{Enabled: false})
	require.NoError(t, err)
	require.Nil(t, svc)
}

func TestParseAddress(t *testing.T) {
	const raw = "350 rue des Lilas Ouest Quebec city Quebec G1L 1B6"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/parse/bpemb-attention", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Contains(t, string(body), raw)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(`{
			"model_type": "bpemb_attention",
			"parsed_addresses": [
				{"` + raw + `": {
					"StreetNumber": "350",
					"StreetName": "rue des lilas",
					"Orientation": "ouest",
					"Municipality": "quebec city",
					"Province": "quebec",
					"PostalCode": "g1l 1b6"
				}}
			],
			"version": "test"
		}`))
		require.NoError(t, err)
	}))
	t.Cleanup(srv.Close)

	svc := &Service{
		client: deepparsego.NewClient(srv.Client(), srv.URL),
		model:  deepparsego.ModelBPEmbAttention,
	}

	addr, err := svc.ParseAddress(context.Background(), raw)
	require.NoError(t, err)
	require.Equal(t, "350 rue des lilas ouest", addr.Line1)
	require.Equal(t, "quebec city", addr.City)
	require.Equal(t, "quebec", addr.State)
	require.Equal(t, "g1l 1b6", addr.PostalCode)
	require.Empty(t, addr.Country)
}

func TestParseAddress_HTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(map[string]string{"detail": "Addresses parameter must not be empty"})
	}))
	t.Cleanup(srv.Close)

	svc := &Service{
		client: deepparsego.NewClient(srv.Client(), srv.URL),
		model:  deepparsego.ModelBPEmb,
	}

	_, err := svc.ParseAddress(context.Background(), "an address")
	require.Error(t, err)
}

func TestMapParsed(t *testing.T) {
	got := mapParsed(deepparsego.ParsedAddress{
		StreetNumber:    "2325",
		StreetName:      "rue de l'université",
		Unit:            "suite 100",
		Municipality:    "québec",
		Province:        "qc",
		PostalCode:      "g1v 0a6",
		GeneralDelivery: "po box 12",
	})
	require.Equal(t, "2325 rue de l'université", got.Line1)
	require.Equal(t, "suite 100, po box 12", got.Line2)
	require.Equal(t, "québec", got.City)
	require.Equal(t, "qc", got.State)
	require.Equal(t, "g1v 0a6", got.PostalCode)
}

func TestIntegrationParseAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping deepparse docker integration in short mode")
	}
	base := os.Getenv("DEEPPARSE_URL")
	if base == "" {
		t.Skip("DEEPPARSE_URL not set")
	}

	svc, err := NewService(log.NewTestLogger(), Config{
		Enabled: true,
		BaseURL: base,
		Model:   string(deepparsego.ModelBPEmbAttention),
		Timeout: 30 * time.Second,
	})
	require.NoError(t, err)
	require.NotNil(t, svc)

	addr, err := svc.ParseAddress(context.Background(), "350 rue des Lilas Ouest Quebec city Quebec G1L 1B6")
	require.NoError(t, err)
	require.Equal(t, "350", firstField(addr.Line1))
	require.NotEmpty(t, addr.PostalCode)
}

func firstField(s string) string {
	for i, r := range s {
		if r == ' ' {
			return s[:i]
		}
	}
	return s
}
