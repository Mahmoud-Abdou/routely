package graphs

import "routely/data"

// Graph reperesents a graph type in memory
type Graph struct {
	AdjList [][]data.Road
}

// NewGraph builds graph from vertices and edges
// O(n*n*log(v))
func NewGraph(vertices []*data.Intersection, edges []*data.Road) *Graph {
	sz := len(vertices)
	G := &Graph{}
	G.AdjList = make([][]data.Road, sz)

	for i := range G.AdjList {
		G.AdjList[i] = make([]data.Road, 0)
	}

	for _, E := range edges {
		if E.Speed == 0 {
			E.Weight = inf
		} else {
			E.Weight = E.Length / E.Speed
		}
		G.AdjList[E.From] = append(G.AdjList[E.From], *E)
		E.From, E.To = E.To, E.From
		G.AdjList[E.From] = append(G.AdjList[E.From], *E)
	}

	return G
}
