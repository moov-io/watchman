package search_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_EmptyQuery(t *testing.T) {
	var query search.Entity[search.Value]
	index := ofactest.FindEntity(t, "47371")

	got := search.Similarity(query, index)
	require.InDelta(t, 0.0, got, 0.001)
}

func TestSimilarityDebug_FromJSON(t *testing.T) {
	query := readEntity(t, "1-query.json").Normalize()
	index := readEntity(t, "1-index.json").Normalize()

	var buf bytes.Buffer
	got := search.DebugSimilarity(&buf, query, index)
	t.Logf("%.2f - %v (%v)", got, query.Name, index.Name)
	fmt.Println(buf.String())

	require.InDelta(t, got, 0.690, 0.001)
}

func readEntity(tb testing.TB, name string) search.Entity[search.Value] {
	tb.Helper()

	bs, err := os.ReadFile(filepath.Join("testdata", name))
	require.NoError(tb, err)

	var entity search.Entity[search.Value]
	err = json.Unmarshal(bs, &entity)
	require.NoError(tb, err)

	return entity
}
