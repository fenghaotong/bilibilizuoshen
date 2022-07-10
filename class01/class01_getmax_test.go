package class01

import (
	"fmt"
	"testing"
)

func TestGetMax(t *testing.T) {
	arr := GenRandIntArr(10, 100)
	fmt.Println(arr)
	fmt.Println(GetMax(arr, 0, len(arr)-1))
}
