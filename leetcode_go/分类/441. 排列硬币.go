package main

func arrangeCoins(n int) int {
	res := 0
	for i := 1;n > 0 && n >= i;i++ {
		n -= i
		res++
	}
	return res
}