package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
	num := 0
	for i := 0;i <= len(grid) - 3;i++ {
		for j := 0;j <= len(grid[0]) - 3;j++ {
			if isMagicSquaresInside(grid, i, j) {
				num++
			}
		}
	}
	return num
}

func isMagicSquaresInside(grid [][]int, a, b int) bool {
	// 行列
	r := []int{}
	c := []int{}
	temp := make([]int, 16)
	for i := 0;i < 3;i++ {
		tempr := 0
		tempc := 0
		for j := 0;j < 3;j++ {
			tempr += grid[i+a][j+b]
			tempc += grid[j+a][i+b]
			temp[grid[i+a][j+b]]++
			if temp[grid[i+a][j+b]] > 1 {
				return false
			}
		}
		r = append(r, tempr)
		c = append(c, tempc)
	}
	if r[0] != r[1] || r[1] != r[2] || c[0] != c[1] || c[1] != c[2] || r[0] != c[0] {
		return false
	}
	// 对角线
	if grid[a][b] + grid[a+2][b+2] != grid[a+2][b] + grid[a][b+2] || grid[a][b] + grid[a+1][b+1] + grid[a+2][b+2] != r[0] {
		return false
	}
	return true
}

func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4,3,8,4},{9,5,1,9},{2,7,6,2}}))
}