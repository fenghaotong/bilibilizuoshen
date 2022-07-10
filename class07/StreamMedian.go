package class07

import (
	"algorithm/datastruct"
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	MaxHeap datastruct.PriorityQueueBig
	MinHeap datastruct.PriorityQueueSmall
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{
		MaxHeap: make(datastruct.PriorityQueueBig, 0),
		MinHeap: make(datastruct.PriorityQueueSmall, 0),
	}
}

func (SM *MedianFinder) AddNum(num int) {
	item := &datastruct.Item{
		Value:    fmt.Sprintf("%d", num),
		Priority: num,
		Index:    num,
	}
	if SM.MaxHeap.Len() == 0 {
		heap.Push(&SM.MaxHeap, item)
	} else {
		priority := SM.MaxHeap.Pull().(*datastruct.Item).Priority
		if priority <= num {
			heap.Push(&SM.MinHeap, item)
		} else {
			heap.Push(&SM.MaxHeap, item)
		}

		if SM.MaxHeap.Len()-SM.MinHeap.Len() >= 2 {
			i := heap.Pop(&SM.MaxHeap).(*datastruct.Item)
			heap.Push(&SM.MinHeap, i)
		} else if SM.MinHeap.Len()-SM.MaxHeap.Len() >= 2 {
			i := heap.Pop(&SM.MinHeap).(*datastruct.Item)
			heap.Push(&SM.MaxHeap, i)
		}
	}

}

func (SM *MedianFinder) PrintPQ() {
	for SM.MaxHeap.Len() > 0 {
		item := heap.Pop(&SM.MaxHeap).(*datastruct.Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}

	fmt.Println(".....")
	for SM.MinHeap.Len() > 0 {
		item := heap.Pop(&SM.MinHeap).(*datastruct.Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
}

func (SM *MedianFinder) FindMedian() float64 {
	if SM.MaxHeap.Len() == SM.MinHeap.Len() {
		return float64(SM.MaxHeap.Pull().(*datastruct.Item).Priority+SM.MinHeap.Pull().(*datastruct.Item).Priority) / 2.0
	} else if SM.MaxHeap.Len() > SM.MinHeap.Len() {
		return float64(SM.MaxHeap.Pull().(*datastruct.Item).Priority)
	} else {
		return float64(SM.MinHeap.Pull().(*datastruct.Item).Priority)
	}
}
