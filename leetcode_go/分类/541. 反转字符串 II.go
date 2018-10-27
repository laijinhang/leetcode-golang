package main

import "fmt"

func reverseStr(s string, k int) string {
	bs := []byte(s)
	i := 2 * k
	for ;i < len(bs);i += 2 * k {
		for n, m := i - 2 * k, i - k - 1;n < m;n, m = n + 1, m - 1 {
			bs[n], bs[m] = bs[m], bs[n]
		}
	}
	if i != len(bs) - 1 {
		i -= 2 * k
	}
	if i < len(bs) {
		n := 0
		m := i - 1
		if len(bs) - i < k {
			n = i
			m = len(bs) - 1
		} else {
			n = i
			m = i + k - 1
		}
		for ;n < m;n, m = n + 1, m - 1 {
			bs[n], bs[m] = bs[m], bs[n]
		}
	}
	return string(bs)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}