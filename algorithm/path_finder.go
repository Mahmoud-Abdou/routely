package algorithm

import (
	"routely/data"
	"routely/graphs"
)

// BestPath finds optimal route in query for graph
// O(E*log(V) + V*log(D))
func BestPath(g *graphs.Graph, query *data.Query, circleFinder *CircleFinder) *data.Path {
	fromVertices := circleFinder.VerticesInCircle(query.From, query.WalkingRadius) // O(V*log(D))
	toVertices := circleFinder.VerticesInCircle(query.To, query.WalkingRadius)     // O(V*log(D))

	bestTime, walkingDistance, drivingDistance := g.Dijkstra(fromVertices, toVertices) // O(E*log(V))

	return &data.Path{
		DrivingDistance: drivingDistance,
		WalkingDistance: walkingDistance,
		Time:            bestTime * 60,
	}
}
