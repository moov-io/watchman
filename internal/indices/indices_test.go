package indices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewIndicies(t *testing.T) {
	indices := New(122, 5)
	require.Len(t, indices, 6)

	expected := []int{0, 25, 50, 74, 98, 122}
	require.Equal(t, expected, indices)
}

func TestProcessSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	originalInput := make([]int, len(input))
	copy(originalInput, input)
	require.Equal(t, input, originalInput)

	fn := func(in int) string {
		return fmt.Sprintf("%d", in*5)
	}
	expected := []string{"5", "10", "15", "20", "25", "30", "35", "40", "45", "50"}

	require.ElementsMatch(t, expected, ProcessSlice(input, 3, fn))
	require.Equal(t, originalInput, input) // input is unchanged

	require.ElementsMatch(t, expected, ProcessSlice(input, 7, fn))
	require.Equal(t, originalInput, input) // input is unchanged

	require.ElementsMatch(t, expected, ProcessSlice(input, 10, fn))
	require.Equal(t, originalInput, input) // input is unchanged
}
