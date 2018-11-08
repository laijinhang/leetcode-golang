package main

import "strings"

func shortestCompletingWord(licensePlate string, words []string) string {
	lpM := make(map[byte]int)
	wM := map[byte]int{}
	licensePlate = strings.ToLower(licensePlate)
	for i := 0;i < len(licensePlate);i++ {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			lpM[licensePlate[i]]++
		}
	}
	res := "aaaaaaaaaaaaaaaaaaa"
	for i := 0;i < len(words);i++ {
		if len(words[i]) < len(res) && len(words[i]) >= len(lpM) {
			wM = make(map[byte]int)
			for j := 0;j < len(words[i]);j++ {
				wM[words[i][j]]++
			}
			flag := true
			for k, v := range lpM {
				if wM[k] < v {
					flag = false
					break
				}
			}
			if flag {
				res = words[i]
			}
		}
	}
	return res
}