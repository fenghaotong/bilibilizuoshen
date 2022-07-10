package class02

import (
	"algorithm/datastruct"
	"container/heap"
	"fmt"
)

func headInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		Swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index, heapSize int) {
	leftChild := index*2 + 1
	largest := index
	for leftChild < heapSize {
		rightChild := leftChild + 1
		if rightChild < heapSize && arr[rightChild] > arr[leftChild] { // 如果存在右孩子，并且右孩子的值大于左孩子的值
			largest = rightChild
		} else {
			largest = leftChild
		}

		if arr[largest] <= arr[index] { // 如果父节点的值大于孩子的值
			largest = index
			return
		}

		Swap(arr, largest, index)
		index = largest
		leftChild = index*2 + 1

	}
}

func HeapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr); i++ {
		headInsert(arr, i)
	}

	heapSize := len(arr) - 1
	Swap(arr, 0, heapSize)
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize--
		Swap(arr, 0, heapSize)
	}

}

func HeapStd() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(datastruct.PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &datastruct.Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &datastruct.Item{
		Value:    "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*datastruct.Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}
