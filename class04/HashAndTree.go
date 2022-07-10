package class04

import (
	"fmt"
)

func HashTable() {
	testMap := make(map[int]string, 0)
	testMap[0] = "1"
	testMap[1] = "2"
	testMap[2] = "3"

	fmt.Println(testMap)
}
