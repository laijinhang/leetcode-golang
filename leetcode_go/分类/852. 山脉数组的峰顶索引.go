package main

func peakIndexInMountainArray(A []int) int {
	res := 0
	for i := 1;i < len(A) - 1;i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			res = i
			break
		}
	}
	return res
}