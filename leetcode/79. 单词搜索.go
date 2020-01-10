package main

import (
	"fmt"
	"time"
)

// 解题关键：记录不可通路径
type coordinate struct {
	x, y int
}

func exist(board [][]byte, word string) bool {
	if board == nil || len(word) > len(board) * len(board[0]) {
		return false
	}
	b := make(map[byte]int)
	w := make(map[byte]int)
	for i := 0;i < len(board);i++ {
		for j := 0;j < len(board[i]);j++ {
			b[board[i][j]]++
		}
	}
	for i := 0;i < len(word);i++ {
		w[word[i]]++
	}
	for k, n := range w {
		if n > b[k] {
			return false
		}
	}
	flag := false
	noGoPath := make([][]coordinate, 0)
	cStack := []coordinate{}
	x, y, wi := 0, 0, 0
	for i := 0;i < len(board);i++ {
		// 清零
		noGoPath = make([][]coordinate, 0)
		for j := 0;j < len(board[0]);j++ {
			wi = 0
			if board[i][j] == word[wi] {
				x, y = i, j
				cStack = append(cStack, coordinate{i,j})
				for len(cStack) != 0 {
					for ;y + 1 < len(board[0]) && !isGo(cStack, coordinate{x,y+1}) && !noGoPathF(noGoPath, append(cStack, coordinate{x,y+1})) && wi + 1 < len(word) && board[x][y+1] == word[wi+1];y, wi = y + 1, wi + 1 {
						cStack = append(cStack, coordinate{x,y+1})
					}    // 向右
					for ;x + 1 < len(board) && !isGo(cStack, coordinate{x+1,y}) && !noGoPathF(noGoPath, append(cStack, coordinate{x+1,y}))  && wi + 1 < len(word) && board[x+1][y] == word[wi+1];x, wi = x + 1, wi + 1 {
						cStack = append(cStack, coordinate{x+1,y})
					}    // 向下
					for ;y - 1 >= 0 && !isGo(cStack, coordinate{x,y-1}) && !noGoPathF(noGoPath, append(cStack, coordinate{x,y-1}))  && wi + 1 < len(word) && board[x][y-1] == word[wi+1];y, wi = y - 1, wi + 1 {  // 向左
						cStack = append(cStack, coordinate{x,y-1})
					}   // 向左
					for ;x - 1 >= 0 && !isGo(cStack, coordinate{x-1,y}) && !noGoPathF(noGoPath, append(cStack, coordinate{x-1,y}))  && wi + 1 < len(word) && board[x-1][y] == word[wi+1];x, wi = x - 1, wi + 1 {  // 向上
						cStack = append(cStack, coordinate{x-1,y})
					}   // 向上
					if wi == len(word) - 1 {
						return true
					}
					// 走不通了，
					if !(y + 1 < len(board[0]) && !isGo(cStack, coordinate{x,y+1}) && !noGoPathF(noGoPath, append(cStack, coordinate{x,y+1})) && wi + 1 < len(word) && board[x][y+1] == word[wi+1] ||
						x + 1 < len(board) && !isGo(cStack, coordinate{x+1,y}) && !noGoPathF(noGoPath, append(cStack, coordinate{x+1,y}))  && wi + 1 < len(word) && board[x+1][y] == word[wi+1] ||
						y - 1 >= 0 && !isGo(cStack, coordinate{x,y-1}) && !noGoPathF(noGoPath, append(cStack, coordinate{x,y-1}))  && wi + 1 < len(word) && board[x][y-1] == word[wi+1]) {
						if !noGoPathF(noGoPath, cStack) {
							noGoPath = append(noGoPath, append([]coordinate{}, cStack...))
						}
						cStack = cStack[:len(cStack)-1]
						if len(cStack) > 0 {
							x = cStack[len(cStack)-1].x
							y = cStack[len(cStack)-1].y
						}
						wi--
					}
				}
			}
		}
	}
	return flag
}

func isGo(cStack []coordinate, n coordinate) bool {
	for i := 0;i < len(cStack);i++ {
		if cStack[i] == n {
			return true
		}
	}
	return false
}

func noGoPathF(noGoPath [][]coordinate, cStack []coordinate) bool {
	flag := false
	for i := 0;i < len(noGoPath);i++ {
		if len(noGoPath[i]) == len(cStack) {
			j := 0
			for ;j < len(noGoPath[i]);j++ {
				if cStack[j] != noGoPath[i][j] {
					break
				}
			}
			if j == len(noGoPath[i]) {
				flag = true
				break
			}
		}
	}
	return flag
}

func main() {
	tb := time.Now()
	fmt.Println(exist([][]byte{
		{'a','a','a','a'},
		{'a','a','a','a'},
		{'a','a','a','a'},
		{'a','a','a','a'},
		{'a','a','a','b'},
	}, "aaaaaaaaaaaaaaaaaaac"))
	fmt.Println("Run ", time.Since(tb))
}