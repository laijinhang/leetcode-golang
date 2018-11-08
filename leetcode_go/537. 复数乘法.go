package main

import (
	"fmt"
	"strings"
	"strconv"
)

func complexNumberMultiply(a string, b string) string {
	s1 := strings.Split(a, "+")
	s2 := strings.Split(b, "+")
	a1, _ := strconv.Atoi(s1[0])
	a2, _ := strconv.Atoi(s1[1][:len(s1[1])-1])
	b1, _ := strconv.Atoi(s2[0])
	b2, _ := strconv.Atoi(s2[1][:len(s2[1])-1])
	return fmt.Sprintf("%d+%di", a1 * b1 - a2 * b2, a1 * b2 + a2 * b1)
}

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}