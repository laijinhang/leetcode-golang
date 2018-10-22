package main

import (
	"time"
	"fmt"
)

func main() {
	a := []int{1,2,3,4,9}
	var l1, l2 int
	bt := time.Now()
	l1 = len(a)
	for i := 0;i < 10000000000;i++ {
		l1 = l1
	}
	fmt.Println("Run ", time.Since(bt))

	bt = time.Now()
	l2 = l2
	for i := 0;i < 10000000000;i++ {
		l2 = len(a)
	}
	fmt.Println("Run ", time.Since(bt))
}