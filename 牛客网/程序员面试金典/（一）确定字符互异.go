/*
题目描述
请实现一个算法，确定一个字符串的所有字符是否全都不同。这里我们要求不允许使用额外的存储结构。

给定一个string iniString，请返回一个bool值,True代表所有字符全都不同，False代表存在相同的字符。保证字符串中的字符为ASCII字符。字符串的长度小于等于3000。

测试样例：
"aeiou"
返回：True
"BarackObama"
返回：False

分析：
根据题目可知，因为字符串中的字符为ASCII字符，所以字符串如果超过256的话，返回False，又因为不允许使用额外的存储结构，即小于256的话，就一个个对比了
 */
package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	res := checkDifferent(str)
	fmt.Println(res)
}

func checkDifferent(str string) string {
	length := len(str)
	if length > 256 {
		return "False"
	}
	for i := 0;i < length;i++ {
		for j := i + 1;j < length;j++ {
			if str[i] == str[j] {
				fmt.Println()
				return "False"
			}
		}
	}
	return "True"
}