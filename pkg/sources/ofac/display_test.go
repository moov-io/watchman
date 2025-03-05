package ofac

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetailsURL(t *testing.T) {
	got := DetailsURL("13058")
	expected := "https://sanctionssearch.ofac.treas.gov/Details.aspx?id=13058"
	require.Equal(t, expected, got)
}
