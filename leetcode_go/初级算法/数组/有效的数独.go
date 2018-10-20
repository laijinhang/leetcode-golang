package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	// 扫描行、列
	tempI := make([]int, 10)
	tempJ := make([]int, 10)
	for i := 0;i < 9;i++ {
		for j := 0;j < 9;j++ {
			if board[i][j] != '.' {
				tempI[board[i][j]-'0']++
				if tempI[board[i][j]-'0'] > 1 {
					return false
				}
			}
			if board[j][i] != '.' {
				tempJ[board[j][i]-'0']++
				if tempJ[board[j][i]-'0'] > 1 {
					return false
				}
			}
		}
		tempI = make([]int, 10)
		tempJ = make([]int, 10)
	}
	// 扫描小宫格
	for i := 0;i < 9;i += 3 {
		for j := 0;j < 9;j += 3 {
			// 扫描小宫格
			for k := i; k < i+3;k++ {
				for f := j;f < j + 3;f++ {
					if board[k][f] != '.' {
						tempI[board[k][f]-'0']++
						if tempI[board[k][f]-'0'] > 1 {
							return false
						}
					}
				}
			}
			tempI = make([]int, 10)
		}
	}
	return true
}

func main()  {
	fmt.Println(isValidSudoku([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}