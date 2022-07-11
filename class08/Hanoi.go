package class08

import "fmt"

func Hanoi(n int) {
	if n > 0 {
		process(n, "左", "右", "中")
	}
}

func process(i int, from, to, other string) {
	if i == 1 {
		fmt.Println(fmt.Sprintf("Move 1 from %s to %s ", from, to))
	} else {
		process(i-1, from, other, to)
		fmt.Println(fmt.Sprintf("Move %d from %s to %s ", i, from, to))
		process(i-1, other, to, from)
	}
}
