package main

/**
 * 动态规划 o(mn)
 * dp数组含义：走到该位置的最小路径
 * 第一行的时候只能向右走
 * 第一列的时候只能向下走
 */
func minPathSum(grid [][]int) int {
    row := len(grid)
    column := len(grid[0])

    for i := 0; i < row; i++ {
        for j := 0; j < column; j++ {
			if i == 0 && j == 0 {

			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j ==0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
        }
	}

	return grid[row-1][column-1]
}

func min(x int, y int) int {
	if x<y {
        return x
    }
    return y
}