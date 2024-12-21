// Implementation of Kruskal's Minimum Spanning Tree algorithm
// https://en.wikipedia.org/wiki/Kruskal%27s_algorithm

package kruskal

import (
	"fmt"
	"sort"

	"github.com/nevernorbo/kruskal-mst/disjoint_set"
)

type Edge struct {
	src    int
	dest   int
	weight int
}

// Undirected weighted graph represented by an array of edges
type Graph struct {
	vertices int
	edges    []Edge
}

// Checks if graph is empty (used at input validation)
func (g *Graph) IsEmpty() bool {
	return len(g.edges) == 0
}

// Creates a new graph with v number of vertices
func NewGraph(v int) *Graph {
	return &Graph{
		vertices: v,
		edges:    make([]Edge, 0),
	}
}

// Adds an edge to the graph
func (g *Graph) AddEdge(src, dest, weight int) {
	g.edges = append(g.edges, Edge{src, dest, weight})
}

/*
Finds the Minimum Spanning Tree using Kruskal's algorithm
 1. Sort edges
 2. Walk through the sorted edges and check if the edge's source and destination is in the same set
    if yes: do nothing, because adding this edge would create a cycle
    if no: append the edge to the result and union the edge components
*/
func (g *Graph) KruskalMST() []Edge {
	// Sort edges by weight
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	// Initialize disjoint set
	ds := disjoint_set.NewDisjointSet(g.vertices)

	// Result will store the MST edges
	result := make([]Edge, 0)

	// Process edges in ascending order of weight
	for _, edge := range g.edges {
		// Check if adding this edge creates a cycle
		if ds.Find(edge.src) != ds.Find(edge.dest) {
			// Add edge to MST and union the sets
			result = append(result, edge)
			ds.Union(edge.src, edge.dest)
		}
	}

	return result
}

// Prints the result of the algorithm
func DisplayKruskal(mst []Edge) {
	totalWeight := 0

	fmt.Println("-------------------------------------")
	fmt.Println("      The Minimum Spanning Tree      ")
	fmt.Println("[source] -- [destination] : [weight]")
	for _, edge := range mst {
		fmt.Printf("%d -- %d : %d\n", edge.src, edge.dest, edge.weight)
		totalWeight += edge.weight
	}

	fmt.Println()
	fmt.Printf("Total MST weight: %d\n", totalWeight)
	fmt.Println("-------------------------------------")
}
