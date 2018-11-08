package main

import "fmt"

type str struct {
	a []byte	// 奇数
	b []byte	// 偶数
}

func numSpecialEquivGroups(A []string) int {
	s := make([]str, len(A))
	for i := 0;i < len(s);i++ {
		s[i].a = make([]byte, 26)
		s[i].b = make([]byte, 26)
		for j := 0;j < len(A[i]);j++ {
			if j % 2 == 0 {
				s[i].b[A[i][j]-'a']++
			} else {
				s[i].a[A[i][j]-'a']++
			}
		}
	}
	goThrough := make(map[int]bool)
	num := 0
	for i := 0;i < len(s);i++ {
		if !goThrough[i] {
			num++
			for j := i + 1; j < len(s); j++ {
				if string(s[i].a) == string(s[j].a) && string(s[i].b) == string(s[j].b) {
					goThrough[j] = true
				}
			}
		}
	}
	return num
}

func main() {
	a := []byte{1,2}
	b := []byte{1,2}
	fmt.Println(string(a) == string(b))
}