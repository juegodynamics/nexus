package graphdb

import (
	"os"
	"testing"
)

type Person struct {
	ID    string   `graphdb:"id"`
	Name  string   `graphdb:"property,name"`
	Age   int      `graphdb:"property,age"`
	Knows []string `graphdb:"relation,knows"`
}

func TestGraphORM(t *testing.T) {
	g := NewGraph()

	// Add structs
	alice := &Person{ID: "1", Name: "Alice", Age: 30, Knows: []string{"2"}}
	bob := &Person{ID: "2", Name: "Bob", Age: 25}

	err := g.AddStruct(alice)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = g.AddStruct(bob)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Get structs
	var retrievedAlice Person
	err = g.GetStruct("1", &retrievedAlice)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if retrievedAlice.Name != "Alice" || retrievedAlice.Age != 30 || len(retrievedAlice.Knows) != 1 || retrievedAlice.Knows[0] != "2" {
		t.Fatalf("retrieved Alice does not match expected, got %+v", retrievedAlice)
	}

	// Save to file
	err = g.SaveToFile("test_graph.json")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Load from file
	g2 := NewGraph()
	err = g2.LoadFromFile("test_graph.json")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validate loaded data
	var loadedAlice Person
	err = g2.GetStruct("1", &loadedAlice)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if loadedAlice.Name != "Alice" || loadedAlice.Age != 30 || len(loadedAlice.Knows) != 1 || loadedAlice.Knows[0] != "2" {
		t.Fatalf("loaded Alice does not match expected, got %+v", loadedAlice)
	}

	// Clean up
	os.Remove("test_graph.json")
}
