package algorithm

import (
	"routely/data"
	"routely/graphs"
)

// BestPath finds optimal route in query for graph
//O(n*m*log(dx*dx + dy*dy))
func BestPath(g *graphs.Graph, query *data.Query, circleFinder *CircleFinder) *data.Path {
	fromVertices := circleFinder.VerticesInCircle(query.From, query.WalkingRadius)
	toVertices := circleFinder.VerticesInCircle(query.To, query.WalkingRadius)

	var WalkingDistance = 1e9
	var DrivingDistance = 1e9
	var BestTime = 1e9
	var walkingTime float64
	var totalTime float64

	for _, from := range fromVertices {
		for _, to := range toVertices {
			walkingTime = from.Distance + to.Distance
			totalTime = g.ShortestTime[from.ID][to.ID] + (walkingTime)/5.0
			if totalTime < BestTime {
				BestTime = totalTime
				WalkingDistance = walkingTime
				DrivingDistance = g.Distances[from.ID][to.ID]
			}
		}
	}

	return &data.Path{
		DrivingDistance: DrivingDistance,
		WalkingDistance: WalkingDistance,
		Time:            BestTime * 60,
	}
}
