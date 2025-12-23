package geocodetest_test

import (
	"context"
	"testing"

	"github.com/moov-io/watchman/internal/geocoding/geocodetest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestRandomGeocoder(t *testing.T) {
	addrs := []search.Address{
		{
			Line1: "123 1st St",
		},
		{
			Line1: "432 2nd St",
		},
	}

	g := &geocodetest.RandomGeocoder{}

	got := g.GeocodeAddresses(context.Background(), addrs)
	require.Len(t, got, 2)

	for i := range got {
		require.NotEmpty(t, got[i].Latitude)
		require.NotEmpty(t, got[i].Longitude)
	}
}
