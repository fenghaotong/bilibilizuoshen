package datastruct

import (
	"container/heap"
)

// A PriorityQueueSmall implements heap.Interface and holds Items.
type PriorityQueueBig []*Item

func (pq PriorityQueueBig) Len() int { return len(pq) }

func (pq PriorityQueueBig) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueueBig) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueueBig) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueBig) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueueBig) Pull() interface{} {
	old := *pq
	item := old[0]
	return item
}

// update modifies the Priority and Value of an Item in the queue.
func (pq *PriorityQueueBig) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
