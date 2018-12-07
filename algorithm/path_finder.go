package algorithm

import (
	"routely/data"
	"routely/graphs"
)

// BestPath finds optimal route in query for graph
func BestPath(g *graphs.Graph, query *data.Query, circleFinder *CircleFinder) *data.Path {
	fromVertices := circleFinder.VerticesInCircle(query.From, query.WalkingRadius)
	toVertices := circleFinder.VerticesInCircle(query.To, query.WalkingRadius)

	var bestWalking = 1e9
	var bestDriving = 1e9
	var bestTime = 1e9

	for _, from := range fromVertices {
		for _, to := range toVertices {
			distance := g.Distance(from.ID, to.ID)
			walkingDistanceSource := getDistance(query.From, from.Point)
			walkingDistanceDestination := getDistance(query.To, to.Point)
			if distance+(walkingDistanceSource+walkingDistanceDestination)/5.0 < bestTime {
				bestTime = distance + (walkingDistanceSource+walkingDistanceDestination)/5.0
				bestWalking = walkingDistanceSource + walkingDistanceDestination
				bestDriving = distance
			}
		}
	}

	return &data.Path{
		DrivingDistance: bestDriving,
		WalkingDistance: bestWalking,
		Length:          bestTime,
	}
}
