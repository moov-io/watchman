package ofac

import (
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func testInputs(tb testing.TB, paths ...string) map[string]io.ReadCloser {
	tb.Helper()

	input := make(map[string]io.ReadCloser)
	for _, path := range paths {
		_, filename := filepath.Split(path)

		fd, err := os.Open(path)
		require.NoError(tb, err)

		input[filename] = fd
	}
	return input
}

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

	e := ToEntity(*sdn)
	require.Equal(t, "MORENO, Daniel", e.Name)
	require.Equal(t, search.EntityPerson, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.NotNil(t, e.Person)
	require.Equal(t, "MORENO, Daniel", e.Person.Name)
	require.Equal(t, "", string(e.Person.Gender))
	require.Equal(t, "1972-10-12T00:00:00Z", e.Person.BirthDate.Format(time.RFC3339))
	require.Nil(t, e.Person.DeathDate)
	require.Len(t, e.Person.GovernmentIDs, 0)

	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
	require.Nil(t, e.Vessel)

	require.Equal(t, "15102", e.SourceData.EntityID)
}

func TestMapper__Vessel(t *testing.T) {
	res, err := Read(testInputs(t, filepath.Join("..", "..", "test", "testdata", "sdn.csv")))
	require.NoError(t, err)

	var sdn *SDN
	for i := range res.SDNs {
		if res.SDNs[i].EntityID == "15036" {
			sdn = res.SDNs[i]
		}
	}
	require.NotNil(t, sdn)

	e := ToEntity(*sdn)
	require.Equal(t, "ARTAVIL", e.Name)
	require.Equal(t, search.EntityVessel, e.Type)
	require.Equal(t, search.SourceUSOFAC, e.Source)

	require.Nil(t, e.Person)
	require.Nil(t, e.Business)
	require.Nil(t, e.Organization)
	require.Nil(t, e.Aircraft)
	require.NotNil(t, e.Vessel)

	require.Equal(t, "ARTAVIL", e.Vessel.Name)
	require.Equal(t, "Malta", e.Vessel.Flag)
	require.Equal(t, "9187629", e.Vessel.IMONumber)
	require.Equal(t, "572469210", e.Vessel.MMSI)

	require.Equal(t, "15036", e.SourceData.EntityID)
}

func TestMapper__Aircraft(t *testing.T) {
	res, err := Read(testInputs(t, filepath.Join("..", "..", "test", "testdata", "sdn.csv")))
	require.NoError(t, err)

	var sdn *SDN
	for i := range res.SDNs {
		if res.SDNs[i].EntityID == "18158" {
			sdn = res.SDNs[i]
		}
	}
	require.NotNil(t, sdn)

	e := ToEntity(*sdn)
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

	require.Equal(t, "18158", e.SourceData.EntityID)
}

func TestParseTime(t *testing.T) {
	t.Run("DOB", func(t *testing.T) {
		tt, _ := parseTime(dobPatterns, "01 Apr 1950")
		require.Equal(t, "1950-04-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "01 Feb 1958 to 28 Feb 1958")
		require.Equal(t, "1958-02-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "1928")
		require.Equal(t, "1928-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "1928 to 1930")
		require.Equal(t, "1928-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "Sep 1958")
		require.Equal(t, "1958-09-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "circa 01 Jan 1961")
		require.Equal(t, "1961-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "circa 1934")
		require.Equal(t, "1934-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "circa 1979-1982")
		require.Equal(t, "1979-01-01", tt.Format(time.DateOnly))
	})
}
