package class02

import (
	"algorithm/utils"
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	utils.SortCallback(MergeSort)
}

func TestSmallSum(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}
	s := &SmallSum{}
	fmt.Println(s.Run(arr))
	arr = []int{1}
	fmt.Println(s.Run(arr))
}

func TestNewQuickSort(t *testing.T) {
	utils.SortCallback(QuickSortVersion1)
	utils.SortCallback(QuickSortVersion2)
	utils.SortCallback(QuickSortVersion3)
}

func TestHeapSort(t *testing.T) {
	utils.SortCallback(HeapSort)
	HeapStd()
}
