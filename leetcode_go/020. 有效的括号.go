package main

import "strings"

func isValid(s string) bool {
	stack := []byte{}
	flag := true
	for i := 0;i < len(s) && flag;i++ {
		if len(stack) == 0 && (s[i] == ')' || s[i] == '}' || s[i] == ']') {
			flag = false
			break
		}
		if s[i] == ')' {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else if s[i] == '}' {
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else if s[i] == ']' {
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
		strings.Index("-", "")
	}
	if len(stack) != 0 {
		flag = false
	}
	return flag
}