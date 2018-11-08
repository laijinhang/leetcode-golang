package main

import "fmt"

func addBinary(a string, b string) string {
	sum := []byte{}
	jz := 0
	i, j := len(a) - 1, len(b) - 1
	for ;i >= 0 && j >= 0;i, j = i - 1, j - 1 {
		fmt.Println(jz)
		if a[i] + b[j] + byte(jz) - '0' == '2' {
			sum = append(sum, '0')
			jz = 1
		} else if a[i] + b[j] + byte(jz) - '0' == '3' {
			sum = append(sum, '1')
			jz = 1
		} else {
			sum = append(sum, a[i] + b[j] + byte(jz) - '0')
			jz = 0
		}
	}
	for ;i >= 0;i-- {
		if a[i] + byte(jz) == '2' {
			sum = append(sum, '0')
			jz = 1
		} else {
			sum = append(sum, a[i] + byte(jz))
			jz = 0
		}
	}
	for ;j >= 0;j-- {
		if b[j] + byte(jz) == '2' {
			sum = append(sum, '0')
			jz = 1
		} else {
			sum = append(sum, b[j] + byte(jz))
			jz = 0
		}
	}
	if jz != 0 {
		sum = append(sum, '1')
	}
	for i, j = 0, len(sum) - 1;i < j;i, j = i + 1, j - 1 {
		sum[i], sum[j] = sum[j], sum[i]
	}
	return string(sum)
}

func main()  {
	fmt.Println(addBinary("1010", "1011"))
}