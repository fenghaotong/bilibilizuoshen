package class08

import "fmt"

func PrintStringPermutations(word string) {
	res := make([]string, 0)
	str := make([]uint8, 0)
	for i := 0; i < len(word); i++ {
		str = append(str, word[i])
	}

	processSP(str, 0, res)
}

func processSP(str []uint8, i int, res []string) {
	if i == len(str) {
		res = append(res, string(str))
		fmt.Println(res)
	}

	visit := make([]bool, 26)
	for j := i; j < len(str); j++ {
		if !visit[int(str[j])-int('a')] {
			visit[int(str[j])-int('a')] = true
			swap(str, i, j)
			processSP(str, i+1, res)
			swap(str, i, j)
		}
	}
}

func swap(str []uint8, i, j int) {
	tmp := str[i]
	str[i] = str[j]
	str[j] = tmp
}
