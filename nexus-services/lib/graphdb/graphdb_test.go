package graphdb

import (
	"os"
	"testing"
)

func TestGraphOperations(t *testing.T) {
	g := NewGraph()

	// Add nodes
	g.AddNode("1", "Person", map[string]interface{}{"name": "Alice"})
	g.AddNode("2", "Person", map[string]interface{}{"name": "Bob"})

	// Add edge
	err := g.AddEdge("1", "2", "knows", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Get node
	node, exists := g.GetNode("1")
	if !exists || node.Properties["name"] != "Alice" {
		t.Fatalf("expected to find node with name Alice, got %v", node)
	}

	// Get edge
	edge, exists := g.GetEdge("1", "2")
	if !exists || edge.Label != "knows" {
		t.Fatalf("expected to find edge with label knows, got %v", edge)
	}

	// Get neighbors
	neighbors := g.GetNeighbors("1")
	if len(neighbors) != 1 || neighbors[0].Properties["name"] != "Bob" {
		t.Fatalf("expected to find Bob as neighbor, got %v", neighbors)
	}

	// Find nodes by property
	nodes := g.FindNodesByProperty("name", "Alice")
	if len(nodes) != 1 || nodes[0].Properties["name"] != "Alice" {
		t.Fatalf("expected to find one node with name Alice, got %v", nodes)
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
	node, exists = g2.GetNode("1")
	if !exists || node.Properties["name"] != "Alice" {
		t.Fatalf("expected to find node with name Alice, got %v", node)
	}

	// Clean up
	os.Remove("test_graph.json")
}
