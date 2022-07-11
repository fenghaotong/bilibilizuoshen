package class01

import (
	"algorithm/utils"
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func ComparatorBS(arr []int, searchValue int) int {
	v2 := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= searchValue
	})
	if v2 > len(arr)-1 {
		v2 = -1
	} else if arr[v2] != searchValue {
		v2 = -1
	}
	return v2
}

func IsEqualBS(v1, v2 int) bool {
	if v1 == v2 {
		return true
	} else {
		return false
	}
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func CallbackBS(f func(arr []int, v int) int) bool {
	maxSize := 10
	maxValue := 10
	testNum := 100
	for i := 0; i < testNum; i++ {
		arr := utils.GenRandIntArr(maxSize, maxValue)
		sort.Ints(arr)
		arrSet := RemoveRepByLoop(arr)
		searchValue := rand.Intn(maxValue)
		v1 := f(arrSet, searchValue)
		v2 := ComparatorBS(arrSet, searchValue)
		if !IsEqualBS(v1, v2) {
			fmt.Println(fmt.Sprintf("Pass: %v %%", float64(i-1)/float64(testNum)*100))
			return false
		}
	}

	fmt.Println(fmt.Sprintf("Pass: %v %%", testNum/testNum*100))
	return true

}

func TestBSExists(t *testing.T) {
	CallbackBS(BSExists)
}
