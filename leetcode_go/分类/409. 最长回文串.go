package main

func longestPalindrome(s string) int {
	charsL := make([]int, 26)
	charsU := make([]int, 26)
	for i := 0;i < len(s);i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			charsL[s[i]-'a']++
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			charsU[s[i]-'A']++
		}
	}
	singular := 0	// 成单的个数
	even := 0 		// 成对的个数
	for i := 0;i < 26;i++ {
		singular += charsL[i] % 2
		singular += charsU[i] % 2
		even += charsL[i] / 2
		even += charsU[i] / 2
	}
	even *= 2
	if singular != 0 {
		even++
	}
	return even
}