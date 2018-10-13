/*
不可以！
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
判断：两个数x、y的正负性。

要求：不可以使用比较运算符，即"<",">","<=",">=","==","!="。



输入
有多组数据，每组数据占一行，每一行两个数x，y。
x、y保证在int范围内。
输出
每组数据输出占一行。
如果两个数是一正一负，输出"Signs are opposite"
如果是同为正或同为负，输出"Signs are not opposot"
如果无法确定，输出"Signs can't be sure"
输出不包括引号
样例输入
1 1
-1 1
样例输出
Signs are not opposot
Signs are opposite
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, y, n int
	var str string
	for {
		if _, err := fmt.Scan(&x, &y);err != nil {
			break
		}
		n = 0
		if strconv.Itoa(x)[0] != '-' {
			n++
		}
		if strconv.Itoa(y)[0] != '-' {
			n++
		}
		if strconv.Itoa(x)[0] == '0' || strconv.Itoa(y)[0] == '0' {
			n = -1
		}
		if n == 0 || n == 2 {	// 两负或两正
			str = "Signs are not opposot"
		} else if n == 1 {	// 一正一负
			str = "Signs are opposite"
		} else if n == -1 {	// 存在0
			str = "Signs can't be sure"
		}
		fmt.Println(str)
	}
}