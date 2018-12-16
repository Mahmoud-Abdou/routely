package reader

import (
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	const map1 = "6 0 0.25 0.33 1 0.73 0.47 2 1.00 0.54 3 0.34 0.60 4 0.81 0.20 5 1.07 0.28 7 0 1 0.50 20 0 3 0.28 80 3 4 0.62 80 1 4 0.28 40 1 2 0.28 40 4 5 0.27 60 2 5 0.27 60"
	const q1 = "1 0.16 0.21 1.09 0.44 300"
	intr, rood, queries := ReadData(strings.NewReader(map1), strings.NewReader(q1))
	if len(intr) != 6 {
		t.Errorf("error size of intersections")
	}
	if len(rood) != 7 {
		t.Errorf("error size of roods")
	}
	if len(queries) != 1 {
		t.Errorf("error size of queries")
	}
	var xs = []float64{0.25, 0.73, 1.00, 0.34, 0.81, 1.07}
	var ys = []float64{0.33, 0.47, 0.54, 0.60, 0.20, 0.28}
	for i := 0; i < len(intr); i++ {
		if intr[i].ID != i {
			t.Errorf("error wrong ID")
		}
		if intr[i].X != xs[i] {
			t.Errorf("error wrong X")
		}
		if intr[i].Y != ys[i] {
			t.Errorf("error wrong Y")
		}
	}
	var from = []int{0, 0, 3, 1, 1, 4, 2}
	var to = []int{1, 3, 4, 4, 2, 5, 5}
	var leng = []float64{0.50, 0.28, 0.62, 0.28, 0.28, 0.27, 0.27}
	var speed = []float64{20, 80, 80, 40, 40, 60, 60}
	for i := 0; i < len(rood); i++ {
		if rood[i].From != from[i] {
			t.Errorf("error wrong from")
		}
		if rood[i].To != to[i] {
			t.Errorf("error wrong from")
		}
		if rood[i].Length != leng[i] {
			t.Errorf("error wrong length")
		}
		if rood[i].Speed != speed[i] {
			t.Errorf("error wrong speed")
		}
	}
	if queries[0].From.X != 0.16 {
		t.Errorf("error wrong q f x")
	}
	if queries[0].From.Y != 0.21 {
		t.Errorf("error wrong q f y")
	}
	if queries[0].To.X != 1.09 {
		t.Errorf("error wrong q t x exp 1.09 found %v", queries[0].To.X)
	}
	if queries[0].To.Y != 0.44 {
		t.Errorf("error wrong q t y")
	}
	if queries[0].WalkingRadius != 300 {
		t.Errorf("error wrong q R")
	}
}
