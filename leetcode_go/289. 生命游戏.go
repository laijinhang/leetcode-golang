package main

import "fmt"

func gameOfLife(board [][]int) [][]int {
	res := make([][]int, len(board))
	for i := 0;i < len(board);i++ {
		res[i] = make([]int, len(board[0]))
		for j := 0;j < len(board[0]);j++ {
			n := numsLife(board, i, j)
			if board[i][j] == 0 {
				if n == 3 {
					res[i][j] = 1
				}
			} else {
				if n < 2 || n > 3 {
					res[i][j] = 0
				} else if n == 2 || n == 3 {
					res[i][j] = 1
				}
			}
		}
	}
	return res
}

func numsLife(board [][]int, i, j int) int {
	nums := 0
	if i - 1 >= 0 && j - 1 >= 0 {	// 左上
		nums++
	}
	if i - 1 >= 0 {	// 正上
		nums++
	}
	if i - 1 >= 0 && j + 1 < len(board[0]) {	// 右上
		nums++
	}
	if j + 1 < len(board[0]) {	// 正右
		nums++
	}
	if i + 1 < len(board) && j + 1 < len(board[0]) {	// 右下
		nums++
	}
	if j + 1 < len(board) {	// 正下
		nums++
	}
	if i + 1 >= 0 && j - 1 >= 0 {	// 正左
		nums++
	}
	if j - 1 >= 0 {	// 正左
		nums++
	}
	return nums
}