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
	distance := math.Sqrt((secondPoint.X-firstPoint.X)*(secondPoint.X-firstPoint.X) + (secondPoint.Y-firstPoint.Y)*(secondPoint.Y-firstPoint.Y))
	return distance
}

// VerticesInCircle returns all vertices of CircleFinder that fit inside the circle given
func (c *CircleFinder) VerticesInCircle(center data.Point, radius float64) []*data.Intersection {
	validVertices := make([]*data.Intersection, 0)
	for _, vertex := range c.Vertices {
		if (getDistance(center, vertex.Point) - radius < 1e-9) {
			validVertices = append(validVertices, vertex)
		}
	}
	return validVertices
}
