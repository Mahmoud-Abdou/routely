package algorithm

import (
	"routely/data"
	"routely/graphs"
)

// BestPath finds optimal route in query for graph
// O(n*m*log(dx*dx + dy*dy))
func BestPath(g *graphs.Graph, query *data.Query, circleFinder *CircleFinder) *data.Path {
	fromVertices := circleFinder.VerticesInCircle(query.From, query.WalkingRadius)
	toVertices := circleFinder.VerticesInCircle(query.To, query.WalkingRadius)

	bestTime, walkingDistance, drivingDistance := g.Dijkstra(fromVertices, toVertices)

	return &data.Path{
		DrivingDistance: drivingDistance,
		WalkingDistance: walkingDistance,
		Time:            bestTime * 60,
	}
}
