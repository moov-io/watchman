package ofactest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	if testing.Short() {
		t.Skip("requires the live OFAC Sanctions List Service")
	}

	cc := NewClient()

	ctx := context.Background()
	results, err := cc.Search(ctx, SearchParams{
		Name:      "Mohammad",
		NameScore: 95,
	})
	require.NoError(t, err)
	require.NotEmpty(t, results)

	for _, res := range results {
		t.Logf("%#v", res)
	}
}
