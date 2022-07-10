package class03

import (
	"algorithm/utils"
	"testing"
)

func TestRadixSort(t *testing.T) {
	utils.SortCallback(RadixSort)
}
