package main

func numSubarrayBoundedMax(A []int, L int, R int) int {
	num := 0
	tempMax := 0
	for i := 0;i < len(A);i++ {
		tempMax = A[i]
		for j := i;j < len(A) && A[j] <= R;j++ {
			if A[j] > tempMax {
				tempMax = A[j]
			}
			if tempMax >= L && tempMax <= R {
				num++
			}
		}
	}
	return num
}