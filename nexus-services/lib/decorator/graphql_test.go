package decorator

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestGraphQLDecorator_Generate(t *testing.T) {
	content := `package main

type MyStruct struct {
	ID   string
	Name string
	Age  int
}`
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse content: %v", err)
	}

	decorator := &GraphQLDecorator{}
	pos := node.Decls[0].Pos()
	generatedCode := decorator.Generate(pos, fset, node)
	expectedCode := `// Generated code for @graphql
func (s *MyStruct) GraphQLSchema() string {
    return ` + "`type MyStruct {\n        ID: string\n        Name: string\n        Age: int\n    }`\n}\n\n" + `func (s *MyStruct) GraphQLQuery() string {
    return ` + "`{\n    mystruct {\n        ID\n        Name\n        Age\n    }\n}`\n}\n\n"

	if !strings.Contains(generatedCode, expectedCode) {
		t.Errorf("Generated code does not match expected code.\nExpected:\n%s\nGot:\n%s", expectedCode, generatedCode)
	}
}

func TestGraphQLDecorator_Pos(t *testing.T) {
	decorator := &GraphQLDecorator{pos: token.Pos(10)}
	if decorator.Pos() != token.Pos(10) {
		t.Errorf("Expected position 10, got %d", decorator.Pos())
	}
}
