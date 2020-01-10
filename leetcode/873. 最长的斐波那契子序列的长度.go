package main

func lenLongestFibSubseq(A []int) int {
	maxLen := 0
	res := 0
	vals := make(map[int]bool)
	for i := 0;i < len(A);i++ {
		vals[A[i]] = true
	}
	a, b := 0, 0
	for i := 0;i < len(A);i++ {
		for j := i + 1;j < len(A);j++ {
			a, b = A[i], A[j]
			maxLen = 0
			for ;vals[a+b];a, b = b, a + b {
				maxLen++
			}
			if maxLen > res {
				res = maxLen
			}
		}
	}
	if res != 0 {
		res += 2
	}
	return res
}