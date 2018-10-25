package main

import "fmt"

func getSum(a int, b int) int {
	sum, carry := a ^ b, a & b << 1
	for carry != 0 {
		sum, carry = sum ^ carry, sum & carry << 1
	}
	return sum
}

func main() {
	fmt.Println(getSum(-1, 3))
}