package graphs

import (
	"container/heap"
	"routely/data"
)

const inf = 10000000000000

// Dijkstra runs dijkstra on an adjacency list from a list of nodes
// O(E*log(V))
func (g *Graph) Dijkstra(sources []*data.CircleVertex, destinations []*data.CircleVertex) (minTime float64, walkingDistance float64, drivingDistance float64) {
	time := make([]float64, len(g.AdjList))
	distance := make([]float64, len(g.AdjList))
	path := make([]int, len(g.AdjList))

	for i := range time { // O(V)
		time[i] = inf
		path[i] = -1
	}

	var pq data.PriorityQueue
	heap.Init(&pq) // O(1)

	// O(Vi * log(Vi))
	// Vi = number of sources
	for _, from := range sources {
		time[from.ID] = from.Distance / 5.0
		distance[from.ID] = from.Distance
		next := &data.DijkstraNode{To: from.ID, Weight: time[from.ID]}
		heap.Push(&pq, next)
	}

	// O(E * log(V))
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*data.DijkstraNode) // O(log(V))

		if time[cur.To] < cur.Weight {
			continue
		}

		// O(E')
		// E' = number of edges per node
		// Amortized O(E)
		for _, e := range g.AdjList[cur.To] {
			speedIndex := int(time[cur.To]) / data.SpeedInterval
			for speedIndex >= data.SpeedCount {
				speedIndex -= data.SpeedCount
			}

			if e.Weights[speedIndex]+time[e.From] < time[e.To] {
				path[e.To] = e.From
				time[e.To] = time[e.From] + e.Weights[speedIndex]
				distance[e.To] = distance[e.From] + e.Length
				heap.Push(&pq, &data.DijkstraNode{
					To:     e.To,
					Weight: time[e.To],
				}) // O(log(V))
			}
		}
	}

	minTime = 1e9
	walkingDistance = 1e9
	drivingDistance = 1e9
	totalDistance := 1e9
	finalDestination := -1
	destinationDistance := 0.0

	// O(Vf)
	// Vi = number of destinations
	for _, to := range destinations {
		if time[to.ID]+to.Distance/5.0 < minTime {
			minTime = time[to.ID] + to.Distance/5.0
			destinationDistance = to.Distance
			finalDestination = to.ID
			totalDistance = distance[to.ID]
		}
	}

	builtPath := g.buildPath(finalDestination, path) // O(V')
	finalSource := builtPath[len(builtPath)-1]       // O(1)

	drivingDistance = totalDistance - distance[finalSource]
	walkingDistance = distance[finalSource] + destinationDistance

	return
}

// O(V')
// V' = number of nodes in path
func (g *Graph) buildPath(des int, path []int) (builtPath []int) {
	builtPath = make([]int, 0)
	cur := des

	for cur != -1 {
		builtPath = append(builtPath, cur)
		cur = path[cur]
	}
	return
}
