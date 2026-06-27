package us_tel

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
)

func TestUSTelDownload_initialDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "initial-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	mk := func(t *testing.T, name string, body string) {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(body), 0600); err != nil {
			t.Fatalf("writing %s: %v", path, err)
		}
	}

	mk(t, "us_tel.json", "file=us_tel.json")

	file, err := DownloadUsTel(context.Background(), log.NewNopLogger(), dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(file) == 0 {
		t.Fatal("no US TEL file")
	}

	for fn, fd := range file {
		if strings.EqualFold("us_tel.json", filepath.Base(fn)) {
			_, err := io.ReadAll(fd)
			if err != nil {
				t.Fatal(err)
			}
		} else {
			t.Fatalf("unknown file: %v", file)
		}
	}
}

func TestUSTelDownload_fromURL(t *testing.T) {
	// prepare temporary file with known content
	f, err := os.CreateTemp("", "us_tel-*.json")
	if err != nil {
		t.Fatal(err)
	}
	content := "hello-us-tel"
	if _, err := f.WriteString(content); err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()

	defer os.Remove(f.Name())
	orig := publicUSTelDownloadUrl
	publicUSTelDownloadUrl = "file://" + f.Name()
	defer func() { publicUSTelDownloadUrl = orig }()

	os.Setenv("US_TEL_URL", "file://"+f.Name())
	defer os.Unsetenv("US_TEL_URL")

	files, err := DownloadUsTel(context.Background(), log.NewNopLogger(), "")
	if err != nil {
		t.Fatalf("expected download success, got %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("expected exactly 1 file, got %d", len(files))
	}

	for _, rc := range files {
		buf, err := io.ReadAll(rc)
		if err != nil {
			t.Fatal(err)
		}
		if string(buf) != content {
			t.Errorf("downloaded content mismatch: %q", string(buf))
		}
	}
}

func TestUSTelDownload_badURL(t *testing.T) {
	orig := publicUSTelDownloadUrl
	defer func() { publicUSTelDownloadUrl = orig }()

	publicUSTelDownloadUrl = "file:///nonexistent/path"

	files, err := DownloadUsTel(context.Background(), log.NewNopLogger(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(files) != 0 {
		t.Errorf("expected 0 files returned on failure, got %d", len(files))
	}
}