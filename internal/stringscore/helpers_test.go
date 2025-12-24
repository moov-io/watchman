package stringscore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloatSlicesEqual(t *testing.T) {
	a := []float64{1.2300000000001}
	b := []float64{1.23}

	require.True(t, floatSlicesEqual(a, b))
}
