package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnglishLegalForm_LtdVsLimited(t *testing.T) {
	ltd := Entity[Value]{
		Name:     "Harrowfield Bearings Limited",
		Type:     EntityBusiness,
		Business: &Business{Name: "Harrowfield Bearings Limited"},
	}.Normalize()
	short := Entity[Value]{
		Name:     "HARROWFIELD BEARINGS LTD.",
		Type:     EntityBusiness,
		Business: &Business{Name: "HARROWFIELD BEARINGS LTD."},
	}.Normalize()

	require.Equal(t, "harrowfield bearings ltd", ltd.PreparedFields.Name)
	require.Equal(t, ltd.PreparedFields.Name, short.PreparedFields.Name)
	got := Similarity(ltd, short)
	require.InDelta(t, 0.855, got, 0.001)
	require.InDelta(t, got, scoreSimilarityFast(ltd, short, SimilarityOpts{}), 0.001)
}

func TestEnglishLegalForm_GmbHVsSdnBhdStaysDistinct(t *testing.T) {
	gmbh := Entity[Value]{
		Name:     "Korvessa Industrial GmbH",
		Type:     EntityBusiness,
		Business: &Business{Name: "Korvessa Industrial GmbH"},
	}.Normalize()
	sdn := Entity[Value]{
		Name:     "Korvessa Industrial Sdn. Bhd.",
		Type:     EntityBusiness,
		Business: &Business{Name: "Korvessa Industrial Sdn. Bhd."},
	}.Normalize()

	require.Contains(t, gmbh.PreparedFields.Name, "gmbh")
	require.Contains(t, sdn.PreparedFields.Name, "sdn")
	require.NotEqual(t, gmbh.PreparedFields.Name, sdn.PreparedFields.Name)
}
