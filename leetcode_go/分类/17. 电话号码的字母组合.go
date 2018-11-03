package main

var refNumLetter = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	strs := []string{}
	combinations(digits, "", 0, &strs)
	return strs
}

func combinations(digits, ct string, i int, strs *[]string) {
	if i >= len(digits) {
		*strs = append(*strs, ct)
		return
	}
	combinations(digits, ct + string(refNumLetter[digits[i]][0]), i + 1, strs)
	combinations(digits, ct + string(refNumLetter[digits[i]][1]), i + 1, strs)
	combinations(digits, ct + string(refNumLetter[digits[i]][2]), i + 1, strs)
	if digits[i] == '7' || digits[i] == '9' {
		combinations(digits, ct + string(refNumLetter[digits[i]][3]), i + 1, strs)
	}
}