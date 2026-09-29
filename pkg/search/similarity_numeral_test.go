package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNameNumeralConflict_NumberedSPV(t *testing.T) {
	three := Entity[Value]{
		Name:     "Brasskey Shipping Three Inc.",
		Type:     EntityBusiness,
		Business: &Business{Name: "Brasskey Shipping Three Inc."},
	}.Normalize()
	five := Entity[Value]{
		Name:     "Brasskey Shipping Five Inc.",
		Type:     EntityBusiness,
		Business: &Business{Name: "Brasskey Shipping Five Inc."},
	}.Normalize()

	require.True(t, hasNameNumeralConflict(three, five))
	got := Similarity(three, five)
	require.Less(t, got, 0.80, "spelled hull/SPV numerals that disagree should fall below the screening line")
	require.InDelta(t, got, scoreSimilarityFast(three, five, SimilarityOpts{}), 0.001)
}

func TestNameNumeralConflict_MissingOnOneSideIsNotConflict(t *testing.T) {
	base := Entity[Value]{
		Name:   "NS LEADER",
		Type:   EntityVessel,
		Vessel: &Vessel{Name: "NS LEADER"},
	}.Normalize()
	ii := Entity[Value]{
		Name:   "NS LEADER II",
		Type:   EntityVessel,
		Vessel: &Vessel{Name: "NS LEADER II"},
	}.Normalize()

	require.False(t, hasNameNumeralConflict(base, ii))
	got := Similarity(base, ii)
	require.Greater(t, got, 0.80)
}

func TestNameNumeralConflict_MatchingIMOStillOne(t *testing.T) {
	query := Entity[Value]{
		Name: "NS LEADER II",
		Type: EntityVessel,
		Vessel: &Vessel{
			Name:      "NS LEADER II",
			IMONumber: "9339301",
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "NS LEADER",
		Type: EntityVessel,
		Vessel: &Vessel{
			Name:      "NS LEADER",
			IMONumber: "9339301",
		},
	}.Normalize()

	got := Similarity(query, index)
	require.InDelta(t, 1.0, got, 0.001)
}

func TestTokenNumeral(t *testing.T) {
	n, ok := tokenNumeral("viii")
	require.True(t, ok)
	require.Equal(t, 8, n)

	n, ok = tokenNumeral("12")
	require.True(t, ok)
	require.Equal(t, 12, n)

	_, ok = tokenNumeral("no")
	require.False(t, ok)

	_, ok = tokenNumeral("leader")
	require.False(t, ok)
}
