package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = a + b, a + b
	fmt.Println(a, b)
}