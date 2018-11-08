package main

import "fmt"

var res []string

func letterCasePermutation(S string) []string {
	res = []string{}
	permutation([]byte(S), 0)
	fmt.Println(res)
	return res
}

func permutation(s []byte, i int) {
	for i == len(s) {
		res = append(res, string(s))
		return
	}
	permutation(s, i + 1)
	if s[i] >= 'a' && s[i] <= 'z' {
		t := append([]byte{}, s...)
		t[i] = s[i] - 'a' + 'A'
		permutation(t, i + 1)
	}
	if s[i] >= 'A' && s[i] <= 'Z' {
		t := append([]byte{}, s...)
		t[i] = s[i] - 'A' + 'a'
		permutation(t, i + 1)
	}
}

func main() {
	letterCasePermutation("s8a")
}