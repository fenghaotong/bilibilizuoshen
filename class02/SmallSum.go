package class02

type SmallSum struct {
}

func (s *SmallSum) Run(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	L := 0
	R := len(arr) - 1
	return s.process(arr, L, R)
}

func (s *SmallSum) process(arr []int, L, R int) int {
	if L == R {
		return 0
	}

	mid := L + (R-L)>>2
	return s.process(arr, L, mid) + s.process(arr, mid+1, R) + s.merge(arr, L, mid, R)
}

func (s *SmallSum) merge(arr []int, L, m, R int) int {
	help := make([]int, (R-L)+1)
	i := 0
	p1 := L
	p2 := m + 1
	res := 0
	for p1 <= m && p2 <= R {
		if arr[p1] < arr[p2] {
			res = res + (R-p2+1)*arr[p1]
			help[i] = arr[p1]
			i++
			p1++
		} else {
			res = res + 0
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
	return res
}
