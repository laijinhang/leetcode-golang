package main

import (
	"strings"
	"fmt"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	var s1, s2 string
	for {
		fmt.Scan(&s1, &s2)
		fmt.Println(strStr("dasd", ""))
	}
}