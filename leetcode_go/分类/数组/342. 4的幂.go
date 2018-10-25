package main

import (
	"strconv"
	"strings"
	"fmt"
)

func isPowerOfFour(num int) bool {
	if num & (num - 1) != 0 || num == 1{
		return false
	}
	s := strconv.FormatInt(int64(num), 2)
	fmt.Println(len(s) - strings.LastIndex(s, "1"))
	if (len(s) - strings.LastIndex(s, "1")) % 2 == 0 &&  strings.Count(s, "1") == 1 {
		return false
	}
	return true
}

func main() {
	isPowerOfFour(1)
}