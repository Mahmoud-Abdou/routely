package algorithm

import (
	"routely/data"
	"math"
)

// CircleFinder object holds vertices in order to find which ones fit in a circle
type CircleFinder struct {
	Vertices []*data.Intersection
}

// NewCircleFinder initializes preprocessing for CircleFinder
func NewCircleFinder(vertices []*data.Intersection) *CircleFinder {
	circlefinder := &CircleFinder{}
	circlefinder.Vertices = vertices

	return circlefinder
}

func getDistance(firstPoint data.Point, secondPoint data.Point) float64 {
	xd := secondPoint.X - firstPoint.X
	yd := secondPoint.Y - firstPoint.Y
	return math.Sqrt(xd*xd + yd*yd)
}

// VerticesInCircle returns all vertices of CircleFinder that fit inside the circle given
func (c *CircleFinder) VerticesInCircle(center data.Point, radius float64) []*data.CircleVertex {
	validVertices := make([]*data.CircleVertex, 0)
	for _, vertex := range c.Vertices {
		dist := getDistance(center, vertex.Point)
		if dist-radius < 1e-9 {
			validVertices = append(validVertices, &data.CircleVertex{
				Distance:     dist,
				Intersection: vertex,
			})
		}
	}
	return validVertices
}
