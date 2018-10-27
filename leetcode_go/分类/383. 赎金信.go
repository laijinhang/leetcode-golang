package main

func canConstruct(ransomNote string, magazine string) bool {
	mr := make(map[byte]int)
	mm := make(map[byte]int)
	for i := 0;i < len(ransomNote);i++ {
		mr[ransomNote[i]]++
	}
	for i := 0;i < len(magazine);i++ {
		mm[magazine[i]]++
	}
	for k, v := range mr {
		if mm[k] < v {
			return false
		}
	}
	return true
}