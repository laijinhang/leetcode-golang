package main

import "strings"

func uncommonFromSentences(A string, B string) []string {
	ms := make(map[string]int)
	s1 := strings.Split(A, " ")
	s2 := strings.Split(B, " ")
	for i := 0;i < len(s1);i++ {
		ms[s1[i]]++
	}
	for i := 0;i < len(s2);i++ {
		ms[s2[i]]++
	}
	res := []string{}
	for k, v := range ms {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}