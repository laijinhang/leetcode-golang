/*
题目描述
利用字符重复出现的次数，编写一个方法，实现基本的字符串压缩功能。比如，字符串“aabcccccaaa”经压缩会变成“a2b1c5a3”。若压缩后的字符串没有变短，则返回原先的字符串。

给定一个string iniString为待压缩的串(长度小于等于10000)，保证串内字符均由大小写英文字母组成，返回一个string，为所求的压缩后或未变化的串。

测试样例
"aabcccccaaa"
返回："a2b1c5a3"
"welcometonowcoderrrrr"
返回："welcometonowcoderrrrr"
 */
package main

import "strconv"

func zipString(iniString string) string {
	var str string
	length := len(iniString)

	for i, j, len1 := 0, 0, 1;i < length;i++ {
		len1 = 1
		for j = i + 1;j < length && iniString[j] == iniString[j-1];j++ {
			len1++
		}
		if len1 > 1 {
			str += string(iniString[i]) + strconv.Itoa(len1)
			i += len1 - 1
		}
	}
	if length == len(str) {
		str = iniString
	}

	return str
}
