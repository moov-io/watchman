package ast

import (
	"cmp"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"slices"
	"strconv"
	"sync"
)

// extractCache holds parsed model values for the life of the process.
// parsedFiles holds one parsed AST per path so several types from models.go
// share a single parse. The embedded file does not change at runtime.
var (
	extractCache sync.Map // path + type -> []string
	parsedFiles  sync.Map // path -> *parsedSource
)

type parsedSource struct {
	once sync.Once
	node *ast.File
	err  error
}

// ExtractVariablesOfType parses a Go source file and finds all variables of the specified type.
func ExtractVariablesOfType(fsys fs.FS, path, typeName string) ([]string, error) {
	key := path + "\x00" + typeName
	if cached, ok := extractCache.Load(key); ok {
		return cloneStrings(cached.([]string)), nil
	}

	values, err := extractVariablesOfType(fsys, path, typeName)
	if err != nil {
		return nil, err
	}
	extractCache.Store(key, values)
	return cloneStrings(values), nil
}

func extractVariablesOfType(fsys fs.FS, path, typeName string) ([]string, error) {
	node, err := parseCached(fsys, path)
	if err != nil {
		return nil, err
	}

	var values []string

	// Walk the AST to find variables of the specified type
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for variable declarations
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.VAR {
			return true
		}

		// Process each variable in the declaration
		for _, spec := range decl.Specs {
			vspec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			// Check the type of the variable
			if vspec.Type != nil {
				ident, ok := vspec.Type.(*ast.Ident)
				if ok && ident.Name == typeName {
					// Add all variable names in this declaration to the result
					for i := range vspec.Names {
						if i < len(vspec.Values) {
							values = append(values, formatNode(vspec.Values[i]))
						}
					}
				}
			}
		}
		return true
	})

	slices.Sort(values)

	return values, nil
}

func parseCached(fsys fs.FS, path string) (*ast.File, error) {
	loaded, _ := parsedFiles.LoadOrStore(path, &parsedSource{})
	entry := loaded.(*parsedSource)
	entry.once.Do(func() {
		entry.node, entry.err = parseSource(fsys, path)
	})
	return entry.node, entry.err
}

func parseSource(fsys fs.FS, path string) (*ast.File, error) {
	fd, err := fsys.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening %s failed: %w", path, err)
	}
	defer fd.Close()

	src, err := io.ReadAll(fd)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, src, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}
	return node, nil
}

func cloneStrings(in []string) []string {
	if len(in) == 0 {
		return nil
	}
	out := make([]string, len(in))
	copy(out, in)
	return out
}

func formatNode(expr ast.Expr) string {
	if expr == nil {
		return ""
	}

	switch v := expr.(type) {
	case *ast.BasicLit: // Literal values like numbers or strings
		value, _ := strconv.Unquote(v.Value)
		return cmp.Or(value, v.Value)

	case *ast.Ident: // Identifiers (e.g., constants or variables)
		value, _ := strconv.Unquote(v.Name)
		return cmp.Or(value, v.Name)

	case *ast.CompositeLit: // Composite literals like structs or arrays
		return "composite literal"

	default:
		return "complex expression"
	}
}
