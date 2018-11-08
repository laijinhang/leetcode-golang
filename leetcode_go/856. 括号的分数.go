package main

import (
	"strings"
	"fmt"
)

/*
参考：https://cloud.tencent.com/developer/article/1152121
 */
func scoreOfParentheses(S string) int {
	S = strings.Replace(S, "()", "1", -1)
	i, n, sum := 0, 0, 0
	stack := make([]int, len(S))
	for ;i < len(S);i++ {
		switch S[i] {
		case ')':
			temp := 0
			for n--;n >= 0 && stack[n] != '(';n-- {
				temp += int(stack[n] - '0')
			}
			temp *= 2
			stack[n] = int(temp) + '0'
			n++
		default:
			stack[n] = int(S[i])
			n++
		}
	}
	for n--;n >= 0;n-- {
		sum += int(stack[n] - '0')
	}
	return sum
}

func main() {
	fmt.Println(scoreOfParentheses("((((((((((((((((((((((((()))))))))))))))))))))))))"))
	fmt.Println(scoreOfParentheses("(((((((((())))))))))"))
}