package main

import (
	"sort"
	"fmt"
)

func fairCandySwap(A []int, B []int) []int {
	a, b := 0, 0
	for i := 0;i < len(A);i++ {
		a += A[i]
	}
	for i := 0;i < len(B);i++ {
		b += B[i]
	}
	res := []int{0,0}
	if a == b {
		return res
	}
	sort.Ints(A)
	sort.Ints(B)
	for i := 0;i < len(A);i++ {
		for j := 0;j < len(B);j++ {
			if ((a > b && A[i] > B[j]) ||
				(a < b && A[i] < B[j])) &&
				(a - A[i] + B[j] == b + A[i] - B[j]){
				res = []int{A[i],B[j]}
				i = len(A)
				j = len(B)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(fairCandySwap([]int{32,38,8,1,9}, []int{68,92}))
}