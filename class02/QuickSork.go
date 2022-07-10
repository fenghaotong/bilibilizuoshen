package class02

import "math/rand"

func QuickSortVersion1(arr []int) {
	length := len(arr)
	if arr == nil || length < 2 {
		return
	}

	process1(arr, 0, length-1)
}

func process1(arr []int, L, R int) {
	if L < R {
		p := partition1(arr, L, R)
		process1(arr, L, p[0]-1)
		process1(arr, p[0]+1, R)
	}
}

//1, 3, 4, 2, 5, 20, 19, 8, 9, 10
//1, 3, 4, 2, 5, 8, 19, 20, 9, 10
//1, 3, 4, 2, 5, 8, 9, 20, 19, 10
//1, 3, 4, 2, 5, 8, 9, 10, 19, 20

func partition1(arr []int, L, R int) []int {
	less := L - 1 // 左边界
	for L < R {
		if arr[L] < arr[R] {
			less++
			Swap(arr, less, L)
			L++
		} else {
			L++
		}
	}
	Swap(arr, less+1, R)
	return []int{less + 1}
}

func Swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func QuickSortVersion2(arr []int) {
	length := len(arr)
	if arr == nil || length < 2 {
		return
	}

	process2(arr, 0, length-1)
}

func process2(arr []int, L, R int) {
	if L < R {
		p := partition2(arr, L, R)
		process2(arr, L, p[0]-1)
		process2(arr, p[1]+1, R)
	}
}

//1, 3, 4, 2, 5, 20, 19, 8, 9, 10
//1, 3, 4, 2, 5, 9, 19, 8, 20, 10
//1, 3, 4, 2, 5, 9, 8, 19, 20, 10
//1, 3, 4, 2, 5, 9, 8, 10, 20, 19
func partition2(arr []int, L, R int) []int {
	less := L - 1 // 左边界
	more := R     // 右边界
	for L < more {
		if arr[L] < arr[R] {
			less++
			Swap(arr, less, L)
			L++
		} else if arr[L] > arr[R] {
			more--
			Swap(arr, more, L)
		} else {
			L++
		}
	}
	Swap(arr, more, R)
	return []int{less + 1, more}
}

func QuickSortVersion3(arr []int) {
	length := len(arr)
	if arr == nil || length < 2 {
		return
	}

	process3(arr, 0, length-1)

}

func process3(arr []int, L, R int) {
	if L < R {
		Swap(arr, L+rand.Intn(R-L+1), R)
		p := partition3(arr, L, R)
		process3(arr, L, p[0]-1)
		process3(arr, p[1]+1, R)
	}
}

func partition3(arr []int, L, R int) []int {
	less := L - 1
	more := R
	for L < more {
		if arr[L] < arr[R] {
			less++
			Swap(arr, less, L)
			L++
		} else if arr[L] > arr[R] {
			more--
			Swap(arr, more, L)
		} else {
			L++
		}
	}
	Swap(arr, more, R)
	return []int{less + 1, more}
}
