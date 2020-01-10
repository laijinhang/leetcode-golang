package main

import (
	"strconv"
)

func maximumSwap(num int) int {
	ns := []byte(strconv.Itoa(num))
	max := 0
	i := 0
	for ;i < len(ns) - 1;i++ {
		if ns[i] < ns[i+1] {
			break
		}
	}
	max = i + 1
	for j := i + 2;j < len(ns);j++ {
		if ns[j] >= ns[max] {
			max = j
		}
	}
	if max < len(ns) {
		for j := 0;j <= i;j++ {
			if ns[j] < ns[max] {
				ns[j], ns[max] = ns[max], ns[j]
				break
			}
		}
	}
	num, _ = strconv.Atoi(string(ns))
	return num
}