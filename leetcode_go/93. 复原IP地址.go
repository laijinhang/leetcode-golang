package main

import "fmt"

func restoreIpAddresses(s string) []string {
	res := []string{}
	for i := 1;i < len(s) && i <= 3 && (len(s[:i]) == 1 || (len(s[:i]) == 2 && s[:i][0] != '0') || (len(s[:i]) == 3 && s[:i][0] != '0')) && (len(s[:i]) != 3 || (len(s[:i]) == 3 && s[:i] <= "255"));i++ {
		for j := i + 1;j < len(s) && j <= 6 && (len(s[i:j]) == 1 || (len(s[i:j]) == 2 && s[i:j][0] != '0') || (len(s[i:j]) == 3 && s[i:j][0] != '0')) && (len(s[i:j]) != 3 || (len(s[i:j]) == 3 && s[i:j] <= "255"));j++ {
			for k := j + 1;k < len(s) && k <= 9 && (len(s[j:k]) == 1 || (len(s[j:k]) == 2 && s[j:k][0] != '0') || (len(s[j:k]) == 3 && s[j:k][0] != '0')) && (len(s[j:k]) != 3 || (len(s[j:k]) == 3 && s[j:k] <= "255"));k++ {
				if (len(s[k:]) == 1 || (len(s[k:]) == 2 && s[k:][0] != '0') || (len(s[k:]) == 3 && s[k:][0] != '0')) && (len(s[k:]) != 3 || (len(s[k:]) == 3 && s[k:] <= "255")) {
					res = append(res, fmt.Sprintf("%s.%s.%s.%s", s[:i], s[i:j], s[j:k], s[k:]))
				}
			}
		}
	}
	return res
}

func main() {
	res := restoreIpAddresses("010010")
	fmt.Println(res)
}
