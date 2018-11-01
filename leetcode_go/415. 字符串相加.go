package main

import (
	"math"
	"fmt"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	sum := make([]byte, int(math.Max(float64(len(num1)), float64(len(num2)))) + 1)
	jw := byte(0)
	i, j, k := len(num1) - 1, len(num2) - 1, 0
	for ;i >= 0 && j >= 0;i, j = i - 1, j - 1 {
		sum[k] = (num1[i] + num2[j] - '0' - '0' + jw) % 10 + '0'
		jw = (num1[i] + num2[j] - '0' - '0' + jw) / 10
		k++
	}
	for ;i >= 0;i-- {
		sum[k] = (num1[i] - '0' + jw) % 10 + '0'
		jw = (num1[i] - '0' + jw) / 10
		k++
	}
	for ;j >= 0;j-- {
		sum[k] = (num2[j] - '0' + jw) % 10 + '0'
		jw = (num2[j] - '0' + jw) / 10
		k++
	}
	sum[k] = jw + '0'
	for i, j := 0, len(sum) - 1;i < j;i, j = i + 1, j - 1 {
		sum[i], sum[j] = sum[j], sum[i]
	}
	s := strings.TrimLeft(string(sum), "0")
	if s == "" {
		s = "0"
	}
	return s
}

func main()  {
	fmt.Println(addStrings("0", "1"))
}