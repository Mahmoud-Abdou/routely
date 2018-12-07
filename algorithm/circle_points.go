package algorithm

import "routely/data"

// CircleFinder object holds vertices in order to find which ones fit in a circle
type CircleFinder struct {
}

// NewCircleFinder initializes preprocessing for CircleFinder
func NewCircleFinder(vertices []*data.Intersection) *CircleFinder {
	return nil
}

// VerticesInCircle returns all vertices of CircleFinder that fit inside the circle given
func (c *CircleFinder) VerticesInCircle(center data.Point, radius float64) []*data.Intersection {
	return nil
}
