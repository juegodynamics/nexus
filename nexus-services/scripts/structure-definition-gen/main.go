package main

import (
	"encoding/json"
	"os"

	"github.com/juegodynamics/nexus/nexus-services/lib/gogen"
)

type Schema struct {
	Discriminator struct {
		Mapping map[string]string `json:"mapping"`
	} `json:"discriminator"`
	Definitions map[string]*Definition `json:"definitions"`
}

type Definition struct {
	Description string               `json:"description"`
	Properties  map[string]*Property `json:"property"`
}

func (definition *Definition) ToGoStruct(name string) *gogen.GoStruct {
	definitionStruct := gogen.GoField()
}

type Property struct {
	Description string   `json:"description"`
	Ref         *string  `json:"$ref,omitempty"`
	Pattern     *string  `json:"pattern,omitempty"`
	Type        *string  `json:"type,omitempty"`
	Required    []string `json:"required,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}

func loadSchema(file string) (*Schema, error) {
	schemaData, err := os.ReadFile("lib/fhir/src/fhir.schema.json")
	if err != nil {
		return nil, err
	}

	schema := &Schema{}
	if err = json.Unmarshal(schemaData, schema); err != nil {
		return nil, err
	}

	return schema, nil
}

func main() {

	schema := loadSchema("lib/fhir/src/fhir.schema.json")

}
