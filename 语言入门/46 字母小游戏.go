/*
字母小游戏
时间限制：1000 ms  |  内存限制：65535 KB
难度：0
描述
给你一个乱序的字符串,里面包含有小写字母(a--z)以及一些特殊符号，请你找出所给字符串里面所有的小写字母的个数， 拿这个数对26取余，输出取余后的数字在子母表中对应的小写字母(0对应z,1对应a，2对应b....25对应y)。
输入
第一行是一个整数n(1<n<1000)表示接下来有n行的字符串m(1<m<200)需要输入
输出
输出对应的小写字母 每个小写字母单独占一行
样例输入
2
asdasl+%$^&ksdhkjhjksd
adklf&(%^(alkha
样例输出
q
j
 */
package main

import "fmt"

func main() {
	var n, num int
	var str string
	fmt.Scan(&n)
	for i := 0;i < n;i++ {
		fmt.Scan(&str)
		num = 0
		for _, s := range str {
			if s >= 'a' && s <= 'z' {
				num++
			}
		}
		num %= 26
		if num == 0 {
			fmt.Println("z")
		} else {
			fmt.Println(string(num + 'a' - 1))
		}
	}
}