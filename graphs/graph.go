package graphs

import "routely/data"

// Graph reperesents a graph type in memory
type Graph struct {
	AdjList      [][]data.Road
	ShortestTime [][]float64
	Distances    [][]float64
}

// NewGraph builds graph from vertices and edges
func NewGraph(vertices []*data.Intersection, edges []*data.Road) *Graph {
	sz := len(vertices)
	G := &Graph{}
	G.AdjList = make([][]data.Road, sz)
	G.ShortestTime = make([][]float64, sz)
	G.Distances = make([][]float64, sz)

	for i := range G.AdjList {
		G.AdjList[i] = make([]data.Road, 0)
		G.ShortestTime[i] = make([]float64, 0)
		G.Distances[i] = make([]float64, 0)
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

	for i := range vertices {
		cost, dist := Dijkstra(G.AdjList, uint(i))
		G.ShortestTime[i] = cost
		G.Distances[i] = dist
	}
	return G
}

// Distance Gets shortest path between 2 nodes
func (g *Graph) GetTime(from, to uint) float64 {
	return g.ShortestTime[from][to]
}
func (g *Graph) GetDrivingDistance(from, to uint) float64 {
	return g.Distances[from][to]
}
