package ofac_test

import (
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/ofac"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMapperPerson__FromSource(t *testing.T) {
	t.Run("48603", func(t *testing.T) {
		found := ofactest.FindEntity(t, "48603")

		require.Equal(t, "Dmitry Yuryevich KHOROSHEV", found.Name)
		require.Equal(t, search.EntityPerson, found.Type)
		require.Equal(t, search.SourceUSOFAC, found.Source)
		require.Equal(t, "48603", found.SourceID)

		require.NotNil(t, found.Person)
		require.Nil(t, found.Business)
		require.Nil(t, found.Organization)
		require.Nil(t, found.Aircraft)
		require.Nil(t, found.Vessel)

		person := found.Person
		require.Equal(t, "Dmitry Yuryevich KHOROSHEV", person.Name)

		expectedAltNames := []string{
			"LOCKBITSUPP", "Dmitriy Yurevich KHOROSHEV",
			"Dmitry YURIEVICH", "Dmitrii Yuryevich KHOROSHEV",
		}
		require.ElementsMatch(t, expectedAltNames, person.AltNames)

		require.Equal(t, search.GenderMale, person.Gender)

		expectedBirthDate := time.Date(1993, time.April, 17, 0, 0, 0, 0, time.UTC)
		require.Equal(t, expectedBirthDate.Format(time.RFC3339), person.BirthDate.Format(time.RFC3339))
		require.Nil(t, person.DeathDate)

		require.Empty(t, person.Titles)

		expectedGovernmentIDs := []search.GovernmentID{
			{
				Type:       search.GovernmentIDPassport,
				Country:    "Russia",
				Identifier: "2018278055",
			},
			{
				Type:       search.GovernmentIDPassport,
				Country:    "Russia",
				Identifier: "2006801524",
			},
			{
				Type:       search.GovernmentIDTax,
				Country:    "Russia",
				Identifier: "366110340670",
			},
		}
		require.ElementsMatch(t, expectedGovernmentIDs, person.GovernmentIDs)

		// "DOB 17 Apr 1993; POB Russian Federation; nationality Russia; citizen Russia; Email Address khoroshev1@icloud.com;
		// alt. Email Address sitedev5@yandex.ru; Gender Male; Digital Currency Address - XBT bc1qvhnfknw852ephxyc5hm4q520zmvf9maphetc9z;
		// Secondary sanctions risk: Ukraine-/Russia-Related Sanctions Regulations, 31 CFR 589.201; Passport 2018278055 (Russia);
		// alt. Passport 2006801524 (Russia); Tax ID No. 366110340670 (Russia); a.k.a. 'LOCKBITSUPP'."

		expectedContact := search.ContactInfo{
			EmailAddresses: []string{"khoroshev1@icloud.com", "sitedev5@yandex.ru"},
		}
		require.Equal(t, expectedContact, found.Contact)
		require.Empty(t, found.Addresses)

		expectedCryptoAddresses := []search.CryptoAddress{
			{Currency: "XBT", Address: "bc1qvhnfknw852ephxyc5hm4q520zmvf9maphetc9z"},
		}
		require.ElementsMatch(t, expectedCryptoAddresses, found.CryptoAddresses)

		require.Empty(t, found.Affiliations)
		require.Nil(t, found.SanctionsInfo)
		require.Empty(t, found.HistoricalInfo)

		sdn, ok := found.SourceData.(ofac.SDN)
		require.True(t, ok)
		require.Equal(t, "48603", sdn.EntityID)
		require.Equal(t, "KHOROSHEV, Dmitry Yuryevich", sdn.SDNName)
	})
}

func TestMapper__Person(t *testing.T) {
	e := ofactest.FindEntity(t, "15102")

	require.Equal(t, "Daniel MORENO", e.Name)
	require.Equal(t, search.EntityPerson, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Person)
	require.Equal(t, "Daniel MORENO", e.Person.Name)
	require.Equal(t, "", string(e.Person.Gender))
	require.Equal(t, "1972-10-12T00:00:00Z", e.Person.BirthDate.Format(time.RFC3339))
	require.Nil(t, e.Person.DeathDate)
	require.Len(t, e.Person.GovernmentIDs, 1)

	passport := e.Person.GovernmentIDs[0]
	require.Equal(t, search.GovernmentIDPassport, passport.Type)
	require.Equal(t, "Belize", passport.Country)
	require.Equal(t, "0291622", passport.Identifier)

	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
	require.Nil(t, e.Vessel)

	sourceData, ok := e.SourceData.(ofac.SDN)
	require.True(t, ok)
	require.Equal(t, "15102", sourceData.EntityID)
}

