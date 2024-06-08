package graphdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"
)

type Node struct {
	ID         string                 `json:"id"`
	Label      string                 `json:"label"`
	Properties map[string]interface{} `json:"properties"`
}

type Edge struct {
	From       string                 `json:"from"`
	To         string                 `json:"to"`
	Label      string                 `json:"label"`
	Properties map[string]interface{} `json:"properties"`
}

type Graph struct {
	Nodes map[string]*Node            `json:"nodes"`
	Edges map[string]map[string]*Edge `json:"edges"`
	mutex sync.RWMutex
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string]map[string]*Edge),
	}
}

func (g *Graph) AddNode(id, label string, properties map[string]interface{}) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.Nodes[id] = &Node{ID: id, Label: label, Properties: properties}
}

func (g *Graph) AddEdge(from, to, label string, properties map[string]interface{}) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, fromExists := g.Nodes[from]; !fromExists {
		return errors.New("from node does not exist")
	}
	if _, toExists := g.Nodes[to]; !toExists {
		return errors.New("to node does not exist")
	}

	if g.Edges[from] == nil {
		g.Edges[from] = make(map[string]*Edge)
	}
	g.Edges[from][to] = &Edge{From: from, To: to, Label: label, Properties: properties}
	return nil
}

func (g *Graph) GetNode(id string) (*Node, bool) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	node, exists := g.Nodes[id]
	return node, exists
}

func (g *Graph) GetEdge(from, to string) (*Edge, bool) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	edge, exists := g.Edges[from][to]
	return edge, exists
}

func (g *Graph) GetNeighbors(id string) []*Node {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	neighbors := []*Node{}
	edges, exists := g.Edges[id]
	if !exists {
		return neighbors
	}
	for to := range edges {
		if node, exists := g.Nodes[to]; exists {
			neighbors = append(neighbors, node)
		}
	}
	return neighbors
}

func (g *Graph) FindNodesByProperty(key string, value interface{}) []*Node {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	result := []*Node{}
	for _, node := range g.Nodes {
		if prop, exists := node.Properties[key]; exists && prop == value {
			result = append(result, node)
		}
	}
	return result
}

func (g *Graph) SaveToFile(filename string) error {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	data, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (g *Graph) LoadFromFile(filename string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, g)
}
