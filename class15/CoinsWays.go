package class15

// 钱的面值是一个数组，表示有多少种面值，aim是需要找零的钱数，可以使用任何面值的钱任何张，有多少种找法

func GetCoinsWays(arr []int, aim int) int {
	return processCW(arr, 0, aim)
}

func processCW(arr []int, index, rest int) int {
	if arr == nil || len(arr) <= 0 {
		return 0
	}
	if index == len(arr) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}

	ways := 0
	for i := 0; arr[index]*i <= rest; i++ {
		ways += processCW(arr, index+1, rest-arr[index]*i)
	}

	return ways
}

func GetCoinsWaysDP1(arr []int, aim int) int {
	dp := make([][]int, 0)
	for i := 0; i <= len(arr); i++ {
		tmp := make([]int, 0)
		for j := 0; j <= aim; j++ {
			tmp = append(tmp, 0)
		}
		dp = append(dp, tmp)
	}
	dp[len(arr)][0] = 1
	return processCWDP1(arr, aim, dp)
}

func processCWDP1(arr []int, aim int, dp [][]int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	N := len(arr)
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for i := 0; arr[index]*i <= rest; i++ {
				ways += dp[index+1][rest-arr[index]*i]
			}
			dp[index][rest] = ways
		}
	}

	return dp[0][aim]
}

func GetCoinsWaysDP2(arr []int, aim int) int {
	dp := make([][]int, 0)
	for i := 0; i <= len(arr); i++ {
		tmp := make([]int, 0)
		for j := 0; j <= aim; j++ {
			tmp = append(tmp, 0)
		}
		dp = append(dp, tmp)
	}
	dp[len(arr)][0] = 1
	return processCWDP2(arr, aim, dp)
}

func processCWDP2(arr []int, aim int, dp [][]int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	N := len(arr)
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			for i := 0; arr[index]*i <= rest; i++ {
				dp[index][rest] = dp[index+1][rest]
				if rest-arr[index] >= 0 {
					dp[index][rest] += dp[index][rest-arr[index]]
				}
			}
		}
	}

	return dp[0][aim]
}
