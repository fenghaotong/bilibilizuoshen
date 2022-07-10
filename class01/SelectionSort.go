package class01

// SelectionSort
// 选择第一小的值放到第一个位置
// 选择第二小的值放到第二个位置
func SelectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	len := len(arr)
	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		Swap(arr, i, minIndex)
	}
}

// Swap
// 这个没有使用位运算是因为，如果说同一个位置的值发生交换，使用异或之后值变为0
func Swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
