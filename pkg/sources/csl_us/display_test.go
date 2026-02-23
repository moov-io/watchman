package csl_us

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetailsURL(t *testing.T) {
	got := DetailsURL("123")
	require.NotEmpty(t, got)
	expected := "https://www.trade.gov/data-visualization/csl-search"
	require.Equal(t, expected, got)
}
