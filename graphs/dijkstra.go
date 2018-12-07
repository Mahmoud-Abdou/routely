package graphs

import (
	"container/heap"
	"routely/data"
)

const inf = 10000000000000

// Dijkstra runs dijkstra on an adjacency list from a certain node
func Dijkstra(adjList [][]data.Road, from uint) []float64 {
	cost := make([]float64, len(adjList))
	for i := range cost {
		cost[i] = inf
	}
	cost[from] = 0

	var pq data.PriorityQueue
	heap.Init(&pq)

	init := &data.Road{From: 0, To: from, Length: 0, Speed: 0, Index: 0, Weight: 0}
	heap.Push(&pq, init)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*data.Road)

		if cost[cur.To] < cur.Weight {
			continue
		}

		for _, e := range adjList[cur.To] {
			if e.Weight+cost[e.From] < cost[e.To] {
				cost[e.To] = cost[e.From] + e.Weight
				var next = e
				e.Weight = cost[e.To]
				heap.Push(&pq, &next)
			}
		}
	}

	return cost
}
