package main

func validPalindrome(s string) bool {
	res := true
	if len(s) <= 2 {
		res = true
	} else {
		for i, j := 0, len(s) - 1;i < j;i, j = i + 1, j - 1 {
			if (s[i] != s[j] && s[i+1] == s[j] && isValid(s[i+1:j+1])) || (s[i] != s[j] && s[i] == s[j-1] && isValid(s[i:j])) {
				break
			} else if s[i] != s[j] {
				res = false
				break
			}
		}
	}
	return res
}

func isValid(s string) bool {
	flag := true
	for i, j := 0, len(s) - 1;i < j;i, j = i + 1, j - 1 {
		if s[i] != s[j] {
			flag = false
			break
		}
	}
	return flag
}