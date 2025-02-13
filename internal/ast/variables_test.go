package ast_test

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/internal/ast"

	"github.com/stretchr/testify/require"
)

func TestExtractVariablesOfType(t *testing.T) {
	modelsPath := filepath.Join("..", "..", "pkg", "search", "models.go")
	found, err := ast.ExtractVariablesOfType(modelsPath, "EntityType")
	require.NoError(t, err)

	expected := []string{"aircraft", "business", "organization", "person", "vessel"}
	require.ElementsMatch(t, expected, found)
}
