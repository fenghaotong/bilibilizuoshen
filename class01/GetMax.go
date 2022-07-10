package class01

import "math"

func GetMax(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}

	mid := L + (R-L)>>2
	leftMax := GetMax(arr, L, mid)
	rightMax := GetMax(arr, mid+1, R)
	return int(math.Max(float64(leftMax), float64(rightMax)))
}
