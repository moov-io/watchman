package decisions

import (
	"os"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestIntegration_SystemOne(t *testing.T) {
	if testing.Short() {
		t.Skip("-short flag provided")
	}

	baseURL := os.Getenv("DECISIONS_BASE_URL")
	if baseURL == "" {
		t.Skip("DECISIONS_BASE_URL not set")
	}

	conf := DefaultConfig()
	conf.Enabled = true
	conf.BaseURL = baseURL
	if model := os.Getenv("DECISIONS_MODEL"); model != "" {
		conf.Model = model
	}
	if key := os.Getenv("DECISIONS_API_KEY"); key != "" {
		conf.APIKey = key
	}
	conf.Timeout = 15 * time.Second

	svc, err := NewService(log.NewTestLogger(), conf)
	require.NoError(t, err)
	require.NotNil(t, svc)

	results := []search.SearchedEntity[search.Value]{
		{
			Entity: search.Entity[search.Value]{
				Name:     "Dmitry Yuryevich KHOROSHEV",
				Type:     search.EntityPerson,
				Source:   search.SourceUSOFAC,
				SourceID: "48603",
				Person: &search.Person{
					Name: "Dmitry Yuryevich KHOROSHEV",
					GovernmentIDs: []search.GovernmentID{{
						Type:       search.GovernmentIDPassport,
						Country:    "RU",
						Identifier: "2018278055",
					}},
				},
			},
			Match: 1.0,
		},
	}
	svc.Annotate(t.Context(), search.Entity[search.Value]{
		Name: "Dmitry Khoroshev",
		Type: search.EntityPerson,
		Person: &search.Person{
			Name: "Dmitry Khoroshev",
			GovernmentIDs: []search.GovernmentID{{
				Type:       search.GovernmentIDPassport,
				Country:    "RU",
				Identifier: "2018278055",
			}},
		},
	}, results)

	require.NotNil(t, results[0].Decision)
	require.Greater(t, results[0].Decision.SameEntity, 0.5)
}
