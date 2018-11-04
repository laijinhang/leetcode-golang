package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	tempStr := ""
	sStack := make([]string, len(s))
	n := 0
	tempI := 0
	for i := 0;i < len(s);i++ {
		if s[i] == '[' {
			sStack[n] = "["
			n++
		} else if s[i] == ']' {
			for ;n - 2 >= 0 && sStack[n-2] != "[";n-- {	// 向前合并
				sStack[n-2] += sStack[n-1]
			}
			tempStr = ""
			tempI, _ := strconv.Atoi(sStack[n-3])
			for i := 0;i < tempI;i++ {
				tempStr += sStack[n-1]
			}
			sStack[n-3] = tempStr
			n -= 2
			for ;n - 2 > 0 && sStack[n-2] != "[";n-- {	// 合并
				sStack[n-2] += sStack[n-1]
			}
		} else {
			if s[i] >= '0' && s[i] <= '9' {
				sStack[n], tempI = getK(s[i:])
				i += tempI
			} else {
				sStack[n], tempI = getStr(s[i:])
				i += tempI
			}
			n++
		}
	}
	res := ""
	for i := 0;i < n;i++ {
		res += sStack[i]
	}
	return res
}

func getK(s string) (string, int) {
	k := 0
	i := 0
	for ;i < len(s) && s[i] != '[';i++ {
		k = k * 10 + int(s[i] - '0')
	}
	return strconv.Itoa(k), i - 1
}

func getStr(s string) (string, int) {
	i := 0
	for ;i < len(s) && !(s[i] >= '0' && s[i] <= '9') && s[i] != '[' && s[i] != ']';i++ {}
	return s[:i], i - 1
}

func main() {
	fmt.Println(decodeString("3[a]2[b4[F]c]"))
}