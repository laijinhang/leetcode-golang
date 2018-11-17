package main

import "fmt"

func solveSudoku(board [][]byte)  {
	c := make(map[int]map[int]bool)	// 行
	l := make(map[int]map[int]bool)	// 列
	sSquare := make(map[int]map[int]bool)	// 小矩阵
	squareRef := []int{
		0:1,3:1,6:1,
		1:2,4:2,7:2,
		2:3,5:3,8:3,
		9:4,12:4,15:4,
		10:5,13:5,16:5,
		11:6,14:6,17:6,
		18:7,21:7,24:7,
		19:8,22:8,25:8,
		20:9,23:9,26:9}
	stack := make([]int, 0)
	for i := 0;i < len(board);i++ {		// 行存储
		c[i] = make(map[int]bool)
		for j := 0;j < len(board[0]);j++ {
			if board[i][j] != '.' {
				c[i][int(board[i][j] - '0')] = true
			}
		}
	}
	for i := 0;i < len(board[0]);i++ {	// 列存储
		l[i] = make(map[int]bool)
		for j := 0;j < len(board);j++ {
			if board[j][i] != '.' {
				l[i][int(board[j][i]-'0')] = true
			}
		}
	}
	for i := 1;i <= 9;i++ {
		sSquare[i] = make(map[int]bool)
	}
	for i := 0;i < len(board);i++ {
		for j := 0;j < len(board[0]);j++ {
			if board[i][j] != '.' {
				sSquare[squareRef[(i*len(board[0])+j)/3]][int(board[i][j]-'0')] = true
			}
		}
	}
	for i := 0;i < len(board);i++ {
		for j := 0;j < len(board[0]);j++ {
			if board[i][j] == '.' {
				k := 1
				for {
					for ;k <= 9 && (c[i][k] || l[j][k] || sSquare[squareRef[(i*len(board[0])+j)/3]][k]);k++ {}
					if k >= 1 && k <= 9 {
						board[i][j] = byte(k) + '0'
						c[i][k] = true
						l[j][k] = true
						sSquare[squareRef[(i*len(board[0])+j)/3]][k] = true
						stack = append(stack, i*len(board[0])+j)
						break
					} else {

						i = stack[len(stack)-1] / len(board[0])
						j = stack[len(stack)-1] % len(board[0])
						k = int(board[i][j]-'0'+1)
						delete(c[i], int(board[i][j]-'0'))
						delete(l[j], int(board[i][j]-'0'))
						delete(sSquare[squareRef[(i*len(board[0])+j)/3]], int(board[i][j]-'0'))
						board[i][j] = '.'
						stack = stack[:len(stack)-1]
					}
				}
			}
		}
	}
}

func main() {
	a := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}}
	solveSudoku(a)
	for i := 0;i < len(a);i++ {
		for j := 0;j < len(a[i]);j++ {
			if a[i][j] != '.' {
				fmt.Print(a[i][j]-'0', " ")
			} else {
				fmt.Print(string(a[i][j]), " ")
			}
		}
		fmt.Println()
	}
}