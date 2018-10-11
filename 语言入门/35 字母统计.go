/*
字母统计
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
现在给你一个由小写字母组成字符串，要你找出字符串中出现次数最多的字母，如果出现次数最多字母有多个那么输出最小的那个。
输入
第一行输入一个正整数T（0<T<25）
随后T行输入一个字符串s,s长度小于1010。
输出
每组数据输出占一行，输出出现次数最多的字符；
样例输入
3
abcd
bbaa
jsdhfjkshdfjksahdfjkhsajkf
样例输出
a
a
j
 */
package main

import "fmt"

func main() {
	var T int
	var s string
	var n[26] int
	fmt.Scan(&T)
	for ;T > 0;T-- {
		n = [26]int{}
		fmt.Scan(&s)
		for i := len(s) - 1;i >= 0;i-- {
			n[s[i]-'a']++
		}
		maxN := 0
		for i := 0;i < 26;i++ {
			if n[i] > n[maxN] {
				maxN = i
			}
		}
		fmt.Println(string(maxN + 'a'))
	}
}