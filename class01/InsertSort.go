package class01

// InsertSort
// 选择一个值，依次比较剩余的所有数据，满足之后交换位置
func InsertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}
}
