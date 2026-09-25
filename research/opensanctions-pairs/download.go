// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	sampleURL = "https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs/resolve/main/sample_1000.json"
	fullURL   = "https://data.opensanctions.org/contrib/training/pairs-20251209.json.gz"
)

func downloadDataset(kind, destDir string) (string, error) {
	if err := os.MkdirAll(destDir, 0o755); err != nil {
		return "", err
	}
	var url, name string
	switch kind {
	case "sample":
		url, name = sampleURL, "sample_1000.json"
	case "full":
		url, name = fullURL, "pairs-20251209.json.gz"
	default:
		return "", fmt.Errorf("unknown download kind %q (use sample or full)", kind)
	}
	dest := filepath.Join(destDir, name)
	if st, err := os.Stat(dest); err == nil && st.Size() > 0 {
		return dest, nil
	}

	client := &http.Client{Timeout: 10 * time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("download %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download %s: %s", url, resp.Status)
	}

	tmp := dest + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		os.Remove(tmp)
		return "", fmt.Errorf("write download: %w", err)
	}
	if err := f.Close(); err != nil {
		os.Remove(tmp)
		return "", err
	}
	if err := os.Rename(tmp, dest); err != nil {
		return "", err
	}
	return dest, nil
}

func writeFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil && filepath.Dir(path) != "." {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
