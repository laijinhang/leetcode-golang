package main

import (
	"sort"
	"fmt"
)

func reorganizeString(S string) string {
	if len(S) == 0 {
		return ""
	}
	chars := make(map[byte]int)
	for i := 0;i < len(S);i++ {
		chars[S[i]]++
	}
	numbers := make([]int, len(chars))
	i := 0
	for _, count := range chars {
		numbers[i] = count
		i++
	}
	sort.Ints(numbers)
	charsSort := make(map[int]byte)
	for char, count := range chars {
		for i := 0;i < len(numbers);i++ {
			if numbers[i] == count {
				charsSort[len(numbers) - i] = char
				numbers[i] = -1
				break
			}
		}
	}
	if chars[charsSort[1]] > len(S) / 2 + len(S) % 2 {
		return ""
	}
	// 拼接
	resb := make([]byte, len(S))
	n := 1
	for i := 0;i < len(S);i += 2 {
		resb[i] = charsSort[n]
		chars[charsSort[n]]--
		if chars[charsSort[n]] <= 0 {
			n++
		}
	}
	for i := 1;i < len(S);i += 2 {
		resb[i] = charsSort[n]
		chars[charsSort[n]]--
		if chars[charsSort[n]] <= 0 {
			n++
		}
	}
	return string(resb)
}

func main() {
	fmt.Println(reorganizeString("aaabbbccc"))
}