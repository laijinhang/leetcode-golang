package main

func partitionDisjoint(A []int) int {
	leftLen := 1
	max, tempMax := 0, 0
	for i := 1;i < len(A);i++ {
		if A[i] > A[tempMax] {
			tempMax = i
		}
		if A[i] < A[max] {
			max = tempMax
			leftLen = i + 1
		}
	}
	return leftLen
}
