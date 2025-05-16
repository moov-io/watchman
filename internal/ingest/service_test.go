package ingest_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

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

	parsedFile, err := svc.ReadEntitiesFromFile(ctx, "fincen-business", fd)
	require.NoError(t, err)
	require.Equal(t, "fincen-business", parsedFile.FileType)

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
			PreparedFields: search.PreparedFields{
				Name:          "acme corp",
				AltNames:      []string{"acme corp llc"},
				NameFields:    []string{"acme", "corp"},
				AltNameFields: [][]string{{"acme", "corp", "llc"}},
				Addresses: []search.PreparedAddress{
					{
						Line1:       "999 south marine corps drive",
						Line1Fields: []string{"999", "south", "marine", "corps", "drive"},
						City:        "anytown",
						CityFields:  []string{"anytown"},
						PostalCode:  "90210",
						State:       "ca",
						Country:     "united states",
					},
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
			PreparedFields: search.PreparedFields{
				Name:       "anvil investments llc",
				NameFields: []string{"anvil", "investments", "llc"},
				Addresses: []search.PreparedAddress{
					{
						Line1:       "79 other way",
						Line1Fields: []string{"79", "other", "way"},
						City:        "othertown",
						CityFields:  []string{"othertown"},
						PostalCode:  "12946",
						State:       "ny",
						Country:     "united states",
					},
				},
			},
		},
	}
	require.ElementsMatch(t, expected, parsedFile.Entities)
}

func TestService_ReadEntitiesFromFile_FincenPerson(t *testing.T) {
	t.Setenv("APP_CONFIG_SECRETS", filepath.Join("testdata", "fincen-config.yml"))

	ctx := context.Background()
	logger := log.NewTestLogger()

	conf, err := config.LoadConfig(logger)
	require.NoError(t, err)

	svc := ingest.NewService(logger, conf.Ingest)

	fd, err := os.Open(filepath.Join("testdata", "fincen-person.csv"))
	require.NoError(t, err)
	t.Cleanup(func() { fd.Close() })

	parsedFile, err := svc.ReadEntitiesFromFile(ctx, "fincen-person", fd)
	require.NoError(t, err)
	require.Equal(t, "fincen-person", parsedFile.FileType)

	expected := []search.Entity[search.Value]{
		{
			Name:     "John Jr K Doe1",
			Type:     search.EntityPerson,
			Source:   search.SourceList("fincen-person"),
			SourceID: "123456",
			Person: &search.Person{
				Name:      "John Jr K Doe1",
				AltNames:  []string{"Johnny K Doe"},
				BirthDate: ptr(time.Date(1988, time.February, 8, 0, 0, 0, 0, time.UTC)),
				GovernmentIDs: []search.GovernmentID{
					{
						Identifier: "123456789",
					},
				},
			},
			Addresses: []search.Address{
				{
					Line1:      "193 Southfield Lane",
					City:       "Anytown",
					State:      "CA",
					PostalCode: "90210",
					Country:    "US",
				},
			},
			PreparedFields: search.PreparedFields{
				Name:          "john jr k doe1",
				AltNames:      []string{"johnny k doe"},
				NameFields:    []string{"john", "jr", "k", "doe"},
				AltNameFields: [][]string{{"johnny", "k", "doe"}},
				Addresses: []search.PreparedAddress{
					{
						Line1:       "193 southfield lane",
						Line1Fields: []string{"193", "southfield", "lane"},
						City:        "anytown",
						CityFields:  []string{"anytown"},
						PostalCode:  "90210",
						State:       "ca",
						Country:     "united states",
					},
				},
			},
		},
		{
			Name:     "Jane K Doe2",
			Type:     search.EntityPerson,
			Source:   search.SourceList("fincen-person"),
			SourceID: "214365",
			Person: &search.Person{
				Name:      "Jane K Doe2",
				AltNames:  []string{"Jane L Doe"},
				BirthDate: ptr(time.Date(1988, time.March, 9, 0, 0, 0, 0, time.UTC)),
				GovernmentIDs: []search.GovernmentID{
					{
						Identifier: "4CC44444",
					},
				},
			},
			Addresses: []search.Address{
				{
					Line1:      "931 Southfield Lane",
					City:       "Anytown",
					State:      "CA",
					PostalCode: "90210",
					Country:    "US",
				},
			},
			PreparedFields: search.PreparedFields{
				Name:          "jane k doe2",
				AltNames:      []string{"jane l doe"},
				NameFields:    []string{"jane", "k", "doe2"},
				AltNameFields: [][]string{{"jane", "l", "doe"}},
				Addresses: []search.PreparedAddress{
					{
						Line1:       "931 southfield lane",
						Line1Fields: []string{"931", "southfield", "lane"},
						City:        "anytown",
						CityFields:  []string{"anytown"},
						PostalCode:  "90210",
						State:       "ca",
						Country:     "united states",
					},
				},
			},
		},
		{
			Name:     "Jose K Doe3",
			Type:     search.EntityPerson,
			Source:   search.SourceList("fincen-person"),
			SourceID: "321654",
			Person: &search.Person{
				Name:      "Jose K Doe3",
				AltNames:  []string{"Joseph M Doe"},
				BirthDate: ptr(time.Date(1988, time.April, 10, 0, 0, 0, 0, time.UTC)),
				GovernmentIDs: []search.GovernmentID{
					{
						Identifier: "987654321",
					},
				},
			},
			Addresses: []search.Address{
				{
					Line1:      "391 Southfield Lane",
					City:       "Anytown",
					State:      "CA",
					PostalCode: "90210",
					Country:    "US",
				},
			},
			PreparedFields: search.PreparedFields{
				Name:          "jose k doe3",
				AltNames:      []string{"joseph m doe"},
				NameFields:    []string{"jose", "k", "doe3"},
				AltNameFields: [][]string{{"joseph", "m", "doe"}},
				Addresses: []search.PreparedAddress{
					{
						Line1:       "391 southfield lane",
						Line1Fields: []string{"391", "southfield", "lane"},
						City:        "anytown",
						CityFields:  []string{"anytown"},
						PostalCode:  "90210",
						State:       "ca",
						Country:     "united states",
					},
				},
			},
		},
	}

	require.Len(t, parsedFile.Entities, len(expected))

	require.Equal(t, expected[0], parsedFile.Entities[0])
	require.Equal(t, expected[1], parsedFile.Entities[1])
	require.Equal(t, expected[2], parsedFile.Entities[2])
}

func ptr[T any](in T) *T {
	return &in
}
