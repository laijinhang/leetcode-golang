package main

func numberOfLines(widths []int, S string) []int {
	h, num := 1, 0
	for i := 0;i < len(S);i++ {
		if num + widths[S[i]-'a'] > 100 {
			num = widths[S[i]-'a']
			h++
		} else {
			num += widths[S[i]-'a']
		}
	}
	return []int{h,num}
}