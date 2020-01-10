package main

import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2;i <= int(math.Sqrt(float64(num)));i++ {
		if num % i == 0 {
			sum = sum + i + num / i
		}
	}
	if sum != num {
		return false
	}
	return true
}