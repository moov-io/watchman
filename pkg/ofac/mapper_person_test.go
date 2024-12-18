package ofac

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMapper__Person(t *testing.T) {
	res, err := Read(testInputs(t, filepath.Join("..", "..", "test", "testdata", "sdn.csv")))
	require.NoError(t, err)

	var sdn *SDN
	for i := range res.SDNs {
		if res.SDNs[i].EntityID == "15102" {
			sdn = res.SDNs[i]
		}
	}
	require.NotNil(t, sdn)

	e := ToEntity(*sdn, nil, nil)
	require.Equal(t, "MORENO, Daniel", e.Name)
	require.Equal(t, search.EntityPerson, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Person)
	require.Equal(t, "MORENO, Daniel", e.Person.Name)
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

	require.Equal(t, "15102", e.SourceData.EntityID)
}

func TestMapper__CompletePerson(t *testing.T) {
	sdn := &SDN{
		EntityID: "26057",
		SDNName:  "AL-ZAYDI, Shibl Muhsin 'Ubayd",
		SDNType:  "individual",
		Remarks:  "DOB 28 Oct 1968; POB Baghdad, Iraq; Additional Sanctions Information - Subject to Secondary Sanctions Pursuant to the Hizballah Financial Sanctions Regulations; alt. Additional Sanctions Information - Subject to Secondary Sanctions; Gender Male; a.k.a. 'SHIBL, Hajji'; nationality Iran; Passport A123456 (Iran) expires 2024; Driver's License No. 04900377 (Moldova) issued 02 Jul 2004; Email Address test@example.com; Phone: +1-123-456-7890; Fax: +1-123-456-7899",
	}

	e := ToEntity(*sdn, nil, nil)
	require.Equal(t, "AL-ZAYDI, Shibl Muhsin 'Ubayd", e.Name)
	require.Equal(t, search.EntityPerson, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	// Person specific fields
	require.NotNil(t, e.Person)
	require.Equal(t, "AL-ZAYDI, Shibl Muhsin 'Ubayd", e.Person.Name)
	require.Equal(t, search.GenderMale, e.Person.Gender)
	require.Equal(t, "1968-10-28T00:00:00Z", e.Person.BirthDate.Format(time.RFC3339))
	require.Nil(t, e.Person.DeathDate)

	// Test alt names
	require.Len(t, e.Person.AltNames, 1)
	require.Equal(t, "SHIBL, Hajji", e.Person.AltNames[0])

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

func TestParseAltNames(t *testing.T) {
	tests := []struct {
		remarks  []string
		expected []string
	}{
		{
			remarks:  []string{"a.k.a. 'SMITH, John'"},
			expected: []string{"SMITH, John"},
		},
		{
			remarks:  []string{"a.k.a. 'SMITH, John'; a.k.a. 'DOE, Jane'"},
			expected: []string{"SMITH, John", "DOE, Jane"},
		},
		{
			remarks:  []string{"Some other remark", "a.k.a. 'SMITH, John'"},
			expected: []string{"SMITH, John"},
		},
		{
			remarks:  []string{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		result := parseAltNames(tt.remarks)
		require.Equal(t, tt.expected, result)
	}
}

func TestMapper__CompletePersonWithRemarks(t *testing.T) {
	sdn := &SDN{
		EntityID: "26057",
		SDNName:  "AL-ZAYDI, Shibl Muhsin 'Ubayd",
		SDNType:  "individual",
		Remarks:  "DOB 28 Oct 1968; POB Baghdad, Iraq; Gender Male; Title: Commander; Former Name: AL-ZAYDI, Muhammad; Linked To: ISLAMIC REVOLUTIONARY GUARD CORPS (IRGC)-QODS FORCE; Additional Sanctions Information - Subject to Secondary Sanctions",
	}

	e := ToEntity(*sdn, nil, nil)

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
	require.Equal(t, []string{"Commander"}, e.Titles)
}

func TestMapper__PersonWithTitle(t *testing.T) {
	sdn := &SDN{
		EntityID: "12345",
		SDNName:  "SMITH, John",
		SDNType:  "individual",
		Title:    "Chief Financial Officer",
		Remarks:  "Title: Regional Director",
	}

	e := ToEntity(*sdn, nil, nil)
	require.Equal(t, "SMITH, John", e.Name)
	require.Equal(t, search.EntityPerson, e.Type)

	// Should have both titles - from SDN field and remarks
	require.Contains(t, e.Titles, "Chief Financial Officer")
	require.Contains(t, e.Titles, "Regional Director")
}
