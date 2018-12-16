package graphs

import (
	"container/heap"
	"routely/data"
)

const inf = 10000000000000

// Dijkstra runs dijkstra on an adjacency list from a certain node
// O(n*log(v))
func (g *Graph) Dijkstra(sources []*data.CircleVertex, destinations []*data.CircleVertex) (minTime float64, walkingDistance float64, drivingDistance float64) {
	time := make([]float64, len(g.AdjList))
	distance := make([]float64, len(g.AdjList))
	path := make([]int, len(g.AdjList))

	for i := range time {
		time[i] = inf
		path[i] = -1
	}

	var pq data.PriorityQueue
	heap.Init(&pq)

	for _, from := range sources {
		time[from.ID] = from.Distance / 5.0
		distance[from.ID] = from.Distance
		next := &data.DijkstraNode{To: from.ID, Weight: time[from.ID]}
		heap.Push(&pq, next)
	}

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*data.DijkstraNode)

		if time[cur.To] < cur.Weight {
			continue
		}

		for _, e := range g.AdjList[cur.To] {
			if e.Weight+time[e.From] < time[e.To] {
				path[e.To] = e.From
				time[e.To] = time[e.From] + e.Weight
				distance[e.To] = distance[e.From] + e.Length
				heap.Push(&pq, &data.DijkstraNode{
					To:     e.To,
					Weight: time[e.To],
				})
			}
		}
	}

	minTime = 1e9
	walkingDistance = 1e9
	drivingDistance = 1e9
	totalDistance := 1e9
	finalDestination := -1
	destinationDistance := 0.0

	for _, to := range destinations {
		if time[to.ID]+to.Distance/5.0 < minTime {
			minTime = time[to.ID] + to.Distance/5.0
			destinationDistance = to.Distance
			finalDestination = to.ID
			totalDistance = distance[to.ID]
		}
	}

	builtPath := g.buildPath(finalDestination, path)
	finalSource := builtPath[len(builtPath)-1]

	drivingDistance = totalDistance - distance[finalSource]
	walkingDistance = distance[finalSource] + destinationDistance

	return
}

func (g *Graph) buildPath(des int, path []int) (builtPath []int) {
	builtPath = make([]int, 0)
	cur := des

	for cur != -1 {
		builtPath = append(builtPath, cur)
		cur = path[cur]
	}
	return
}
