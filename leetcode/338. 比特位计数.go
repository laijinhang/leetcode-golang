package main

func countBits(num int) []int {
	res := []int{}
	for i := 0;i <= num;i++ {
		x := 0
		for n := i;n != 0;n >>= 1 {
			if n & 1 == 1 {
				x++
			}
		}
		res = append(res, x)
	}
	return res
}