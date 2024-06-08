package gogen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// GoStruct represents a Go struct type.
type GoStruct struct {
	Comment string
	Name    string
	Fields  []*GoField
}

func (goStruct *GoStruct) String() string {
	var sb strings.Builder

	// Print any comments
	if goStruct.Comment != "" {
		for _, commentLine := range SplitStringByCharLimit(goStruct.Comment, 77) {
			sb.WriteString(fmt.Sprintf("// %s\n", commentLine))
		}
	}

	// Print the name
	sb.WriteString(fmt.Sprintf("type %s struct {\n", goStruct.Name))

	// Print the fields
	for _, field := range goStruct.Fields {
		sb.WriteString(fmt.Sprintf("\t%s\n", field.String()))
	}

	sb.WriteString("}")
	return sb.String()
}

// SplitStringByCharLimit splits a string into a slice of strings,
// where each string has a maximum length of charLimit.
func SplitStringByCharLimit(text string, charLimit int) []string {
	var result []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return result
	}

	var currentLine string
	for _, word := range words {
		if len(currentLine)+len(word)+1 > charLimit {
			result = append(result, strings.TrimSpace(currentLine))
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if currentLine != "" {
		result = append(result, currentLine)
	}

	return result
}

// GoField represents a field in a Go struct.
type GoField struct {
	Comment string
	Name    string
	Type    *GoFieldType
	Tags    []*GoFieldTag
}

func (goField *GoField) String() string {
	var sb strings.Builder
	if goField.Comment != "" {
		for _, commentLine := range SplitStringByCharLimit(goField.Comment, 77) {
			sb.WriteString(fmt.Sprintf("// %s\n", commentLine))
		}
	}
	sb.WriteString(goField.Name)
	sb.WriteString(" ")
	sb.WriteString(goField.Type.String())
	if len(goField.Tags) > 0 {
		sb.WriteString(" `")
		for i, tag := range goField.Tags {
			if i > 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(tag.String())
		}
		sb.WriteString("`")
	}
	return sb.String()
}

// GoFieldType represents the type of a field in a Go struct.
type GoFieldType struct {
	Name      string
	IsPointer bool
	IsSlice   bool
}

func (goFieldType *GoFieldType) String() string {
	var sb strings.Builder
	if goFieldType.IsSlice {
		sb.WriteString("[]")
	}
	if goFieldType.IsPointer {
		sb.WriteString("*")
	}
	sb.WriteString(goFieldType.Name)
	return sb.String()
}

type GoFieldTag struct {
	Key    string
	Values []string
}

func (goFieldTag *GoFieldTag) String() string {
	return fmt.Sprintf("%s:\"%s\"", goFieldTag.Key, strings.Join(goFieldTag.Values, ","))
}

func Parse(data []byte) ([]*GoStruct, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", string(data), parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return extractStructs(node), nil
}

// extractStructs extracts information about structs from an AST node.
func extractStructs(node ast.Node) []*GoStruct {
	var structs []*GoStruct

	ast.Inspect(node, func(n ast.Node) bool {
		// Find type declarations
		genDecl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			goStruct := &GoStruct{
				Name: typeSpec.Name.Name,
			}

			for _, field := range structType.Fields.List {
				for _, fieldName := range field.Names {
					goField := &GoField{
						Name: fieldName.Name,
						Type: extractFieldType(field.Type),
					}
					goStruct.Fields = append(goStruct.Fields, goField)
				}
			}

			structs = append(structs, goStruct)
		}

		return false
	})

	return structs
}

// extractFieldType extracts information about a field type from an AST expression.
func extractFieldType(expr ast.Expr) *GoFieldType {
	switch t := expr.(type) {
	case *ast.Ident:
		return &GoFieldType{Name: t.Name}
	case *ast.StarExpr:
		innerType := extractFieldType(t.X)
		innerType.IsPointer = true
		return innerType
	case *ast.ArrayType:
		innerType := extractFieldType(t.Elt)
		innerType.IsSlice = true
		return innerType
	default:
		return &GoFieldType{Name: fmt.Sprintf("%T", expr)}
	}
}
