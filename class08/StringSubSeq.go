package class08

import (
	"fmt"
)

func PrintStringSubSeq(word string) {
	res := make([]uint8, 0)
	processStringSubSeq(word, 0, res)
}

func processStringSubSeq(word string, i int, res []uint8) {
	if i == len(word) {
		printList(res)
		return
	}
	resKeep := res
	resKeep = append(resKeep, word[i])
	processStringSubSeq(word, i+1, resKeep)
	resNoInclude := res
	processStringSubSeq(word, i+1, resNoInclude)
}

func printList(res []uint8) {
	fmt.Println(string(res))
}
