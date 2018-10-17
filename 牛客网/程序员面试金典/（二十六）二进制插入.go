/*
题目描述
有两个32位整数n和m，请编写算法将m的二进制数位插入到n的二进制的第j到第i位,其中二进制的位数从低位数到高位且以0开始。

给定两个数int n和int m，同时给定int j和int i，意义如题所述，请返回操作后的数，保证n的第j到第i位均为零，且m的二进制位数小于等于i-j+1。

测试样例：
1024，19，2，6
返回：1100
 */
 package main

import (
	"strconv"
	"fmt"
)

func binInsert(n, m, j, i int) int {
	// 获取n的二进制表示
	bi := []byte(strconv.FormatInt(int64(n), 2))
	// 获取m的二进制表示
	mi := []byte(strconv.FormatInt(int64(m), 2))
	// n取反
	for i, j := 0, len(bi) - 1;i < j;i, j = i + 1, j - 1 {
		bi[i], bi[j] = bi[j], bi[i]
	}
	// m取反
	for i, j := 0, len(mi) - 1;i < j;i, j = i + 1, j - 1 {
		mi[i], mi[j] = mi[j], mi[i]
	}
	// 替换
	for k := 0;k < len(mi);k++ {
		bi[j+k] = mi[k]
	}
	// 反转
	for i, j := 0, len(bi) - 1;i < j;i, j = i + 1, j - 1 {
		bi[i], bi[j] = bi[j], bi[i]
	}
	r, _ := strconv.ParseInt(string(bi), 2, 0)
	return int(r)
}

func main() {
	fmt.Println(binInsert(1024, 19, 2, 6))
}
