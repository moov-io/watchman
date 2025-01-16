package ofac_test

import (
	"sort"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/ofac"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMapperBusiness__FromSource(t *testing.T) {
	t.Run("33151 - crypto addresses", func(t *testing.T) {
		found := ofactest.FindEntity(t, "33151")
		require.Equal(t, "SUEX OTC, S.R.O.", found.Name)
		require.Equal(t, search.EntityBusiness, found.Type)
		require.Equal(t, search.SourceUSOFAC, found.Source)
		require.Equal(t, "33151", found.SourceID)

		require.Nil(t, found.Person)
		require.NotNil(t, found.Business)
		require.Nil(t, found.Organization)
		require.Nil(t, found.Aircraft)
		require.Nil(t, found.Vessel)

		business := found.Business
		require.Equal(t, "SUEX OTC, S.R.O.", business.Name)
		require.ElementsMatch(t, []string{"SUCCESSFUL EXCHANGE"}, business.AltNames)

		createdAt := time.Date(2018, time.September, 25, 0, 0, 0, 0, time.UTC)
		require.Equal(t, createdAt.Format(time.RFC3339), business.Created.Format(time.RFC3339))
		require.Nil(t, business.Dissolved)

		expectedGovernmentIDs := []search.GovernmentID{
			{Type: search.GovernmentIDBusinessRegisration, Country: "Czech Republic", Identifier: "07486049"},
			{Type: search.GovernmentIDBusinessRegisration, Country: "Czech Republic", Identifier: "5299007NTWCC3U23WM81"},
		}
		require.ElementsMatch(t, expectedGovernmentIDs, business.GovernmentIDs)

		expectedContact := search.ContactInfo{
			Websites: []string{"suex.io"},
		}
		require.Equal(t, expectedContact, found.Contact)
		require.Empty(t, found.Addresses)

		expectedCryptoAddresses := []search.CryptoAddress{
			{Currency: "XBT", Address: "12HQDsicffSBaYdJ6BhnE22sfjTESmmzKx"},
			{Currency: "XBT", Address: "1L4ncif9hh9TnUveqWq77HfWWt6CJWtrnb"},
			{Currency: "XBT", Address: "13mnk8SvDGqsQTHbiGiHBXqtaQCUKfcsnP"},
			{Currency: "XBT", Address: "1Edue8XZCWNoDBNZgnQkCCivDyr9GEo4x6"},
			{Currency: "XBT", Address: "1ECeZBxCVJ8Wm2JSN3Cyc6rge2gnvD3W5K"},
			{Currency: "XBT", Address: "1J9oGoAiHeRfeMZeUnJ9W7RpV55CdKtgYE"},
			{Currency: "XBT", Address: "1295rkVyNfFpqZpXvKGhDqwhP1jZcNNDMV"},
			{Currency: "XBT", Address: "1LiNmTUPSJEd92ZgVJjAV3RT9BzUjvUCkx"},
			{Currency: "XBT", Address: "1LrxsRd7zNuxPJcL5rttnoeJFy1y4AffYY"},
			{Currency: "XBT", Address: "1KUUJPkyDhamZXgpsyXqNGc3x1QPXtdhgz"},
			{Currency: "XBT", Address: "1CF46Rfbp97absrs7zb7dFfZS6qBXUm9EP"},
			{Currency: "XBT", Address: "1Df883c96LVauVsx9FEgnsourD8DELwCUQ"},
			{Currency: "XBT", Address: "bc1qdt3gml5z5n50y5hm04u2yjdphefkm0fl2zdj68"},
			{Currency: "XBT", Address: "1B64QRxfaa35MVkf7sDjuGUYAP5izQt7Qi"},
			{Currency: "ETH", Address: "0x2f389ce8bd8ff92de3402ffce4691d17fc4f6535"},
			{Currency: "ETH", Address: "0x19aa5fe80d33a56d56c78e82ea5e50e5d80b4dff"},
			{Currency: "ETH", Address: "0xe7aa314c77f4233c18c6cc84384a9247c0cf367b"},
			{Currency: "ETH", Address: "0x308ed4b7b49797e1a98d3818bff6fe5385410370"},
			{Currency: "USDT", Address: "0x2f389ce8bd8ff92de3402ffce4691d17fc4f6535"},
			{Currency: "USDT", Address: "0x19aa5fe80d33a56d56c78e82ea5e50e5d80b4dff"},
			{Currency: "USDT", Address: "1KUUJPkyDhamZXgpsyXqNGc3x1QPXtdhgz"},
			{Currency: "USDT", Address: "1CF46Rfbp97absrs7zb7dFfZS6qBXUm9EP"},
			{Currency: "USDT", Address: "1LrxsRd7zNuxPJcL5rttnoeJFy1y4AffYY"},
			{Currency: "USDT", Address: "1Df883c96LVauVsx9FEgnsourD8DELwCUQ"},
			{Currency: "USDT", Address: "16iWn2J1McqjToYLHSsAyS6En3QA8YQ91H"},
		}
		require.ElementsMatch(t, expectedCryptoAddresses, found.CryptoAddresses)

		require.Empty(t, found.Affiliations)
		require.Nil(t, found.SanctionsInfo)
		require.Empty(t, found.HistoricalInfo)

		sdn, ok := found.SourceData.(ofac.SDN)
		require.True(t, ok)
		require.Equal(t, "33151", sdn.EntityID)
		require.Equal(t, "SUEX OTC, S.R.O.", sdn.SDNName)
	})

	t.Run("12685", func(t *testing.T) {
		found := ofactest.FindEntity(t, "12685")
		require.Equal(t, "GADDAFI INTERNATIONAL CHARITY AND DEVELOPMENT FOUNDATION", found.Name)

		expectedContact := search.ContactInfo{
			EmailAddresses: []string{"info@gicdf.org"},
			PhoneNumbers:   []string{"(218) (0)214778301", "(022) 7363030"},
			FaxNumbers:     []string{"(218) (0)214778766", "(022) 7363196"},
			Websites:       []string{"www.gicdf.org"},
		}
		require.Equal(t, expectedContact, found.Contact)
	})

	t.Run("44525", func(t *testing.T) {
		found := ofactest.FindEntity(t, "44525")
		require.Equal(t, "BEL-KAP-STEEL LLC", found.Name)

		t.Logf("%#v", found.Business)

		expectedGovernmentIDs := []search.GovernmentID{
			{Type: "tax-id", Country: "United States", Identifier: "52-2083095"},
			{Type: "business-registration", Country: "Connecticut", Identifier: "0582030"},
			{Type: "business-registration", Country: "Florida", Identifier: "M99000000961"},
		}
		require.ElementsMatch(t, expectedGovernmentIDs, found.Business.GovernmentIDs)
	})

	t.Run("50544", func(t *testing.T) {
		found := ofactest.FindEntity(t, "50544")

		require.Equal(t, "AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG REGIONS", found.Name)
		require.Equal(t, search.EntityBusiness, found.Type)
		require.Equal(t, search.SourceUSOFAC, found.Source)
		require.Equal(t, "50544", found.SourceID)

		require.Nil(t, found.Person)
		require.NotNil(t, found.Business)
		require.Nil(t, found.Organization)
		require.Nil(t, found.Aircraft)
		require.Nil(t, found.Vessel)

		business := found.Business
		require.Equal(t, "AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG REGIONS", business.Name)

		expectedAltNames := []string{
			"DIALOGUE REGIONS",
			"DIALOG REGIONY",
			"DIALOGUE",
			"AUTONOMOUS NON-PROFIT ORGANIZATION FOR THE DEVELOPMENT OF DIGITAL PROJECTS IN THE FIELD OF PUBLIC RELATIONS AND COMMUNICATIONS DIALOG REGIONS",
			"ANO DIALOG REGIONS",
			"AVTONOMNAYA NEKOMMERCHESKAYA ORGANIZATSIYA PO RAZVITIYU TSIFROVYKH PROEKTOV V SFERE OBSHCHESTEVENNYKH SVYAZEI I KOMMUNIKATSII DIALOG REGIONY",
		}
		require.ElementsMatch(t, expectedAltNames, business.AltNames)

		createdAt := time.Date(2020, time.July, 21, 0, 0, 0, 0, time.UTC)
		require.Equal(t, createdAt.Format(time.RFC3339), business.Created.Format(time.RFC3339))
		require.Nil(t, business.Dissolved)

		expectedGovernmentIDs := []search.GovernmentID{
			{Type: search.GovernmentIDBusinessRegisration, Country: "Russia", Identifier: "1207700248030"},
			{Type: search.GovernmentIDTax, Country: "Russia", Identifier: "9709063550"},
		}
		require.ElementsMatch(t, expectedGovernmentIDs, business.GovernmentIDs)

		expectedContact := search.ContactInfo{
			Websites: []string{"www.dialog.info", "www.dialog-regions.ru"},
		}
		require.Equal(t, expectedContact, found.Contact)
		require.Empty(t, found.Addresses)
		require.Empty(t, found.CryptoAddresses)

		expectedAffiliations := []search.Affiliation{
			{EntityName: "AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG", Type: "Linked To", Details: ""},
		}
		require.ElementsMatch(t, expectedAffiliations, found.Affiliations)

		require.Nil(t, found.SanctionsInfo)
		require.Empty(t, found.HistoricalInfo)

		sdn, ok := found.SourceData.(ofac.SDN)
		require.True(t, ok)
		require.Equal(t, "50544", sdn.EntityID)
		require.Equal(t, "AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG REGIONS", sdn.SDNName)

	})
}

