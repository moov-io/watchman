package address

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAddress(t *testing.T) {
	input := "123 First St Anytown CA 90210"

	got := ParseAddress(context.Background(), input)
	require.Equal(t, "123 first st", strings.ToLower(got.Line1))
	require.Equal(t, "anytown", strings.ToLower(got.City))
	require.Equal(t, "90210", strings.ToLower(got.PostalCode))
	require.Equal(t, "ca", strings.ToLower(got.State))
}
