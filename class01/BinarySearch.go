package class01

func BSExists(arr []int, v int) int {
	if arr == nil {
		return -1
	}

	//mid := len(arr) / 2
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)>>2
		if arr[mid] == v {
			return mid
		} else if arr[mid] > v {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return -1
}
