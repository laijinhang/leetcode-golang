/*
心急的C小加
时间限制：1000 ms  |  内存限制：65535 KB
难度：4
描述
C小加有一些木棒，它们的长度和质量都已经知道，需要一个机器处理这些木棒，机器开启的时候需要耗费一个单位的时间，如果第i+1个木棒的重量和长度都大于等于第i个处理的木棒，那么将不会耗费时间，否则需要消耗一个单位的时间。因为急着去约会，C小加想在最短的时间内把木棒处理完，你能告诉他应该怎样做吗？

输入
第一行是一个整数T(1<T<1500)，表示输入数据一共有T组。
每组测试数据的第一行是一个整数N（1<=N<=5000）,表示有N个木棒。接下来的一行分别输入N个木棒的L，W（0 < L ,W <= 10000），用一个空格隔开，分别表示木棒的长度和质量。
输出
处理这些木棒的最短时间。
样例输入
3
5
4 9 5 2 2 1 3 5 1 4
3
2 2 1 1 2 2
3
1 3 2 2 3 1
样例输出
2
1
3
 */
package main

import (
	"fmt"
	"sort"
)

type stick struct {
	l, w int
}

type newStick []stick

func main() {
	var T, N int
	fmt.Scan(&T)
	for i := 0;i < T;i++ {
		fmt.Scan(&N)
		tempS := stick{}
		s := []stick{}
		for i := 0;i < N;i++ {
			fmt.Scan(&tempS.l, &tempS.w)
			s = append(s, tempS)
		}
		sort.Sort(newStick(s))
		num := 0
		for i := 0;i < len(s);i++ {
			if s[i].w != 0 {
				tw := s[i].w
				s[i].w = 0
				num++
				for j := i + 1;j < len(s);j++ {
					if s[j].w >= tw {
						tw = s[j].w
						s[j].w = 0
					}
				}
			}
		}
		fmt.Println(num)
	}
}

func (s newStick) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s newStick) Less(i, j int) bool {
	if s[i].l <= s[j].l && s[i].w <= s[j].w {
		return true
	}
	return false
}

func (s newStick) Len() int {
	return len(s)
}