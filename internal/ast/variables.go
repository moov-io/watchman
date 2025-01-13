package ast

import (
	"cmp"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"slices"
	"strconv"
)

// ExtractVariablesOfType parses a Go source file and finds all variables of the specified type.
func ExtractVariablesOfType(path, typeName string) ([]string, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, src, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
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
