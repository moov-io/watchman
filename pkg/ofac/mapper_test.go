package ofac

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMapper(t *testing.T) {
	res, err := Read(filepath.Join("..", "..", "test", "testdata", "sdn.csv"))
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
