package main

import "math"

func judgeSquareSum(c int) bool {
	flag := false
	for i := 0;i <= int(math.Sqrt(float64(c)));i++ {
		temp := c - i * i
		if (int(math.Sqrt(float64(temp))) * int(math.Sqrt(float64(temp)))) == temp {
			flag = true
			break
		}
	}
	return flag
}