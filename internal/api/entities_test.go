package api_test

import (
	"net/http"
	"testing"

	"github.com/moov-io/watchman/internal/api"

	"github.com/stretchr/testify/require"
)

func TestChooseEntityFormat(t *testing.T) {
	cases := []struct {
		name string

		headers    map[string]string
		queryParam string

		expectedFormat    api.EntityFormat
		expectedSubformat string
	}{
		{
			name:              "blank defaults",
			expectedFormat:    api.EntityWatchman,
			expectedSubformat: "json",
		},
		{
			name: "header senzing",
			headers: map[string]string{
				"accept": "application/json, senzing/json",
			},
			expectedFormat:    api.EntitySenzing,
			expectedSubformat: "json",
		},
		{
			name:              "query senzing",
			queryParam:        "senzing",
			expectedFormat:    api.EntitySenzing,
			expectedSubformat: "",
		},
		{
			name: "header senzing jsonl",
			headers: map[string]string{
				"accept": "senzing/jsonl, application/json",
			},
			expectedFormat:    api.EntitySenzing,
			expectedSubformat: "jsonl",
		},
		{
			name:              "query senzing jsonl",
			queryParam:        "senzing/jsonl",
			expectedFormat:    api.EntitySenzing,
			expectedSubformat: "jsonl",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			headers := make(http.Header)
			for k, v := range tc.headers {
				headers.Set(k, v)
			}

			format, subformat := api.ChooseEntityFormat(headers, tc.queryParam)
			require.Equal(t, tc.expectedFormat, format)
			require.Equal(t, tc.expectedSubformat, subformat)
		})
	}
}
