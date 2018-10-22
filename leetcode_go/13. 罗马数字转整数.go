package main

import "fmt"

var rm = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000}

func romanToInt(s string) int {
	n := 0
	sLen := len(s)

	for i := 0;i < sLen;i++ {
		if i != sLen - 1 {
			if s[i] == 'I' && s[i+1] == 'V' {
				n += 4
				i++
			} else if s[i] == 'I' && s[i+1] == 'X' {
				n += 9
				i++
			} else if s[i] == 'X' && s[i+1] == 'L' {
				n += 40
				i++
			} else if s[i] == 'X' && s[i+1] == 'C' {
				n += 90
				i++
			} else if s[i] == 'C' && s[i+1] == 'D' {
				n += 400
				i++
			} else if s[i] == 'C' && s[i+1] == 'M' {
				n += 900
				i++
			} else {
				n += rm[s[i]]
			}
			continue
		}
		n += rm[s[i]]
	}
	return n
}

func main()  {
	fmt.Println(romanToInt("MCMXCIV"))
}