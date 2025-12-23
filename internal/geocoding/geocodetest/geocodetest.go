package geocodetest

import (
	"context"
	"math/rand/v2"

	"github.com/moov-io/watchman/pkg/search"
)

type RandomGeocoder struct{}

func (g *RandomGeocoder) GeocodeAddresses(_ context.Context, addresses []search.Address) []search.Address {
	for idx := range addresses {
		addresses[idx].Latitude = rand.Float64()
		addresses[idx].Longitude = rand.Float64()
	}
	return addresses
}
