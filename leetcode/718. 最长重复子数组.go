package main

func findLength(A []int, B []int) int {
	res := 0
	for i := 0;i < len(A);i++ {
		for j := 0;j < len(B);j++ {
			l := 0
			for tempI, tempJ := i, j;tempI < len(A) && tempJ < len(B) && A[tempI] == B[tempJ];tempI, tempJ = tempI + 1, tempJ + 1 {l++}
			if l > res {
				res = l
			}
		}
	}
	return res
}