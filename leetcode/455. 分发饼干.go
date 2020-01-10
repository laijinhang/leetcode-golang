package main

import (
	"sort"
	"fmt"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	for i := 0;i < len(s);i++ {
		for j := len(g) - 1;j >= 0;j-- {
			if g[j] <= s[i] && g[j] != -1 {
				g[j] = -1
				s[i] = -1
				break
			}
		}
	}
	num := 0
	for j := 0;j < len(g);j++ {
		if g[j] == -1 {
			num++
		}
	}
	return num
}

func main() {
	n := findContentChildren([]int{1,2}, []int{1,1})
	fmt.Println(n)
}