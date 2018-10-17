/*
题目描述
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。这里规定大小写为不同字符，且考虑字符串中的空格。

给定一个string stringA和一个string stringB，请返回一个bool，代表两串是否重新排列后可相同。保证两串的长度都小于等于5000。

测试样例：
"This is nowcoder","is This nowcoder"
返回：true
"Here you are","Are you here"
返回：false

分析：
1）如果长度不想等，则直接返回
2）我这边的方法是第二个字符串第j个字符与第一个字符串第i个进行比较，
      i是从0开始
      j是从第i个开始，
            如果str1[i] == str2[j]，交换str2[i]，str2[j]，之后i++,j=i
            如果j 等于字符串长度则说明，str1和str2有不同，则可以返回false。
 */
package main

import "fmt"

func checkSam(str1 , str2 string) bool {
	length1, length2 := len(str1), len(str2)
	b1 := []byte(str1)
	b2 := []byte(str2)
	if length1 != length2 {
		return false
	}
	for i, j := 0, 0;i < length1;i++ {
		for j = i; j < length1; j++ {
			if b1[i] == b2[j] {
				b2[i], b2[j] = b2[j], b2[i]
				break
			}
		}
		if j == length1 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "est1 t"
	str2 := "test 1"
	if checkSam(str1, str2) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
