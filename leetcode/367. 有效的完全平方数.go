package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	res := false
	for i := 2;i <= num / 2;i++ {
		if i * i == num {
			res = true
			break
		}
	}
	return res
}