package ofac

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMapper__Vessel(t *testing.T) {
	res, err := Read(testInputs(t, filepath.Join("testdata", "sdn.csv")))
	require.NoError(t, err)

	var sdn *SDN
	for i := range res.SDNs {
		if res.SDNs[i].EntityID == "15036" {
			sdn = &res.SDNs[i]
		}
	}
	require.NotNil(t, sdn)

	e := ToEntity(*sdn, nil, nil, nil)
	require.Equal(t, "ARTAVIL", e.Name)
	require.Equal(t, search.EntityVessel, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.Nil(t, e.Person)
	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
	require.NotNil(t, e.Vessel)

	require.Equal(t, "ARTAVIL", e.Vessel.Name)
	require.Equal(t, "Iran", e.Vessel.Flag)
	require.Equal(t, "9187629", e.Vessel.IMONumber)
	require.Equal(t, "572469210", e.Vessel.MMSI)

	sourceData, ok := e.SourceData.(SDN)
	require.True(t, ok)
	require.Equal(t, "15036", sourceData.EntityID)
}

func TestMapper__Aircraft(t *testing.T) {
	res, err := Read(testInputs(t, filepath.Join("testdata", "sdn.csv")))
	require.NoError(t, err)

	var sdn *SDN
	for i := range res.SDNs {
		if res.SDNs[i].EntityID == "18158" {
			sdn = &res.SDNs[i]
		}
	}
	require.NotNil(t, sdn)

	e := ToEntity(*sdn, nil, nil, nil)
	require.Equal(t, "MSN 550", e.Name)
	require.Equal(t, search.EntityAircraft, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.Nil(t, e.Person)
	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.NotNil(t, e.Aircraft)
	require.Nil(t, e.Vessel)

	require.Equal(t, "MSN 550", e.Aircraft.Name)
	require.Equal(t, "1995-01-01", e.Aircraft.Built.Format(time.DateOnly))
	require.Equal(t, "Airbus A321-131", e.Aircraft.Model)
	require.Equal(t, "550", e.Aircraft.SerialNumber)

	sourceData, ok := e.SourceData.(SDN)
	require.True(t, ok)
	require.Equal(t, "18158", sourceData.EntityID)
}

func TestMapper__CompleteVessel(t *testing.T) {
	sdn := &SDN{
		EntityID: "67890",
		SDNName:  "CARGO VESSEL X",
		SDNType:  "vessel",
		Remarks:  "Vessel Type Cargo; Other Vessel Flag Malta; IMO 9999999; MMSI 123456789; Tonnage 50,000",
	}

	e := ToEntity(*sdn, nil, nil, nil)
	require.Equal(t, "CARGO VESSEL X", e.Name)
	require.Equal(t, search.EntityVessel, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Vessel)
	require.Equal(t, "CARGO VESSEL X", e.Vessel.Name)
	require.Equal(t, search.VesselTypeCargo, e.Vessel.Type)
	require.Equal(t, "Malta", e.Vessel.Flag)
	require.Equal(t, "9999999", e.Vessel.IMONumber)
	require.Equal(t, "123456789", e.Vessel.MMSI)
	require.Equal(t, 50000, e.Vessel.Tonnage)

	// Verify other entity types are nil
	require.Nil(t, e.Person)
	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
}

func TestMapper__CompleteAircraft(t *testing.T) {
	sdn := &SDN{
		EntityID: "54321",
		SDNName:  "AIRCRAFT Y",
		SDNType:  "aircraft",
		Remarks:  "Aircraft Type Cargo; Flag United States; Aircraft Model Boeing 747; Manufacture Date 01 Jan 1995; Serial Number (MSN) 12345; ICAO Code B744",
	}

	e := ToEntity(*sdn, nil, nil, nil)
	require.Equal(t, "AIRCRAFT Y", e.Name)
	require.Equal(t, search.EntityAircraft, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Aircraft)
	require.Equal(t, "AIRCRAFT Y", e.Aircraft.Name)
	require.Equal(t, search.AircraftCargo, e.Aircraft.Type)
	require.Equal(t, "United States", e.Aircraft.Flag)
	require.Equal(t, "Boeing 747", e.Aircraft.Model)
	require.Equal(t, "1995-01-01", e.Aircraft.Built.Format(time.DateOnly))
	require.Equal(t, "12345", e.Aircraft.SerialNumber)
	require.Equal(t, "B744", e.Aircraft.ICAOCode)

	// Verify other entity types are nil
	require.Nil(t, e.Person)
	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Vessel)
}

func TestNormalizeVesselType(t *testing.T) {
	tests := []struct {
		input    string
		expected search.VesselType
	}{
		{"cargo", search.VesselTypeCargo},
		{"Cargo", search.VesselTypeCargo},
		{"CARGO", search.VesselTypeCargo},
		{"unknown", search.VesselTypeUnknown},
		{"", search.VesselTypeUnknown},
	}

	for _, tt := range tests {
		result := normalizeVesselType(tt.input)
		require.Equal(t, tt.expected, result)
	}
}

func TestNormalizeAircraftType(t *testing.T) {
	tests := []struct {
		input    string
		expected search.AircraftType
	}{
		{"cargo", search.AircraftCargo},
		{"Cargo", search.AircraftCargo},
		{"CARGO", search.AircraftCargo},
		{"unknown", search.AircraftTypeUnknown},
		{"", search.AircraftTypeUnknown},
	}

	for _, tt := range tests {
		result := normalizeAircraftType(tt.input)
		require.Equal(t, tt.expected, result)
	}
}

func TestMapper__CompleteVesselWithAllFields(t *testing.T) {
	sdn := &SDN{
		EntityID:               "67890",
		SDNName:                "CARGO VESSEL X",
		SDNType:                "vessel",
		CallSign:               "ABC123",
		GrossRegisteredTonnage: "25,000",
		VesselOwner:            "SHIPPING CORP",
		Remarks:                "Vessel Type Cargo; Flag Malta; IMO 9999999; MMSI 123456789; Tonnage 50,000",
	}

	e := ToEntity(*sdn, nil, nil, nil)
	require.Equal(t, "CARGO VESSEL X", e.Name)
	require.Equal(t, search.EntityVessel, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Vessel)
	require.Equal(t, "CARGO VESSEL X", e.Vessel.Name)
	require.Equal(t, search.VesselTypeCargo, e.Vessel.Type)
	require.Equal(t, "Malta", e.Vessel.Flag)
	require.Equal(t, "9999999", e.Vessel.IMONumber)
	require.Equal(t, "123456789", e.Vessel.MMSI)
	require.Equal(t, 50000, e.Vessel.Tonnage)

	// Test new fields
	require.Equal(t, "ABC123", e.Vessel.CallSign)
	require.Equal(t, 25000, e.Vessel.GrossRegisteredTonnage)
	require.Equal(t, "SHIPPING CORP", e.Vessel.Owner)
}
