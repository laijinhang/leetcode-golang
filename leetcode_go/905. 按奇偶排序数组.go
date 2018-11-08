package main

func sortArrayByParity(A []int) []int {
	aLen := len(A)
	a1 := make([]int, 0, aLen)
	a2 := make([]int, 0, aLen)
	for i := 0;i < aLen;i++ {
		if A[i] % 2 == 0 {
			a1 = append(a1, A[i])
		} else {
			a2 = append(a2, A[i])
		}
	}
	return append(a1, a2...)
}