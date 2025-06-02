// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/cslustest"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/csl_us"

	"github.com/stretchr/testify/require"
)

func TestMapBusiness(t *testing.T) {
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
