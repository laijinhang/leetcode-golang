package main

func islandPerimeter(grid [][]int) int {
	num := 0
	for i := 0;i < len(grid);i++ {
		for j := 0;j < len(grid[0]);j++ {
			if grid[i][j] == 1 {
				num += 4 - findNum(grid, i, j)
			}
		}
	}
	return num
}

func findNum(grid [][]int, i, j int) int {
	num := 0
	if i - 1 >= 0 && grid[i-1][j] == 1 {	// 正上
		num++
	}
	if j + 1 < len(grid[0]) && grid[i][j+1] == 1 {	// 正右
		num++
	}
	if i + 1 < len(grid) && grid[i+1][j] == 1 {	// 正下
		num++
	}
	if j - 1 >= 0 && grid[i][j-1] == 1 {		// 正左
		num++
	}
	return num
}