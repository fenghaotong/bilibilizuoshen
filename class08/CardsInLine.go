package class08

import "math"

func Win(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	return int(math.Max(float64(f(arr, 0, len(arr)-1)), float64(s(arr, 0, len(arr)-1))))
}

// 先手函数
func f(arr []int, L int, R int) int {
	if L == R {
		return arr[L]
	}

	return int(math.Max(float64(arr[L]+s(arr, L+1, R)), float64(arr[R]+s(arr, L, R-1))))
}

// 后手函数
func s(arr []int, L, R int) int {
	if L == R {
		return 0
	}

	return int(math.Min(float64(f(arr, L+1, R)), float64(f(arr, L, R-1))))
}
