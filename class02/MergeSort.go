package class02

func MergeSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	L := 0
	R := len(arr) - 1
	process(arr, L, R)

}

func process(arr []int, L, R int) {
	if L == R {
		return
	}

	mid := L + (R-L)>>2
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, m, R int) {
	help := make([]int, (R-L)+1)
	i := 0
	p1 := L
	p2 := m + 1

	for p1 <= m && p2 <= R {
		if arr[p1] < arr[p2] {
			help[i] = arr[p1]
			i++
			p1++
		} else {
			help[i] = arr[p2]
			i++
			p2++
		}
	}

	for p1 <= m {
		help[i] = arr[p1]
		i++
		p1++
	}

	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}

	for i = 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
}
