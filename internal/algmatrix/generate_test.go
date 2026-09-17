package algmatrix

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteAlgorithmComparisonDoc(t *testing.T) {
	md := Render()
	if !strings.Contains(md, "jaro-winkler") {
		t.Fatal("missing algorithm legend")
	}
	if !strings.Contains(md, "dominguez") {
		t.Fatal("missing false-positive row")
	}

	root := repoRoot(t)
	path := filepath.Join(root, "docs", "algorithm-comparison.md")
	if err := os.WriteFile(path, []byte(md), 0o644); err != nil {
		t.Fatal(err)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 6; i++ {
		if _, err := os.Stat(filepath.Join(dir, "docs", "_data", "docs-menu.yml")); err == nil {
			return dir
		}
		dir = filepath.Dir(dir)
	}
	t.Fatal("watchman repo root not found")
	return ""
}
