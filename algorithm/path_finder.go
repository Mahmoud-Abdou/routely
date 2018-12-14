package algorithm

import (
	"routely/data"
	"routely/graphs"
)

// BestPath finds optimal route in query for graph
func BestPath(g *graphs.Graph, query *data.Query, circleFinder *CircleFinder) *data.Path {
	fromVertices := circleFinder.VerticesInCircle(query.From, query.WalkingRadius)
	toVertices := circleFinder.VerticesInCircle(query.To, query.WalkingRadius)

	var WalkingDistance = 1e9
	var DrivingDistance = 1e9
	var BestTime = 1e9

	for _, from := range fromVertices {
		for _, to := range toVertices {
			time := g.GetTime(from.ID, to.ID)
			walkingDistanceSource := getDistance(query.From, from.Point)
			walkingDistanceDestination := getDistance(query.To, to.Point)
			if time + (walkingDistanceSource + walkingDistanceDestination)/5.0 < BestTime {
				BestTime = time + (walkingDistanceSource + walkingDistanceDestination)/5.0
				WalkingDistance = walkingDistanceSource + walkingDistanceDestination
				DrivingDistance = g.GetDrivingDistance(from.ID, to.ID)
			}
		}
	}

	return &data.Path{
		DrivingDistance: DrivingDistance,
		WalkingDistance: WalkingDistance,
		Time:          	 BestTime*60,
	}
}
