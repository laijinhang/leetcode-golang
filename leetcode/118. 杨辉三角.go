package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	yh := make([][]int, numRows)
	yh[0] = []int{1}
	for i := 1;i < numRows;i++ {
		yh[i] = make([]int, len(yh[i-1])+1)
		for j := 0;j < len(yh[i]);j++ {
			if j != 0 && j != len(yh[i]) - 1 {
				yh[i][j] = yh[i-1][j-1] + yh[i-1][j]
			} else if j == len(yh[i]) - 1 {
				yh[i][j] = yh[i-1][j-1]
			} else {
				yh[i][j] = yh[i-1][0]
			}
		}
	}
	return yh
}

func main()  {
	fmt.Println(generate(5))
}