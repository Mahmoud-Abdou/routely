package graphs

import "routely/data"
import "fmt"

// Graph reperesents a graph type in memory
type Graph struct {
	AdjList       [][]data.Road
	ShortestPaths [][]float64
}

// NewGraph builds graph from vertices and edges
func NewGraph(vertices []*data.Intersection, edges []*data.Road) *Graph {
	sz := len(vertices)
	G := &Graph{}
	G.AdjList = make([][]data.Road, sz)
	G.ShortestPaths = make([][]float64, sz)

	for i := range G.AdjList {
		G.AdjList[i] = make([]data.Road, 0)
		G.ShortestPaths[i] = make([]float64, 0)
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
		cost := Dijkstra(G.AdjList, uint(i))
		for j := range cost{
				G.ShortestPaths[i] = append(G.ShortestPaths[i],cost[j])
		}
	}
	return G
}

// Distance Gets shortest path between 2 nodes
func (g *Graph) Distance(from, to uint) float64 {
	fmt.Println("from: ",from)
	fmt.Println("len: ", to)
	fmt.Println("shortest: ", len(g.ShortestPaths))
	fmt.Println("shortest from: ", len(g.ShortestPaths[from]))
	fmt.Println("--------------------------------------")
	return g.ShortestPaths[from][to]
}
