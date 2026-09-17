package address

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

type stubParser struct {
	addr search.Address
	err  error
}

func (s stubParser) ParseAddress(ctx context.Context, input string) (search.Address, error) {
	return s.addr, s.err
}

func TestParse_UsesParser(t *testing.T) {
	want := search.Address{Line1: "from parser", City: "montreal"}
	got := Parse(context.Background(), stubParser{addr: want}, "ignored")
	require.Equal(t, want, got)
}

func TestParse_FallsBackOnError(t *testing.T) {
	got := Parse(context.Background(), stubParser{err: errors.New("boom")}, "123 First St Anytown CA 90210")
	require.Equal(t, "123 first st", strings.ToLower(got.Line1))
	require.Equal(t, "anytown", strings.ToLower(got.City))
}

func TestParse_NilParser(t *testing.T) {
	got := Parse(context.Background(), nil, "123 First St Anytown CA 90210")
	require.Equal(t, "123 first st", strings.ToLower(got.Line1))
}
