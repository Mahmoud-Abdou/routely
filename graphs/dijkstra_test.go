package graphs

import (
	"math"
	"routely/data"
	"testing"
)

func TestDijkstra(t *testing.T) {
	adjList := [][]data.Road{
		[]data.Road{
			data.Road{From: 0, To: 1, Length: 0.5, Speed: 20, Weight: 0.5 / 20, Index: 0},
			data.Road{From: 0, To: 3, Length: 0.28, Speed: 80, Weight: 0.28 / 80, Index: 2},
		},
		[]data.Road{
			data.Road{From: 1, To: 0, Length: 0.5, Speed: 20, Weight: 0.5 / 20, Index: 1},
			data.Road{From: 1, To: 4, Length: 0.28, Speed: 40, Weight: 0.28 / 40, Index: 6},
			data.Road{From: 1, To: 2, Length: 0.28, Speed: 40, Weight: 0.28 / 40, Index: 8},
		},
		[]data.Road{
			data.Road{From: 2, To: 1, Length: 0.28, Speed: 40, Weight: 0.28 / 40, Index: 9},
			data.Road{From: 2, To: 5, Length: 0.27, Speed: 60, Weight: 0.27 / 60, Index: 12},
		},
		[]data.Road{
			data.Road{From: 3, To: 0, Length: 0.28, Speed: 80, Weight: 0.28 / 80, Index: 3},
			data.Road{From: 3, To: 4, Length: 0.62, Speed: 80, Weight: 0.62 / 80, Index: 4},
		},
		[]data.Road{
			data.Road{From: 4, To: 3, Length: 0.62, Speed: 80, Weight: 0.62 / 80, Index: 5},
			data.Road{From: 4, To: 1, Length: 0.28, Speed: 40, Weight: 0.28 / 40, Index: 7},
			data.Road{From: 4, To: 5, Length: 0.27, Speed: 60, Weight: 0.27 / 60, Index: 10},
		},
		[]data.Road{
			data.Road{From: 5, To: 4, Length: 0.27, Speed: 60, Weight: 0.27 / 60, Index: 11},
			data.Road{From: 5, To: 2, Length: 0.27, Speed: 60, Weight: 0.27 / 60, Index: 13},
		},
	}

	expectedResult := [][]float64{
		[]float64{
			0.000000,
			0.018250,
			0.020250,
			0.003500,
			0.011250,
			0.015750,
		},
		[]float64{
			0.018250,
			0.000000,
			0.007000,
			0.014750,
			0.007000,
			0.011500,
		},
		[]float64{
			0.020250,
			0.007000,
			0.000000,
			0.016750,
			0.009000,
			0.004500,
		},
		[]float64{
			0.003500,
			0.014750,
			0.016750,
			0.000000,
			0.007750,
			0.012250,
		},
		[]float64{
			0.011250,
			0.007000,
			0.009000,
			0.007750,
			0.000000,
			0.004500,
		},
		[]float64{
			0.015750,
			0.011500,
			0.004500,
			0.012250,
			0.004500,
			0.000000,
		},
	}

	var result [][]float64
	for i := range expectedResult {
		time, _ := Dijkstra(adjList, uint(i))
		result = append(result, time)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if math.Abs(result[i][j]-expectedResult[i][j]) > 1e-9 {
				t.Errorf("Invalid Shortest Path between %v and %v, abs error %v", i, j, math.Abs(result[i][j]-expectedResult[i][j]))
			}
		}
	}
}
