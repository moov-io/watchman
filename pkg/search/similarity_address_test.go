package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddress_findCountryCode(t *testing.T) {
	require.Equal(t, "US", findCountryCode("US"))
	require.Equal(t, "US", findCountryCode("USA"))
	require.Equal(t, "US", findCountryCode("UNITED STATES"))
	require.Equal(t, "US", findCountryCode("united states of america"))
}
