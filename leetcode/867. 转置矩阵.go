package main

import "fmt"

func transpose(A [][]int) [][]int {
	a1Len := len(A)
	a2Len := len(A[0])
	B := [][]int{}
	for i := 0;i < a2Len;i++ {
		temp := []int{}
		for j := 0;j < a1Len;j++ {
			temp = append(temp, A[j][i])
		}
		B = append(B, temp)
	}
	return B
}