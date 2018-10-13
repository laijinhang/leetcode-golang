/*
数字分隔（二）
时间限制：1000 ms  |  内存限制：65535 KB
难度：3
描述
在一个遥远的国家，银行为了更快更好的处理用户的订单，决定将一整串的数字按照一定的规则分隔开来，分隔规则如下：

1、实数的整数部分按照每三个数字用逗号分隔开（整数部分的高位有多余的0时，需先将多余的0过滤后，再进行数字分隔，如：0001234567 输出结果为1,234,567.00）

2、小数部分保留两位小数(四舍五入)

3、如果该数是负的，则在输出时需用括号将分隔后的数字括起来，例如：-10005.1645的输出结果为(10,005.16)

输入
多组测试数据(以eof结尾)，每行输入一个实数n（n的位数小于100）
输出
输出分隔后的结果
样例输入
0001234567
0.0000
-10005.1645
样例输出
1,234,567.00
0.00
(10,005.16)
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, s string
	var e float64
	for {
		fmt.Scan(&n)
		e, _ = strconv.ParseFloat(n, 0)
		if e < 0 {
			s = fmt.Sprintf("(%.2f)", -e)
		} else {
			s = fmt.Sprintf("%.2f", e)
		}
		// 去掉首尾
		b := 0
		if s[0] == '(' {b = 1}
		e := strings.IndexRune(s, '.')
		temps1 := []byte(s[b:e])
		// 反转
		for i, j := 0, len(temps1) - 1;i < j;i, j = i + 1, j - 1 {
			temps1[i], temps1[j] = temps1[j], temps1[i]
		}
		temps2 := []byte{}
		// 添加 ,
		for i := 0;i < len(temps1);i++ {
			if i % 3 == 0 && i != 0 {
				temps2 = append(temps2, ',')
			}
			temps2 = append(temps2, temps1[i])
		}
		// 反转
		for i, j := 0, len(temps2) - 1;i < j;i, j = i + 1, j - 1 {
			temps2[i], temps2[j] = temps2[j], temps2[i]
		}
		// 加首尾
		n = ""
		if b != 0 {
			n = "("
		}
		n += string(temps2)
		n += s[e:]

		fmt.Println(n)
	}
}
