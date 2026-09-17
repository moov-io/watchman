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
	if !strings.Contains(md, "algorithm-comparison.csv") {
		t.Fatal("missing CSV download link")
	}
	if !strings.Contains(md, "csvq") {
		t.Fatal("missing csvq tip")
	}

	csvBody := RenderCSV()
	if !strings.HasPrefix(csvBody, "query,index,note,") {
		t.Fatalf("csv header: %q", csvBody[:min(80, len(csvBody))])
	}
	if strings.Count(csvBody, "\n") < len(pairs)+1 {
		t.Fatal("csv missing rows")
	}

	root := repoRoot(t)
	if err := os.WriteFile(filepath.Join(root, "docs", "algorithm-comparison.md"), []byte(md), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "docs", "algorithm-comparison.csv"), []byte(csvBody), 0o644); err != nil {
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
