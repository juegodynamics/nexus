package decorator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"regexp"
	"strings"
)

// Decorator interface
type Decorator interface {
	Generate(pos token.Pos, fset *token.FileSet, node *ast.File) string
	Pos() token.Pos
}

// Registry for decorators
var decoratorRegistry = make(map[string]Decorator)

// RegisterDecorator registers a custom decorator
func RegisterDecorator(name string, decorator Decorator) {
	decoratorRegistry[name] = decorator
}

// ResetRegistry resets the decorator registry (useful for testing)
func ResetRegistry() {
	decoratorRegistry = make(map[string]Decorator)
}

// ParseFile parses the Go source file and identifies decorators
func ParseFile(filename string) ([]Decorator, *ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, nil, err
	}

	var decorators []Decorator
	for _, comment := range node.Comments {
		for _, c := range comment.List {
			match := regexp.MustCompile(`@(\w+)`).FindStringSubmatch(c.Text)
			if match != nil {
				if decorator, exists := decoratorRegistry[match[1]]; exists {
					decoratorRegistry[match[1]].(interface {
						SetPos(token.Pos)
					}).SetPos(c.Pos())
					decorators = append(decorators, decorator)
				}
			}
		}
	}
	return decorators, node, fset, nil
}

// GenerateDecoratorMethods generates methods for the registered decorators
func GenerateDecoratorMethods(decorators []Decorator, node *ast.File, fset *token.FileSet) string {
	var generatedCode strings.Builder

	for _, dec := range decorators {
		generatedCode.WriteString(dec.Generate(dec.Pos(), fset, node))
	}
	return generatedCode.String()
}

// RemoveGeneratedCode removes previously generated code from the source
func RemoveGeneratedCode(src string) string {
	re := regexp.MustCompile(`(?s)// Generated code for @\w+.*?}\n\n`)
	return re.ReplaceAllString(src, "")
}

// InsertGeneratedCode inserts generated code into the source after the decorator's position
func InsertGeneratedCode(src, generatedCode string, decorators []Decorator, fset *token.FileSet) string {
	lines := strings.Split(src, "\n")
	generatedLines := strings.Split(generatedCode, "\n")

	for _, dec := range decorators {
		pos := fset.Position(dec.Pos())
		insertIndex := pos.Line
		lines = append(lines[:insertIndex], append(generatedLines, lines[insertIndex:]...)...)
	}
	return strings.Join(lines, "\n")
}

// ProcessFile processes the file, generating and inserting decorator methods
func ProcessFile(filename string) error {
	decorators, node, fset, err := ParseFile(filename)
	if err != nil {
		return fmt.Errorf("error parsing file: %v", err)
	}

	// Read the original file
	srcBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	src := string(srcBytes)

	// Remove previous generated code
	src = RemoveGeneratedCode(src)

	// Generate new decorator methods
	generatedCode := GenerateDecoratorMethods(decorators, node, fset)

	// Insert generated code into the source
	updatedSrc := InsertGeneratedCode(src, generatedCode, decorators, fset)

	// Write the updated source back to the file
	err = ioutil.WriteFile(filename, []byte(updatedSrc), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Println("Decorators processed and file updated successfully.")
	return nil
}
