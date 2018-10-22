package main

func isMonotonic(A []int) bool {
	flag := 0
	for i := 1;i < len(A);i++ {
		if A[i] - A[i-1] > 0 {
			if flag != 0 && flag != 1 {
				return false
			}
			flag = 1
		} else if A[i] - A[i-1] < 0 {
			if flag != 0 && flag != -1 {
				return false
			}
			flag = -1
		}
	}
	return true
}