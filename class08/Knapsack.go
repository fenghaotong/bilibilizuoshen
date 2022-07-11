package class08

import "math"

// 两个数组weights和values，背包装的重量不变，如何装才能获取最大价值

func GetMaxValue(weights []int, values []int, bag int) int {
	if weights == nil || len(weights) != len(values) {
		return 0
	}
	return processGMV(weights, values, 0, 0, bag)
}

func processGMV(weights []int, values []int, i, alreadyWeight, bag int) int {
	if alreadyWeight > bag {
		return math.MinInt
	}

	if i == len(weights) {
		return 0
	}

	return int(math.Max(
		float64(processGMV(weights, values, i+1, alreadyWeight, bag)),
		float64(values[i]+processGMV(weights, values, i+1, alreadyWeight+weights[i], bag)),
	))
}
