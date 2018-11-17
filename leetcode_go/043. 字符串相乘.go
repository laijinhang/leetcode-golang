package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	// 用短的数字乘长的数字
	tempRes := make([]byte, len(num1) + len(num2))
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	for i := len(num1) - 1;i >= 0;i-- {
		tempI := len(num1) - i - 1
		ajw, cjw := 0, 0
		for j := len(num2) - 1;j >= 0;j-- {
			tempRes[tempI], ajw = byte(((int(num1[i] - '0') * int(num2[j] - '0')) + cjw) % 10 + int(tempRes[tempI]) + ajw) % 10,
				((int(num1[i] - '0') * int(num2[j] - '0') + cjw) % 10 + int(tempRes[tempI]) + ajw) / 10
			cjw = (int(num1[i] - '0') * int(num2[j] - '0') + cjw) / 10
			tempI++
		}
		for ;tempI < len(tempRes);tempI++ {
			tempRes[tempI], ajw = byte(cjw + int(tempRes[tempI]) + ajw) % 10, (cjw + int(tempRes[tempI]) + ajw) / 10
			cjw = cjw / 10
		}
	}
	i, j := 0, len(tempRes) - 1
	for ;j >= 0 && tempRes[j] == byte(0);j-- {}
	tempRes = tempRes[:j+1]
	for ;i <= j;i, j = i + 1,j - 1 {
		tempRes[i], tempRes[j] = tempRes[j] + '0', tempRes[i] + '0'
	}
	if len(tempRes) == 0 {
		tempRes = []byte{'0'}
	}
	return string(tempRes)
}

func main() {
	fmt.Println(multiply("123456789", "987654321"))
}