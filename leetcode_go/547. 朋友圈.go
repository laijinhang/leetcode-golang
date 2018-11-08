package main

import "fmt"

func findCircleNum(M [][]int) int {
	num := 0
	for i := 0;i < len(M);i++ {
		for j := 0;j < len(M[0]);j++ {
			if M[i][j] > 0 {
				num++
				findCircleDFS(M, -num, i, j)
			}
		}
	}
	return num
}

func findCircleDFS(M [][]int, v, i, j int) {
	for t := 0;t < len(M[0]) && i >= 0;t++ {
		if M[i][t] > 0 {
			M[i][t] = v
			findCircleDFS(M, v, i, t)
		}
	}
	for t := 0;t < len(M) && j >= 0;t++ {
		if M[t][j] > 0 {
			M[t][j] = v
			findCircleDFS(M, v, t, j)
		}
	}
}

func main() {
	/*
	{1,1,0},
	{1,1,1},
	{0,1,1}
	 */
	fmt.Println(findCircleNum([][]int{{1,0,0,1},{0,1,1,0},
		{0,1,1,1},
		{1,0,1,1}}))
}