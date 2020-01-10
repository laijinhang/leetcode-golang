package main

import "fmt"

func partitionLabels(S string) []int {
	pt := []int{}
	for i := 0;i < len(S);i++ {
		temp := i
		b := i
		for {
			j := temp + 1
			for ; j < len(S); j++ {
				if S[j] == S[b] {
					temp = j
				}
			}
			if temp == i { // 该字符在后面字符串中没有出现
				pt = append(pt, 1)
				break
			} else if b == temp {
				pt = append(pt, temp - i + 1)
				break
			}
			b++
		}
		i = b
	}
	return pt
}

func main() {
	pt := partitionLabels("ababcbaca")
	fmt.Println(pt)
}