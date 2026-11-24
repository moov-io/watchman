package stringscore

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeepCombinationScore(t *testing.T) {
	t.Parallel()

	require.False(t, keepCombinationScore(math.NaN(), 0, false))
	require.False(t, keepCombinationScore(math.NaN(), 0.5, true))
	require.True(t, keepCombinationScore(0, 0, false), "keep a first finite zero score")
	require.False(t, keepCombinationScore(0, 0.5, true))
	require.True(t, keepCombinationScore(0.9, 0.5, true))
}
