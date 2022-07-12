package class10

func CountIsLands(arr [][]int) int {
	if arr == nil || arr[0] == nil {
		return 0
	}

	row := len(arr)
	col := len(arr[0])
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if arr[i][j] == 1 {
				res++
				infect(arr, i, j, row, col)
			}
		}
	}
	return res
}

func infect(arr [][]int, i, j, row, col int) {
	if i < 0 || i >= row || j < 0 || j >= col || arr[i][j] != 1 {
		return
	}

	arr[i][j] = 2
	infect(arr, i-1, j, row, col)
	infect(arr, i+1, j, row, col)
	infect(arr, i, j-1, row, col)
	infect(arr, i, j+1, row, col)
}
