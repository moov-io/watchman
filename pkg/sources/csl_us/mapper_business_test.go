// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us_test

import (
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/cslustest"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/csl_us"

	"github.com/stretchr/testify/require"
)

func TestMapBusiness_Sinobright(t *testing.T) {
	entity := cslustest.FindEntity(t, "233a4d725770c81fb561ffe3842c14010c2201971d6be62eca1e613b")

	require.Equal(t, "Sinobright Import and Export Company", entity.Name)
	require.Equal(t, search.EntityBusiness, entity.Type)
	require.Equal(t, search.SourceUSCSL, entity.Source)
	require.Equal(t, "233a4d725770c81fb561ffe3842c14010c2201971d6be62eca1e613b", entity.SourceID)

	require.Nil(t, entity.Person)
	require.NotNil(t, entity.Business)
	require.Nil(t, entity.Organization)
	require.Nil(t, entity.Aircraft)
	require.Nil(t, entity.Vessel)

	require.Equal(t, search.ContactInfo{}, entity.Contact)
	require.Empty(t, entity.Addresses)
	require.Empty(t, entity.CryptoAddresses)

	// Sanctions Info
	require.NotNil(t, entity.SanctionsInfo)
	expected := []string{"INKSNA"}
	require.ElementsMatch(t, expected, entity.SanctionsInfo.Programs)

	// Prepared fields
	require.Equal(t, "sinobright import and export company", entity.PreparedFields.Name)

	expected = []string{"sinobright", "import", "export", "company"}
	require.ElementsMatch(t, expected, entity.PreparedFields.NameFields)

	// Source data
	us, ok := entity.SourceData.(csl_us.SanctionsEntry)
	require.True(t, ok)

	require.Equal(t, "233a4d725770c81fb561ffe3842c14010c2201971d6be62eca1e613b", us.ID)
	require.Equal(t, "Nonproliferation Sanctions (ISN) - State Department", us.Source)
	require.Equal(t, "INKSNA", us.Programs)
	require.Equal(t, "Sinobright Import and Export Company; and any successor, sub-unit, or subsidiary thereof", us.Name)
}

func TestMapBusiness_ChinaElectronicsTechnologyGroup(t *testing.T) {
	entity := cslustest.FindEntity(t, "99ed1b052fdb09695fae1ba87516d8a0882b0df04652e07aeb5ce7be")

	require.Equal(t, "China Electronics Technology Group Corporation 13th Research Institute", entity.Name)
	require.Equal(t, search.EntityBusiness, entity.Type)
	require.Equal(t, search.SourceUSCSL, entity.Source)
	require.Equal(t, "99ed1b052fdb09695fae1ba87516d8a0882b0df04652e07aeb5ce7be", entity.SourceID)

	require.Nil(t, entity.Person)
	require.NotNil(t, entity.Business)
	require.Nil(t, entity.Organization)
	require.Nil(t, entity.Aircraft)
	require.Nil(t, entity.Vessel)

	// Business Fields
	require.Equal(t, "China Electronics Technology Group Corporation 13th Research Institute", entity.Business.Name)

	expected := []string{
		"CETC 13",
		"Hebei Semiconductor Research Institute",
		"HSRI",
		"Hebei Institute of Semiconductors",
		"Hebei Semiconductor Institute",
		"Hebei Semiconductor",
		"CETC Research Institute 13",
	}
	require.ElementsMatch(t, expected, entity.Business.AltNames)

	createdAt := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC)
	require.Equal(t, createdAt, *entity.Business.Created)

	// Misc Common Fields
	require.Equal(t, search.ContactInfo{}, entity.Contact)
	require.Empty(t, entity.CryptoAddresses)

	expectedAddresses := []search.Address{
		{Line1: "113 Hezuo Road", City: "Hebei", Country: "China"},
		{Line1: "21 Changsheng Street", City: "Hebei", Country: "China"},
		{Line1: "21 Changsheng Road", Line2: "Shijiazhuang Hebei China", City: "Hebei Province", Country: "China"},
	}
	require.ElementsMatch(t, expectedAddresses, entity.Addresses)

	// Sanctions Info
	require.NotNil(t, entity.SanctionsInfo)
	require.Empty(t, entity.SanctionsInfo.Programs)

	// Prepared fields
	require.Equal(t, "china electronics technology group corporation 13th research institute", entity.PreparedFields.Name)

	expected = []string{"china", "electronics", "technology", "group", "corporation", "13th", "research", "institute"}
	require.ElementsMatch(t, expected, entity.PreparedFields.NameFields)

	// Source data
	us, ok := entity.SourceData.(csl_us.SanctionsEntry)
	require.True(t, ok)

	require.Equal(t, "99ed1b052fdb09695fae1ba87516d8a0882b0df04652e07aeb5ce7be", us.ID)
	require.Equal(t, "Entity List (EL) - Bureau of Industry and Security", us.Source)
	require.Empty(t, us.Programs)
	require.Equal(t, "China Electronics Technology Group Corporation 13th Research Institute (CETC 13)", us.Name)
}
