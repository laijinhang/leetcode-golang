package main

import "strconv"

func findComplement(num int) int {
	b1 := []byte(strconv.FormatInt(int64(num), 2))
	for i := 0;i < len(b1);i++ {
		if b1[i] == '0' {
			b1[i] += 1
		} else {
			b1[i] -= 1
		}
	}
	res, _ := strconv.ParseInt(string(b1), 2, 0)
	return int(res)
}