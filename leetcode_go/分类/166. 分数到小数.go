package main

import (
	"strconv"
	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ""
	}
	if numerator == 0 {
		return "0"
	}
	res := ""
	if numerator * denominator < 0 {
		res += "-"
		if numerator < 0 {
			numerator = -numerator
		}
		if denominator < 0 {
			denominator = -denominator
		}
	}
	rm := make(map[int]int)
	res += strconv.Itoa(numerator / denominator)
	r := []int{}
	for _, ok := rm[numerator % denominator];numerator % denominator != 0 && !ok;_, ok = rm[numerator % denominator] {
		rm[numerator % denominator] = len(r)
		numerator = numerator % denominator * 10
		r = append(r, numerator / denominator)
	}
	if len(r) != 0 {
		res += "."
		if numerator % denominator != 0 {
			for i := 0; i < rm[numerator%denominator]; i++ {
				res += strconv.Itoa(r[i])
			}
			if rm[numerator%denominator] < len(r) {
				res += "("
			}
			for i := rm[numerator%denominator]; i < len(r); i++ {
				res += strconv.Itoa(r[i])
			}
			if rm[numerator%denominator] < len(r) {
				res += ")"
			}
		} else {
			for i := 0; i < len(r); i++ {
				res += strconv.Itoa(r[i])
			}
		}
	}
	return res
}

func main() {
	fmt.Println(fractionToDecimal(-50, 8))
}