package class11

import (
	"fmt"
	"math"
)

func LongestStringofSubroutines(str string) int {
	if len(str) == 0 {
		return 0
	}
	arr := manacherString(str)
	fmt.Println(arr)
	return processLSOF(arr)
}

// manacherString
// abc -> #a#b#c#
func manacherString(str string) []int32 {
	arr := make([]int32, len(str)*2+1)
	i := 0
	for _, v := range str {
		arr[2*i] = '#'
		arr[2*i+1] = v
		i++
	}
	arr[2*len(str)] = '#'
	return arr
}

func processLSOF(arr []int32) int {
	res := math.MinInt
	R := -1                       // 最长回文子串的最右边界
	C := -1                       // 最长回文子串的中心节点
	pArr := make([]int, len(arr)) // 存放每个位置回文子串的长度

	for i := 0; i != len(arr); i++ {
		if i >= R {
			pArr[i] = 1
		} else {
			pArr[i] = int(math.Min(float64(R-i), float64(pArr[2*C-i])))
		}

		for i+pArr[i] < len(arr) && i-pArr[i] > -1 {
			if arr[i+pArr[i]] == arr[i-pArr[i]] {
				pArr[i] += 1
			} else {
				break
			}
		}

		if i+pArr[i] > R {
			R = pArr[i] + i
			C = i
		}
		res = int(math.Max(float64(res), float64(pArr[i])))
	}

	return res - 1
}
