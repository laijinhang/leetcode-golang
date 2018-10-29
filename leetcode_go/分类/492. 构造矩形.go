package main

import (
	"fmt"
)

func constructRectangle(area int) []int {
	L, W := area, 1
	minL := area - 1
	i, j := area / 2, 2
	for ;i >= j;i-- {
		for j := 2;j <= i;j++ {
			if i * j > area {
				break
			} else if i * j == area {
				if minL > i - j {
					minL = i - j
					L = i
					W = j
				}
				break
			}
		}
	}
	return []int{L,W}
}

func main() {
	fmt.Println(constructRectangle(9999998))
}