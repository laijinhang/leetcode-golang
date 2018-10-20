package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	dLen := len(digits)
	y := 0
	digits[dLen-1] += 1
	for i := dLen - 1;i >= 0;i-- {
		digits[i] += y
		y = digits[i] / 10
		if y == 0 {
			break
		}
		digits[i] %= 10
	}
	if y != 0 {
		digits = append([]int{y}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9}))
}