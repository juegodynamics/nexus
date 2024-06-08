package gogen

import (
	"fmt"
	"strings"

	"github.com/juegodynamics/nexus/nexus-services/lib/fhir/resources"
)

type StructureBuilder struct {
	SourceDefinitions map[string]*resources.StructureDefinition
	BuiltDefinitions  map[string]*GoStruct
	queue             chan string
}

func NewStructureBuilder() *StructureBuilder {
	return &StructureBuilder{
		SourceDefinitions: map[string]*resources.StructureDefinition{},
		BuiltDefinitions:  map[string]*GoStruct{},
		queue:             make(chan string, 300),
	}
}

func (structureBuilder *StructureBuilder) Push(typeName string) {
	structureBuilder.queue <- typeName
}

func (structureBuilder *StructureBuilder) Pop() string {
	nextTypeName := <-structureBuilder.queue
	return nextTypeName
}

func (builder *StructureBuilder) Build(definition *resources.StructureDefinition) error {
	builder.SourceDefinitions[definition.Name] = definition

	goStruct := &GoStruct{
		Comment: fmt.Sprintf("%s: %s", definition.Name, definition.Description),
		Name:    definition.Name,
		Fields:  []*GoField{},
	}

	for _, element := range definition.Snapshot.Element {
		field, err := builder.generateElement(element)
		if err != nil {
			return err
		}
		goStruct.Fields = append(goStruct.Fields, field)
	}

	builder.BuiltDefinitions[goStruct.Name] = goStruct

	return nil
}

func (builder *StructureBuilder) generateElement(element *resources.StructureDefinitionSnapshotElement) (*GoField, error) {
	elementType := &GoFieldType{}

	if len(element.Type) > 0 {
		elementType.Name = element.Type[0].Code
		elementType.IsPointer = true
		if element.Max == "*" {
			elementType.IsSlice = true
		}
	} else {
		elementType.Name = "interface{}"
	}

	fieldName := builder.toCamelCase(element.Path)
	// tag := fmt.Sprintf("`json:\"%s,omitempty\"`", element.Path)
	tag := &GoFieldTag{Key: "json", Values: []string{fieldName, "omitempty"}}

	comment := []string{element.Definition}

	if len(element.Comment) > 0 {
		comment = append(comment, element.Comment)
	}

	return &GoField{
		Name:    fieldName,
		Type:    elementType,
		Tags:    []*GoFieldTag{tag},
		Comment: strings.Join(comment, "\n\n"),
	}, nil
}

func (builder *StructureBuilder) toCamelCase(input string) string {
	parts := strings.Split(input, ".")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
