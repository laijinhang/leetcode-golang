package main

import "fmt"

func reverseVowels(s string) string {
	vowelM := map[byte]bool {
		'a': true, 'A': true,
		'e': true, 'E': true,
		'i': true, 'I': true,
		'o': true, 'O': true,
		'u': true, 'U': true,
	}
	bs := []byte(s)
	for i, j := 0, len(bs) - 1;i < j;i, j = i + 1, j - 1 {
		for ;i < j && !vowelM[bs[i]];i++ {}
		for ;j > i && !vowelM[bs[j]];j-- {}
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}