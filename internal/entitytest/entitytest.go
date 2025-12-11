package entitytest

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func Equal[T any](tb testing.TB, e1, e2 search.Entity[T]) {
	require.Equal(tb, e1.Name, e2.Name)
	require.Equal(tb, e1.Type, e2.Type)
	require.Equal(tb, e1.Source, e2.Source)

	require.Equal(tb, e1.SourceID, e2.SourceID)

	if e1.Person != nil && e2.Person != nil {
		require.Equal(tb, e1.Person.Name, e2.Person.Name)
		require.ElementsMatch(tb, e1.Person.AltNames, e2.Person.AltNames)

		require.Equal(tb, e1.Person.Gender, e2.Person.Gender)

		require.Equal(tb, e1.Person.BirthDate, e2.Person.BirthDate)
		require.Equal(tb, e1.Person.PlaceOfBirth, e2.Person.PlaceOfBirth)
		require.Equal(tb, e1.Person.DeathDate, e2.Person.DeathDate)

		require.ElementsMatch(tb, e1.Person.Titles, e2.Person.Titles)
		require.ElementsMatch(tb, e1.Person.GovernmentIDs, e2.Person.GovernmentIDs)
	}

	if e1.Business != nil && e2.Business != nil {
		require.Equal(tb, e1.Business.Name, e2.Business.Name)
		require.ElementsMatch(tb, e1.Business.AltNames, e2.Business.AltNames)

		require.Equal(tb, e1.Business.Created, e2.Business.Created)
		require.Equal(tb, e1.Business.Dissolved, e2.Business.Dissolved)

		require.ElementsMatch(tb, e1.Business.GovernmentIDs, e2.Business.GovernmentIDs)
	}

	if e1.Organization != nil && e2.Organization != nil {
		require.Equal(tb, e1.Organization.Name, e2.Organization.Name)
		require.ElementsMatch(tb, e1.Organization.AltNames, e2.Organization.AltNames)

		require.Equal(tb, e1.Organization.Created, e2.Organization.Created)
		require.Equal(tb, e1.Organization.Dissolved, e2.Organization.Dissolved)

		require.ElementsMatch(tb, e1.Organization.GovernmentIDs, e2.Organization.GovernmentIDs)
	}

	if e1.Aircraft != nil && e2.Aircraft != nil {
		require.Equal(tb, e1.Aircraft.Name, e2.Aircraft.Name)
		require.ElementsMatch(tb, e1.Aircraft.AltNames, e2.Aircraft.AltNames)

		require.Equal(tb, e1.Aircraft.Type, e2.Aircraft.Type)
		require.Equal(tb, e1.Aircraft.Flag, e2.Aircraft.Flag)
		require.Equal(tb, e1.Aircraft.Built, e2.Aircraft.Built)
		require.Equal(tb, e1.Aircraft.ICAOCode, e2.Aircraft.ICAOCode)
		require.Equal(tb, e1.Aircraft.Model, e2.Aircraft.Model)
		require.Equal(tb, e1.Aircraft.SerialNumber, e2.Aircraft.SerialNumber)
	}

	if e1.Vessel != nil && e2.Vessel != nil {
		require.Equal(tb, e1.Vessel.Name, e2.Vessel.Name)
		require.ElementsMatch(tb, e1.Vessel.AltNames, e2.Vessel.AltNames)

		require.Equal(tb, e1.Vessel.IMONumber, e2.Vessel.IMONumber)
		require.Equal(tb, e1.Vessel.Type, e2.Vessel.Type)
		require.Equal(tb, e1.Vessel.Flag, e2.Vessel.Flag)
		require.Equal(tb, e1.Vessel.Built, e2.Vessel.Built)
		require.Equal(tb, e1.Vessel.Model, e2.Vessel.Model)
		require.Equal(tb, e1.Vessel.Tonnage, e2.Vessel.Tonnage)
		require.Equal(tb, e1.Vessel.MMSI, e2.Vessel.MMSI)
		require.Equal(tb, e1.Vessel.CallSign, e2.Vessel.CallSign)
		require.Equal(tb, e1.Vessel.GrossRegisteredTonnage, e2.Vessel.GrossRegisteredTonnage)
		require.Equal(tb, e1.Vessel.Owner, e2.Vessel.Owner)
	}

	require.ElementsMatch(tb, e1.Contact.EmailAddresses, e2.Contact.EmailAddresses)
	require.ElementsMatch(tb, e1.Contact.PhoneNumbers, e2.Contact.PhoneNumbers)
	require.ElementsMatch(tb, e1.Contact.FaxNumbers, e2.Contact.FaxNumbers)
	require.ElementsMatch(tb, e1.Contact.Websites, e2.Contact.Websites)

	require.ElementsMatch(tb, e1.Addresses, e2.Addresses)
	require.ElementsMatch(tb, e1.CryptoAddresses, e2.CryptoAddresses)

	require.ElementsMatch(tb, e1.Affiliations, e2.Affiliations)
	require.Equal(tb, e1.SanctionsInfo, e2.SanctionsInfo)
	require.ElementsMatch(tb, e1.HistoricalInfo, e2.HistoricalInfo)

	// require.Equal(tb, e1.PreparedFields, e2.PreparedFields) // TODO(adam): want to check these?
	require.Equal(tb, e1.SourceData, e2.SourceData)
}
