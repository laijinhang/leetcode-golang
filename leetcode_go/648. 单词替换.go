package main

import (
	"sort"
	"fmt"
	"strings"
)

type strs []string

func (s strs) Less(i, j int) bool {
	if len(s[i]) >= len(s[j]) {
		return false
	}
	return true
}

func (s strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s strs) Len() int {
	return len(s)
}

func replaceWords(dict []string, sentence string) string {
	sort.Sort(strs(dict))
	words := strings.Split(sentence, " ")
	for i := 0;i < len(words);i++ {
		for j := 0;j < len(dict);j++ {
			if len(words[i]) > len(dict[j]) && words[i][:len(dict[j])] == dict[j] {
				words[i] = dict[j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	replaceWords([]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa")
}