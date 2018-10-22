package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	N := make([][]int, len(M))
	for i := 0;i < len(M);i++ {
		N[i] = make([]int, len(M[0]))
	}
	for i := 0;i < len(M);i++ {
		for j := 0;j < len(M[0]);j++ {
			N[i][j] = smoother(M, i, j)
		}
	}
	return N
}

func smoother(M [][]int, i, j int) int {
	m := 0
	n := 0
	if i - 1 >= 0 {
		if j - 1 >= 0 {
			m += M[i-1][j-1]
			n++
		}
		m += M[i-1][j]
		n++
		if j + 1 < len(M[0]) {
			m += M[i-1][j+1]
			n++
		}
 	}

 	if j - 1 >= 0 {
 		m += M[i][j-1]
 		n++
	}
	m += M[i][j]
	n++
	if j + 1 < len(M[0]) {
		m += M[i][j+1]
		n++
	}

	if i + 1 < len(M) {
		if j - 1 >= 0 {
			m += M[i+1][j-1]
			n++
		}
		m += M[i+1][j]
		n++
		if j + 1 < len(M[0]) {
			m += M[i+1][j+1]
			n++
		}
	}
	return m / n
}

func main() {
	N := imageSmoother([][]int{{1,1,1},{1,1,1},{1,1,1}})
	fmt.Println(N)
}