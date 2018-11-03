package main

/*
	回文分两种，偶数个和奇数个的这两种情况
 */
func countSubstrings(s string) int {
	count := 0
	num := 0
	for i := 0;i < len(s);i++ {
		for num = 0;i - num >= 0 && i + num < len(s) && s[i-num] == s[i+num];num++ {		// 奇数
			count++
		}
		for num = 0;i - num >= 0 && i + num + 1 < len(s) && s[i-num] == s[i+num+1];num++ {	// 偶数
			count++
		}
	}
	return count
}
