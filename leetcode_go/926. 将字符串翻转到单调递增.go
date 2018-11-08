package main

import (
	"strings"
	"fmt"
)

type MonoIncr struct {
	num int	// 次数
	char byte // 字符
}

func minFlipsMonoIncr(S string) int {
	S = strings.TrimLeft(S, "0")
	sLen := len(S)
	if sLen == 0 {
		return 0
	}
	m := []MonoIncr{}
	m = append(m, MonoIncr{1,S[0]})
	for i := 1;i < sLen;i++ {
		if S[i] == m[len(m)-1].char {
			m[len(m)-1].num++
		} else {
			m = append(m, MonoIncr{1, S[i]})
		}
	}
	num := sLen
	tempNum := 0
	for i := 0;i < len(m);i++ {
		// 0 -> 1, 1 -> 1全部翻转
		{
			temp := tempNum
			for j := i;j < len(m);j++ {
				if m[j].char == '0' {
					temp += m[j].num
				}
			}
			if temp < num {
				num = temp
			}
		}
		// 0->0, 1->0 翻转部分
		{
			if m[i].char == '1' {
				tempNum += m[i].num
			}
		}
	}
	if tempNum < num {
		num = tempNum
	}
	return num
}

func main() {
	fmt.Println(minFlipsMonoIncr("00110"))
}