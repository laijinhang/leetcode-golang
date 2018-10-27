package main

func surfaceArea(grid [][]int) int {
	totalArea := 0
	shadowArea := 0
	// 总面积
	for i := 0;i < len(grid);i++ {
		for j := 0;j < len(grid[0]);j++ {
			totalArea += grid[i][j]
		}
	}
	totalArea *= 6
	// 阴影部分面积
	for i := 0;i < len(grid);i++ {
		for j := 0;j < len(grid[0]);j++ {
			if grid[i][j] > 0 {
				shadowArea += (grid[i][j] - 1) * 2
				if i-1 >= 0 {
					shadowArea += min(grid[i][j], grid[i-1][j])
				}
				if i+1 < len(grid) {
					shadowArea += min(grid[i][j], grid[i+1][j])
				}
				if j-1 >= 0 {
					shadowArea += min(grid[i][j], grid[i][j-1])
				}
				if j+1 < len(grid[0]) {
					shadowArea += min(grid[i][j], grid[i][j+1])
				}
			}
		}
	}
	return totalArea - shadowArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}