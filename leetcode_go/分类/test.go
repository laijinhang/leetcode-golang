package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{1,1,4,2}
	b := a
	a = append(a, 1)
	sort.Ints(a)
	fmt.Println(a, b)
}