package main

/*
这题题目说是s的长度小于等于100，所以写代码的时候要把长度为0的时候也要计算上去，要考虑 s = ""这种情况
 */
func isSubsequence(s string, t string) bool {
	var i, j int
	sLen := len(s)
	tLen := len(t)
	if sLen == 0 {
		return true
	}
	for i = 0;i < tLen;i++ {
		if t[i] == s[j] {
			j++
			if j == sLen {
				break
			}
		}
	}
	if j != sLen {
		return false
	}
	return true
}