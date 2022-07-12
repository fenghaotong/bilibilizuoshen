package class10

import (
	"fmt"
	"testing"
)

func TestCountIsLands(t *testing.T) {
	arr := [][]int{
		{0, 1, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(CountIsLands(arr))
}

func TestNewUnionFindSet(t *testing.T) {
	m := make(map[int]string, 0)
	fmt.Println(m[0] == "")
}
