package cslustest_test

import (
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/cslustest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestCSLUSTest_FindEntity(t *testing.T) {
	found := cslustest.FindEntity(t, "9651")

	require.Equal(t, "Fathi Mohammed QAR'AWI", found.Name)
	require.Equal(t, search.EntityPerson, found.Type)
	require.Equal(t, "9651", found.SourceID)

	require.NotNil(t, found.Person)
	require.Nil(t, found.Business)

	require.Equal(t, "Fathi Mohammed QAR'AWI", found.Person.Name)

	altNames := []string{"Mohammed Fathi QARAWI"}
	require.ElementsMatch(t, altNames, found.Person.AltNames)

	birthDate := time.Date(1958, time.January, 1, 0, 0, 0, 0, time.UTC)
	require.Equal(t, birthDate.Format(time.RFC3339), found.Person.BirthDate.Format(time.RFC3339))

	require.NotNil(t, found.SanctionsInfo)

	programs := []string{"NS-PLC"}
	require.ElementsMatch(t, programs, found.SanctionsInfo.Programs)
}
