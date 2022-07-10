package class01

import "fmt"

func PrintEvenTimes(arr []int) {
	eor := 0
	for _, i := range arr {
		eor = eor ^ i
	}
	fmt.Println(eor)

}

func PrintOddTimes(arr []int) {
	eor := 0
	for _, i := range arr {
		eor = eor ^ i
	}

	rightOne := eor & (^eor + 1)
	eor2 := 0
	for _, i := range arr {
		if (i & rightOne) == 0 {
			eor2 = eor2 ^ i
		}
	}
	fmt.Println(eor^eor2, eor2)
}
