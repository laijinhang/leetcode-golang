package main

import (
	"strconv"
	"math"
	"fmt"
)

func reverse(x int) int {
	b := []byte{}
	fh := 1
	if x < 0 {
		fh = -1
		x = -x
	}
	for x != 0 {
		b = append(b, byte(x % 10 + '0'))
		x /= 10
	}
	val, _ := strconv.ParseInt(string(b), 10, 0)
	val *= int64(fh)
	if val < math.MinInt32 || val > math.MaxInt32 {
		return 0
	}
	return int(val)
}

func main()  {
	fmt.Println(reverse(-123))
}