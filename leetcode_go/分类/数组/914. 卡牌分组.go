package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	// 1）压缩
	// 2）找出个数有1和次数最少的那个
	// 3）重组
	// 4）计算个数是否满足
	m := make(map[int]int)
	for i := 0;i < len(deck);i++ {
		m[deck[i]]++
	}
	fmt.Println(m)
	min := len(deck)
	for i := 0;i < len(deck);i++ {
		if m[deck[i]] == 1 {
			return false
		}
		if min > deck[i] {
			min = deck[i]
		}
 	}

 	c := make(map[int]int)
 	for _, count := range m {
 		for i := 2;i <= min;i++ {
 			if count % i == 0 {
 				c[i]++
			}
			fmt.Println(count % i)
		}
	}
	fmt.Println(c)
	for i := 2;i <= min;i++ {
		if len(deck) == c[i] * i {
			return true
		}
	}
	return false
}