package api_test

import (
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/watchman/internal/api"
	"github.com/stretchr/testify/require"
)

func TestErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	err := api.ErrorResponse(w, errors.New("bad error"))
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Equal(t, "application/json", w.Header().Get("Content-Type"))

	expected := `{"error":"bad error"}` + "\n"
	require.Equal(t, expected, w.Body.String())
}

type Object struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Male  bool    `json:"male"`
	Score float64 `json:"score"`
}

func TestJsonResponse(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := api.JsonResponse(w, Object{
			Name: "Jane Doe",
			Age:  23,
			Male: false,
		})
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "application/json", w.Header().Get("Content-Type"))
		require.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))

		expected := `{"name":"Jane Doe","age":23,"male":false,"score":0}` + "\n"
		require.Equal(t, expected, w.Body.String())
	})

	t.Run("NaN", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := api.JsonResponse(w, Object{
			Name:  "Jane Doe",
			Age:   23,
			Male:  false,
			Score: math.NaN(),
		})
		require.ErrorContains(t, err, "json: unsupported value: NaN")

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Equal(t, "application/json", w.Header().Get("Content-Type"))
		require.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))

		expected := `{"error":"problem rendering json: json: unsupported value: NaN"}` + "\n"
		require.Equal(t, expected, w.Body.String())
	})
}
