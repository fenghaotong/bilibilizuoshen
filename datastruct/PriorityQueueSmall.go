package datastruct

import (
	"container/heap"
)

// A PriorityQueueSmall implements heap.Interface and holds Items.
type PriorityQueueSmall []*Item

func (pq PriorityQueueSmall) Len() int { return len(pq) }

func (pq PriorityQueueSmall) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueueSmall) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueueSmall) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueSmall) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueueSmall) Pull() interface{} {
	old := *pq
	item := old[0]
	return item
}

// update modifies the Priority and Value of an Item in the queue.
func (pq *PriorityQueueSmall) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
