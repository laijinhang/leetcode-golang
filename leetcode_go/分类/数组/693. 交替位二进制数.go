package main

import (
	"strconv"
	"fmt"
)

func hasAlternatingBits(n int) bool {
	res := true
	s := strconv.FormatInt(int64(n), 2)
	for i := 0;i < len(s) - 1;i++ {
		if s[i] == s[i+1] {
			res = false
			break
		}
	}
	return res
}
