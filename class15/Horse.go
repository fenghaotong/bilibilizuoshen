package class15

// 象棋棋盘从(0, 0)位置出发经过k步到达(x,y)位置，求多少种走法

func GetWay(x, y, k int) int {
	return process(x, y, k)
}

func process(x, y, k int) int {
	if x < 0 || x > 8 || y < 0 || y > 9 {
		return 0
	}
	if k == 0 {
		if x == 0 && y == 0 {
			return 1
		} else {
			return 0
		}
	}

	return process(x-1, y+2, k-1) +
		process(x-1, y-2, k-1) +
		process(x+1, y+2, k-1) +
		process(x+1, y-2, k-1) +
		process(x-2, y-1, k-1) +
		process(x-2, y+1, k-1) +
		process(x+2, y-1, k-1) +
		process(x+2, y+1, k-1)
}

func DPWay1(x, y, k int) int {
	var dp [9][10][100]int
	dp[0][0][0] = 1
	return processDP1(x, y, k, dp)
}

func processDP1(x, y, k int, dp [9][10][100]int) int {

	if x < 0 || x > 8 || y < 0 || y > 9 || k < 0 {
		return 0
	}

	if dp[x][y][k] != 0 {
		return dp[x][y][k]
	}

	if k == 0 {
		if x == 0 && y == 0 {
			dp[x][y][k] = 1
		} else {
			dp[x][y][k] = 1
		}
		return dp[x][y][k]
	}

	dp[x][y][k] = process(x-1, y+2, k-1) +
		process(x-1, y-2, k-1) +
		process(x+1, y+2, k-1) +
		process(x+1, y-2, k-1) +
		process(x-2, y-1, k-1) +
		process(x-2, y+1, k-1) +
		process(x+2, y-1, k-1) +
		process(x+2, y+1, k-1)
	return dp[x][y][k]
}

func DPWay(x, y, k int) int {
	if x < 0 || x > 8 || y < 0 || y > 9 || k < 0 {
		return 0
	}
	var dp [9][10][100]int
	dp[0][0][0] = 1
	for h := 1; h <= k; h++ {
		for i := 0; i < 9; i++ {
			for j := 0; j < 10; j++ {
				dp[i][j][h] += getValue(dp, i-1, j+2, h-1)
				dp[i][j][h] += getValue(dp, i-1, j-2, h-1)
				dp[i][j][h] += getValue(dp, i+1, j+2, h-1)
				dp[i][j][h] += getValue(dp, i+1, j-2, h-1)
				dp[i][j][h] += getValue(dp, i-2, j+1, h-1)
				dp[i][j][h] += getValue(dp, i-2, j-1, h-1)
				dp[i][j][h] += getValue(dp, i+2, j+1, h-1)
				dp[i][j][h] += getValue(dp, i+2, j-1, h-1)
			}

		}
	}
	return dp[x][y][k]
}

func getValue(dp [9][10][100]int, x, y, k int) int {
	if x < 0 || x > 8 || y < 0 || y > 9 {
		return 0
	}
	return dp[x][y][k]
}
