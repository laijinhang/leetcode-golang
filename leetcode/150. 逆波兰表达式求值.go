package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	n := 0
	for i := 0;i < len(tokens);i++ {
		switch tokens[i] {
		case "+":
			stack[n-2] += stack[n-1]
			n--
		case "-":
			stack[n-2] -= stack[n-1]
			n--
		case "*":
			stack[n-2] *= stack[n-1]
			n--
		case "/":
			stack[n-2] /= stack[n-1]
			n--
		default:
			stack[n], _ = strconv.Atoi(tokens[i])
			n++
		}
	}
	return stack[0]
}
