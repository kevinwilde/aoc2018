package priorityqueue

import "container/heap"

// A stringHeap is a min-heap of strings
type stringHeap []string

func (h stringHeap) Len() int           { return len(h) }
func (h stringHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h stringHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *stringHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(string))
}

func (h *stringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type PriorityQueue struct {
	h *stringHeap
}

func NewPriorityQueue() PriorityQueue {
	h := make(stringHeap, 0)
	return PriorityQueue{&h}
}

func (pq *PriorityQueue) Push(s string) {
	heap.Push(pq.h, s)
}

func (pq *PriorityQueue) Pop() string {
	return heap.Pop(pq.h).(string)
}

func (pq *PriorityQueue) Len() int {
	return len(*pq.h)
}
