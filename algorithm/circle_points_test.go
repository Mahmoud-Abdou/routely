package algorithm

import (
	"routely/data"
	"testing"
	"fmt"
)

func TestVerticesInCircle(t *testing.T) {
	center := data.Point{
		X: 0.16,
		Y: 0.21,
	}
	vertices := []*data.Intersection{{0, data.Point{0.25, 0.33}}, {1,data.Point{0.73, 0.47}}, {2,data.Point{1.00, 0.54}}, {3,data.Point{0.34, 0.60}}, {4, data.Point{0.81, 0.20}}, {5,data.Point{1.07, 0.28}}}
	cf := NewCircleFinder(vertices)

	result :=  cf.VerticesInCircle(center, 300)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
