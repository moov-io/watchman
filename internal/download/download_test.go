package download_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/geocoding/geocodetest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestDownloader_RefreshAll_InitialDir(t *testing.T) {
	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
		},
	}

	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)

	require.Greater(t, len(stats.Entities), 100)

	name := string(search.SourceUSOFAC)
	require.Greater(t, stats.Lists[name], 100)
	require.NotEmpty(t, stats.ListHashes[name])
}

func TestDownloader_RefreshAll_UKCSL(t *testing.T) {
	// Skip HTML index parsing by providing a placeholder URL
	// The actual file will be loaded from InitialDataDirectory
	t.Setenv("UK_SANCTIONS_LIST_URL", "https://example.com/UK_Sanctions_List.ods")

	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUKCSL,
		},
	}

	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)

	// Verify entities were loaded
	require.Greater(t, len(stats.Entities), 100, "expected more than 100 UK CSL entities")

	// Verify list tracking
	name := string(search.SourceUKCSL)
	require.Greater(t, stats.Lists[name], 100)
	require.NotEmpty(t, stats.ListHashes[name], "expected list hash to be computed")

	// Count entity types and collect data for verification
	var persons, businesses, vessels int
	var entitiesWithAddresses, entitiesWithAltNames int

	for _, entity := range stats.Entities {
		require.Equal(t, search.SourceUKCSL, entity.Source, "entity should have UK CSL source")
		require.NotEmpty(t, entity.SourceID, "entity should have source ID")
		require.NotEmpty(t, entity.Name, "entity should have name")

		// Verify SourceData is preserved
		require.NotNil(t, entity.SourceData, "entity should have SourceData")

		// Count entities with addresses
		if len(entity.Addresses) > 0 {
			entitiesWithAddresses++
		}

		switch entity.Type {
		case search.EntityPerson:
			persons++
			require.NotNil(t, entity.Person, "person entity should have Person field")
			if len(entity.Person.AltNames) > 0 {
				entitiesWithAltNames++
			}
		case search.EntityBusiness:
			businesses++
			require.NotNil(t, entity.Business, "business entity should have Business field")
			if len(entity.Business.AltNames) > 0 {
				entitiesWithAltNames++
			}
		case search.EntityVessel:
			vessels++
			require.NotNil(t, entity.Vessel, "vessel entity should have Vessel field")
			if len(entity.Vessel.AltNames) > 0 {
				entitiesWithAltNames++
			}
		}
	}

	// Verify we have a mix of entity types
	require.Greater(t, persons, 0, "expected some Person entities")
	require.Greater(t, businesses, 0, "expected some Business entities")

	// Verify data quality - some entities should have addresses and alt names
	require.Greater(t, entitiesWithAddresses, 0, "expected some entities with addresses")
	require.Greater(t, entitiesWithAltNames, 0, "expected some entities with alt names")

	t.Logf("UK CSL loaded: %d persons, %d businesses, %d vessels", persons, businesses, vessels)
	t.Logf("Data quality: %d with addresses, %d with alt names", entitiesWithAddresses, entitiesWithAltNames)
}

func TestDownloader_RefreshAll_UKCSL_EntityDetails(t *testing.T) {
	// Skip HTML index parsing by providing a placeholder URL
	t.Setenv("UK_SANCTIONS_LIST_URL", "https://example.com/UK_Sanctions_List.ods")

	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUKCSL,
		},
	}

	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)

	// Find entities and verify their structure
	var foundPerson, foundBusiness, foundVessel bool

	for _, entity := range stats.Entities {
		// Check a Person entity
		if entity.Type == search.EntityPerson && !foundPerson {
			require.NotNil(t, entity.Person)
			require.NotEmpty(t, entity.Person.Name, "person should have name")
			// PreparedFields should be populated after Normalize()
			require.NotEmpty(t, entity.PreparedFields.NameFields, "person should have prepared name fields")
			foundPerson = true
		}

		// Check a Business entity
		if entity.Type == search.EntityBusiness && !foundBusiness {
			require.NotNil(t, entity.Business)
			require.NotEmpty(t, entity.Business.Name, "business should have name")
			require.NotEmpty(t, entity.PreparedFields.NameFields, "business should have prepared name fields")
			foundBusiness = true
		}

		// Check a Vessel entity
		if entity.Type == search.EntityVessel && !foundVessel {
			require.NotNil(t, entity.Vessel)
			require.NotEmpty(t, entity.Vessel.Name, "vessel should have name")
			require.NotEmpty(t, entity.PreparedFields.NameFields, "vessel should have prepared name fields")
			foundVessel = true
		}

		if foundPerson && foundBusiness && foundVessel {
			break
		}
	}

	require.True(t, foundPerson, "should find at least one Person entity")
	require.True(t, foundBusiness, "should find at least one Business entity")
	require.True(t, foundVessel, "should find at least one Vessel entity")
}

func TestDownloader_Geocode(t *testing.T) {
	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
		},
	}

	g := &geocodetest.RandomGeocoder{}

	dl, err := download.NewDownloader(logger, conf, g)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)
	require.Greater(t, len(stats.Entities), 100)

	for _, entity := range stats.Entities {
		for _, addr := range entity.Addresses {
			require.NotEmpty(t, addr.Latitude)
			require.NotEmpty(t, addr.Longitude)
		}
	}
}
