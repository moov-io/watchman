package ingest_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestService_ReadEntitiesFromFile_FincenBusiness(t *testing.T) {
	t.Setenv("APP_CONFIG_SECRETS", filepath.Join("testdata", "fincen-config.yml"))

	ctx := context.Background()
	logger := log.NewTestLogger()

	conf, err := config.LoadConfig(logger)
	require.NoError(t, err)

	svc := ingest.NewService(logger, conf.Ingest)

	fd, err := os.Open(filepath.Join("testdata", "fincen-business.csv"))
	require.NoError(t, err)
	t.Cleanup(func() { fd.Close() })

	entities, err := svc.ReadEntitiesFromFile(ctx, "fincen-business", fd)
	require.NoError(t, err)

	expected := []search.Entity[search.Value]{
		{
			Name:     "Acme Corp",
			Type:     search.EntityBusiness,
			Source:   search.SourceList("fincen-business"),
			SourceID: "123456",
			Business: &search.Business{
				Name:     "Acme Corp",
				AltNames: []string{"Acme Corp LLC"},
			},
			Addresses: []search.Address{
				{
					Line1:      "999 South Marine Corps Drive",
					City:       "Anytown",
					State:      "CA",
					PostalCode: "90210",
					Country:    "US",
				},
			},
		},
		{
			Name:     "Anvil Investments LLC",
			Type:     search.EntityBusiness,
			Source:   search.SourceList("fincen-business"),
			SourceID: "654321",
			Business: &search.Business{
				Name: "Anvil Investments LLC",
				GovernmentIDs: []search.GovernmentID{
					{
						Identifier: "992288844",
					},
				},
			},
			Addresses: []search.Address{
				{
					Line1:      "79 Other Way",
					City:       "Othertown",
					State:      "NY",
					PostalCode: "12946",
					Country:    "US",
				},
			},
		},
	}
	require.ElementsMatch(t, expected, entities)
}
