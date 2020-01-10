package main

func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	for i := 0;i < len(words);i++ {
		w := make(map[byte]byte)
		p := make(map[byte]byte)
		j := 0
		for ;j < len(words[i]);j++ {
			if _, ok := w[words[i][j]];ok {
				if w[words[i][j]] != pattern[j] {
					break
				}
			} else {
				w[words[i][j]] = pattern[j]
			}
			if _, ok := p[pattern[j]];ok {
				if p[pattern[j]] != words[i][j] {
					break
				}
			} else {
				p[pattern[j]] = words[i][j]
			}
		}
		if j == len(words[i]) {
			res = append(res, words[i])
		}
	}
	return res
}