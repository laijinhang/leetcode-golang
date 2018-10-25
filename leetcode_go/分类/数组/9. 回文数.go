package main

import "strconv"

func isPalindrome(x int) bool {
	b1 := []byte(strconv.Itoa(x))
	b2 := make([]byte, len(b1))
	for i := 0;i < len(b1);i++ {
		b2[len(b1)-i-1] = b1[i]
	}
	return string(b1) == string(b2)
}
