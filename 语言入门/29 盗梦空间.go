/*
盗梦空间
时间限制：3000 ms  |  内存限制：65535 KB
难度：2
描述
《盗梦空间》是一部精彩的影片，在这部电影里，Cobb等人可以进入梦境之中，梦境里的时间会比现实中的时间过得快得多，这里假设现实中的3分钟，在梦里就是1小时。

然而，Cobb他们利用强效镇静剂，可以从第一层梦境进入第二层梦境，甚至进入三层，四层梦境，每层梦境都会产生同样的时间加速效果。那么现在给你Cobb在各层梦境中经历的时间，你能算出现实世界过了多长时间吗？

比如，Cobb先在第一层梦境待了1个小时，又在第二层梦境里待了1天，之后，返回第一层梦境之后立刻返回了现实。

那么在现实世界里，其实过了396秒（6.6分钟）

输入
第一行输入一个整数T（0<=T<=100)，表示测试数据的组数。
每组测试数据的第一行是一个数字M(3<=M<=100)
随后的M行每行的开头是一个字符串，该字符串如果是"IN" 则Cobb向更深层的梦境出发了，如果是字符串"OUT"则表示Cobb从深层的梦回到了上一层。如果是首字符串是"STAY"则表示Cobb在该层梦境中停留了一段时间，本行随后将是一个整数S表示在该层停留了S分钟（1<=S<=10000000)。数据保证在现实世界中，时间过了整数秒。
输出
对于每组测试数据，输出现实世界过的时间（以秒为单位）。
样例输入
1
6
IN
STAY 60
IN
STAY 1440
OUT
OUT
样例输出
396
 */
package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {
	var M, T, t, layer int
	fmt.Scan(&M)
	for ;M > 0;M-- {
		fmt.Scan(&T)
		layer = 0
		for ;T > 0;T-- {
			var inputReader *bufio.Reader
			inputReader = bufio.NewReader(os.Stdin)
			str, err := inputReader.ReadString('\n')
			if err != nil {
				break
			}
			str = str[:len(str)-1]
			if str == "IN" {
				layer++
			} else if str == "OUT" {
				layer--
			} else {
				tempT, _ := strconv.Atoi(strings.Split(str, " ")[1])
				tempT *= 60	// 转化成秒
				for i := layer;i > 0;i-- {
					tempT /= 20
				}
				t += tempT
			}
		}
		fmt.Println(t)
	}
}