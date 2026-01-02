package senzing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSenzing_US_Congress(t *testing.T) {
	fd, err := os.Open(filepath.Join("..", "..", "..", "test", "testdata", "opensanctions", "us_congress_senzing.json"))
	require.NoError(t, err)

	t.Cleanup(func() { fd.Close() })

	entities, err := ReadEntities(fd, search.SourceList("opensanctions_us_congress"))
	require.NoError(t, err)
	require.Len(t, entities, 1319)

	found := entities[13]
	require.Equal(t, "Michael Lawler", found.Name)
	require.Equal(t, search.EntityPerson, found.Type)
	require.Equal(t, search.SourceList("opensanctions_us_congress"), found.Source)
	require.Equal(t, "Q105179052", found.SourceID)

	require.NotNil(t, found.Person)
	require.Nil(t, found.Business)
	require.Nil(t, found.Organization)
	require.Nil(t, found.Aircraft)
	require.Nil(t, found.Vessel)

	require.Equal(t, 1986, found.Person.BirthDate.Year()) // no other data in file
}
