package main

import (
	"time"
	"fmt"
)

func sumSubarrayMins(A []int) int {
	t1 := time.Now()
	sum := 0
	min := append([]int{}, A...)
	for i, j := 1, 0;i <= len(A);i++ {
		for j = 0;j + i <= len(A);j++ {
			if min[j] > A[i+j-1] {
				min[j] = A[i+j-1]
			}
			sum += min[j]
		}
	}
	fmt.Println(time.Since(t1))
	return sum % 1000000007
}

func main() {
	t1 := time.Now()
	a := []int{}
	for i := 0;i < 100000;i++ {
		a = append(a, i)
	}
	fmt.Println(time.Since(t2))
	t2 := time.Now()
	fmt.Println(sumSubarrayMins([]int{3,1,2,4}))
	fmt.Println(time.Since(t2))
}