/*
题目描述
请实现一个算法，在不使用额外数据结构和储存空间的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

给定一个string iniString，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

测试样例：
"This is nowcoder"
返回："redocwon si sihT"

 */
package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	str = reverseString(str)
	fmt.Println(str)
}

// 严格来说，这边会创建新的存储空间
func reverseString(str string) string {
	runes := []rune(str)
	length := len(str)

	for from, to := 0, length - 1;from < to;from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return str
}