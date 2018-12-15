package graphs

import (
	"container/heap"
	"routely/data"
)

const inf = 10000000000000

// Dijkstra runs dijkstra on an adjacency list from a certain node
// O(n*log(v))
func Dijkstra(adjList [][]data.Road, from uint) ([]float64, []float64) {
	time := make([]float64, len(adjList))
	for i := range time {
		time[i] = inf
	}
	time[from] = 0

	distance := make([]float64, len(adjList))
	distance[from] = 0

	var pq data.PriorityQueue
	heap.Init(&pq)

	init := &data.DijkstraNode{To: from, Weight: 0}
	heap.Push(&pq, init)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*data.DijkstraNode)

		if time[cur.To] < cur.Weight {
			continue
		}

		for _, e := range adjList[cur.To] {
			if e.Weight+time[e.From] < time[e.To] {
				time[e.To] = time[e.From] + e.Weight
				distance[e.To] = distance[e.From] + e.Length
				heap.Push(&pq, &data.DijkstraNode{
					To:     e.To,
					Weight: time[e.To],
				})
			}
		}
	}

	return time, distance
}
