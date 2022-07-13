package class12

import (
	"fmt"
	"testing"
)

func TestGetSlidingWindowMaxArray(t *testing.T) {
	arr := []int{5, 1, 3, 6}
	window := 3
	fmt.Println(GetSlidingWindowMaxArray(arr, window))
}
