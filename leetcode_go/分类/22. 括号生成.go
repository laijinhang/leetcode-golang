package main

func generateParenthesis(n int) []string {
	res := []string{}
	gp(n, 0, 0, "", &res)
	return res
}

func gp(n, l, r int, str string, strs *[]string) {
	if l == n && r == n {
		*strs = append(*strs, str)
		return
	}
	if l < n {
		gp(n, l + 1, r, str + "(", strs)
	}
	if r < n && r < l {
		gp(n, l, r + 1, str + ")", strs)
	}
}