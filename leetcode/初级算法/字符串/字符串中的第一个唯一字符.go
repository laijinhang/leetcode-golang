package main

func firstUniqChar(s string) int {
	temp := [26]int{}
	sLen := len(s)
	for i := 0;i < sLen;i++ {
		temp[s[i]-'a']++
	}
	for i := 0;i < sLen;i++ {
		if temp[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}