package main

import (
	"strconv"
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")	// 去掉前后空格
	sLen := len(str)
	if  sLen == 0 || str[0] != '-' && str[0] != '+' && (len(str) == 0 || str[0] < '0' || str[0] > '9') {	// 去掉最前面的空格之后的字符串首字符非数字，则返回0
		return 0
	}
	temp := []byte{str[0]}
	for i := 1;i < sLen;i++ {
		if str[i] >= '0' && str[i] <= '9' {
			temp = append(temp, str[i])
		} else {
			break
		}
	}
	val, _ := strconv.Atoi(string(temp))
	if val < math.MinInt32 {
		val = math.MinInt32
	} else if val > math.MaxInt32 {
		val = math.MaxInt32
	}
	return int(val)
}

func main()  {
	fmt.Println(myAtoi("   +42-"))
}