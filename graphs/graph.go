package graphs

import "routely/data"

// Graph reperesents a graph type in memory
type Graph struct {
	AdjList [][]data.Road
}

// NewGraph builds graph from vertices and edges
// O(V+E)
func NewGraph(vertices []*data.Intersection, edges []*data.Road) *Graph {
	sz := len(vertices)
	G := &Graph{}
	G.AdjList = make([][]data.Road, sz)

	for i := range G.AdjList {
		G.AdjList[i] = make([]data.Road, 0)
	}

	for _, E := range edges {
		E.Weights = make([]float64, len(E.Speeds))

		for idx := 0; idx < data.SpeedCount; idx++ {
			if E.Speeds[idx] == 0 {
				E.Weights[idx] = inf
			} else {
				E.Weights[idx] = E.Length / E.Speeds[idx]
			}
		}

		G.AdjList[E.From] = append(G.AdjList[E.From], *E)
		E.From, E.To = E.To, E.From
		G.AdjList[E.From] = append(G.AdjList[E.From], *E)
	}

	return G
}
