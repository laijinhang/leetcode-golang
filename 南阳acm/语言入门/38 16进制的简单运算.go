/*
16进制的简单运算
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
现在给你一个16进制的加减法的表达式，要求用8进制输出表达式的结果。
输入
第一行输入一个正整数T（0<T<100000）
接下来有T行，每行输入一个字符串s（长度小于15)字符串中有两个数和一个加号或者一个减号，且表达式合法并且所有运算的数都小于31位
输出
每个表达式输出占一行，输出表达式8进制的结果。
样例输入
3
29+4823
18be+6784
4ae1-3d6c
样例输出
44114
100102
6565
 */
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var T int
	var s string
	fmt.Scan(&T)
	for ;T > 0;T-- {
		fmt.Scan(&s)
		if strings.Index(s, "+") != -1 {
			number := strings.Split(s, "+")
			a1, _ := strconv.ParseInt(number[0], 16, 0)
			a2, _ := strconv.ParseInt(number[1], 16, 0)
			fmt.Println(strconv.FormatInt(a1 + a2, 8))
		} else {
			number := strings.Split(s, "-")
			a1, _ := strconv.ParseInt(number[0], 16, 0)
			a2, _ := strconv.ParseInt(number[1], 16, 0)
			fmt.Println(strconv.FormatInt(a1 - a2, 8))
		}
	}
}