package graphgorm

import (
	"fmt"
	"reflect"
	"strings"
)

// GenerateGraphQLSchema generates GraphQL schema from struct definitions
func GenerateGraphQLSchema(v interface{}) string {
	return generateGraphQLType(reflect.TypeOf(v), "")
}

// generateGraphQLType recursively generates GraphQL types from struct fields
func generateGraphQLType(t reflect.Type, prefix string) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := "type " + prefix + t.Name() + " {\n"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		graphqlTag := field.Tag.Get("graphql")
		if graphqlTag == "" {
			graphqlTag = field.Name
		}
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			// Generate nested types for struct fields
			nestedSchema := generateGraphQLType(fieldType, prefix+field.Name)
			schema += nestedSchema
			schema += fmt.Sprintf("  %s: %s%s\n", field.Name, prefix, field.Name)
		} else {
			schema += fmt.Sprintf("  %s: %s\n", field.Name, graphqlTag)
		}
	}
	schema += "}\n"
	return schema
}

// GenerateGraphQLQuery generates GraphQL queries for a given struct
func GenerateGraphQLQuery(v interface{}) string {
	return generateGraphQLQuery(reflect.TypeOf(v))
}

// generateGraphQLQuery recursively generates GraphQL queries for struct fields
func generateGraphQLQuery(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	query := fmt.Sprintf("query get%s($id: ID!) {\n  get%s(id: $id) {\n", t.Name(), t.Name())
	query += generateGraphQLQueryFields(t)
	query += "  }\n}"
	return query
}

// generateGraphQLQueryFields generates fields for a GraphQL query
func generateGraphQLQueryFields(t reflect.Type) string {
	fields := ""
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			fields += fmt.Sprintf("    %s {\n", field.Name)
			fields += generateGraphQLQueryFields(fieldType)
			fields += "    }\n"
		} else {
			fields += fmt.Sprintf("    %s\n", field.Name)
		}
	}
	return fields
}

// GenerateGraphQLMutation generates GraphQL mutations for a given struct
func GenerateGraphQLMutation(v interface{}) string {
	t := reflect.TypeOf(v)
	mutation := fmt.Sprintf("mutation add%s($input: Add%sInput!) {\n  add%s(input: $input) {\n    %s {\n", t.Name(), t.Name(), t.Name(), strings.ToLower(t.Name()))
	mutation += generateGraphQLMutationFields(t)
	mutation += "    }\n  }\n}"
	return mutation
}

// generateGraphQLMutationFields generates fields for a GraphQL mutation
func generateGraphQLMutationFields(t reflect.Type) string {
	fields := ""
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			fields += fmt.Sprintf("      %s {\n", field.Name)
			fields += generateGraphQLMutationFields(fieldType)
			fields += "      }\n"
		} else {
			fields += fmt.Sprintf("      %s\n", field.Name)
		}
	}
	return fields
}

// GenerateGraphQLInput generates GraphQL input types for a given struct
func GenerateGraphQLInput(v interface{}) string {
	return generateGraphQLInputType(reflect.TypeOf(v), "")
}

// generateGraphQLInputType recursively generates GraphQL input types from struct fields
func generateGraphQLInputType(t reflect.Type, prefix string) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	input := "input Add" + prefix + t.Name() + "Input {\n"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		graphqlTag := field.Tag.Get("graphql")
		if graphqlTag == "" {
			graphqlTag = field.Name
		}
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			// Generate nested input types for struct fields
			nestedInput := generateGraphQLInputType(fieldType, prefix+field.Name)
			input += nestedInput
			input += fmt.Sprintf("  %s: %s%s\n", field.Name, prefix, field.Name)
		} else {
			input += fmt.Sprintf("  %s: %s\n", field.Name, graphqlTag)
		}
	}
	input += "}\n"
	return input
}

// GenerateGraphQLEnum generates GraphQL enums from struct definitions
func GenerateGraphQLEnum(v interface{}) string {
	return generateGraphQLEnumType(reflect.TypeOf(v))
}

// generateGraphQLEnumType generates GraphQL enums from struct fields
func generateGraphQLEnumType(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	enum := "enum " + t.Name() + " {\n"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		enum += fmt.Sprintf("  %s\n", field.Name)
	}
	enum += "}\n"
	return enum
}
