package mock

import "github.com/graphql-go/graphql"

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

// @graphql:end

type Planet struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
