package ofac

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapper__SwiftBICFromTestdata(t *testing.T) {
	res, err := Read(testInputs(t, filepath.Join("testdata", "sdn.csv")))
	require.NoError(t, err)

	want := map[string][]string{
		"906":   {"HAVIGB2L"},
		"12028": {"HNRBBY2X"},
		"17016": {"GAZPRUMM"},
		"19033": {"VECARU21", "VVBKRU2Y"},
		"25417": {"AFABAFKA"},
	}

	found := 0
	withBIC := 0
	for i := range res.SDNs {
		e := ToEntity(res.SDNs[i], nil, nil, nil)
		if len(e.SwiftBICs) > 0 {
			withBIC++
		}
		expect, ok := want[e.SourceID]
		if !ok {
			continue
		}
		require.Equal(t, expect, e.SwiftBICs, e.SourceID)
		require.Equal(t, expect, e.PreparedFields.SwiftBICs, e.SourceID)
		found++
	}
	require.Equal(t, len(want), found)
	require.Greater(t, withBIC, 50)
}
