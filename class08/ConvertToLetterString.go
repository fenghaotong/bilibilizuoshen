package class08

func Convert(arr []uint8) int {
	if len(arr) == 0 {
		return 0
	}

	return processConvert(arr, 0)
}

func processConvert(arr []uint8, i int) int {
	if i == len(arr) {
		return 1
	}

	if arr[i] == '0' {
		return 0
	}

	if arr[i] == '1' {
		res := 0
		res += processConvert(arr, i+1)
		if i+1 < len(arr) {
			res += processConvert(arr, i+2)
		}
		return res
	}

	if arr[i] == '2' {
		res := 0
		res += processConvert(arr, i+1)
		if i+1 < len(arr) && arr[i+1] > '0' && arr[i+1] < '6' {
			res += processConvert(arr, i+2)
		}
		return res
	}

	return processConvert(arr, i+1)
}
