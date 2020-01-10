package main

import (
	"strings"
	"fmt"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.Replace(paragraph, "!", " ", -1)
	paragraph = strings.Replace(paragraph, "?", " ", -1)
	paragraph = strings.Replace(paragraph, "'", " ", -1)
	paragraph = strings.Replace(paragraph, ",", " ", -1)
	paragraph = strings.Replace(paragraph, ";", " ", -1)
	paragraph = strings.Replace(paragraph, ".", " ", -1)
	bannedCont := make(map[string]bool)
	for i := 0;i < len(banned);i++ {
		bannedCont[banned[i]] = true
	}
	paragraph = strings.ToLower(paragraph)
	pgs := strings.Split(paragraph, " ")
	temps := []string{}
	for i := 0;i < len(pgs);i++ {
		if pgs[i] != "" {
			temps = append(temps, pgs[i])
		}
	}
	pgs = temps
	words := make(map[string]int)
	max := 0
	for i := 0;i < len(pgs);i++ {
		for i := 0;i < len(pgs);i++ {
			pgs[i] = strings.Trim(pgs[i], " !?',;.")
			if !bannedCont[pgs[i]] {
				words[pgs[i]]++
				if words[pgs[i]] > max {
					max = words[pgs[i]]
				}
			}
		}
	}
	res := ""
	for i := 0;i < len(pgs);i++ {
		for i := 0;i < len(pgs);i++ {
			if words[pgs[i]] == max {
				res = pgs[i]
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println()
}