package ast_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/internal/ast"

	"github.com/stretchr/testify/require"
)

func TestExtractVariablesOfType(t *testing.T) {
	fsys := os.DirFS(filepath.Join("..", ".."))
	modelsPath := "pkg/search/models.go" // DirFS always expects forward slashes
	found, err := ast.ExtractVariablesOfType(fsys, modelsPath, "EntityType")
	require.NoError(t, err)

	expected := []string{"aircraft", "business", "organization", "person", "vessel"}
	require.ElementsMatch(t, expected, found)
}