func TestMapper__CompletePerson(t *testing.T) {
	sdn := ofac.SDN{
		EntityID: "26057",
		SDNName:  "AL-ZAYDI, Shibl Muhsin 'Ubayd",
		SDNType:  "individual",
		Remarks:  "DOB 28 Oct 1968; POB Baghdad, Iraq; Additional Sanctions Information - Subject to Secondary Sanctions Pursuant to the Hizballah Financial Sanctions Regulations; alt. Additional Sanctions Information - Subject to Secondary Sanctions; Gender Male; a.k.a. 'SHIBL, Hajji'; nationality Iran; Passport A123456 (Iran) expires 2024; Driver's License No. 04900377 (Moldova) issued 02 Jul 2004; Email Address test@example.com; Phone: +1-123-456-7890; Fax: +1-123-456-7899",
	}

	e := ofac.ToEntity(sdn, nil, nil, nil)
	require.Equal(t, "Shibl Muhsin 'Ubayd AL-ZAYDI", e.Name)
	require.Equal(t, search.EntityPerson, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	// Person specific fields
	require.NotNil(t, e.Person)
	require.Equal(t, "Shibl Muhsin 'Ubayd AL-ZAYDI", e.Person.Name)
	require.Equal(t, search.GenderMale, e.Person.Gender)
	require.Equal(t, "1968-10-28T00:00:00Z", e.Person.BirthDate.Format(time.RFC3339))
	require.Nil(t, e.Person.DeathDate)

	// Test alt names
	require.Len(t, e.Person.AltNames, 1)
	require.Equal(t, "Hajji SHIBL", e.Person.AltNames[0])

	// Test government IDs
	require.Len(t, e.Person.GovernmentIDs, 2)
	var passport, license *search.GovernmentID
	for i := range e.Person.GovernmentIDs {
		if e.Person.GovernmentIDs[i].Type == search.GovernmentIDPassport {
			passport = &e.Person.GovernmentIDs[i]
		} else {
			license = &e.Person.GovernmentIDs[i]
		}
	}
	require.NotNil(t, passport)
	require.Equal(t, "Iran", passport.Country)
	require.Equal(t, "A123456", passport.Identifier)

	require.NotNil(t, license)
	require.Equal(t, "Moldova", license.Country)
	require.Equal(t, "04900377", license.Identifier)

	// Verify other entity types are nil
	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
	require.Nil(t, e.Vessel)
}

func TestMapper__CompletePersonWithRemarks(t *testing.T) {
	sdn := ofac.SDN{
		EntityID: "26057",
		SDNName:  "Shibl Muhsin Ubayd al-Zaydi",
		SDNType:  "individual",
		Remarks:  "DOB 28 Oct 1968; POB Baghdad, Iraq; Gender Male; Title: Commander; Former Name: AL-ZAYDI, Muhammad; Linked To: ISLAMIC REVOLUTIONARY GUARD CORPS (IRGC)-QODS FORCE; Additional Sanctions Information - Subject to Secondary Sanctions",
	}

	e := ofac.ToEntity(sdn, nil, nil, nil)

	// Test affiliations
	require.Len(t, e.Affiliations, 1)
	require.Equal(t, "ISLAMIC REVOLUTIONARY GUARD CORPS (IRGC)-QODS FORCE", e.Affiliations[0].EntityName)
	require.Equal(t, "Linked To", e.Affiliations[0].Type)

	// Test sanctions info
	require.NotNil(t, e.SanctionsInfo)
	require.True(t, e.SanctionsInfo.Secondary)
	require.Equal(t, "Subject to Secondary Sanctions", e.SanctionsInfo.Description)

	// Test historical info
	require.Len(t, e.HistoricalInfo, 1)
	require.Equal(t, "Former Name", e.HistoricalInfo[0].Type)
	require.Equal(t, "AL-ZAYDI, Muhammad", e.HistoricalInfo[0].Value)

	// Test titles
	require.Equal(t, []string{"Commander"}, e.Person.Titles)
}

func TestMapper__PersonWithTitle(t *testing.T) {
	sdn := ofac.SDN{
		EntityID: "12345",
		SDNName:  "SMITH, John",
		SDNType:  "individual",
		Title:    "Chief Financial Officer",
		Remarks:  "Title: Regional Director",
	}

	e := ofac.ToEntity(sdn, nil, nil, nil)
	require.Equal(t, "John SMITH", e.Name)
	require.Equal(t, search.EntityPerson, e.Type)

	// Should have both titles - from SDN field and remarks
	require.Contains(t, e.Person.Titles, "Chief Financial Officer")
	require.Contains(t, e.Person.Titles, "Regional Director")
}
