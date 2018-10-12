/*
大小写互换
时间限制：1000 ms  |  内存限制：65535 KB
难度：0
描述
      现在给出了一个只包含大小写字母的字符串，不含空格和换行，要求把其中的大写换成小写，小写换成大写，然后输出互换后的字符串。
输入
第一行只有一个整数m（m<=10),表示测试数据组数。
接下来的m行，每行有一个字符串（长度不超过100）。
输出
输出互换后的字符串，每组输出占一行。
样例输入
2
Acm
ACCEPTED
样例输出
aCM
accepted
 */
package main

import "fmt"

func main() {
	var m int
	var str string
	fmt.Scan(&m)
	for i := 0;i < m;i++ {
		fmt.Scan(&str)
		for _, s := range str {
			if s >= 'a' && s <= 'z' {
				fmt.Print(string(s - 'a' + 'A'))
			} else if s >= 'A' && s <= 'Z' {
				fmt.Print(string(s - 'A' + 'a'))
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println()
	}
}