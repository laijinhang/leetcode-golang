package main

import "fmt"

func customSortString(S string, T string) string {
	ref := make(map[int]byte)
	m := make(map[byte]bool)
	for i := 0;i < len(S);i++ {
		ref[i] = S[i]
		m[S[i]] = true
	}
	res := make([]byte, len(T))
	n := 0
	s := make(map[byte]int)
	for i := 0;i < len(T);i++ {
		if !m[T[i]] {
			res[n] = T[i]
			n++
		} else {
			s[T[i]]++
		}
	}
	for i := 0;i < len(ref);i++ {
		for j := 0;j < s[ref[i]];j++ {
			res[n] = ref[i]
			n++
		}
	}
	return string(res)
}

func main() {
	fmt.Println(customSortString("bca", "bacd"))
}