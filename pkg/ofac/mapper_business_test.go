package ofac

import (
	"sort"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMapper__CompleteBusiness(t *testing.T) {
	sdn := &SDN{
		EntityID: "12345",
		SDNName:  "ACME CORPORATION",
		SDNType:  "-0-",
		Remarks:  "Business Registration Number 51566843 (Hong Kong); Commercial Registry Number CH-020.1.066.499-9 (Switzerland); Company Number 05527424 (United Kingdom)",
	}

	e := ToEntity(*sdn, nil, nil)
	require.Equal(t, "ACME CORPORATION", e.Name)
	require.Equal(t, search.EntityBusiness, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Business)
	require.Equal(t, "ACME CORPORATION", e.Business.Name)
	require.Len(t, e.Business.Identifier, 3)

	// Sort the identifiers to ensure consistent ordering for tests
	identifiers := e.Business.Identifier
	sort.Slice(identifiers, func(i, j int) bool {
		return identifiers[i].Country < identifiers[j].Country
	})

	// Verify identifiers
	require.Equal(t, "Hong Kong", identifiers[0].Country)
	require.Equal(t, "Business Registration Number", identifiers[0].Name)
	require.Equal(t, "51566843", identifiers[0].Identifier)

	require.Equal(t, "Switzerland", identifiers[1].Country)
	require.Equal(t, "Commercial Registry Number", identifiers[1].Name)
	require.Equal(t, "CH-020.1.066.499-9", identifiers[1].Identifier)

	require.Equal(t, "United Kingdom", identifiers[2].Country)
	require.Equal(t, "Company Number", identifiers[2].Name)
	require.Equal(t, "05527424", identifiers[2].Identifier)

	// Verify other entity types are nil
	require.Nil(t, e.Person)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
	require.Nil(t, e.Vessel)
}

func TestMapper__CompleteBusinessWithRemarks(t *testing.T) {
	sdn := &SDN{
		EntityID: "12345",
		SDNName:  "ACME CORPORATION",
		SDNType:  "-0-",
		Remarks:  "Business Registration Number 51566843 (Hong Kong); Subsidiary Of: PARENT CORP; Former Name: OLD ACME LTD; Additional Sanctions Information - Subject to Secondary Sanctions",
	}

	e := ToEntity(*sdn, nil, nil)

	// Test affiliations
	require.Len(t, e.Affiliations, 1)
	require.Equal(t, "PARENT CORP", e.Affiliations[0].EntityName)
	require.Equal(t, "Subsidiary Of", e.Affiliations[0].Type)

	// Test sanctions info
	require.NotNil(t, e.SanctionsInfo)
	require.True(t, e.SanctionsInfo.Secondary)

	// Test historical info
	require.Len(t, e.HistoricalInfo, 1)
	require.Equal(t, "Former Name", e.HistoricalInfo[0].Type)
	require.Equal(t, "OLD ACME LTD", e.HistoricalInfo[0].Value)
}
