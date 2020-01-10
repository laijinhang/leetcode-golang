package main

import "fmt"

func rotate(matrix [][]int)  {
	var i, j int
	n := len(matrix)
	for ;i < n / 2 && i != n - i - 1;i++ {
		// 上下翻转
		for j = 0;j < n;j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	// 沿主对角线翻转
	for i = 0;i < n;i++ {
		for j = i + 1;j < n;j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main()  {
	rotate([][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9}})
}