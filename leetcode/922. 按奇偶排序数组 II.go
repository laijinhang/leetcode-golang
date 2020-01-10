package main

func sortArrayByParityII(A []int) []int {
	B := append([]int{}, A...)
	i1, i2 := 0, 1
	for i1 < len(B) || i2 < len(B) {
		for ;i1 < len(B) && B[i1] % 2 == 0;i1 += 2 {}
		for ;i2 < len(B) && B[i2] % 2 == 0;i2 += 2 {}
		if i1 < len(B) && i2 < len(B) {
			B[i1], B[i2] = B[i2], B[i1]
		}
	}
	return B
}