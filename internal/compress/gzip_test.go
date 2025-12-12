package compress_test

import (
	"compress/gzip"
	"io"
	"strings"
	"testing"

	"github.com/moov-io/watchman/internal/compress"

	"github.com/stretchr/testify/require"
)

func TestGzip(t *testing.T) {
	content := strings.NewReader("hello, world")

	file := compress.GzipTestFile(t, content)

	r, err := gzip.NewReader(file)
	require.NoError(t, err)

	bs, err := io.ReadAll(r)
	require.NoError(t, err)

	require.Equal(t, "hello, world", string(bs))
}
