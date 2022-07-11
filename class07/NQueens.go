package class07

import "math"

// N皇后问题，有多少种摆法

func NQueens1(num int) int {
	if num < 1 || num > 32 {
		return 0
	}

	record := make([]int, num) // 表示i行的皇后放到了第几列

	return processRecur1(0, record, num)
}

func processRecur1(i int, record []int, n int) int {
	if i == n {
		return 1
	}

	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) { // 第i行，j列是否有效
			record[i] = j
			res += processRecur1(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k)) {
			return false
		}
	}
	return true
}

func NQueens2(num int) int {
	if num < 1 || num > 32 {
		return 0
	}

	limit := 0
	if num == 32 {
		limit = -1
	} else {
		limit = (1 << num) - 1
	}

	return processRecur2(limit, 0, 0, 0)
}

func processRecur2(limit int, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == limit {
		return 1
	}

	validPos := limit & (^(colLim | leftDiaLim | rightDiaLim))
	mostRightOne := 0
	res := 0
	for validPos != 0 {
		mostRightOne = validPos & (^validPos + 1)
		validPos = validPos - mostRightOne
		res += processRecur2(limit, colLim|mostRightOne, (leftDiaLim|mostRightOne)<<1, (rightDiaLim|mostRightOne)>>1)
	}
	return res
}
