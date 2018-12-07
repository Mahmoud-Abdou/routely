package graphs

import (
	"container/heap"
	"routely/data"
)

const inf = 10000000000000

// Dijkstra runs dijkstra on graph from a certain node
func Dijkstra(graph Graph, from uint) []float64 {
	cost := make([]float64, len(graph.AdjList))
	for i := range cost {
		cost[i] = inf
	}
	cost[from] = 0

	var pq data.PriorityQueue
	heap.Init(&pq)

	init := data.Road{From: 0, To: from, Length: 0, Speed: 0, Index: 0, Weight: 0}
	heap.Push(&pq, init)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*data.Road)

		if cost[cur.To] < cur.Weight {
			continue
		}

		for _, e := range graph.AdjList[cur.To] {
			if e.Weight+cost[e.From] < cost[e.To] {
				cost[e.To] = cost[e.From] + e.Weight
				var next = e
				e.Weight = cost[e.To]
				heap.Push(&pq, next)
			}
		}
	}

	return cost
}
