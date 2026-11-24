package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOriginalAltNamesLastTypeWins(t *testing.T) {
	t.Parallel()

	e := Entity[Value]{
		Person: &Person{AltNames: []string{"from-person"}},
		Vessel: &Vessel{AltNames: []string{"from-vessel"}},
	}
	require.Equal(t, []string{"from-vessel"}, originalAltNames(e))
}

func TestScoreTrail_CoverageReasonsOnlyWhenCollecting(t *testing.T) {
	t.Parallel()

	trail := scoreTrail{collect: false}
	applyPenaltiesAndBonuses(&trail, 0.50, coverage{ratio: 0.10, criticalRatio: 0.10}, entityFields{
		required: 1,
		hasName:  true,
	})
	require.Empty(t, trail.steps)
	require.Greater(t, trail.score, 0.0)
	require.Less(t, trail.score, 0.50)

	collected := scoreTrail{collect: true}
	applyPenaltiesAndBonuses(&collected, 0.50, coverage{ratio: 0.10, criticalRatio: 0.10}, entityFields{
		required: 1,
		hasName:  true,
	})
	require.NotEmpty(t, collected.steps)
	var sawCoverage bool
	for _, step := range collected.steps {
		if step.Stage == "low-coverage" || step.Stage == "critical-coverage" {
			require.NotEmpty(t, step.Reason)
			sawCoverage = true
		}
	}
	require.True(t, sawCoverage)
}
