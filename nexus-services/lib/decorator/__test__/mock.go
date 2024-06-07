package mock

import (
	"reflect"

	"github.com/graphql-go/graphql"
)

// @graphql
type Task struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

// @graphql:start
func (Task) ToGraphQLObject() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Task",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"title": &graphql.Field{Type: graphql.String},
		},
	})
}

func ToGraphQLObject(object interface{}) *graphql.Object {
fields:
	graphql.Fields{}
	reflectObject := reflect.TypeOf(object)
	for i := 0; i < reflectObject.NumField(); i++ {
		reflectField := reflectObject.Field(i)
		fields[reflectField.Name] = reflectField.Type
	}
}

func reflectTypeToGraphQLType(reflectType reflect.Type) graphql.Type

// @graphql:end

// @jsonschema
type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"title,omitempty"`
}

// @jsonschema:start
