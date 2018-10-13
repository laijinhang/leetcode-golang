/*
关于521
时间限制：1000 ms  |  内存限制：65535 KB
难度：2
描述
Acm队的流年对数学的研究不是很透彻，但是固执的他还是想一头扎进去。

浏览网页的流年忽然看到了网上有人用玫瑰花瓣拼成了521三个数字，顿时觉得好浪漫，因为每个男生都会不经意的成为浪漫的制造者。此后，流年走到哪里都能看到5、2、1三个数字，他怒了，现在他想知道在连续的数中有多少数全部包含了这三个数字。例如12356就算一个，而5111就不算。特别的，如果他看到了521三个数连续出现，会特别的愤怒。例如35210。

输入
多组测试数据：
一行给定两个数a，b（0<a，b<1000000），表示数字的开始和结束。
输出
一行显示他想要知道的数有几个及显示有多少个数字令他特别的愤怒。用空格隔开。
样例输入
200 500
300 900
1 600
样例输出
Case 1:2 0
Case 2:2 1
Case 3:6 1
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b, n1, n2 int
	for i := 1;;i++ {
		if _, err := fmt.Scan(&a, &b); err != nil {
			break
		}
		for n1, n2 = 0, 0;a <= b;a++ {
			s := strconv.Itoa(a)
			if strings.IndexRune(s, '1') != -1 &&
				strings.IndexRune(s, '2') != -1 &&
				strings.IndexRune(s, '5') != -1 {
				n1++
			}
			if strings.Index(s, "521") != -1 {
				n2++
			}
		}
		fmt.Printf("Case %d:%d %d\n", i, n1, n2)
	}
}