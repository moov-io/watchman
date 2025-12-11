package search_test

import (
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/cslustest"
	"github.com/moov-io/watchman/internal/entitytest"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

var (
	johnDoe = search.Entity[search.Value]{
		Name:     "John Doe",
		Type:     search.EntityPerson,
		Source:   search.SourceUSOFAC,
		SourceID: "12345",

		Person: &search.Person{
			Name:   "John Doe",
			Gender: search.GenderMale,
		},

		Contact: search.ContactInfo{
			EmailAddresses: []string{"john.doe@example.com"},
			PhoneNumbers:   []string{"123.456.7890"},
		},

		Addresses: []search.Address{
			{
				Line1:      "123 First St",
				City:       "Anytown",
				State:      "CA",
				PostalCode: "90210",
				Country:    "US",
			},
		},

		CryptoAddresses: []search.CryptoAddress{
			{
				Currency: "BTC",
				Address:  "be503b97-a5ec-4494-aacd-dc97c70293f3",
			},
		},
	}

	johnnyBirthDate = time.Date(1971, time.March, 26, 0, 0, 0, 0, time.UTC)
	johnnyDoe       = search.Entity[search.Value]{
		Name:     "Johnny Doe",
		Type:     search.EntityPerson,
		Source:   search.SourceUSOFAC,
		SourceID: "12345",

		Person: &search.Person{
			Name:      "Johnny Doe",
			BirthDate: &johnnyBirthDate,
			GovernmentIDs: []search.GovernmentID{
				{
					Type:       search.GovernmentIDPassport,
					Country:    "US",
					Identifier: "1981204918019",
				},
			},
		},

		Contact: search.ContactInfo{
			PhoneNumbers:   []string{"123.456.7890"},
			EmailAddresses: []string{"johnny.doe@example.com"},
			Websites:       []string{"http://johnnydoe.com"},
		},

		Addresses: []search.Address{
			{
				Line1:      "123 First St",
				Line2:      "Unit 456",
				City:       "Anytown",
				State:      "CA",
				PostalCode: "90210",
				Country:    "US",
			},
		},
	}

	johnJohnnyMerged = search.Entity[search.Value]{
		Name:     "John Doe",
		Type:     search.EntityPerson,
		Source:   search.SourceUSOFAC,
		SourceID: "12345",

		Person: &search.Person{
			Name:      "John Doe",
			AltNames:  []string{"Johnny Doe"},
			Gender:    "male",
			BirthDate: &johnnyBirthDate,
			GovernmentIDs: []search.GovernmentID{
				{Type: "passport", Country: "US", Identifier: "1981204918019"},
			},
		},

		Contact: search.ContactInfo{
			EmailAddresses: []string{"john.doe@example.com", "johnny.doe@example.com"},
			PhoneNumbers:   []string{"123.456.7890"},
			Websites:       []string{"http://johnnydoe.com"},
		},

		Addresses: []search.Address{
			{
				Line1:      "123 First St",
				City:       "Anytown",
				PostalCode: "90210",
				State:      "CA",
				Country:    "US",
			},
			{
				Line1:      "123 First St",
				Line2:      "Unit 456",
				City:       "Anytown",
				PostalCode: "90210",
				State:      "CA",
				Country:    "US",
			},
		},

		CryptoAddresses: []search.CryptoAddress{
			{
				Currency: "BTC",
				Address:  "be503b97-a5ec-4494-aacd-dc97c70293f3",
			},
		},
	}
)

func TestMerge(t *testing.T) {
	input := []search.Entity[search.Value]{johnDoe, johnnyDoe}
	merged := search.Merge(input)
	require.Len(t, merged, 1)

	require.Equal(t, johnJohnnyMerged.Name, merged[0].Name)
	require.Equal(t, johnJohnnyMerged.Type, merged[0].Type)
	require.Equal(t, johnJohnnyMerged.Source, merged[0].Source)

	require.Equal(t, johnJohnnyMerged.SourceID, merged[0].SourceID)

	require.Equal(t, johnJohnnyMerged.Person, merged[0].Person)
	require.Equal(t, johnJohnnyMerged.Business, merged[0].Business)
	require.Equal(t, johnJohnnyMerged.Organization, merged[0].Organization)
	require.Equal(t, johnJohnnyMerged.Aircraft, merged[0].Aircraft)
	require.Equal(t, johnJohnnyMerged.Vessel, merged[0].Vessel)

	require.Equal(t, johnJohnnyMerged.Contact, merged[0].Contact)
	require.ElementsMatch(t, johnJohnnyMerged.Addresses, merged[0].Addresses)
	require.ElementsMatch(t, johnJohnnyMerged.CryptoAddresses, merged[0].CryptoAddresses)

	require.ElementsMatch(t, johnJohnnyMerged.Affiliations, merged[0].Affiliations)
	require.Equal(t, johnJohnnyMerged.SanctionsInfo, merged[0].SanctionsInfo)
	require.ElementsMatch(t, johnJohnnyMerged.HistoricalInfo, merged[0].HistoricalInfo)
}

func TestMerge__Identity(t *testing.T) {
	// mostly to check that the code doens't panic or empty out fields

	t.Run("person", func(t *testing.T) {
		entity := ofactest.FindEntity(t, "10278")
		require.NotNil(t, entity.Person)

		got := search.Merge([]search.Entity[search.Value]{entity})
		require.Len(t, got, 1)

		entitytest.Equal(t, entity, got[0])
	})

	t.Run("business", func(t *testing.T) {
		entity := ofactest.FindEntity(t, "12685")
		require.NotNil(t, entity.Business)

		got := search.Merge([]search.Entity[search.Value]{entity})
		require.Len(t, got, 1)

		entitytest.Equal(t, entity, got[0])
	})

	t.Run("organization", func(t *testing.T) {
		entity := cslustest.FindEntity(t, "18283")
		require.NotNil(t, entity.Organization)

		got := search.Merge([]search.Entity[search.Value]{entity})
		require.Len(t, got, 1)

		entitytest.Equal(t, entity, got[0])
	})

	t.Run("aircraft", func(t *testing.T) {
		entity := ofactest.FindEntity(t, "20540")
		require.NotNil(t, entity.Aircraft)

		got := search.Merge([]search.Entity[search.Value]{entity})
		require.Len(t, got, 1)

		entitytest.Equal(t, entity, got[0])
	})

	t.Run("vessel", func(t *testing.T) {
		entity := ofactest.FindEntity(t, "40716")
		require.NotNil(t, entity.Vessel)

		got := search.Merge([]search.Entity[search.Value]{entity})
		require.Len(t, got, 1)

		entitytest.Equal(t, entity, got[0])
	})
}
