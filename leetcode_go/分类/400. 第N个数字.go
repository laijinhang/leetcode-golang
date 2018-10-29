package main

import (
	"strconv"
	"fmt"
)

func findNthDigit(n int) int {
	res := 0
	if n >= 1 && n <= 9 {
		res = n
	} else {
		mul := 1
		sub := 0
		num := 0
		if n >= 10 && n < 190 {
			mul = 10
			sub = 10
			num = 2
		} else if n >= 190 && n < 2890 {
			mul = 100
			sub = 190
			num = 3
		} else if n >= 2890 && n < 38890 {
			mul = 1000
			sub = 2890
			num = 4
		} else if n >= 38890 && n < 488890 {
			mul = 10000
			sub = 38890
			num = 5
		} else if n >= 488890 && n < 6088890 {
			mul = 100000
			sub = 488890
			num = 6
		} else if n >= 6088890 && n < 69088890 {
			mul = 1000000
			sub = 6088890
			num = 7
		} else {
			mul = 10000000
			sub = 69088890
			num = 8
		}
		// 定值
		val := mul + (n - sub) / num
		// 定位置
		position := (n - sub) % num
		fmt.Println(val, position, (n - sub) % num)
		res = int(strconv.Itoa(val)[position] - '0')
	}
	return res
}

func main() {
	fmt.Println(findNthDigit(10000000))
}