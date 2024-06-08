package graphdb

import (
	"fmt"
	"reflect"
	"strings"
)

func (g *Graph) AddStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("input must be a pointer to a struct")
	}
	val = val.Elem()
	typ := val.Type()

	idField := ""
	label := typ.Name()
	properties := make(map[string]interface{})
	relations := make(map[string][]string)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		tag := field.Tag.Get("graphdb")
		if tag == "" {
			continue
		}

		tagParts := strings.Split(tag, ",")
		switch tagParts[0] {
		case "id":
			idField = fieldValue.String()
		case "label":
			if len(tagParts) > 1 {
				label = tagParts[1]
			}
		case "property":
			if len(tagParts) > 1 {
				properties[tagParts[1]] = fieldValue.Interface()
			} else {
				properties[field.Name] = fieldValue.Interface()
			}
		case "relation":
			if len(tagParts) > 1 {
				relations[tagParts[1]] = append(relations[tagParts[1]], fieldValue.String())
			}
		}
	}

	if idField == "" {
		return fmt.Errorf("struct must have an id field")
	}

	g.AddNode(idField, label, properties)

	for label, targets := range relations {
		for _, target := range targets {
			err := g.AddEdge(idField, target, label, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Graph) GetStruct(id string, s interface{}) error {
	node, exists := g.GetNode(id)
	if !exists {
		return fmt.Errorf("node with id %s not found", id)
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("output must be a pointer to a struct")
	}
	val = val.Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("graphdb")
		if tag == "" {
			continue
		}

		tagParts := strings.Split(tag, ",")
		switch tagParts[0] {
		case "id":
			val.Field(i).SetString(node.ID)
		case "property":
			propName := field.Name
			if len(tagParts) > 1 {
				propName = tagParts[1]
			}
			if propValue, exists := node.Properties[propName]; exists {
				val.Field(i).Set(reflect.ValueOf(propValue))
			}
		case "relation":
			if len(tagParts) > 1 {
				relatedNodes := g.GetNeighbors(node.ID)
				relatedIDs := []string{}
				for _, n := range relatedNodes {
					relatedIDs = append(relatedIDs, n.ID)
				}
				val.Field(i).Set(reflect.ValueOf(relatedIDs))
			}
		}
	}

	return nil
}
