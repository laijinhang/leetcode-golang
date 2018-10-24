package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0;i < n;i++ {
		res[i] = make([]int, n)
	}
	num := 1
	i, j := -1, -1
	for num <= n * n {
		// 向右
		for i, j = i + 1, j + 1;j < n && res[i][j] == 0;j++ {
			res[i][j] = num
			num++
		}
		// 向下
		for j, i = j - 1, i + 1;i < n && res[i][j] == 0;i++ {
			res[i][j] = num
			num++
		}
		// 向左
		for i, j = i - 1, j - 1;j >= 0 && res[i][j] == 0;j-- {
			res[i][j] = num
			num++
		}
		// 向上
		for i, j = i - 1, j + 1;i >= 0 && res[i][j] == 0;i-- {
			res[i][j] = num
			num++
		}
	}
	return res
}

func main()  {
	fmt.Println(generateMatrix(5))
}