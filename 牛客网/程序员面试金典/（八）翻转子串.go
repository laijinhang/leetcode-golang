/*
题目描述
假定我们都知道非常高效的算法来检查一个单词是否为其他字符串的子串。请将这个算法编写成一个函数，给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成，要求只能调用一次检查子串的函数。

给定两个字符串s1,s2,请返回bool值代表s2是否由s1旋转而成。字符串中字符为英文字母和空格，区分大小写，字符串长度小于等于1000。

测试样例：
"Hello world","worldhello "
返回：false
"waterbottle","erbottlewat"
返回：true

分析：如果str2是str1的翻转子串，则可以推出以下两个结果
1）str2的长度等于str1
2）str2一定是str1+str1的子串
     abcd的所有翻转子串：abcd -> bcda -> cdab -> dabc
                                          abcdabcd
 */
package main

import (
	"strings"
	"fmt"
)

func checkReverseEqual(str1, str2 string) bool {
	len1, len2 := len(str1), len(str2)
	if len1 != len2 {
		return false
	}
	str1 += str1
	return strings.Contains(str1, str2)
}

func main()  {
	fmt.Println(checkReverseEqual("abc", "bca"))
}