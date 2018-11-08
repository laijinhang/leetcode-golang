package main

import (
	"fmt"
	"math"
)

func titleToNumber(s string) int {
	n := 0
	for i := 0;i < len(s) - 1;i++ {
		n += int(s[i] - 'A' + 1) * int(math.Pow(26.0, float64(len(s) - i - 1)))
	}
	n += int(s[len(s)-1] - 'A' + 1)
	return n
}

func main()  {
	fmt.Println(titleToNumber("ZA"), 27 * 26)
}