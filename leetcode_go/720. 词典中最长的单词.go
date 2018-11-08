package main

import (
	"sort"
	"fmt"
)

/*
// 1）存入map中
// 2）对words排序
// 3）前缀匹配
 */
type strs []string

func (s strs) Len() int {
	return len(s)
}

func (s strs) Less(i, j int) bool {
	if len(s[i]) > len(s[j]) {
		return false
	}
	if len(s[i]) == len(s[j]) && s[i] > s[j] {
		return false
	}
	return true
}

func (s strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func longestWord(words []string) string {
	// 1）存入map中
	wordsM := make(map[string]bool)
	for i := 0;i < len(words);i++ {
		wordsM[words[i]] = true
	}
	// 2）对words排序
	sort.Sort(strs(words))
	// 3）前缀匹配
	maxLenWords := ""
	for i := len(words) - 1;i >= 0;i-- {
		j := 1
		for ;j < len(words[i]);j++ {
			if !wordsM[words[i][:j]] {
				break
			}
		}
		fmt.Println(words[i], j)
		if j == len(words[i]) && (maxLenWords == "" || len(words[i]) >= len(maxLenWords)) {
			maxLenWords = words[i]
		}
	}
	return maxLenWords
}

/*
type strs []string

func (s strs) Len() int {
	return len(s)
}

func (s strs) Less(i, j int) bool {
	if len(s[i]) > len(s[j]) {
		return false
	}
	if len(s[i]) == len(s[j]) && s[i] > s[j] {
		return false
	}
	return true
}

func (s strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func longestWord(words []string) string {
	wordGroup := make([][]string, 31)
	// 字典序
	wordsNumber := make(map[string]int)
	// 1）按单词个数分组
	for i := 0;i < len(words);i++ {
		wordsNumber[words[i]] = i
		wordGroup[len(words[i])] = append(wordGroup[len(words[i])], words[i])
	}
	// 排序
	sort.Sort(strs(words))
	fmt.Println(words)
	// 2）判断，找出第一个符合的字符串
	maxLenWorld := ""
	for i := len(words) - 1;i >= 0;i-- {
		flag := true
		for j := 0;j < len(wordGroup);j++ {
			if len(wordGroup[j]) > 0 {
				tempFlag := true
				for k := 0; k < len(wordGroup[j]) && len(wordGroup[j][k]) < len(words[i]); k++ {
					if strings.Index(words[i], wordGroup[j][k]) != -1 {
						tempFlag = true
						break
					} else {
						tempFlag = false
					}
				}
				if tempFlag == false {
					flag = false
					j = len(wordGroup)
				}
			}
		}
		if flag {
			if len(words[i]) >= len(maxLenWorld) && (maxLenWorld == "" || wordsNumber[maxLenWorld] > wordsNumber[words[i]]) {
				maxLenWorld = words[i]
			}
			if i > 0 && len(words[i]) > len(words[i-1]) {
				break
			}
		}
	}
	return maxLenWorld
}
*/