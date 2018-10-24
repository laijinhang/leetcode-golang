package main

import (
	"math"
	"fmt"
)

type coordinate struct {
	x, y int
}

type val struct {
	min, v int
}

var minSum map[coordinate]val

func minPathSum(grid [][]int) int {
	minSum = make(map[coordinate]val)
	pathSum(grid, len(grid) - 1, len(grid[0]) - 1)
	fmt.Println(minSum)
	return minSum[coordinate{len(grid) - 1,len(grid[0]) - 1}].min + minSum[coordinate{len(grid) - 1,len(grid[0]) - 1}].v
}

func pathSum(grid [][]int, i, j int) val {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return val{math.MaxInt32, 0}
	}
	if _, ok := minSum[coordinate{i,j}];ok {	// 如果该值已经求出，则返回
		return  minSum[coordinate{i,j}]
	}
	if i == 0 && j == 0 {	// 左上角
		minSum[coordinate{i,j}] = val{0, grid[i][j]}
	} else if i == 0 && j == len(grid[0]) - 1 {	// 右上角
		if _, ok := minSum[coordinate{i,j-1}];!ok {
			minSum[coordinate{i,j-1}] = pathSum(grid, i, j - 1)
		}
		minSum[coordinate{i,j}] = val{
			min(minSum[coordinate{i, j - 1}].min+minSum[coordinate{i, j - 1}].v),grid[i][j]}
	} else if i == len(grid) - 1 && j == 0 {	// 左下角
		if _, ok := minSum[coordinate{i-1,j}];!ok {
			minSum[coordinate{i-1,j}] = pathSum(grid, i - 1, j)
		}
		minSum[coordinate{i,j}] = val{
			min(minSum[coordinate{i - 1, j}].min+minSum[coordinate{i - 1, j}].v),grid[i][j]}
	} else if i == len(grid) - 1 && j == len(grid[0]) - 1 {	// 右下角
		if _, ok := minSum[coordinate{i-1,j}];!ok {
			minSum[coordinate{i-1,j}] = pathSum(grid, i - 1, j)
		}
		if _, ok := minSum[coordinate{i, j+1}];!ok {
			minSum[coordinate{i,j - 1}] = pathSum(grid, i, j - 1)
		}
		minSum[coordinate{i,j}] = val{
			min(minSum[coordinate{i-1, j}].min+minSum[coordinate{i-1, j}].v,
				minSum[coordinate{i,j-1}].min+minSum[coordinate{i,j-1}].v),
			grid[i][j]}
	} else if i == 0 && i != len(grid) - 1 {	// 上边
		if _, ok := minSum[coordinate{i, j-1}];!ok {
			minSum[coordinate{i,j-1}] = pathSum(grid, i, j - 1)
		}
		minSum[coordinate{i,j}] = val{
				minSum[coordinate{i,j - 1}].min+minSum[coordinate{i, j - 1}].v, grid[i][j]}
	} else if i != 0 && i == len(grid) - 1 {	// 下边
		if _, ok := minSum[coordinate{i-1,j}];!ok {
			minSum[coordinate{i-1,j}] = pathSum(grid, i - 1,j)
		}
		if _, ok := minSum[coordinate{i, j-1}];!ok {
			minSum[coordinate{i,j-1}] = pathSum(grid, i, j - 1)
		}
		minSum[coordinate{i,j}] = val{
			min(minSum[coordinate{i-1, j}].min+minSum[coordinate{i-1,j}].v,
				minSum[coordinate{i,j - 1}].min+minSum[coordinate{i, j - 1}].v),
			grid[i][j]}
	} else if j == 0 {	// 左边
		if _, ok := minSum[coordinate{i-1,j}];!ok {
			minSum[coordinate{i-1,j}] = pathSum(grid, i - 1,j)
		}
		minSum[coordinate{i,j}] = val{
			min(minSum[coordinate{i-1, j}].min+minSum[coordinate{i-1,j}].v),
			grid[i][j]}
	} else if j == len(grid[0]) - 1 {	// 右边
		if _, ok := minSum[coordinate{i-1,j}];!ok {
			minSum[coordinate{i-1,j}] = pathSum(grid, i - 1,j)
		}
		if _, ok := minSum[coordinate{i, j-1}];!ok {
			minSum[coordinate{i,j-1}] = pathSum(grid, i, j - 1)
		}
		minSum[coordinate{i,j}] = val{
			min(minSum[coordinate{i-1, j}].min+minSum[coordinate{i-1,j}].v,
				minSum[coordinate{i,j-1}].min+minSum[coordinate{i,j-1}].v),
			grid[i][j]}
	} else {
		if _, ok := minSum[coordinate{i-1,j}];!ok {
			minSum[coordinate{i-1,j}] = pathSum(grid, i - 1,j)
		}
		if _, ok := minSum[coordinate{i, j-1}];!ok {
			minSum[coordinate{i,j-1}] = pathSum(grid, i, j - 1)
		}
		minSum[coordinate{i,j}] = val{
			min(minSum[coordinate{i-1, j}].min+minSum[coordinate{i-1,j}].v,
				minSum[coordinate{i,j-1}].min+minSum[coordinate{i,j-1}].v),
			grid[i][j]}
	}
	return minSum[coordinate{i,j}]
}

func min(s ...int) int {
	m := s[0]
	for i := 1;i < len(s);i++ {
		if m > s[i] {
			m = s[i]
		}
	}
	return m
}

func main() {
	fmt.Println(minPathSum([][]int{{1,3,1}, {1,5,1}, {4,2,1}}))
}