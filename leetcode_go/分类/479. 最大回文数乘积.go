package main

import (
	"strconv"
	"fmt"
	"time"
)

func largestPalindrome(n int) int {
	i, j := 1, 0
	for ;n > 1;n-- {
		i *= 10
	}
	e := i
	minI := 0
	max := int64(0)
	for i = i * 10 - 1;i >= e;i-- {
		if i < minI && int64(i * (e * 10 - 1)) < max {
			break
		}
		for j = e * 10 - 1;j > minI && j >= e;j-- {
			if isValid(strconv.FormatInt(int64(i * j), 10)) {
				max = int64(i * j)
				minI = j
				break
			}
		}
	}
	return int(max % 1337)
}

func isValid(s string) bool {
	flag := true
	for i, j := 0, len(s) - 1;i < j;i, j = i + 1, j - 1 {
		if s[i] != s[j] {
			flag = false
			break
		}
	}
	return flag
}

func main() {
	for i := 1;i < 9;i++ {
		fmt.Println(i)
		b := time.Now()
		fmt.Println(largestPalindrome(i))
		fmt.Println("Run ", time.Since(b))
		fmt.Println("--------------------")
	}
}