package main

import (
	"strings"
	"fmt"
)

func findSubstring(s string, words []string) []int {
	indexStr := make(map[int]string)
	strCount, tempStrCount := make(map[string]int), make(map[string]int)
	for i := 0;i < len(words);i++ {
		strCount[words[i]]++
		for index := -1;; {
			indexT := strings.Index(s[index+1:], words[i]) + index + 1
			if index == indexT {
				break
			}
			index = indexT
			indexStr[index] = words[i]
		}
	}
	res := []int{}
	for index, _ := range indexStr {
		flag := true
		tempStrCount = make(map[string]int)
		for i := 0;i < len(words);i++ {
			if _, ok := indexStr[index + i * len(words[0])];ok {
				tempStrCount[indexStr[index + i * len(words[0])]]++
				if tempStrCount[indexStr[index + i * len(words[0])]] > strCount[indexStr[index + i * len(words[0])]] {
					flag = false
					break
				}
			} else {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, index)
		}
	}
	return res
}

func main() {
	findSubstring("wordgoodstudentgoodword", []string{"word","student"})
}