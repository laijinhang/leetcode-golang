package main

import (
	"strconv"
	"fmt"
)

func toHex(num int) string {
	t := uint32(num)
	return strconv.FormatInt(int64(t), 16)
}

func main() {
	fmt.Println(toHex(-1))
}