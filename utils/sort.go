package utils

import (
	"fmt"
	"math/rand"
	"sort"
)

func GenRandIntArr(maxSize int, maxValue int) []int {
	rand.Seed(rand.Int63())
	length := rand.Intn(maxSize)
	nums := make([]int, length+1)
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(maxValue)
	}
	return nums
}

func Comparator(arr []int) {
	sort.Ints(arr)
}

func IsEqual(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func SortCallback(f func(arr []int)) {
	maxSize := 10
	maxValue := 100
	testNum := 100
	for i := 0; i < testNum; i++ {
		arr := GenRandIntArr(maxSize, maxValue)
		fmt.Println("before:", arr)
		arr1 := make([]int, len(arr), maxSize)
		arr2 := make([]int, len(arr), maxSize)
		copy(arr1, arr)
		copy(arr2, arr)
		f(arr1)
		Comparator(arr2)
		fmt.Println("after(my):", arr1)
		fmt.Println("after(std):", arr2)
		if !IsEqual(arr1, arr2) {
			fmt.Println(fmt.Sprintf("Pass: %v %%", i/testNum*100))
			return
		}
	}

	fmt.Println(fmt.Sprintf("Pass: %v %%", testNum/testNum*100))
}
