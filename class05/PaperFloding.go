package class05

import "fmt"

func PrintPaperFolding(N int) {
	processPF(1, N, true)
}

// processPF
// true: 凹, false: 凸
func processPF(i, N int, down bool) {
	if i > N {
		return
	}
	processPF(i+1, N, true)
	if down {
		fmt.Println("凹")
	} else {
		fmt.Println("凸")
	}

	processPF(i+1, N, false)
}
