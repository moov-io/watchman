package us_non_sdn_test

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/sources/ofac"
	"github.com/moov-io/watchman/pkg/sources/us_non_sdn"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	ctx := context.Background()
	logBuf, logger := log.NewBufferLogger()
	initialDir := filepath.Join("..", "..", "..", "test", "testdata")

	files, err := us_non_sdn.Download(ctx, logger, initialDir)
	require.NoError(t, err)
	require.NotEmpty(t, files)

	results, err := ofac.Read(files)
	require.NoError(t, err)

	require.Len(t, results.SDNs, 442)
	require.Len(t, results.Addresses, 442)
	require.Len(t, results.AlternateIdentities, 395)
	require.Len(t, results.SDNComments, 10)

	require.Equal(t, "6c789e23b49f2f276908b1f89bab88fb21623f708a1d14fc5db6d66f563652b2", results.ListHash)

	// Verify logs
	if testing.Verbose() {
		fmt.Printf("\n%s\n", logBuf.String())
	}

	output := strings.TrimSpace(logBuf.String())
	lines := strings.Split(output, "\n")
	require.Len(t, lines, 4)

	require.Contains(t, logBuf.String(), `msg="found cons_add.csv"`)
	require.Contains(t, logBuf.String(), `msg="found cons_alt.csv"`)
	require.Contains(t, logBuf.String(), `msg="found cons_comments.csv"`)
	require.Contains(t, logBuf.String(), `msg="found cons_prim.csv"`)
}
