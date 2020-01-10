package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	sLen := len(strs)
	if sLen == 0 {
		return ""
	}
	tempS := ""
	i, j := 0, 0
	for {
		if j >= len(strs[0]) {
			break
		}
		temp := strs[0][j]
		for i = 0; i < sLen; i++ {
			if j >= len(strs[i]) || strs[i][j] != temp {
				break
			}
		}
		if i == sLen {
			tempS += string(temp)
		} else {
			break
		}
		j++
	}
	return tempS
}

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
}