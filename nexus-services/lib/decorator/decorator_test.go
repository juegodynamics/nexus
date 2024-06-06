package decorator

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestRegisterDecorator(t *testing.T) {
	ResetRegistry()
	decorator := &GraphQLDecorator{}
	RegisterDecorator("graphql", decorator)
	if _, exists := decoratorRegistry["graphql"]; !exists {
		t.Error("Failed to register decorator")
	}
}

func TestParseFile(t *testing.T) {
	ResetRegistry()
	RegisterDecorator("graphql", &GraphQLDecorator{})
	content := `package main

// @graphql
type MyStruct struct {
	ID   string
	Name string
	Age  int
}`
	filename := "testfile.go"
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	defer os.Remove(filename)

	decorators, _, _, err := ParseFile(filename)
	if err != nil {
		t.Fatalf("Failed to parse file: %v", err)
	}

	if len(decorators) != 1 {
		t.Fatalf("Expected 1 decorator, found %d", len(decorators))
	}
}

func TestGenerateDecoratorMethods(t *testing.T) {
	ResetRegistry()
	RegisterDecorator("graphql", &GraphQLDecorator{})
	content := `package main

// @graphql
type MyStruct struct {
	ID   string
	Name string
	Age  int
}`
	filename := "testfile.go"
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	defer os.Remove(filename)

	decorators, node, fset, err := ParseFile(filename)
	if err != nil {
		t.Fatalf("Failed to parse file: %v", err)
	}

	generatedCode := GenerateDecoratorMethods(decorators, node, fset)
	expectedCode := `// Generated code for @graphql
func (s *MyStruct) GraphQLSchema() string {
    return ` + "`type MyStruct {\n        ID: string\n        Name: string\n        Age: int\n    }`\n}\n\n" + `func (s *MyStruct) GraphQLQuery() string {
    return ` + "`{\n    mystruct {\n        ID\n        Name\n        Age\n    }\n}`\n}\n\n"

	if !strings.Contains(generatedCode, expectedCode) {
		t.Errorf("Generated code does not match expected code.\nExpected:\n%s\nGot:\n%s", expectedCode, generatedCode)
	}
}

func TestRemoveGeneratedCode(t *testing.T) {
	content := `package main

// Generated code for @graphql
func (s *MyStruct) GraphQLSchema() string {
    return ` + "`type MyStruct {\n        ID: string\n        Name: string\n        Age: int\n    }`\n}\n\n" + `func (s *MyStruct) GraphQLQuery() string {
    return ` + "`{\n    mystruct {\n        ID\n        Name\n        Age\n    }\n}`\n}\n\n"

	expected := `package main

`
	result := RemoveGeneratedCode(content)
	if result != expected {
		t.Errorf("Failed to remove generated code.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestInsertGeneratedCode(t *testing.T) {
	content := `package main

type MyStruct struct {
	ID   string
	Name string
	Age  int
}`
	generatedCode := `// Generated code for @graphql
func (s *MyStruct) GraphQLSchema() string {
    return ` + "`type MyStruct {\n        ID: string\n        Name: string\n        Age: int\n    }`\n}\n\n" + `func (s *MyStruct) GraphQLQuery() string {
    return ` + "`{\n    mystruct {\n        ID\n        Name\n        Age\n    }\n}`\n}\n\n"

	expected := `package main

type MyStruct struct {
	ID   string
	Name string
	Age  int
}
// Generated code for @graphql
func (s *MyStruct) GraphQLSchema() string {
    return ` + "`type MyStruct {\n        ID: string\n        Name: string\n        Age: int\n    }`\n}\n\n" + `func (s *MyStruct) GraphQLQuery() string {
    return ` + "`{\n    mystruct {\n        ID\n        Name\n        Age\n    }\n}`\n}\n\n"

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse content: %v", err)
	}

	decorators := []Decorator{
		&GraphQLDecorator{pos: node.Decls[0].Pos()},
	}
	result := InsertGeneratedCode(content, generatedCode, decorators, fset)
	if result != expected {
		t.Errorf("Failed to insert generated code.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}
