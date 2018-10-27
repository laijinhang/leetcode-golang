package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	res := 0
	for i := 1;;i++ {
		if i * i <= x && x < (i + 1) * (i + 1) {
			res = i
			break
		}
	}
	return res
}