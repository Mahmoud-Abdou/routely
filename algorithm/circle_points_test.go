package algorithm

import (
	"routely/data"
	"testing"
)

func TestVerticesInCircle(t *testing.T) {
	center := data.Point{
		X: 2.00,
		Y: 2.00,
	}
	vertices := []*data.Intersection{{0, data.Point{1.00, 1.00}}, {1,data.Point{3.00, 1.00}}, {2,data.Point{1.00, 4.00}}, {3,data.Point{3.00, 4.00}}}
	cf := NewCircleFinder(vertices)
	result :=  cf.VerticesInCircle(center, 1.414213562)
	answer := []*data.Intersection{{0, data.Point{1.00, 1.00}}, {1,data.Point{3.00, 1.00}}}
	for i := 0; i < len(answer); i++ {
		if result[i].X != answer[i].X && result[i].Y != answer[i].Y  {
			t.Errorf("points was incorrect, got: %f, want: %f.", result[i], answer[i])
		}
	}
}
