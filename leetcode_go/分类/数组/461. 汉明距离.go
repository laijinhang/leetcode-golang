package main

import (
	"strconv"
	"fmt"
)

func hammingDistance(x int, y int) int {
	if x == y {
		return 0
	}
	if x < y {
		x, y = y, x
	}
	s1 := strconv.FormatInt(int64(x), 2)
	s2 := strconv.FormatInt(int64(y), 2)
	num := 0
	i, j := len(s1) - 1, len(s2) - 1
	for ;i >= 0 && j >= 0;i, j = i - 1, j - 1 {
		if s1[i] != s2[j] {
			num++
		}
	}
	for ;i >= 0;i-- {
		if s1[i] != '0' {
			num++
		}
	}
	return num
}

func main() {
	fmt.Println(hammingDistance(1, 2))
}