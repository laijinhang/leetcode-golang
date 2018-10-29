package main

import "strings"

func wordPattern(pattern string, str string) bool {
	patt1 := make(map[byte]string)
	patt2 := make(map[string]byte)
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i := 0;i < len(pattern);i++ {
		if _, ok := patt1[pattern[i]];!ok {
			patt1[pattern[i]] = strs[i]
		} else if patt1[pattern[i]] != strs[i] {
			return false
		}
		if _, ok := patt2[strs[i]];!ok {
			patt2[strs[i]] = pattern[i]
		} else if patt2[strs[i]] != pattern[i] {
			return false
		}
	}
	return true
}