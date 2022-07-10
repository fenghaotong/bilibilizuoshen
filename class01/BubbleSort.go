package class01

func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	len := len(arr)
	for i := len - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				SwapBit(arr, j, j+1)
			}
		}

	}
}

func SwapBit(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
