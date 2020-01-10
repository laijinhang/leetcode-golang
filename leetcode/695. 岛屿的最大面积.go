package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	area := 0
	rLen := len(grid)
	cLen := len(grid[0])
	for i := 0;i < rLen;i++ {
		for j := 0;j < cLen;j++ {
			if grid[i][j] == 1 {
				area = 0
				find(grid, i, j, &area)
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func find(grid [][]int, x, y int, area *int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
		return
	}
	grid[x][y] = -1
	*area++
	find(grid, x - 1, y, area)	// 往左
	find(grid, x, y - 1, area)	// 往上
	find(grid, x, y + 1, area)	// 往下
	find(grid, x + 1, y, area)	// 往右
}

func main() {
	a := [][]int{{1,1,0,1,1},{1,0,0,0,0},{0,0,0,0,1},{1,1,0,1,1}}
	fmt.Println(maxAreaOfIsland(a))
}