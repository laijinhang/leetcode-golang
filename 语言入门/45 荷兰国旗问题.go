/*
荷兰国旗问题
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述

荷兰国旗有三横条块构成，自上到下的三条块颜色依次为红、白、蓝。现有若干由红、白、蓝三种颜色的条块序列，要将它们重新排列使所有相同颜色的条块在一起。本问题要求将所有红色的条块放最左边、所有白色的条块放中间、所有蓝色的条块放最右边。
输入
第1行是一个正整数n（n<100），表示有n组测试数据。接下来有n行，每行有若干个由R，W，B三种字符构成的字符串序列，其中R，W和B分别表示红、白、蓝三种颜色的条块,每行最多有1000个字符。
输出
对输入中每行上由R，W，B三种字符构成的字符串序列，将它们重新排列使所有相同颜色的条块在一起，满足前述要求。
样例输入
3
BBRRWBWRRR
RRRWWRWRB
RBRW
样例输出
RRRRRWWBBB
RRRRRWWWB
RRWB
 */
package main

import "fmt"

func main() {
	var N int
	var str string
	xl := [3]int{}
	xs := []string{"R", "W", "B"}
	fmt.Scan(&N)
	for i := 0;i < N;i++ {
		fmt.Scan(&str)
		xl = [3]int{}
		for _, s := range str {
			switch s {
			case 'R':
				xl[0]++
			case 'W':
				xl[1]++
			case 'B':
				xl[2]++
			}
		}
		for i := 0;i < 3;i++ {
			for j := 0;j < xl[i];j++ {
				fmt.Print(xs[i])
			}
		}
		fmt.Println()
	}
}