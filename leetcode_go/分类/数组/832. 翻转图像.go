package main

import "math"

func flipAndInvertImage(A [][]int) [][]int {
	B := make([][]int, len(A))	// 不该变原有A的内容
	for i := 0;i < len(A);i++ {
		B[i] = append([]int{}, A[i]...)
	}
	for i := 0;i < len(B);i++ {
		for j := 0;j < len(B[0]) / 2;j++ {
			B[i][j], B[i][len(B[0])-j-1] = B[i][len(B[0])-j-1], B[i][j]		// 翻转
			B[i][j], B[i][len(B[0])-j-1] = int(math.Abs(float64(B[i][j]-1))), int(math.Abs(float64(B[i][len(B[0])-j-1]-1)))	// 反转
		}
		if len(B[0]) % 2 == 1 {
			B[i][len(B[0])/2] = int(math.Abs(float64(B[i][len(B[0])/2]-1)))
		}
	}
	return B
}