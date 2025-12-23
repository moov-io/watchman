package geocoding

import (
	"context"

	"github.com/moov-io/watchman/pkg/search"
)

// Coordinates represents geographic coordinates returned by a geocoding provider.
type Coordinates struct {
	Latitude  float64
	Longitude float64
	// Accuracy indicates the precision level of the geocoding result.
	// Possible values: "rooftop", "street", "city", "state", "country", "approximate"
	Accuracy string
}

// Geocoder is the interface that geocoding providers must implement.
type Geocoder interface {
	// Geocode converts an address to geographic coordinates.
	// Returns nil coordinates if the address cannot be geocoded.
	Geocode(ctx context.Context, address search.Address) (*Coordinates, error)

	// Name returns the provider name for logging and metrics.
	Name() string
}

// NoOpGeocoder is a geocoder that does nothing, used when geocoding is disabled.
type NoOpGeocoder struct{}

func (n *NoOpGeocoder) Geocode(ctx context.Context, address search.Address) (*Coordinates, error) {
	return nil, nil
}

func (n *NoOpGeocoder) Name() string {
	return "noop"
}