func TestMapper__CompleteBusiness(t *testing.T) {
	sdn := &ofac.SDN{
		EntityID: "12345",
		SDNName:  "ACME CORPORATION",
		SDNType:  "-0-",
		Remarks:  "Business Registration Number 51566843 (Hong Kong); Commercial Registry Number CH-020.1.066.499-9 (Switzerland); Company Number 05527424 (United Kingdom)",
	}

	e := ofac.ToEntity(*sdn, nil, nil, nil)
	require.Equal(t, "ACME CORPORATION", e.Name)
	require.Equal(t, search.EntityBusiness, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Business)
	require.Equal(t, "ACME CORPORATION", e.Business.Name)
	require.Len(t, e.Business.GovernmentIDs, 3)

	// Sort the identifiers to ensure consistent ordering for tests
	govIDs := e.Business.GovernmentIDs
	sort.Slice(govIDs, func(i, j int) bool {
		return govIDs[i].Country < govIDs[j].Country
	})

	// Verify identifiers
	require.Equal(t, "Hong Kong", govIDs[0].Country)
	require.Equal(t, search.GovernmentIDBusinessRegisration, govIDs[0].Type)
	require.Equal(t, "51566843", govIDs[0].Identifier)

	require.Equal(t, "Switzerland", govIDs[1].Country)
	require.Equal(t, search.GovernmentIDCommercialRegistry, govIDs[1].Type)
	require.Equal(t, "CH-020.1.066.499-9", govIDs[1].Identifier)

	require.Equal(t, "United Kingdom", govIDs[2].Country)
	require.Equal(t, search.GovernmentIDBusinessRegisration, govIDs[2].Type)
	require.Equal(t, "05527424", govIDs[2].Identifier)

	// Verify other entity types are nil
	require.Nil(t, e.Person)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
	require.Nil(t, e.Vessel)
}

func TestMapper__CompleteBusinessWithRemarks(t *testing.T) {
	sdn := &ofac.SDN{
		EntityID: "12345",
		SDNName:  "ACME CORPORATION",
		SDNType:  "-0-",
		Remarks:  "Business Registration Number 51566843 (Hong Kong); Subsidiary Of: PARENT CORP; Former Name: OLD ACME LTD; Additional Sanctions Information - Subject to Secondary Sanctions",
	}

	e := ofac.ToEntity(*sdn, nil, nil, nil)

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
