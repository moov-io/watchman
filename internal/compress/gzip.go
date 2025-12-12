package compress

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func GzipToFile(dir string, content io.Reader) (*os.File, error) {
	fd, err := os.CreateTemp(dir, "gzip-*")
	if err != nil {
		return nil, fmt.Errorf("creating temp file: %w", err)
	}

	var success bool
	defer func() {
		if !success {
			fd.Close()
			os.Remove(fd.Name())
		}
	}()

	// Write content
	w := gzip.NewWriter(fd)
	_, err = io.Copy(w, content)
	if err != nil {
		return nil, fmt.Errorf("gzip write: %w", err)
	}
	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("gzip close: %w", err)
	}

	// Prep file for read
	err = fd.Sync()
	if err != nil {
		return nil, fmt.Errorf("temp file sync: %w", err)
	}
	_, err = fd.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("seek reset: %w", err)
	}

	success = true

	return fd, nil
}

func GzipTestFile(tb testing.TB, content io.Reader) *os.File {
	tb.Helper()

	fd, err := GzipToFile(tb.TempDir(), content)
	require.NoError(tb, err)

	tb.Cleanup(func() {
		fd.Close()
		os.Remove(fd.Name())
	})

	return fd
}
