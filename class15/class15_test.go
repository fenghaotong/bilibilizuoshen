package class15

import (
	"fmt"
	"testing"
)

func TestDPWay(t *testing.T) {
	x := 7
	y := 7
	k := 10
	fmt.Println(GetWay(x, y, k))
	fmt.Println(DPWay1(x, y, k))
	fmt.Println(DPWay(x, y, k))
}

func TestGetCoinsWayst(t *testing.T) {
	arr := []int{3, 5, 7, 10}
	aim := 100

	fmt.Println(GetCoinsWays(arr, aim))
	fmt.Println(GetCoinsWaysDP1(arr, aim))
	fmt.Println(GetCoinsWaysDP2(arr, aim))
}
