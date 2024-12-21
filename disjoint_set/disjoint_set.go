// Disjoint-set data structure to calculate the Minimum Spanning Tree of a graph using Kruskal's algorithm
// https://en.wikipedia.org/wiki/Disjoint-set_data_structure#Applications

package disjoint_set

type DisjointSet struct {
	parent []int
	rank   []int
}

// Creates a new DSU with n elements
// https://en.wikipedia.org/wiki/Disjoint-set_data_structure#Making_new_sets
func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)

	// Initialize each element as its own set
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &DisjointSet{
		parent: parent,
		rank:   rank,
	}
}

/*
Recursively loops through the parents of x
Returns the representative element of the set containing x (x in our specific case is a vertex of a graph)
Uses path compression, making all nodes in path point to root
https://en.wikipedia.org/wiki/Disjoint-set_data_structure#Finding_set_representatives
*/
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

/*
Merges the sets containing elements x and y
Uses union by rank for optimization
https://en.wikipedia.org/wiki/Disjoint-set_data_structure#Union_by_rank
Very nice visualization (https://en.wikipedia.org/wiki/Disjoint-set_data_structure#/media/File:UnionFindKruskalDemo.gif)
*/
func (ds *DisjointSet) Union(x, y int) {
	// Replace nodes by roots
	xRoot := ds.Find(x)
	yRoot := ds.Find(y)

	if xRoot == yRoot {
		return
	}

	// Union by rank: attach smaller rank tree under root of higher rank tree
	if ds.rank[xRoot] < ds.rank[yRoot] {
		ds.parent[xRoot] = yRoot
	} else if ds.rank[xRoot] > ds.rank[yRoot] {
		ds.parent[yRoot] = xRoot
	} else {
		ds.parent[yRoot] = xRoot
		ds.rank[xRoot]++
	}
}

// Checks if x and y are in the same set
func (ds *DisjointSet) Connected(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}
