package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	strm := []map[byte]int{}
	for i, j := 0, 0;i < len(strs);i++ {
		m := make(map[byte]int)
		for j = 0;j < len(strs[i]);j++ {
			m[strs[i][j]]++
		}
		for j = 0;j < len(res);j++ {
			if len(strm[j]) == len(m) {
				flag := true
				for k, v := range strm[j] {
					if m[k] != v {
						flag = false
						break
					}
				}
				if flag {
					res[j] = append(res[j], strs[i])
					break
				}
			}
		}
		if j == len(res) {
			res = append(res, append([]string{}, strs[i]))
			strm = append(strm, m)
		}
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}