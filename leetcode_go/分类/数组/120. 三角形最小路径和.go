package main

func minimumTotal(triangle [][]int) int {
	minNum := make([]int, len(triangle))

}

func minnum(t [][]int, minNum []int, i int) {
	if i == 0 {
		minNum[i] = t[0][0]
	}
	minnum(t, minNum, i - 1)

}