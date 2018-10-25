package main

import (
	"strconv"
	"strings"
	"fmt"
)

// 10^6 = 11110100001001000000	20位二进制，所以2的个数最多是20
var m = map[int]bool{
	2:true,3:true,
	5:true,7:true,
	11:true,13:true,
	17:true,19:true,
}

func countPrimeSetBits(L int, R int) int {
	num := 0
	for ;L <= R;L++ {
		s := strconv.FormatInt(int64(L), 2)
		if m[strings.Count(s, "1")] {
			num++
		}
	}
	return num
}

func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}