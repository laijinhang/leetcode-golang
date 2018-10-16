package main

import (
	"fmt"
	"strconv"
)

//var maxNum int64

func getVal(a [][]int) int {
	line := len(a)
	columns := len(a[0])
	total := 0
	for i := 0; i < line; i++ {
		ts := ""
		for j := 0; j < columns; j++ {
			ts += string('0' + a[i][j])
		}
		v, _ := strconv.ParseInt(ts, 2, 0)
		total += int(v)
	}
	return total
}

func reversal(a [][]int, i, j int, maxNum *int) {
	t := getVal(a)
	if t > *maxNum {
		*maxNum = t
	}
	if i >= len(a) || j >= len(a[0]) {	// 结束
		return
	}
	tl, tc, i, j := func(a [][]int, i, j int) ([][]int, [][]int, int, int) {
		// 深度拷贝
		t1 := make([][]int, len(a))
		t2 := make([][]int, len(a))
		for i := 0;i < len(a);i++ {
			t1[i] = append(t1[i], a[i]...)
			t2[i] = append(t2[i], a[i]...)
		}
		// 所在行翻转
		for k := 0;k < len(t1[0]);k++ {
			if t1[i][k] == 0 {
				t1[i][k] = 1
			} else {
				t1[i][k] = 0
			}
		}
		// 所在列翻转
		for k := 0;k < len(t1);k++ {
			if t2[k][j] == 0 {
				t2[k][j] = 1
			} else {
				t2[k][j] = 0
			}
		}
		j++
		if j == len(t1[0]) {
			j = 0
			i++
		}
		return t1, t2, i, j
	}(a, i, j)
	// 行翻转
	reversal(tl, i, j, maxNum)
	// 列翻转
	reversal(tc, i, j, maxNum)
	// 不翻转
	reversal(a, i, j, maxNum)
}

func matrixScore(a [][]int) int {
	var maxN int
	reversal(a, 0, 0, &maxN)
	return maxN
}

func main() {
	var a [][]int
	a = [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	a = [][]int{{1, 1},{1,1},{1,1},{1,0},{1,0},{0,0},{0,0}}
	a = [][]int{{0,0},{0,1},{1,1},{0,1},{1,1},{0,0},{1,1},{0,0}}
	fmt.Println(matrixScore(a))
}
