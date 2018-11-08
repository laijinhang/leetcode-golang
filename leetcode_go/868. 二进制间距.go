package main

func binaryGap(N int) int {
	maxL := 0
	n := -1
	for ;N != 0;N >>= 1 {
		if N & 1 == 1 {
			if maxL < n {
				maxL = n
			}
			n = 0
		}
		if n != -1 {
			n++
		}
	}
	return maxL
}