package ingest

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/moov-io/watchman/internal/compress"

	"github.com/stretchr/testify/require"
)

func TestDecompressBody(t *testing.T) {
	body := compress.GzipTestFile(t, strings.NewReader(strings.Repeat("hello world", 25)))

	cases := []struct {
		input    io.Reader
		expected string
	}{
		{
			input:    strings.NewReader(""),
			expected: "",
		},
		{
			input:    strings.NewReader("hello"),
			expected: "hello",
		},
		{
			input:    strings.NewReader("hello, world"),
			expected: "hello, world",
		},
		{
			input:    strings.NewReader(strings.Repeat("hello, world", 100)),
			expected: strings.Repeat("hello, world", 100),
		},
		{
			input:    body,
			expected: strings.Repeat("hello world", 25),
		},
	}
	for idx, tc := range cases {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			got := maybeDecompressBody(tc.input)

			gotBytes, err := io.ReadAll(got)
			require.NoError(t, err)

			require.Equal(t, tc.expected, string(gotBytes))
		})
	}
}
