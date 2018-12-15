package data

// PriorityQueue of roads
type PriorityQueue []*DijkstraNode

// Len gets the length of the PriorityQueue
// O(1)
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two items
// O(1)
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

// Pop pops an element from the PriorityQueue
// O(log(n))
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// Push pushes an element to the PriorityQueue
// O(log(n))
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*DijkstraNode)
	item.Index = n
	*pq = append(*pq, item)
}

// O(1)
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
