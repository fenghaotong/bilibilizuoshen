package class01

import (
	"algorithm/utils"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	utils.SortCallback(SelectionSort)
}

func TestBubbleSort(t *testing.T) {
	utils.SortCallback(BubbleSort)
}

func TestInsertSort(t *testing.T) {
	utils.SortCallback(InsertSort)
}

func TestPrintEvenTimesOddTime(t *testing.T) {
	arr := []int{1, 1, 2, 3, 3, 4, 4, 4, 4}
	PrintEvenTimes(arr)
	PrintOddTimes(arr)
}
