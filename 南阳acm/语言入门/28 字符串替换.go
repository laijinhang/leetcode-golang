/*
字符串替换
时间限制：3000 ms  |  内存限制：65535 KB
难度：2
描述
编写一个程序实现将字符串中的所有"you"替换成"we"
输入
输入包含多行数据

每行数据是一个字符串，长度不超过1000
数据以EOF结束
输出
对于输入的每一行，输出替换后的字符串
样例输入
you are what you do
样例输出
we are what we do
 */
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	for {
		var inputReader *bufio.Reader
		inputReader = bufio.NewReader(os.Stdin)
		str, err := inputReader.ReadString('\n')
		str = str[:len(str)-1]
		if err != nil {
			break
		}
		fmt.Printf("%s\n", strings.Replace(str, "you", "we", -1))
	}
}