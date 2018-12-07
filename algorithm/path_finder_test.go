package algorithm

import (
	"routely/data"
	"routely/graphs"
	"testing"
)

func TestBestPath(t *testing.T) {
	graph := &graphs.Graph{
		ShortestPaths: [][]float64{
			[]float64{
				1, 2, 3, 4,
			},
			[]float64{
				2, 1, 4, 3,
			},
			[]float64{
				4, 3, 2, 1,
			},
			[]float64{
				3, 4, 1, 2,
			},
		},
	}

	query := &data.Query{
		From:          data.Point{X: 1, Y: 1},
		To:            data.Point{X: 3, Y: 3},
		WalkingRadius: 1.0,
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

	circleFinder := NewCircleFinder(vertices)

	result := BestPath(graph, query, circleFinder)
	if result.DrivingDistance != 4 {
		t.Errorf("Driving distance, expected 4, found %f", result.DrivingDistance)
	}
	if result.WalkingDistance != 1 {
		t.Errorf("Walking distance, expected 1, found %f", result.WalkingDistance)
	}
	if result.Length != 5 {
		t.Errorf("Total distance, expected 5, found %f", result.Length)
	}
}
