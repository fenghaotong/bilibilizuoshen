package class03

import (
	"fmt"
	"math"
)

func RadixSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1, maxBits(arr))
}

func process(arr []int, L, R, digit int) {
	radix := 10
	bucket := make([]int, R-L+1)
	for d := 1; d <= digit; d++ {
		count := make([]int, radix) // 数字频率
		for i := L; i <= R; i++ {
			n := getDigit(arr[i], d)
			count[n]++
		}

		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}

		// 放入桶中
		for i := R; i >= L; i-- {
			n := getDigit(arr[i], d)
			bucket[count[n]-1] = arr[i]
			count[n]--
		}

		j := 0
		for i := L; i <= R; i++ {
			arr[i] = bucket[j]
			j++
		}
		fmt.Println("-----", arr)
	}

}

// maxBits
// 获取最大数字的位数
func maxBits(arr []int) int {
	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		max = int(math.Max(float64(max), float64(arr[i])))
	}
	res := 0
	for max != 0 {
		res++
		max = max / 10
	}
	return res
}

// getDigit
// data: 数字
// digit: 位数
// return: 返回data第digit位的数字
func getDigit(data int, digit int) int {
	return (data / int(math.Pow10(digit-1))) % 10
}
