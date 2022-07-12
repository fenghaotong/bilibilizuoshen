package class10

func GetIndexOf(big, small string) int {
	if len(big) < len(small) || len(big) == 0 || len(small) == 0 {
		return -1
	}

	bigArr := make([]int32, 0)
	smallArr := make([]int32, 0)
	bigPtr := 0
	smallPtr := 0

	for _, v := range big {
		bigArr = append(bigArr, v)
	}

	for _, v := range small {
		smallArr = append(smallArr, v)
	}

	nexts := getNextArray(smallArr)

	for bigPtr != len(bigArr) && smallPtr != len(smallArr) {
		if bigArr[bigPtr] == smallArr[smallPtr] {
			smallPtr++
			bigPtr++
		} else if smallPtr == 0 {
			bigPtr++
		} else {
			smallPtr = nexts[smallPtr]
		}
	}

	if smallPtr == len(smallArr) {
		return bigPtr - smallPtr
	} else {
		return -1
	}

}

func getNextArray(str []int32) []int {
	if len(str) == 1 {
		return []int{-1}
	}

	nextArr := make([]int, len(str))
	nextArr[0] = -1
	nextArr[1] = 0
	cn := 0
	for i := 2; i < len(str); {
		if str[cn] == str[i-1] {
			cn = cn + 1
			nextArr[i] = cn
			i++
		} else if cn > 0 {
			cn = nextArr[cn]
		} else {
			nextArr[i] = 0
			i++
		}
	}
	return nextArr
}
