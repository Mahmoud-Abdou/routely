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
	vertices := []*data.Intersection{
		&data.Intersection{
			ID:    0,
			Point: data.Point{X: 1.00, Y: 1.00},
		},
		&data.Intersection{
			ID:    1,
			Point: data.Point{X: 3.00, Y: 1.00},
		},
		&data.Intersection{
			ID:    2,
			Point: data.Point{X: 1.00, Y: 4.00},
		},
		&data.Intersection{
			ID:    3,
			Point: data.Point{X: 3.00, Y: 4.00},
		},
	}
	cf := NewCircleFinder(vertices)
	result := cf.VerticesInCircle(center, 1.414213562)
	answer := []*data.Intersection{
		{
			ID:    0,
			Point: data.Point{X: 1.00, Y: 1.00},
		},
		{
			ID:    1,
			Point: data.Point{X: 3.00, Y: 1.00},
		},
	}
	for i := 0; i < len(answer); i++ {
		if result[i].X != answer[i].X && result[i].Y != answer[i].Y {
			t.Errorf("points was incorrect, got: %+v, want: %+v.", result[i], answer[i])
		}
	}
}
