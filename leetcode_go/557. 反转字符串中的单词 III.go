package main

import (
	"strings"
	"math"
)

func reverseWords(s string) string {
	ss := strings.Fields(s)
	for key, str := range ss {
		ss[key] = reverse(str)
	}
	return strings.Join(ss, " ")
}

func reverse(s string) string {
	bs := []byte(s)
	for i, j := 0, len(s) - 1;i < j;i, j = i + 1, j - 1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func main() {
}