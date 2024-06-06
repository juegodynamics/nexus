package decorator

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

// GraphQLDecorator implementation
type GraphQLDecorator struct {
	pos token.Pos
}

func (d *GraphQLDecorator) Generate(pos token.Pos, fset *token.FileSet, node *ast.File) string {
	var structName string
	var structFields []string

	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		switch x := n.(type) {
		case *ast.GenDecl:
			for _, spec := range x.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						if fset.Position(x.Pos()).Line == fset.Position(pos).Line {
							structName = typeSpec.Name.Name
							for _, field := range structType.Fields.List {
								fieldType := field.Type.(*ast.Ident).Name
								for _, fieldName := range field.Names {
									structFields = append(structFields, fmt.Sprintf("%s: %s", fieldName.Name, fieldType))
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	if structName == "" {
		return ""
	}

	var generatedCode strings.Builder
	generatedCode.WriteString(fmt.Sprintf("// Generated code for @graphql\n"))
	generatedCode.WriteString(fmt.Sprintf("func (s *%s) GraphQLSchema() string {\n", structName))
	generatedCode.WriteString(fmt.Sprintf("    return `type %s {\n", structName))
	for _, field := range structFields {
		generatedCode.WriteString(fmt.Sprintf("        %s\n", field))
	}
	generatedCode.WriteString("    }`\n")
	generatedCode.WriteString("}\n\n")
	generatedCode.WriteString(fmt.Sprintf("func (s *%s) GraphQLQuery() string {\n", structName))
	generatedCode.WriteString(fmt.Sprintf("    return `{\n    %s {\n", strings.ToLower(structName)))
	for _, field := range structFields {
		fieldName := strings.Split(field, ":")[0]
		generatedCode.WriteString(fmt.Sprintf("        %s\n", fieldName))
	}
	generatedCode.WriteString("    }\n}`\n")
	generatedCode.WriteString("}\n\n")

	return generatedCode.String()
}

func (d *GraphQLDecorator) Pos() token.Pos {
	return d.pos
}

func (d *GraphQLDecorator) SetPos(pos token.Pos) {
	d.pos = pos
}

func init() {
	RegisterDecorator("graphql", &GraphQLDecorator{})
}
