package graphs

import "routely/data"
import "testing"

func TestNewGraph(t *testing.T) {

	vert := []*data.Intersection{}
	for i := 0; i < 5; i++ {
		var inter data.Intersection
		inter.ID = i
		inter.X = 1
		inter.Y = 2
		vert = append(vert, &inter)
	}
	var edg []*data.Road
	e1 := data.Road{
		From:   0,
		To:     1,
		Length: 1,
		Speed:  2,
	}
	e2 := data.Road{
		From:   1,
		To:     4,
		Length: 1,
		Speed:  2,
	}
	e3 := data.Road{
		From:   3,
		To:     0,
		Length: 1,
		Speed:  2,
	}
	e4 := data.Road{
		From:   2,
		To:     1,
		Length: 1,
		Speed:  2,
	}

	edg = append(edg, &e1)
	edg = append(edg, &e2)
	edg = append(edg, &e3)
	edg = append(edg, &e4)

	var g *Graph
	g = NewGraph(vert, edg)

	if g.AdjList[0][0].To != 1 || g.AdjList[0][1].To != 3 {
		t.Errorf("Incorrect adjacency list, at row 0 found %d %d", g.AdjList[0][0].To, g.AdjList[0][1].To)
	}
	if g.AdjList[1][0].To != 0 || g.AdjList[1][1].To != 4 || g.AdjList[1][2].To != 2 {
		t.Errorf("Incorrect adjacency list, at row 1 found %d %d %d", g.AdjList[1][0].To, g.AdjList[1][1].To, g.AdjList[1][2].To)
	}
	if g.AdjList[2][0].To != 1 {
		t.Errorf("Incorrect adjacency list, at row 0 found %d", g.AdjList[2][0].To)
	}
}
