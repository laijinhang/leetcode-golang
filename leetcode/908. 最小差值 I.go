package main

func smallestRangeI(A []int, K int) int {
	max, min := A[0], A[0]
	for i := 1;i < len(A);i++ {
		if max < A[i] {
			max = A[i]
		}
		if min > A[i] {
			min = A[i]
		}
	}
	res := max - min - 2 * K
	if res < 0 {
		res = 0
	}
	return res
}