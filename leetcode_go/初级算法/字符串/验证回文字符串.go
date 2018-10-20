package main

import (
	"strings"
	"fmt"
)

func isPalindrome(s string) bool {
	sLen := len(s)
	s = strings.ToUpper(s)
	for i, j := 0, sLen - 1;i < j; {
		for ;i < j && ((s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9')) == false;i++ {}
		for ;j > i && ((s[j] >= 'A' && s[j] <= 'Z') || (s[j] >= '0' && s[j] <= '9')) == false;j-- {}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main()  {
	fmt.Println(isPalindrome("race a car"))
}