package search_test

import (
	"context"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestMockClient(t *testing.T) {
	ctx := context.Background()
	mc := search.NewMockClient()

	mc.Index = append(mc.Index, search.Entity[search.Value]{
		Name: "Acme corp",
		Type: search.EntityBusiness,
	})
	mc.Index = append(mc.Index, search.Entity[search.Value]{
		Name: "Jane Doe",
		Type: search.EntityPerson,
	})
	mc.Index = append(mc.Index, search.Entity[search.Value]{
		Name: "Storage Savers",
		Type: search.EntityBusiness,
	})

	query := search.Entity[search.Value]{
		Name: "John Doe",
		Type: search.EntityPerson,
	}
	opts := search.SearchOpts{
		Limit: 2,
	}
	resp, err := mc.SearchByEntity(ctx, query, opts)
	require.NoError(t, err)
	require.Len(t, resp.Entities, 2)

	first := resp.Entities[0]
	require.Equal(t, "Jane Doe", first.Name)
	require.InDelta(t, 0.6667, first.Match, 0.001)
}
